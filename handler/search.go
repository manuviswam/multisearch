package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	m "github.com/manuviswam/multisearch/model"
	s "github.com/manuviswam/multisearch/service"
)

func HandleSearch(google *s.GoogleSearch, duckDuckGo *s.DuckDuckGoSearch, twitter *s.TwitterSearch) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		encodedQuery := url.QueryEscape(query)
		googleResult := google.Search(encodedQuery)
		duckDuckGoResult := duckDuckGo.Search(encodedQuery)
		twitterResult := twitter.Search(encodedQuery)

		response := m.SearchResponse{
			Query: query,
			Results: m.SearchResults{
				Google:     googleResult,
				Duckduckgo: duckDuckGoResult,
				Twitter:    twitterResult,
			},
		}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}
