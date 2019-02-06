package forge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Metadata struct {
}

// Get the metadata of a URN
func (cl *Client) Metadata(u string) (Metadata, error) {
	meta := Metadata{}

	if u == "" {
		return meta, fmt.Errorf("Must have an URN")
	}

	sEnc := base64.StdEncoding.EncodeToString([]byte(u))

	p := cl.Path(
		fmt.Sprintf("/modelderivative/v2/designdata/%s/metadata", sEnc),
	)

	req, err := http.NewRequest("GET", p, nil)
	if err != nil {
		return meta, err
	}
	req.Header.Add("Authorization", cl.jwt.GetAuthHeader())
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.client.Do(req)
	if err != nil {
		return meta, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return meta, err
	}

	if resp.StatusCode >= 400 {
		log.Printf("Failed to get metadata, got http error %d", resp.StatusCode)
		return meta, fmt.Errorf(string(body))
	}

	log.Printf(string(body))
	return meta, nil

	err = json.Unmarshal(body, &meta)
	if err != nil {
		return meta, err
	}

	return meta, nil
}
