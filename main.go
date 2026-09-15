package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func daysUntilNewYear(now time.Time) int {
	newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	return int(newYear.Sub(now).Hours() / 24)
}

type response struct {
	Date string `json:"date"`
	Days int    `json:"days"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func daysHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()

	dateParam := r.URL.Query().Get("date")
	if dateParam != "" {
		parsed, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: "invalid date format, expected YYYY-MM-DD"})
			return
		}
		now = parsed
	}

	resp := response{
		Date: now.Format("2006-01-02"),
		Days: daysUntilNewYear(now),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/days", daysHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
