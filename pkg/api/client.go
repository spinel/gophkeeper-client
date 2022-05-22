package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type client struct {
	http http.Client
	url  string
}

func NewClient(url string) *client {
	return &client{
		http: http.Client{},
		url:  url,
	}
}

func (c *client) Post(path string, st interface{}, target interface{}, token string) error {
	body, err := json.Marshal(st)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.url, path), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	r, err := c.http.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(r.Body).Decode(target)
}
