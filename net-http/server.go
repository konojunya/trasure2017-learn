package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := r.FormValue("url")
		if url == "" {
			http.Error(w, "url is blank", http.StatusBadRequest)
			return
		}

		urls := strings.Split(url, ",")
		pages := GetTitleAndDescription(urls)

		w.Header().Set("Content-Type", "application/json")
		encorder := json.NewEncoder(w)

		if err := encorder.Encode(pages); err != nil {
			http.Error(w, "encoding failed", http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
