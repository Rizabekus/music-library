package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Rizabekus/music-library/internal/models"
)

func (handler *Handlers) MockInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	mock_data := GenerateData()
	fmt.Println("MOCK DATA: ", mock_data)
	json.NewEncoder(w).Encode(mock_data)
}

func GenerateData() models.Info {
	collection := []models.Info{
		{ReleaseDate: "2024-01-01", Text: "Release 1 Description\nRelease 2 Description\nRelease 3 Description", Link: "https://example.com/1"},
		{ReleaseDate: "2024-02-15", Text: "Release 2 DescriptionRelease 1 Description\nRelease 1 Description\nRelease 1 Description\nRelease 1 Description\n", Link: "https://example.com/2"},
		{ReleaseDate: "2024-03-20", Text: "Release 3 DescriptionRelease12313 1 Description\nRelease12 1 Description\nRelease 1qwe Description\n", Link: "https://example.com/3"},
		{ReleaseDate: "2024-04-25", Text: "Release 4 DescriptionRelease 1 Description\nRelease 1 Description\n", Link: "https://example.com/4"},
		{ReleaseDate: "2024-05-30", Text: "Release 5 Description Release 1 Description\n asdads\n asdads", Link: "https://example.com/5"},
		{ReleaseDate: "2024-06-15", Text: "Release 6 Description\nRelease 2 Description\nRelease 3 Description\nRelease 2 Description\nRelease 150 Description", Link: "https://example.com/6"},
		{ReleaseDate: "2024-07-10", Text: "Release 7 Description\nRelease 2 Description\nRelease 3 Description", Link: "https://example.com/7"},
		{ReleaseDate: "2024-08-01", Text: "Release 8 Description\nRelease 2 Description\nRelease 3 Description\nRelease 2 Description\nRelease 3 Description", Link: "https://example.com/8"},
		{ReleaseDate: "2024-09-05", Text: "Release 9 Description\nqwe\nasd\nxcv\noui\nqwe\nasd\nxcv\noui\nqwe\nasd\nxcv\noui\n", Link: "https://example.com/9"},
		{ReleaseDate: "2024-10-20", Text: "Release 10 Description\nqwe\nasd\nxcv\noui\nqwe\nasd\nxcv\noui\nqwe\nasd\nxcv\noui\nqwe\nasd\nxcv\noui\nqwe\nasd\nxcv\noui", Link: "https://example.com/10"},
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(collection))
	return collection[randomIndex]

}
