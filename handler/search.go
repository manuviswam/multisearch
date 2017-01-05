package handler

import (
	"fmt"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
