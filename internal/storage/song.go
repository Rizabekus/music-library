package storage

import (
	"database/sql"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/Rizabekus/music-library/internal/models"
	"github.com/Rizabekus/music-library/pkg/customerrors"
)

type SongDB struct {
	DB *sql.DB
}

func CreateSongStorage(db *sql.DB) *SongDB {
	return &SongDB{DB: db}
}

func (sdb *SongDB) DoesExistByID(songID string) (bool, error) {
	if songID == "" {
		return false, fmt.Errorf("empty ID provided")
	}

	query := "SELECT EXISTS(SELECT 1 FROM musiclib WHERE id = $1)"

	var exists bool
	err := sdb.DB.QueryRow(query, songID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence: %v", err)
	}

	return exists, nil
}
func (sdb *SongDB) DeleteByID(songID string) error {
	query := "DELETE FROM musiclib WHERE id = $1"

	result, err := sdb.DB.Exec(query, songID)
	if err != nil {
		return fmt.Errorf("error deleting song: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return customerrors.ErrNotFound
	}

	return nil
}
func (sdb *SongDB) DoesExist(songInput models.SongInput) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1 FROM musiclib
			WHERE "group" = $1 AND song = $2
		)
	`

	var exists bool
	err := sdb.DB.QueryRow(query, songInput.Group, songInput.Song).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence: %v", err)
	}

	return exists, nil
}
func (sdb *SongDB) AddSong(s models.AddSong) error {

	sqlStatement := `
		INSERT INTO musiclib ("group", song, releasedate, text, link)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := sdb.DB.Exec(sqlStatement, s.Group, s.Song, s.Release_date, s.Text, s.Link)
	if err != nil {

		return fmt.Errorf("error executing SQL statement in AddSong: %v", err)
	}

	return nil
}
func (sdb *SongDB) UpdateSong(update models.UpdateSong, songID string) error {
	var setValues []string
	var values []interface{}

	typ := reflect.TypeOf(update)

	val := reflect.ValueOf(update)

	for i := 0; i < typ.NumField(); i++ {

		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()

		if reflect.Zero(typ.Field(i).Type).Interface() == fieldValue {
			continue
		}

		setValues = append(setValues, fmt.Sprintf("%s = $%d", strings.ToLower(fieldName), len(values)+1))
		values = append(values, fieldValue)
	}

	if len(setValues) == 0 {

		return nil
	}
	join := strings.Join(setValues, ", ")
	if strings.Contains(join, "group =") {
		join = strings.Replace(join, "group =", "\"group\" =", 1)
	}

	query := fmt.Sprintf("UPDATE musiclib SET %s WHERE id = $%d", join, len(values)+1)

	stmt, err := sdb.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	values = append(values, songID)

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}
func (sdb *SongDB) FilteredSearch(queries url.Values) ([]models.Song, error) {
	var filteredSongs []models.Song

	query := "SELECT * FROM musiclib WHERE true"
	params := make([]interface{}, 0)

	if group, ok := queries["group"]; ok && len(group) > 0 {
		query += " AND \"group\" = $1"
		params = append(params, group[0])
	}

	if song, ok := queries["song"]; ok && len(song) > 0 {
		query += " AND song = $" + strconv.Itoa(len(params)+1)
		params = append(params, song[0])
	}

	if releasedate, ok := queries["releasedate"]; ok && len(releasedate) > 0 {
		query += " AND releasedate = $" + strconv.Itoa(len(params)+1)
		params = append(params, releasedate[0])
	}

	if text, ok := queries["text"]; ok && len(text) > 0 {
		query += " AND text = $" + strconv.Itoa(len(params)+1)
		params = append(params, text[0])
	}

	if link, ok := queries["link"]; ok && len(link) > 0 {
		query += " AND link = $" + strconv.Itoa(len(params)+1)
		params = append(params, link[0])
	}

	rows, err := sdb.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.Song

		err := rows.Scan(
			&song.Id,
			&song.Group,
			&song.Song,
			&song.Release_date,
			&song.Text,
			&song.Link,
		)
		if err != nil {
			return nil, err
		}
		filteredSongs = append(filteredSongs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return filteredSongs, nil
}
func (sdb *SongDB) SongPagination(pageQuery string, pageSizeQuery string, songs []models.Song) ([]models.Song, error) {
	return []models.Song{}, nil
}
func (sdb *SongDB) CoupletPagination(pageQuery string, pageSizeQuery string, couplets []string) ([]string, error) {
	return []string{}, nil
}
func (sdb *SongDB) GetSongDataByID(id string) (models.Song, error) {
	var song models.Song

	row := sdb.DB.QueryRow("SELECT * FROM musiclib WHERE id = $1", id)

	err := row.Scan(&song.Id, &song.Group, &song.Song, &song.Release_date, &song.Text, &song.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			return song, fmt.Errorf("song not found with ID %s", id)
		}
		return song, fmt.Errorf("failed to scan song data: %w", err)
	}

	return song, nil
}

//почему-то ноет что нет pagination() в слое storage, добавил ради временного решения
