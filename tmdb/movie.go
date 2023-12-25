package tmdb

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MovieObject struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type SearchMovieResponse struct {
	Page         int           `json:"page"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
	Results      []MovieObject `json:"results"`
}

type SearchMovieRequest struct {
	IncludeAdult       bool   `json:"include_adult"`
	Language           string `json:"language"`
	PrimaryReleaseYear string `json:"primary_release_year"`
	Page               int32  `json:"page"`
	Region             string `json:"region"`
	Year               string `json:"year"`
}

type MovieDetail struct {
	MovieObject

	Budget   int64  `json:"budget"`
	IMDbID   string `json:"imdb_id"`
	Homepage string `json:"homepage"`

	Revenue int64 `json:"revenue"`
	Runtime int   `json:"runtime"`

	Status  string `json:"status"`
	Tagline string `json:"tagline"`

	// BelongsToCollection string `json:"belongs_to_collection"`
	BelongsToCollection struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection"`

	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`

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

	SpokenLanguages []struct {
		ISO639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
	} `json:"spoken_languages"`
}

type MovieDetailRequest struct {
	Language         string `json:"language"`
	AppendToResponse string `json:"append_to_response"`
}

type Member struct {
	Audlt              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float32 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
}

type CastMember struct {
	Member

	Order int `json:"order"`
}

type CrewMember struct {
	Member

	Department string `json:"department"`
	Job        string `json:"job"`
}

type MovieCreditsRequest struct {
	Language string `json:"language"`
}

type MovieCreditsResponse struct {
	ID   int          `json:"id"`
	Cast []CastMember `json:"cast"`
	Crew []CrewMember `json:"crew"`
}

// Search for movies by their original, translated and alternative titles.
// https://developer.themoviedb.org/reference/search-movie
func (client *Client) SearchMovie(query string, opts *SearchMovieRequest) (res *SearchMovieResponse, err error) {
	if opts == nil {
		opts = &SearchMovieRequest{}
	}
	if opts.Page < 1 {
		opts.Page = 1
	}
	data, err := client.get("/search/movie", map[string]string{
		"query":                query,
		"page":                 fmt.Sprint(opts.Page),
		"year":                 opts.Year,
		"region":               opts.Region,
		"language":             opts.Language,
		"primary_release_year": opts.PrimaryReleaseYear,
		"include_adult":        strconv.FormatBool(opts.IncludeAdult),
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}

// Get the top level details of a movie by ID.
// https://developer.themoviedb.org/reference/movie-details
func (client *Client) GetMovieDetail(id int, opts *MovieDetailRequest) (detail *MovieDetail, err error) {
	if opts == nil {
		opts = &MovieDetailRequest{}
	}
	data, err := client.get(fmt.Sprintf("/movie/%d", id), map[string]string{
		"language": opts.Language,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &detail)
	return
}

// Get the cast and crew for a movie.
// https://developer.themoviedb.org/reference/movie-credits
func (client *Client) GetMovieCredits(id int, opts *MovieCreditsRequest) (credits *MovieCreditsResponse, err error) {
	if opts == nil {
		opts = &MovieCreditsRequest{}
	}
	data, err := client.get(fmt.Sprintf("/movie/%d/credits", id), map[string]string{
		"language": opts.Language,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &credits)
	return
}
