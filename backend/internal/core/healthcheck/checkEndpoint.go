package healthcheck

import (
	"encoding/json"
	"fmt"
	"io"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/config"
	"johgo-search-engine/internal/core/coreModels"
	"net/http"
	"time"
)

func SpawnHealthCheck() {
	time.Sleep(time.Duration(config.HealthInterval) * time.Minute)
	for {
		healthCheckSuccessful, reason := CheckEndpoint()
		if !healthCheckSuccessful {
			logger.ApiErrorLogger.Printf("failed to check endpoint: %v", reason)
			err, emailSent := SendEmail(reason)
			if err != nil {
				logger.ApiErrorLogger.Printf("failed to send email: %v", err)
			}
			if !emailSent {
				logger.ApiErrorLogger.Printf("email not sent")
			} else {
				logger.ApiErrorLogger.Printf("email sent")
				time.Sleep(time.Hour * 3)
			}
		} else {
			logger.ApiInfoLogger.Printf("endpoint healthcheck successful")
		}
		time.Sleep(time.Duration(config.HealthInterval) * time.Minute)
	}

}

func CheckEndpoint() (healthCheckSuccessful bool, reason error) {

	req, _ := http.NewRequest("GET", config.SiteSearchEndpoint, nil)

	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add(config.Header, config.HeaderVal)

	res, err := http.DefaultClient.Do(req)

	if err == nil {

		defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return false, fmt.Errorf("failed to read in bytes: %v", err.Error())
		}

		var response coreModels.HealthCheckResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return false, fmt.Errorf("failed to parse json: %v", err.Error())
		}

		fmt.Println(response.Error)
		fmt.Println(response.Success)

		if !response.Success || res.StatusCode != http.StatusOK {
			logger.ApiErrorLogger.Printf("api response indicates failure. Error: %s", response.Error)
			return false, fmt.Errorf("api response indicates failure. Error: %s", response.Error)
		}

		return true, nil

	} else {
		return false, fmt.Errorf("failed to make the request: %v", err.Error())
	}
}
