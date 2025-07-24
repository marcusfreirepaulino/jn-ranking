package main

import (
	"encoding/json"
	"fmt"
	"io"
	"jn-api/types"
	"log"
	"net/http"
	"net/url"
)

const REQ_COOKIE = "jovem-nerd-cookie-consent=true; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBOYW1lIjoiQXBwUG9ydGFsIiwiZW52IjoicHJkIiwiZGF0ZSI6MTc1MDY4ODYxMSwiZXhwIjoxNzUwNzc1MDExfQ.gnCDITrk1ZeHAJE5QCshn38Qqik8eOqGL6dFim2PcEM"
const PAGE_SIZE = 50

func main() {
	var totalPages int
	page := 1

	for {
		fullUrl := getUrl(page)

		req, err := http.NewRequest("GET", fullUrl, nil)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Cookie", REQ_COOKIE)

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

		for _, s := range result.Data.Podcasts {
			fmt.Println(s.Title, "-", s.Metadata.PodcastEpisode)
		}

		if result.Data.PaginateData.CurrentPage >= totalPages {
			break
		}

		page++
	}
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
