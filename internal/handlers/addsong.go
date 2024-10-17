package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Rizabekus/music-library/internal/models"
	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"
	"github.com/go-playground/validator"
)

// AddSong godoc
// @Summary Adds song to DB
// @Description Adds song to DB
// @Accept json
// @Produce json
// @Param song body models.SongInput true "SongInput struct"
// @Success 201 {object} models.Song
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 422 {object} models.Response
// @Router /music [post]
func (handler *Handlers) AddSong(w http.ResponseWriter, r *http.Request) {

	var newSong models.SongInput
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to unmarshal JSON")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return

	}
	validate := validator.New()

	if err := validate.Struct(newSong); err != nil {

		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusBadRequest, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Problem with validating")
		utils.SendResponse(err.Error(), w, http.StatusBadRequest)
		return

	}

	exist, err := handler.Service.SongService.DoesExist(newSong)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	if exist {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusUnprocessableEntity, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Song already is in Database", "Song already is in Database")
		utils.SendResponse("Song already exists", w, http.StatusUnprocessableEntity)
		return
	}
	formatted_group := utils.Query_Formatter(newSong.Group)
	formatted_song := utils.Query_Formatter(newSong.Song)

	resp, err := http.Get(fmt.Sprintf(os.Getenv("HOST")+os.Getenv("PORT")+"/info?group=%s&song=%s", formatted_group, formatted_song))

	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to make Request to /Info path")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to make Request to /Info path")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	fmt.Println("BODY: ", string(body))
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to make Request to /Info path")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}

	var info_resp models.Info

	err = json.Unmarshal(body, &info_resp)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Failed to make Request to /Info path")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	songData := models.AddSong{
		Group:        newSong.Group,
		Song:         newSong.Song,
		Release_date: info_resp.ReleaseDate,
		Text:         info_resp.Text,
		Link:         info_resp.Link,
	}
	err = handler.Service.SongService.AddSong(songData)
	if err != nil {
		file, line, _ := utils.GetCallerInfo()
		loggers.ErrorLog(file, line+1, r.Method, r.URL.Path, http.StatusInternalServerError, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), err.Error(), "Database related problem")
		utils.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		return
	}
	file, line, _ := utils.GetCallerInfo()
	loggers.DebugLog(file, line+1, r.Method, r.URL.Path, http.StatusCreated, strings.Split(r.RemoteAddr, ":")[0], r.Header.Get("Content-Type"), r.Header.Get("User-Agent"), "Successfully added song")
	utils.SendResponse("Successfully added song", w, http.StatusCreated)
}
