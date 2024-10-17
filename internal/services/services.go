package services

import (
	"github.com/Rizabekus/music-library/internal/models"
	"github.com/Rizabekus/music-library/internal/storage"
)

type Services struct {
	SongService models.SongService
}

func ServiceInstance(storage *storage.Storage) *Services {
	return &Services{
		SongService: CreateSongService(storage.SongStorage),
	}
}
