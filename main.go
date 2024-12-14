package main

import (
	"log"
	"net/http"
)

type InMemoryMovieStore struct {
	movies map[int]Movie
}

func (i InMemoryMovieStore) GetMovieById(id int) Movie {
	return i.movies[id]
}

func NewInMemoryMovieStore() *InMemoryMovieStore {
	return &InMemoryMovieStore{
		movies: make(map[int]Movie),
	}
}

func main() {
	server := &MovieServer{
		NewInMemoryMovieStore(),
	}

	log.Fatal(http.ListenAndServe(":8080", server))
}
