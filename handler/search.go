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

const (
	SearchTImeoutInSeconds = 1
	SearchParameterKey     = "q"
)

func HandleSearch(google *s.GoogleSearch, duckDuckGo *s.DuckDuckGoSearch, twitter *s.TwitterSearch) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		query := r.URL.Query().Get(SearchParameterKey)
		encodedQuery := url.QueryEscape(query)

		googleChannel := searchWithTimeout(encodedQuery, google.Search)
		duckDuckGoChannel := searchWithTimeout(encodedQuery, duckDuckGo.Search)
		twitterChannel := searchWithTimeout(encodedQuery, twitter.Search)
		defer func() {
			close(googleChannel)
			close(duckDuckGoChannel)
			close(twitterChannel)
		}()

		response := m.SearchResponse{
			Query: query,
			Results: m.SearchResults{
				Google:     <-googleChannel,
				Duckduckgo: <-duckDuckGoChannel,
				Twitter:    <-twitterChannel,
			},
		}

		jsonResponse, _ := json.Marshal(response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonResponse)

		fmt.Println("Elapsed time ", time.Since(start))
	}
}

func searchWithTimeout(query string, search func(string, chan m.SearchResult)) chan m.SearchResult {
	c := make(chan m.SearchResult)
	go search(query, c)

	timeout := make(chan bool)
	go func() {
		time.Sleep(SearchTImeoutInSeconds * time.Second)
		timeout <- true
	}()

	outChannel := make(chan m.SearchResult)
	go func() {
		select {
		case result := <-c:
			outChannel <- result
		case <-timeout:
			outChannel <- m.SearchResult{
				Error: "Search timed out",
			}
		}
	}()

	return outChannel
}
