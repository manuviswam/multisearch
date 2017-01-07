package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	m "github.com/manuviswam/multisearch/model"
)

const (
	duckDuckGoUrl = "http://api.duckduckgo.com/?format=json&q=%s"
)

type DuckDuckGoSearch struct {
}

func (d *DuckDuckGoSearch) Search(query string, c chan m.SearchResult) {
	defer close(c)
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf(duckDuckGoUrl, query))
	fmt.Println("Elapsed time for duckduckgo ", time.Since(start))
	defer resp.Body.Close()
	if err != nil {
		c <- m.SearchResult{
			Error: err.Error(),
		}
		return
	}

	duckDuckGoResponse := m.DuckDuckGoResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&duckDuckGoResponse)
	if err != nil {
		c <- m.SearchResult{
			Error: err.Error(),
		}
		return
	}

	if len(duckDuckGoResponse.RelatedTopics) < 1 {
		c <- m.SearchResult{
			Error: "No response obtained",
		}
		return
	}
	c <- m.SearchResult{
		Url:  duckDuckGoResponse.RelatedTopics[0].FirstURL,
		Text: duckDuckGoResponse.RelatedTopics[0].Text,
	}
}
