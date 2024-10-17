package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Rizabekus/music-library/internal/models"
	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

// AddSong godoc
// @Summary Updates song info
// @Description Updates song info
// @Accept json
// @Produce json
// @Param id path string true "ID of a song"
// @Param song body models.UpdateSong true "UpdateSong struct"
// @Success 202 {object} models.Song
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /music/{id} [put]
func (handler *Handlers) UpdateSong(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	songID := vars["id"]
	exist, err := handler.Service.SongService.DoesExistByID(songID)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	if !exist {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Song does not exist", "Bad Request")
		utils.SendResponse("Internal Server Error", w, http.StatusBadRequest)
		return
	}
	var updatedValues models.UpdateSong
	err = json.NewDecoder(r.Body).Decode(&updatedValues)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	err = validate.Struct(updatedValues)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Validation related error")
		utils.SendResponse(err.Error(), w, http.StatusBadRequest)
		return
	}
	err = handler.Service.SongService.UpdateSong(updatedValues, songID)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	file, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(file, line+1, r.Method, r.URL.Path, http.StatusAccepted, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully added changes")
	utils.SendResponse("Successfully added changes", w, http.StatusAccepted)
}
