package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemonNameOrId *string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + *pokemonNameOrId
	var reqBody []byte

	entry, exist := c.cache.Get(url)

	if exist {
		fmt.Println("Getting from cache...")
		reqBody = entry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			return Pokemon{}, fmt.Errorf("pokemon not found")
		}

		reqBody, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, reqBody)
	}
	var pokemon Pokemon
	if err := json.Unmarshal(reqBody, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
