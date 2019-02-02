package forge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type InputJob struct {
	Input  Input  `json:"input"`
	Output Output `json:"output"`
}

type OutputJob struct {
	Result         string `json:"result"`
	URN            urn    `json:"urn"`
	AcceptableJobs struct {
		Output Output `json:"output"`
	} `json:"acceptedJobs"`
}

type Input struct {
	URN           string `json:"urn"`
	CompressedURN bool   `json:"compressedUrn"`
	RootFileName  string `json:"rootFilename"`
}

type Output struct {
	Formats []Format `json:"formats"`
}

type Format struct {
	Type  string   `json:"type"`
	Views []string `json:"views"`
}

// Process job for viewing
func (cl *Client) ConvertFile(input *InputJob) (OutputJob, error) {
	out := OutputJob{}

	if input.Input.URN == "" {
		return out, fmt.Errorf("Must have an URN")
	}
	if input.Output.Formats[0].Type == "" {
		return out, fmt.Errorf("Must have an output format type")
	}
	if len(input.Output.Formats[0].Views) == 0 {
		return out, fmt.Errorf("Must have at least one output format type view")
	}
	if !cl.CheckScope("data:write") && !cl.CheckScope("viewables:read") {
		return out, fmt.Errorf("Incorrect scope: missing data:write or viewables:read")
	}

	u := cl.Path("/modelderivative/v2/designdata/job")

	j, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(j))
	if err != nil {
		return out, err
	}
	req.Header.Add("Authorization", cl.jwt.GetAuthHeader())
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.client.Do(req)
	if err != nil {
		return out, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	if resp.StatusCode >= 400 {
		log.Printf("Failed to create job, got http error %d", resp.StatusCode)
		return out, fmt.Errorf(string(body))
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}
