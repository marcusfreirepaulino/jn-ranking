package types

// Database Models

type EpisodeModel struct {
	Id          int
	Title       string
	Slug        string
	Episode     string
	ReleaseDate string
}

type ParticipantModel struct {
	Slug             string
	Name             string
	ParticipantPhoto *string
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

type PodcastThemeDto struct {
	Name string `json:"name"`
}

type PodcastGuestMetadataDto struct {
	GuestPhoto string `json:"guest_photo"`
}

type PodcastGuestDto struct {
	Name     string                  `json:"name"`
	Slug     string                  `json:"slug"`
	Metadata PodcastGuestMetadataDto `json:"metadata"`
}

type PodcastCategoriesDto struct {
	Themes []PodcastThemeDto `json:"podcast_theme"`
	Guests []PodcastGuestDto `json:"podcast_guest"`
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
