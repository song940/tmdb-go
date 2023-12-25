package tmdb

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type TVObject struct {
	ID               int      `json:"id"`
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float32  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Name             string   `json:"name"`
	VoteAverage      float32  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

type SearchTVResponse struct {
	Page         int        `json:"page"`
	TotalPages   int        `json:"total_pages"`
	TotalResults int        `json:"total_results"`
	Results      []TVObject `json:"results"`
}

type SearchTVRequest struct {
	FirstAirDateYear string `json:"first_air_date_year"`
	IncludeAdult     bool   `json:"include_adult"`
	Language         string `json:"language"`
	Page             int32  `json:"page"`
	Year             string `json:"year"`
}

type TVDetail struct {
	TVObject

	CreatedBy []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`

	EpisodeRunTime []int  `json:"episode_run_time"`
	InProduction   bool   `json:"in_production"`
	Homepage       string `json:"homepage"`

	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`

	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
	} `json:"last_episode_to_air"`
	NextEpisodeToAir string `json:"next_episode_to_air"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int `json:"number_of_episodes"`
	NumberOfSeasons     int `json:"number_of_seasons"`
	ProductionCompanies []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`

	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons"`
	SpokenLanguages []struct {
		ISO639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
	} `json:"spoken_languages"`

	Status  string `json:"status"`
	Tagline string `json:"tagline"`
	Type    string `json:"type"`
}

type TVDetailRequest struct {
	Language string `json:"language"`
}

type TVCreditsRequest struct {
	Language string `json:"language"`
}

// Search for TV shows by their original, translated and also known as names.
// https://developer.themoviedb.org/reference/search-tv
func (client *Client) SearchTV(query string, opts *SearchTVRequest) (res *SearchTVResponse, err error) {
	if opts == nil {
		opts = &SearchTVRequest{}
	}
	if opts.Page < 1 {
		opts.Page = 1
	}
	data, err := client.get("/search/tv", map[string]string{
		"query":               query,
		"year":                opts.Year,
		"language":            opts.Language,
		"first_air_date_year": opts.FirstAirDateYear,
		"page":                fmt.Sprint(opts.Page),
		"include_adult":       strconv.FormatBool(opts.IncludeAdult),
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}

// Get the details of a TV show.
// https://developer.themoviedb.org/reference/tv-series-details
func (client *Client) GetTVDetail(id int, opts *TVDetailRequest) (detail *TVDetail, err error) {
	if opts == nil {
		opts = &TVDetailRequest{}
	}
	data, err := client.get(fmt.Sprintf("/tv/%d", id), map[string]string{
		"language": opts.Language,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &detail)
	return
}

// Get the latest season credits of a TV show.
// https://developer.themoviedb.org/reference/tv-series-credits
func (client *Client) GetTVCredits(id int, opts *TVCreditsRequest) (credits *MovieCreditsResponse, err error) {
	if opts == nil {
		opts = &TVCreditsRequest{}
	}
	data, err := client.get(fmt.Sprintf("/tv/%d/credits", id), map[string]string{
		"language": opts.Language,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &credits)
	return
}
