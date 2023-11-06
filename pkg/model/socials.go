package model

type SocialLink struct {
	Platform string `json:"platform"`
	User     string `json:"user"`
	Link     string `json:"link"`
	ArtistID uint   `json:"artistID"`
}
