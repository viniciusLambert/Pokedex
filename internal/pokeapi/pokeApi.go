package pokeApi

// Manage pokemon API

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getMaps() []string {
	mapsSerial, err := makeMapRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mapsSerial)
	return []string{}
}

func makeMapRequest() ([]byte, error) {
	url := "https://pokeapi.co/api/v2/location-area/"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
