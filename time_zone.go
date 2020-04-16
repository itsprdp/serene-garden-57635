package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type timeResponse struct {
	CurrentTime string `json:current_time`
	TimeZone    string `json:time_zone`
	Message     string `json:message`
}

func displayTime(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var currentTime string
	message := "ok"

	timeZone := req.URL.Query().Get("time_zone")
	if timeZone != "" {
		location, err := time.LoadLocation(timeZone)

		if err != nil {
			message = "Invalid TimeZone! Please provide the value in the format 'America/Los_Angeles'. Refer to this list of valid time zones https://en.wikipedia.org/wiki/List_of_tz_database_time_zones"
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			currentTime = time.Now().In(location).Format("15:04:05")
			w.WriteHeader(http.StatusOK)
		}
	} else {
		location, _ := time.LoadLocation("")
		timeZone = location.String()
		currentTime = time.Now().In(location).Format("15:04:05")
	}

	json.NewEncoder(w).Encode(&timeResponse{
		CurrentTime: currentTime,
		Message:     message,
		TimeZone:    timeZone,
	})
}

func main() {
	http.HandleFunc("/", displayTime)

	http.ListenAndServe(":8080", nil)
}
