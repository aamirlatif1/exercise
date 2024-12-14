package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetMovies(t *testing.T) {
	store := StubMovieStore{
		map[int]Movie{
			1: aMovie(1),
			2: aMovie(2),
		},
	}
	server := &MovieServer{&store}
	t.Run("return movies by give id 1", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/v1/movies/1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertMovieResponse(t, parseMovieResponse(t, response), aMovie(1))
	})
	t.Run("return movies by give id 2", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/v1/movies/2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertMovieResponse(t, parseMovieResponse(t, response), aMovie(2))
	})
}

// --------------------------------
func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected want code %d, got %d", want, got)
	}
}

func parseMovieResponse(t testing.TB, response *httptest.ResponseRecorder) Movie {
	t.Helper()
	var got Movie
	err := json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Fatalf("unable to parse response %q into movie: %v", response.Body, err)
	}
	return got
}

func assertMovieResponse(t testing.TB, got, want Movie) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v, want %+v", got, want)
	}
}

func aMovie(id int) Movie {
	want := Movie{
		ID:      id,
		Title:   "Moana",
		Year:    2016,
		Runtime: "107 mins",
		Genres:  []string{"animation", "adventure"},
		Version: 1,
	}
	return want
}

// ------ stub
type StubMovieStore struct {
	movies map[int]Movie
}

func (s *StubMovieStore) GetMovieById(id int) Movie {
	return s.movies[id]
}
