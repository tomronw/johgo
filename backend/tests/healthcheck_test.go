package tests

import (
	"johgo-search-engine/internal/core/healthCheck"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	wasSuccessful, err := healthcheck.CheckEndpoint()
	if err != nil {
		t.Errorf("error checking endpoint: %v", err.Error())
	}
	if !wasSuccessful {
		t.Errorf("endpoint failed health check!")
	}

}
