package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/api/routes"
	"johgo-search-engine/config"
	"net"
	"net/http"
	"strings"
)

func ServeRouter() {
	r := routes.Initialize()
	crawlRoutes(r)
	logger.ApiInfoLogger.Printf("Launched JohGo api on port %s...", config.APIPort)

	err := http.ListenAndServe(config.APIPort, r)
	if err != nil {
		logger.ApiInfoLogger.Printf("Error serving api:  %s", err.Error())
	}
}

func crawlRoutes(r *chi.Mux) {
	// crawls api routes and prints them to the console with addresses

	routes := make([]struct {
		Method, Route string
	}, 0)
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		routes = append(routes, struct {
			Method, Route string
		}{Method: method, Route: route})
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	// Get the network addresses
	networkAddresses := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("error obtaining network addresses: %s\n", err.Error())
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				networkAddresses = append(networkAddresses, ipnet.IP.String())
			}
		}
	}

	// Print routes and addresses
	fmt.Println("Registered routes:")
	for _, route := range routes {
		cleanRoute := strings.Replace(route.Route, "*", "", -1)
		cleanRoute = strings.Replace(cleanRoute, "//", "/", -1)
		fmt.Printf("%s %s\n", route.Method, cleanRoute)
		for _, networkAddress := range networkAddresses {
			fmt.Printf("http://%s%s%s\n", networkAddress, config.APIPort, cleanRoute)
		}
		fmt.Printf("http://localhost%s%s\n", config.APIPort, cleanRoute)
		fmt.Println() // Empty line for separation
	}

}
