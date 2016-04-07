package godt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

const (
	version           = "0.0.1"
	ua                = "godt/" + version
	mediaType         = "application/json"
	envHubAPIVersion  = "GODT_HUB_API_VERSION"
	envHubURL         = "GODT_HUB_URL"
	defaultAPIVersion = "1"
	defaultHubURL     = "https://registry.hub.docker.com"
)

// Client for querying docker hub API
type Client struct {
	client     *http.Client
	UserAgent  string
	APIVersion string
	HubURL     *url.URL
}

// NewHTTPClient create new godt client
func NewHTTPClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	apiVersion := os.Getenv(envHubAPIVersion)
	if apiVersion == "" {
		apiVersion = defaultAPIVersion
	}

	hubURLStr := os.Getenv(envHubURL)
	if hubURLStr == "" {
		hubURLStr = defaultHubURL
	}
	hubURL, err := url.Parse(hubURLStr)
	if err != nil {
		panic(err)
	}

	c := &Client{
		client:     client,
		UserAgent:  ua,
		APIVersion: apiVersion,
		HubURL:     hubURL,
	}

	return c
}

// NewRequest create new http request
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	relPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.HubURL.ResolveReference(relPath)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do make an http request
func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
