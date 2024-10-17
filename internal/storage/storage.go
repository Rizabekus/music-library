package storage

import (
	"database/sql"

	"github.com/Rizabekus/music-library/internal/models"
)

type Storage struct {
	SongStorage models.SongStorage
}

func StorageInstance(db *sql.DB) *Storage {
	return &Storage{SongStorage: CreateSongStorage(db)}
}
