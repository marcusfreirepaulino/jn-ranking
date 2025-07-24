package types

// Database Models

type EpisodeModel struct {
	Id          int
	Title       string
	Slug        string
	ReleaseDate string
}

// Jovem Nerd API Dto's
type RequestPaginateDataDto struct {
	TotalPages  int `json:"totalPages"`
	CurrentPage int `json:"currentPage"`
	LastHitSort int `json:"lastHitSort"`
}

type PodcastMetadataDto struct {
	PodcastEpisode string `json:"podcast_episode"`
}

type PodcastTheme struct {
	Name string `json:"name"`
}

type PodcastGuest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type PodcastCategoriesDto struct {
	Themes []PodcastTheme `json:"podcast_theme"`
	Guests []PodcastGuest `json:"podcast_guest"`
}

type PodcastDataDto struct {
	Id         int                  `json:"id"`
	Title      string               `json:"title"`
	Date       string               `json:"date"`
	Slug       string               `json:"slug"`
	Categories PodcastCategoriesDto `json:"categories"`
	Metadata   PodcastMetadataDto   `json:"metadata"`
}

type RequestDataDto struct {
	PaginateData RequestPaginateDataDto `json:"paginateData"`
	Podcasts     []PodcastDataDto       `json:"podcasts"`
}

type RequestResponseDto struct {
	Success bool           `json:"success"`
	Data    RequestDataDto `json:"data"`
}
