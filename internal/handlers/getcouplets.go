package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"
	"github.com/gorilla/mux"
)

// GetCouplets godoc
// @Summary Lists couplets of lyrics of a specific song
// @Description Lists couplets of lyrics of a specific song based on Id
// @Produce json
// @Param id path string true "ID of a song"
// @Param page query string false "Page of the search for pagination"
// @Param pageSize query string false "Number of couplets in one page for pagination"
// @Success 200 {object} []string
// @Failure 500 {string} models.Response
// @Failure 400 {string} models.Response
// @Router /music/{id} [get]
func (handler *Handlers) GetCouplets(w http.ResponseWriter, r *http.Request) {
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
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Id does not exist")
		utils.SendResponse("Id does not exist", w, http.StatusBadRequest)
		return
	}
	songData, err := handler.Service.SongService.GetSongDataByID(songID)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	finalData, err := handler.Service.SongService.CoupletPagination(r.URL.Query().Get("page"), r.URL.Query().Get("pageSize"), strings.Split(songData.Text, "\n"))
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Bad Request")
		utils.SendResponse("Wrong query inputs for pagination", w, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(finalData)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to marshal JSON")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	file, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(file, line+1, r.Method, r.URL.Path, http.StatusOK, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully listed couplets")

}
