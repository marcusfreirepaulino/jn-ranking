package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"jn-api/types"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const PAGE_SIZE = 50

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cookie := getCookie()

	episodes := fetchJnApiData(cookie)
	insertEpisodes(db, episodes)
}

func fetchJnApiData(cookie string) []types.PodcastDataDto {
	var totalPages int
	var episodesRequestData []types.PodcastDataDto

	page := 1

	for {
		fullUrl := getUrl(page)

		req, err := http.NewRequest("GET", fullUrl, nil)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Cookie", cookie)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var result types.RequestResponseDto

		json.Unmarshal(data, &result)

		if totalPages == 0 {
			totalPages = result.Data.PaginateData.TotalPages
		}

		episodesRequestData = append(episodesRequestData, result.Data.Podcasts...)

		if result.Data.PaginateData.CurrentPage >= totalPages {
			break
		}

		page++
	}

	log.Printf("Fetched %v episodes.", len(episodesRequestData))
	return episodesRequestData
}

func getUrl(page int) string {
	baseUrl := "https://jovemnerd.com.br/api/pagination"

	params := url.Values{}
	params.Add("endpoint", "/podcast/list?category_podcast_product_slug=nerdcast")
	params.Add("page", fmt.Sprint(page))
	params.Add("page_size", fmt.Sprint(PAGE_SIZE))
	params.Add("offset", "0")

	return fmt.Sprintf("%s?%s", baseUrl, params.Encode())
}

func getCookie() string {
	var cookie string

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "cookie=") {
			cookie = strings.TrimPrefix(arg, "cookie=")
		}
	}

	if cookie == "" {
		log.Fatal("Failed to get cookie. Inform cookie using the 'cookie=' flag.")
	}

	return cookie
}

func insertEpisodes(db *sql.DB, episodes []types.PodcastDataDto) {
	tx, _ := db.Begin()
	insertEpisodestatement, err := tx.Prepare("INSERT INTO episodes(title, slug, episode, release_date) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Error at insertEpisodeStatement: %v", err)
	}
	defer insertEpisodestatement.Close()

	upsertParticipantStatement, err := tx.Prepare(`
		INSERT OR REPLACE INTO participants(slug, name, participant_photo) VALUES(?, ?, ?)
	`)
	if err != nil {
		log.Fatalf("Error at upsertParticipantStatement: %v", err)
	}
	defer upsertParticipantStatement.Close()

	insertEpisodesParticipantsStatement, err := tx.Prepare(`
		INSERT INTO episodes_participants(episode_id, participant_key) VALUES(?, ?)
	`)
	if err != nil {
		log.Fatalf("Error at insertEpisodesParticipantsStatement: %v", err)
	}
	defer insertEpisodesParticipantsStatement.Close()

	for _, episode := range episodes {
		res, _ := insertEpisodestatement.Exec(episode.Title, episode.Slug, episode.Metadata.PodcastEpisode, episode.Date)

		episodeId, _ := res.LastInsertId()

		for _, guests := range episode.Categories.Guests {
			upsertParticipantStatement.Exec(guests.Slug, guests.Name, guests.GuestPhoto)
			insertEpisodesParticipantsStatement.Exec(episodeId, guests.Slug)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Error at insertEpisodes: %v", err)
	}
}
