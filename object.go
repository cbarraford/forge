package forge

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type urn string

type Object struct {
	BucketName  string `json:"bucketKey"`
	Id          urn    `json:"objectId"`
	Name        string `json:"objectKey"`
	Sha1        string `json:"sha1"`
	ContentSize int    `json:"size"`
	ContentType string `json:"contentType"`
	Location    string `json:"location"`
}

// Upload a file
func (cl *Client) ObjectUpload(obj *Object, file io.Reader) error {
	if obj.BucketName == "" {
		return fmt.Errorf("Must have a bucket name")
	}
	if obj.Name == "" {
		return fmt.Errorf("Must have a file name")
	}
	if obj.ContentType == "" {
		return fmt.Errorf("Must have a content type")
	}
	if obj.ContentSize == 0 {
		return fmt.Errorf("Must have a content size greater than zero")
	}
	if !cl.CheckScope("data:write") && !cl.CheckScope("data:create") {
		return fmt.Errorf("Incorrect scope: missing data:write or data:create")
	}

	u := cl.Path(
		fmt.Sprintf("/oss/v2/buckets/%s/objects/%s", obj.BucketName, obj.Name),
	)

	req, err := http.NewRequest("PUT", u, file)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", cl.jwt.GetAuthHeader())
	req.Header.Add("Content-Type", obj.ContentType)
	req.Header.Add("Content-Length", strconv.Itoa(obj.ContentSize))
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
		log.Printf("Failed to create object, got http error %d", resp.StatusCode)
		return fmt.Errorf(string(body))
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}

	return nil
}
