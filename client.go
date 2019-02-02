package forge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Autodesk API Client. Use this to authenticate and make API calls to the
// Autodesk API
type Client struct {
	client       *http.Client
	clientId     string
	clientSecret string
	baseURL      string
	jwt          JWT
	scopes       []string
}

// Returns a client, assumes client id and client secret are stored as
// environment variables as FORGE_CLIENT_ID and FORGE_CLIENT_SECRET
func New() (c Client, err error) {
	var clientId string
	var clientSecret string

	clientId = os.Getenv("FORGE_CLIENT_ID")
	clientSecret = os.Getenv("FORGE_CLIENT_SECRET")

	return NewWithCreds(clientId, clientSecret)
}

// Returns a client with given client id and client secret
func NewWithCreds(clientId, clientSecret string) (Client, error) {
	if clientId == "" || clientSecret == "" {
		return Client{}, fmt.Errorf("A forge client must have a client id and client secret")
	}

	return Client{
		clientId:     clientId,
		clientSecret: clientSecret,
		baseURL:      "https://developer.api.autodesk.com", // TODO: don't hard code the base URL
		client: &http.Client{
			Timeout: time.Second * 10, // TODO: don't hard code the timeout
		},
	}, nil
}

// Build full URL with given path
func (cl *Client) Path(p string) string {
	return fmt.Sprintf("%s%s", cl.baseURL, p)
}

// Authenticate our forge client to be able to make API calls
func (cl *Client) Authenticate(scopes []string) error {
	if len(scopes) == 0 {
		return fmt.Errorf("Must have at least one scope defined (ie `data:read` and/or `bucket:create`)")
	}

	grant_type := "client_credentials" // TODO: don't hard code this, but make it default
	u := cl.Path("/authentication/v1/authenticate")
	values := url.Values{
		"grant_type":    {grant_type},
		"client_id":     {cl.clientId},
		"client_secret": {cl.clientSecret},
		"scope":         scopes,
	}

	resp, err := cl.client.PostForm(u, values)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &cl.jwt)
	if err != nil {
		return err
	}
	cl.jwt.SetExpiration()

	cl.scopes = scopes
	return nil
}
