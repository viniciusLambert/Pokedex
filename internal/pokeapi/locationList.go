// Package pokeapi Manage pokemon API
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	var reqBody []byte

	if pageURL != nil {
		url = *pageURL
	}

	entry, exist := c.cache.Get(url)
	if exist {
		reqBody = entry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer res.Body.Close()

		reqBody, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, reqBody)
	}

	var locationList RespShallowLocations
	if err := json.Unmarshal(reqBody, &locationList); err != nil {
		return RespShallowLocations{}, err
	}

	return locationList, nil
}
