package handler

import (
	"encoding/json"
	"net/http"

	m "github.com/manuviswam/multisearch/model"
	s "github.com/manuviswam/multisearch/service"
)

func HandleSearch(google *s.GoogleSearch) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		googleResult, _ := google.Search(query)

		response := m.SearchResponse{
			Query: query,
			Results: m.SearchResults{
				Google: googleResult,
			},
		}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}
