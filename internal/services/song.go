package services

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/Rizabekus/music-library/internal/models"
)

type SongService struct {
	storage models.SongStorage
}

func CreateSongService(storage models.SongService) *SongService {
	return &SongService{storage: storage}
}

func (ss *SongService) DoesExistByID(songID string) (bool, error) {
	return ss.storage.DoesExistByID(songID)
}
func (ss *SongService) DeleteByID(songID string) error {
	return ss.storage.DeleteByID(songID)
}
func (ss *SongService) DoesExist(data models.SongInput) (bool, error) {
	return ss.storage.DoesExist(data)
}
func (ss *SongService) AddSong(data models.AddSong) error {

	return ss.storage.AddSong(data)
}
func (ss *SongService) UpdateSong(update models.UpdateSong, songID string) error {
	return ss.storage.UpdateSong(update, songID)
}
func (ss *SongService) FilteredSearch(queryParams url.Values) ([]models.Song, error) {
	return ss.storage.FilteredSearch(queryParams)
}

func (ss *SongService) SongPagination(pageQuery string, pageSizeQuery string, songs []models.Song) ([]models.Song, error) {
	var page, pageSize int
	var err error
	if pageQuery == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			return nil, err
		}
		if page < 1 {
			page = 1
		}
	}
	if pageSizeQuery == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(pageSizeQuery)
		if err != nil {
			return nil, err
		}
		if pageSize < 1 {
			pageSize = 10
		}
	}
	if pageSize >= len(songs) {
		return songs, nil
	}
	startIdx := (page - 1) * pageSize
	endIdx := startIdx + pageSize
	if startIdx >= len(songs) {
		return nil, errors.New("Page not found")
	}
	if endIdx > len(songs) {
		endIdx = len(songs)
	}
	return songs[startIdx:endIdx], nil
}
func (ss *SongService) CoupletPagination(pageQuery string, pageSizeQuery string, couplets []string) ([]string, error) {
	var page, pageSize int
	var err error
	if pageQuery == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			return nil, err
		}
		if page < 1 {
			page = 1
		}
	}
	if pageSizeQuery == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(pageSizeQuery)
		if err != nil {
			return nil, err
		}
		if pageSize < 1 {
			pageSize = 10
		}
	}
	if pageSize >= len(couplets) {
		return couplets, nil
	}
	startIdx := (page - 1) * pageSize
	endIdx := startIdx + pageSize
	if startIdx >= len(couplets) {
		return nil, errors.New("Page not found")
	}
	if endIdx > len(couplets) {
		endIdx = len(couplets)
	}
	return couplets[startIdx:endIdx], nil
}
func (ss *SongService) GetSongDataByID(id string) (models.Song, error) {
	return ss.storage.GetSongDataByID(id)

}
