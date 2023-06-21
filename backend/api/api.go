package api

import (
	"fmt"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/api/routes"
	"johgo-search-engine/config"
	"net/http"
)

func ServeRouter() {
	r := routes.Initialize()

	logger.ApiInfoLogger.Printf("Launching JohGo api...")
	fmt.Println(config.APIPort)
	err := http.ListenAndServe(config.APIPort, r)
	if err != nil {
		logger.ApiInfoLogger.Printf("Error serving api:  %s", err.Error())
	}
}
