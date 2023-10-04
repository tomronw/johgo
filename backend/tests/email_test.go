package tests

import (
	"errors"
	"johgo-search-engine/internal/core/healthCheck"
	"testing"
)

func TestEmailSending(t *testing.T) {
	err := errors.New("bad request")
	err, emailSent := healthcheck.SendEmail(err)
	if err != nil {
		t.Errorf("Error sending email: %v", err)
	}
	if !emailSent {
		t.Errorf("Email not sent")
	}

}
