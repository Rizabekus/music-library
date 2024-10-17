package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"
)

// GetSongs godoc
// @Summary Lists songs with pagination and filtering
// @Description Lists songs with pagination and filtering
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song name"
// @Param releaseDate query string false "Release date of a song"
// @Param text query string false "Lyrics"
// @Param link query string false "Link of a song"
// @Param page query string false "Page of the search for pagination"
// @Param pageSize query string false "Page size for pagination"
// @Success 200 {object} []models.Song
// @Failure 500 {string} models.Response
// @Failure 400 {string} models.Response
// @Router /music [get]
func (handler *Handlers) GetSongs(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fmt.Println(queryParams)
	songs, err := handler.Service.SongService.FilteredSearch(queryParams)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}

	finalData, err := handler.Service.SongService.SongPagination(r.URL.Query().Get("page"), r.URL.Query().Get("pageSize"), songs)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Bad Request")
		utils.SendResponse("Wrong query inputs for pagination", w, http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(finalData)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to marshal JSON")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	file, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(file, line+1, r.Method, r.URL.Path, http.StatusOK, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully listed songs")

}
