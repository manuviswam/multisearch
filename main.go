package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/manuviswam/multisearch/config"
	h "github.com/manuviswam/multisearch/handler"
	s "github.com/manuviswam/multisearch/service"
)

func main() {
	conf, err := c.ReadFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	google := &s.GoogleSearch{
		APIKey: conf.GoogleAPIKey,
	}

	duckDuckGo := &s.DuckDuckGoSearch{}

	twitter := &s.TwitterSearch{}

	twitter.SetBearerToken(conf.EncodedTwitterKey)

	http.HandleFunc("/", h.HandleSearch(google, duckDuckGo, twitter))
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}
