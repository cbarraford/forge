package forge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Manifest struct {
	Type     string `json:"type"`
	Status   string `json:"status"`
	Progress string `json:"progress"`
	// TODO: add more manifest attributes
}

// Get the manifest of a job
func (cl *Client) Manifest(urn string) (Manifest, error) {
	mani := Manifest{}

	if urn == "" {
		return mani, fmt.Errorf("Must have an URN")
	}

	u := cl.Path(
		fmt.Sprintf("/modelderivative/v2/designdata/%s/manifest", urn),
	)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return mani, err
	}
	req.Header.Add("Authorization", cl.jwt.GetAuthHeader())
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.client.Do(req)
	if err != nil {
		return mani, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return mani, err
	}

	if resp.StatusCode >= 400 {
		log.Printf("Failed to create job, got http error %d", resp.StatusCode)
		return mani, fmt.Errorf(string(body))
	}

	log.Printf("Manifest: %s", string(body))
	return mani, nil

	err = json.Unmarshal(body, &mani)
	if err != nil {
		return mani, err
	}

	return mani, nil
}
