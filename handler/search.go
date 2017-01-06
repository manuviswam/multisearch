package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	m "github.com/manuviswam/multisearch/model"
	s "github.com/manuviswam/multisearch/service"
)

func HandleSearch(google *s.GoogleSearch, duckDuckGo *s.DuckDuckGoSearch, twitter *s.TwitterSearch) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		query := r.URL.Query().Get("q")
		encodedQuery := url.QueryEscape(query)

		c1 := make(chan m.SearchResult)
		c2 := make(chan m.SearchResult)
		c3 := make(chan m.SearchResult)

		go google.Search(encodedQuery, c1)
		go duckDuckGo.Search(encodedQuery, c2)
		go twitter.Search(encodedQuery, c3)

		response := m.SearchResponse{
			Query: query,
			Results: m.SearchResults{
				Google:     <-c1,
				Duckduckgo: <-c2,
				Twitter:    <-c3,
			},
		}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)

		fmt.Println("Elapsed time ", time.Since(start))
	}
}
