package http

import (
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Build simple HTTP client for scraper

func ScraperHttpclient(proxy string) *http.Client {

	cli := createClient(proxy)

	return cli

}

func createClient(proxy string) *http.Client {

	if proxy == "" {
		client := &http.Client{
			Timeout: time.Duration(15) * time.Second,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   20 * time.Second,
					KeepAlive: 20 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   20 * time.Second,
				ResponseHeaderTimeout: 20 * time.Second,
			},
		}
		return client
	} else {

		host, username, password := splitProxy(proxy)

		proxyUrl := &url.URL{
			Scheme: "http",
			User:   url.UserPassword(username, password),
			Host:   host,
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
				Dial: (&net.Dialer{
					Timeout:   15 * time.Second,
					KeepAlive: 15 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   15 * time.Second,
				ResponseHeaderTimeout: 15 * time.Second,
			},

			Timeout: time.Duration(7) * time.Second,
		}
		return client
	}

}

func splitProxy(proxy string) (returnedProxyHost string, returnedUsername string, returnedPassword string) {

	s := strings.Split(proxy, ":")

	returnedProxyHost, returnedProxyHostTwo, returnedUsername, returnedPassword := s[0], s[1], s[2], s[3]

	returnedProxyHost = returnedProxyHost + ":" + returnedProxyHostTwo

	return returnedProxyHost, returnedUsername, returnedPassword

}
