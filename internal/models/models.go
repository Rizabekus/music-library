package models

import "net/url"

type SongService interface {
	DoesExistByID(songID string) (bool, error)
	DeleteByID(songID string) error
	DoesExist(songInput SongInput) (bool, error)
	AddSong(s AddSong) error
	UpdateSong(update UpdateSong, songID string) error
	FilteredSearch(queryParams url.Values) ([]Song, error)
	SongPagination(pageQuery string, pageSizeQuery string, songs []Song) ([]Song, error)
	CoupletPagination(pageQuery string, pageSizeQuery string, couplets []string) ([]string, error)
	GetSongDataByID(id string) (Song, error)
}
type SongStorage interface {
	DoesExistByID(songID string) (bool, error)
	DeleteByID(songID string) error
	DoesExist(songInput SongInput) (bool, error)
	AddSong(s AddSong) error
	UpdateSong(update UpdateSong, songID string) error
	FilteredSearch(queryParams url.Values) ([]Song, error)
	SongPagination(pageQuery string, pageSizeQuery string, songs []Song) ([]Song, error)
	CoupletPagination(pageQuery string, pageSizeQuery string, songs []string) ([]string, error)
	GetSongDataByID(id string) (Song, error)
}
type SongInput struct {
	Group string `json:"group" validate:"required,min=1,max=255,ascii"`
	Song  string `json:"song" validate:"required,min=1,max=255,ascii"`
}
type Song struct {
	Id           int    `json:"id"`
	Group        string `json:"group"`
	Song         string `json:"song"`
	Release_date string `json:"release_date"`
	Text         string `json:"text"`
	Link         string `json:"link"`
}
type AddSong struct {
	Group        string `json:"group"`
	Song         string `json:"song"`
	Release_date string `json:"release_date"`
	Text         string `json:"text"`
	Link         string `json:"link"`
}
type UpdateSong struct {
	Group       string `json:"group" validate:"omitempty,min=1,max=255,ascii"`
	Song        string `json:"song" validate:"omitempty,min=1,max=255,ascii"`
	Releasedate string `json:"release_date" validate:"omitempty,min=1,max=255,ascii"`
	Text        string `json:"text" validate:"omitempty,min=1,ascii"`
	Link        string `json:"link" validate:"omitempty,min=1,max=512,ascii"`
}
type Info struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
type Response struct {
	Message string `json:"message"`
}
