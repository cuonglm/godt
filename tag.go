package godt

import (
	"fmt"
	"net/http"
)

var tagsPath = map[string]string{
	"1": "/v1/repositories/%s/tags",
	"2": "/v2/%s/tags/list",
}

// ListTags list all image tags
func (c *Client) ListTags(image string) (*http.Response, error) {
	path := fmt.Sprintf(tagsPath[c.APIVersion], image)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
