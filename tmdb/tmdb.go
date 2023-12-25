package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TMDBResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

type TMDBClientConfig struct {
	API         string
	APIKey      string
	AccessToken string
}

type TMDBClient struct {
	config *TMDBClientConfig
	http   http.Client
}

func NewClient(config *TMDBClientConfig) (client *TMDBClient, err error) {
	client = &TMDBClient{config: config}
	client.http = *http.DefaultClient
	if client.config.API == "" {
		client.config.API = "https://api.themoviedb.org/3"
	}
	return
}

func (client *TMDBClient) request(method string, path string, body io.Reader) (data []byte, err error) {
	req, err := http.NewRequest(method, client.config.API+path, body)
	if err != nil {
		return
	}
	if client.config.AccessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.config.AccessToken))
	}
	res, err := client.http.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}

func (client *TMDBClient) get(path string, params map[string]string) (data []byte, err error) {
	var qs = url.Values{}
	if client.config.APIKey != "" {
		qs.Add("api_key", client.config.APIKey)
	}
	for k, v := range params {
		if v != "" {
			qs.Add(k, v)
		}
	}
	return client.request(http.MethodGet, path+"?"+qs.Encode(), nil)
}

func (client *TMDBClient) Authentication() (resp *TMDBResponse, err error) {
	data, err := client.get("/authentication", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}
