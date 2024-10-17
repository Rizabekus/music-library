package utils

import (
	"encoding/json"
	"net/http"
	"runtime"
	"strings"

	"github.com/Rizabekus/music-library/internal/models"
)

func Query_Formatter(text string) string {

	exchanges := map[string]string{
		" ":  "%20",
		"\"": "%22",
		"#":  "%23",
		"%":  "%25",
		"&":  "%26",
		"'":  "%27",
		"/":  "%2F",
		":":  "%3A",
		";":  "%3B",
		"=":  "%3D",
		"?":  "%3F",
		"@":  "%40",
		"\\": "%5C",
		"+":  "%2B",
		"[":  "%5B",
		"]":  "%5D",
	}
	for key, val := range exchanges {
		text = strings.ReplaceAll(text, key, val)

	}
	return text
}
func Reversed_Query_Formatter(text string) string {

	exchanges := map[string]string{
		" ":  "%20",
		"\"": "%22",
		"#":  "%23",
		"%":  "%25",
		"&":  "%26",
		"'":  "%27",
		"/":  "%2F",
		":":  "%3A",
		";":  "%3B",
		"=":  "%3D",
		"?":  "%3F",
		"@":  "%40",
		"\\": "%5C",
		"+":  "%2B",
		"[":  "%5B",
		"]":  "%5D",
	}
	for key, val := range exchanges {
		text = strings.ReplaceAll(text, val, key)

	}
	return text
}
func GetCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "Unknown", 0, "Unknown"
	}

	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return file, line, functionName
}
func SendResponse(msg string, w http.ResponseWriter, statusCode int) {
	response := models.Response{Message: msg}

	responseJSON, err := json.Marshal(response)
	if err != nil {

		resp := models.Response{Message: "Internal Server Error"}
		internalErrorJSON, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
