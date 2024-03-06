package main

import (
	"log"
	"os"

	"github.com/song940/tmdb-go/persistent"
	"github.com/song940/tmdb-go/tmdb"
)

func main() {
	// command := os.Args[1]
	config := &tmdb.Config{
		APIKey:      os.Getenv("TMDB_API_KEY"),
		AccessToken: os.Getenv("TMDB_ACCESS_TOKEN"),
	}
	client, err := persistent.NewClient(&persistent.Config{
		Config: *config,
	})
	if err != nil {
		panic(err)
	}

	searchResult, err := client.SearchMovie("The Matrix", nil)
	if err != nil {
		panic(err)
	}
	for _, movie := range searchResult.Results {
		log.Println(movie.ID, movie.Title)
	}

	detail, err := client.GetTVSeason(1399, 1, nil)
	if err != nil {
		panic(err)
	}
	log.Println(detail.Name)
	for _, episode := range detail.Episodes {
		log.Println(episode.Name)
	}

	movieDetail, err := client.GetMovieDetail(603, nil)
	if err != nil {
		panic(err)
	}
	log.Println(movieDetail.Title)
	// case "get_movie_credits":
	// 	res, err := client.GetMovieCredits(603, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for _, cast := range res.Cast {
	// 		log.Println(cast.Name, "=>", cast.Character, client.GetImage(cast.ProfilePath, ""))
	// 	}
	// case "get_tv":
	// 	res, err := client.GetTVDetail(1399, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(res.Name)
	// case "get_tv_credits":
	// 	res, err := client.GetTVCredits(1399, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for _, cast := range res.Cast {
	// 		log.Println(cast.Name, "=>", cast.Character, client.GetImage(cast.ProfilePath, ""))
	// 	}
	// default:
	// 	log.Println("command not found:", command)
	// }
}
