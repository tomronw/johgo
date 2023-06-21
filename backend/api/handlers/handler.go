package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/elastic"
	"net/http"
)

var (
	EmptyQueryError = "Empty query provided"
	ECClientError   = "Error with EC client query"
)

type APIResponse struct {
	Success bool
	Data    json.RawMessage
	Error   string
}

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/search", ReturnQueryResults)
	return r
}

func ReturnQueryResults(rw http.ResponseWriter, r *http.Request) {

	ec, err := elastic.CreateClient("")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-Type", "application/json")
		e := APIResponse{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		}
		json.NewEncoder(rw).Encode(e)
		return
	}

	q := r.URL.Query().Get("query")

	if q == "" {
		rw.WriteHeader(http.StatusForbidden)
		rw.Header().Set("Content-Type", "application/json")
		e := APIResponse{
			Success: false,
			Error:   EmptyQueryError,
		}
		json.NewEncoder(rw).Encode(e)

		return
	}

	logger.ApiInfoLogger.Printf("Getting search query: [%s] for %s", q, r.Host)
	err, successful, result := ec.Query(q)

	if err != nil || !successful {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-Type", "application/json")
		e := APIResponse{
			Success: false,
			Error:   ECClientError,
		}
		json.NewEncoder(rw).Encode(e)
		return
	} else {

		if len(result) > 25 {
			rw.WriteHeader(http.StatusOK)
			rw.Header().Set("Content-Type", "application/json")
			e := APIResponse{
				Success: true,
				Data:    result,
			}
			json.NewEncoder(rw).Encode(e)
			return
		} else {
			rw.WriteHeader(http.StatusNoContent)
			logger.ApiWarningLogger.Printf("No results for search query: %s", q)
			return

		}
	}
}
