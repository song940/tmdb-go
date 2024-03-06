package persistent

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/song940/tmdb-go/tmdb"
)

type Config struct {
	tmdb.Config

	PersistentPath string
}

type Client struct {
	*tmdb.Client
	*Config
}

func NewClient(config *Config) (*Client, error) {
	if config.PersistentPath == "" {
		userConfigDir, _ := os.UserConfigDir()
		config.PersistentPath = filepath.Join(userConfigDir, "tmdb")
	}
	if os.MkdirAll(config.PersistentPath, 0755) != nil {
		return nil, fmt.Errorf("failed to create persistent path: %s", config.PersistentPath)
	}
	client, err := tmdb.NewClient(&config.Config)
	return &Client{
		Config: config,
		Client: client,
	}, err
}

func (client *Client) SearchMovie(query string, opts *tmdb.SearchMovieRequest) (res *tmdb.SearchMovieResponse, err error) {
	key := fmt.Sprintf("movie-search-%s.json", url.QueryEscape(query))
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&res)
		return
	}
	res, err = client.Client.SearchMovie(query, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(res); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) GetMovieDetail(id int, opts *tmdb.MovieDetailRequest) (detail *tmdb.MovieDetail, err error) {
	key := fmt.Sprintf("movie-%d.json", id)
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&detail)
		return
	}
	detail, err = client.Client.GetMovieDetail(id, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(detail); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) GetMovieCredits(id int, opts *tmdb.MovieCreditsRequest) (credits *tmdb.MovieCredits, err error) {
	key := fmt.Sprintf("movie-credits-%d.json", id)
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&credits)
		return
	}
	credits, err = client.Client.GetMovieCredits(id, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(credits); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) SearchTV(query string, opts *tmdb.SearchTVRequest) (res *tmdb.SearchTVResponse, err error) {
	key := fmt.Sprintf("tv-search-%s.json", url.QueryEscape(query))
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&res)
		return
	}
	res, err = client.Client.SearchTV(query, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(res); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) GetTVDetail(id int, opts *tmdb.TVDetailRequest) (detail *tmdb.TVDetail, err error) {
	key := fmt.Sprintf("tv-%d.json", id)
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&detail)
		return
	}
	detail, err = client.Client.GetTVDetail(id, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(detail); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) GetTVCredits(id int, opts *tmdb.TVCreditsRequest) (credits *tmdb.MovieCredits, err error) {
	key := fmt.Sprintf("tv-credits-%d.json", id)
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&credits)
		return
	}
	credits, err = client.Client.GetTVCredits(id, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(credits); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) GetTVSeason(id int, season int, opts *tmdb.TVDetailRequest) (detail *tmdb.TVSeasonDetail, err error) {
	key := fmt.Sprintf("tv-season-%d-%d.json", id, season)
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&detail)
		return
	}
	detail, err = client.Client.GetTVSeason(id, season, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(detail); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}

func (client *Client) GetTVEpisode(seriesId int, seasonNumber int, episodeNumber int, opts *tmdb.TVDetailRequest) (detail *tmdb.TVEpisodeDetail, err error) {
	key := fmt.Sprintf("tv-episode-%d-%d-%d.json", seriesId, seasonNumber, episodeNumber)
	filename := filepath.Join(client.PersistentPath, key)
	if f, e := os.Open(filename); e == nil {
		err = json.NewDecoder(f).Decode(&detail)
		return
	}
	detail, err = client.Client.GetTVEpisode(seriesId, seasonNumber, episodeNumber, opts)
	if err != nil {
		return
	}
	if data, err := json.Marshal(detail); err == nil {
		os.WriteFile(filename, data, 0644)
	}
	return
}
