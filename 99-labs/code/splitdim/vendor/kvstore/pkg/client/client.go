package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"kvstore/pkg/api" // This is needed to use our own key-value store API.
)

type client struct {
	url string
}

func NewClient(addr string) Client {
	return &client{url: "http://" + addr}
}

// Get returns the value and version stored for the given key, or an error if something goes wrong.
func (c *client) Get(key string) (api.VersionedValue, error) {
	// this will package the requested key into the body of the HTTP request
	body := []byte(fmt.Sprintf("\"%s\"", key))

	r, err := http.Post(c.url+"/api/get", "application/json", bytes.NewReader(body))
	if err != nil {
		err2 := errors.New("couldn't call the endpoint")
		return api.VersionedValue{}, err2
	}
	if r.StatusCode != 200 {
		err2 := errors.New("couldn't call the endpoint")
		return api.VersionedValue{}, err2
	}
	var versionedValue api.VersionedValue
	err = json.NewDecoder(r.Body).Decode(&versionedValue)
	if err != nil {
		err2 := errors.New("couldn't decode json object")
		return api.VersionedValue{}, err2
	}
	return versionedValue, nil
}

// Put tries to insert the given key-value pair with the specified version into the store.
func (c *client) Put(vkv api.VersionedKeyValue) error {
	json, err := json.Marshal(vkv)
	if err != nil {
		return fmt.Errorf("cannot marshall object: %w", err)
	}
	response, err := http.Post(c.url+"/api/put", "application/json", bytes.NewReader(json))
	if err != nil || response.StatusCode != 200 {
		return fmt.Errorf("cannot call the endpoint correctly: %w", err)
	}
	return nil
}

// List returns all values stored in the database.
func (c *client) List() ([]api.VersionedKeyValue, error) {
	response, err := http.Get(c.url + "/api/list")
	if err != nil || response.StatusCode != 200 {
		return []api.VersionedKeyValue{}, fmt.Errorf("something went wrong during api call: %w", err)
	}
	var values []api.VersionedKeyValue
	err = json.NewDecoder(response.Body).Decode(&values)
	if err != nil {
		return []api.VersionedKeyValue{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return values, nil
}

// Reset removes all key-value pairs.
func (c *client) Reset() error {
	response, err := http.Get(c.url + "/api/reset")
	if err != nil || response.StatusCode != 200 {
		return fmt.Errorf("couldn't call reset: %w", err)
	}
	return nil
}
