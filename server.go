package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MovieStore interface {
	GetMovieById(id int) Movie
}

type MovieServer struct {
	store MovieStore
}

func (ms *MovieServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/v1/movies/"))

	log.Println(id)
	movie := ms.store.GetMovieById(id)
	json.NewEncoder(w).Encode(movie)
}
