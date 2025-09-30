package handlers

import (
	"encoding/json"
	"net/http"

	"fem.com/movie-site/data"
	"fem.com/movie-site/logger"
	"fem.com/movie-site/models"
)

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("JSON Encoding error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		http.Error(w, "Internal Error", 500)
		h.Logger.Error("Get Top Movies Error", err)
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker Random",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
		{
			ID:          2,
			TMDB_ID:     181,
			Title:       "The Dark Knight Random",
			ReleaseYear: 2008,
			Genres:      []models.Genre{{ID: 2, Name: "Drama"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 2, FirstName: "Christian"}},
		},
	}
	h.writeJSONResponse(w, movies)
}
