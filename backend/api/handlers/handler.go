package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/elastic"
	"net/http"
	"strconv"
)

var (
	ECClientError = "error with EC client query"
	NoneBool      = "singles query parameter must be a boolean"
)

type APIResponse struct {
	Success bool
	Data    json.RawMessage
	Error   string
}

// mount routes for api
func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/search", ReturnQueryResults)
	return r
}

// return search results from elastic
func ReturnQueryResults(rw http.ResponseWriter, r *http.Request) {
	// create elastic client
	ec, err := elastic.CreateClient("")
	// if error creating client, return error
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
	// get query and filter_singles params
	q := r.URL.Query().Get("query")
	b := r.URL.Query().Get("filter_singles")
	// if either are empty, return bad request
	if q == "" || b == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Header().Set("Content-Type", "application/json")
		e := APIResponse{
			Success: false,
			Data:    nil,
			Error:   "bad request params",
		}
		json.NewEncoder(rw).Encode(e)
		return
	}
	// parse singles param
	includeSingles, err := strconv.ParseBool(b)
	// if error parsing singles param, return bad request
	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		rw.Header().Set("Content-Type", "application/json")
		e := APIResponse{
			Success: false,
			Error:   NoneBool,
		}
		json.NewEncoder(rw).Encode(e)

		return
	}
	logger.ApiInfoLogger.Printf("Getting search query: [%s] for %s", q, r.Host)
	// query elastic and return results or error
	err, successful, result := ec.Query(q, includeSingles)

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
		// sometimes elastic returned no results but true to successful search
		if len(result) > 25 {
			// successful search with results, encode and return
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
