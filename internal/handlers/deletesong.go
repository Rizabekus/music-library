package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"
	"github.com/gorilla/mux"
)

// DeleteSong godoc
// @Summary Deletes song from DB
// @Description Deletes song from DB
// @Produce json
// @Param id path string true "ID of a song"
// @Success 204 {object} models.Response
// @Failure 500 {string} models.Response
// @Failure 400 {string} models.Response
// @Router /music/{id} [delete]
func (handler *Handlers) DeleteSong(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	songID := vars["id"]
	_, err := strconv.Atoi(songID)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "ID should be numeric", "ID should be numeric")
		utils.SendResponse("ID should be numeric", w, http.StatusBadRequest)
		return
	}
	exist, err := handler.Service.SongService.DoesExistByID(songID)

	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return

	}

	if !exist {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "ID does not exist", "ID does not exist")
		utils.SendResponse("Bad Request", w, http.StatusBadRequest)
		return
	}

	err = handler.Service.SongService.DeleteByID(songID)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)

		return

	}

	file, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(file, line+1, r.Method, r.URL.Path, http.StatusNoContent, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully deleted song")
	utils.SendResponse("Successfully deleted song", w, http.StatusNoContent)

}
