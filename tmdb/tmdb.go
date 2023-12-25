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

type Config struct {
	API         string `yaml:"api"`
	APIKey      string `yaml:"api_key"`
	AccessToken string `yaml:"access_token"`
	ImageURL    string `yaml:"image_url"`
}

type Client struct {
	config *Config
	http   *http.Client
}

func NewClient(config *Config) (client *Client, err error) {
	client = &Client{config: config}
	client.http = http.DefaultClient
	if client.config.API == "" {
		client.config.API = "https://api.themoviedb.org/3"
	}
	if client.config.ImageURL == "" {
		client.config.ImageURL = "https://image.tmdb.org/t/p/"
	}
	return
}

func (client *Client) request(method string, path string, body io.Reader) (data []byte, err error) {
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
	var resp TMDBResponse
	json.Unmarshal(data, &resp)
	if resp.StatusCode != 0 {
		err = fmt.Errorf("%s (code: %d)", resp.StatusMessage, resp.StatusCode)
	}
	return
}

func (client *Client) get(path string, query map[string]string) (data []byte, err error) {
	var qs = url.Values{}
	if client.config.APIKey != "" {
		qs.Add("api_key", client.config.APIKey)
	}
	for k, v := range query {
		if v != "" {
			qs.Add(k, v)
		}
	}
	return client.request(http.MethodGet, path+"?"+qs.Encode(), nil)
}

func (client *Client) Authentication() (resp *TMDBResponse, err error) {
	data, err := client.get("/authentication", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

func (client *Client) GetImage(path string, size string) string {
	if path == "" {
		return ""
	}
	if size == "" {
		size = "original"
	}
	return fmt.Sprintf("%s%s%s", client.config.ImageURL, size, path)
}
