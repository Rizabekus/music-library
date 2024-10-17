package app

import (
	"log"
	"net/http"
	"os"

	_ "github.com/Rizabekus/music-library/docs"
	"github.com/Rizabekus/music-library/internal/handlers"
	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()

	r.HandleFunc("/music", h.GetSongs).Methods("GET")
	r.HandleFunc("/music/{id}", h.GetCouplets).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/music", h.AddSong).Methods("POST")
	r.HandleFunc("/music/{id}", h.DeleteSong).Methods("DELETE")
	r.HandleFunc("/music/{id}", h.UpdateSong).Methods("PUT")
	r.HandleFunc("/info", h.MockInfo).Methods("GET")
	file, line, _ := utils.GetCallerInfo()
	loggers.InfoLog(file, line, "Started the server")
	defer loggers.CloseLogFile()
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))

}
