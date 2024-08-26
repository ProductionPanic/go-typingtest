package api

import (
	"encoding/json"
	"io"
	"net/http"
)

const api_url = "https://api.quotable.io/random?minLength=100"

type ApiResponse struct {
	Id         string   `json:"_id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	Tags       []string `json:"tags"`
	AuthorSlug string   `json:"authorSlug"`
	Length     int      `json:"length"`
}

// GetRandomQuote returns a random quote from the API
func GetRandomQuote() (ApiResponse, error) {
	var response ApiResponse
	resp, err := http.Get(api_url)
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
