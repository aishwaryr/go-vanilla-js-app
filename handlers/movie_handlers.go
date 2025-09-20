package handlers

import (
	"encoding/json"
	"net/http"

	"fem.com/movie-site/models"
)

type MovieHandler struct{}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMBD_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
		{
			ID:          2,
			TMBD_ID:     181,
			Title:       "The Dark Knight",
			ReleaseYear: 2008,
			Genres:      []models.Genre{{ID: 2, Name: "Drama"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 2, FirstName: "Christian"}},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO: Log error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
