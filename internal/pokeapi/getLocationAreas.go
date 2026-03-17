package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaData(locationNameOrId *string) (LocationAreaInfo, error) {
	url := baseURL + "/location-area/" + *locationNameOrId
	var reqBody []byte

	entry, exist := c.cache.Get(url)

	if exist {
		fmt.Println("Getting from cache...")
		reqBody = entry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaInfo{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaInfo{}, err
		}
		defer res.Body.Close()

		reqBody, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaInfo{}, err
		}
		c.cache.Add(url, reqBody)
	}
	var locationAreaInfo LocationAreaInfo
	if err := json.Unmarshal(reqBody, &locationAreaInfo); err != nil {
		return LocationAreaInfo{}, err
	}

	return locationAreaInfo, nil
}
