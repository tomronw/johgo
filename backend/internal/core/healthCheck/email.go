package healthcheck

import (
	"encoding/json"
	"errors"
	"fmt"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/config"
	"johgo-search-engine/internal/core/coreModels"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

func SendEmail(reason error) (error, bool) {

	email, pword, err := getApi()
	if err != nil {
		return err, false
	}
	rand.Seed(time.Now().UnixNano())
	currentDate := time.Unix(time.Now().Unix(), 0)

	msg := fmt.Sprintf("From: %s\n"+
		"To: %s\n"+
		"Subject: %s failed a healthcheck 🫣\n\n"+
		"%s failed a health check at: %s\n\nreason: `%s` ", email, email, config.Sitename,
		config.Sitename, currentDate.Format("2006-01-02 15:04:05"), reason.Error())

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", email, pword, "smtp.gmail.com"),
		email, []string{email}, []byte(msg))

	if err != nil {
		logger.ApiInfoLogger.Printf("smtp error: %s", err)
		return err, false
	}
	return nil, true
}

func getApi() (pass string, email string, e error) {

	url := fmt.Sprintf("%s/v1/%s", config.SitesAPIURL, config.EmailURI)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		logger.ApiInfoLogger.Printf("Failed to make the request: %v", err)
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		logger.ApiInfoLogger.Printf("Received non-200 response: %d", res.StatusCode)
		return "", "", errors.New("bad status code")
	}

	var data coreModels.ApiEmail
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		logger.ApiInfoLogger.Printf("Failed to decode the response: %v", err)
		return "", "", err
	}

	return data.Email, data.Password, nil
}
