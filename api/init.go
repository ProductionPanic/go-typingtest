package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const api_url = "https://api.quotable.io/random"

type ApiArguments struct {
	MinLength int `json:"minLength"`
	MaxLength int `json:"maxLength"`
}

type ApiResponse struct {
	Id         string   `json:"_id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	Tags       []string `json:"tags"`
	AuthorSlug string   `json:"authorSlug"`
	Length     int      `json:"length"`
}

// GetRandomQuote returns a random quote from the API
func GetRandomQuote(
	args ApiArguments,
) (ApiResponse, error) {
	var response ApiResponse
	url := createUrl(api_url, args)
	resp, err := http.Get(url)
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

func createUrl(url string, args ApiArguments) string {
	has_min := args.MinLength != -1
	has_max := args.MaxLength != -1
	if has_min && has_max {
		return fmt.Sprintf("%s?minLength=%d&maxLength=%d", url, args.MinLength, args.MaxLength)
	}
	if has_min {
		return fmt.Sprintf("%s?minLength=%d", url, args.MinLength)
	}
	if has_max {
		return fmt.Sprintf("%s?maxLength=%d", url, args.MaxLength)
	}
	return url
}
