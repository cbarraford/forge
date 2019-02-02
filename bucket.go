package forge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A struct used for Buckets on Autodesk Management API
// https://forge.autodesk.com/en/docs/data/v2/reference/http/buckets-POST/
type Bucket struct {
	Name        string              `json:"bucketKey"`
	Policy      string              `json:"policyKey"`
	Owner       string              `json:"bucketOwner"`
	CreatedDate int                 `json:"createdDate"`
	Permissions []map[string]string `json:"permissions"`
}

// Create a bucket with given name and policy.
func (cl *Client) CreateBucket(b *Bucket) error {
	if b.Name == "" {
		return fmt.Errorf("Must have a bucket name")
	}
	if b.Policy == "" {
		return fmt.Errorf("Must have a policy name")
	}
	if b.Policy != "transient" && b.Policy != "temporary" && b.Policy != "persistent" {
		return fmt.Errorf("Bucket policy must be either 'transient', 'temporary', 'persistent'. For more information, see https://forge.autodesk.com/en/docs/data/v2/developers_guide/retention-policy/")
	}
	if !cl.CheckScope("bucket:create") {
		return fmt.Errorf("Incorrect scope: missing bucket:create")
	}

	u := cl.Path("/oss/v2/buckets")
	j, err := json.Marshal(b)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", cl.jwt.GetAuthHeader())
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		log.Printf("Failed to create bucket, got http error %d", resp.StatusCode)
		return fmt.Errorf(string(body))
	}

	err = json.Unmarshal(body, b)
	if err != nil {
		return err
	}

	return nil
}

// List all buckets.
func (cl *Client) ListBuckets() error {
	// https://forge.autodesk.com/en/docs/data/v2/reference/http/buckets-GET/
	return notImplemented
}

// Get bucket
func (cl *Client) GetBucket(name string) (Bucket, error) {
	// https://forge.autodesk.com/en/docs/data/v2/reference/http/buckets-:bucketKey-details-GET/
	return Bucket{}, notImplemented
}
