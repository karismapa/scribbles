package main

import (
	"net/http"
	"strconv"

	"log"
)

func main() {
	h := NewHandler()

	http.HandleFunc("/edit_number", h.editNumber)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Handler struct {
	number int
}

func NewHandler() *Handler {
	return &Handler{1}
}

func (h *Handler) editNumber(w http.ResponseWriter, r *http.Request) {
	log.Printf("number before action: %d", h.number)

	num := r.URL.Query()["number"]
	if len(num) > 0 {
		h.number, _ = strconv.Atoi(num[0])
	}

	log.Printf("number after action: %d", h.number)

	return
}
