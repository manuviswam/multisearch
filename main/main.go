package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/manuviswam/multisearch/config"
	h "github.com/manuviswam/multisearch/handler"
)

func main() {
	conf, err := c.ReadFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", h.Search)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}
