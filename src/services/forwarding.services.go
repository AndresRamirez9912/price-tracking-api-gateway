package services

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"price-tracking-api-gateway/src/utils"
	"strings"
)

func GetTargetRoute(r *http.Request) (*url.URL, error) {
	// Read the Endpoint register
	enpointList, err := utils.ReadEndpointList()
	if err != nil {
		return nil, err
	}

	// Determine the desired host based on the received path
	_, path, found := strings.Cut(r.URL.Path, "/api")
	if !found {
		log.Println("Error cutting the received path")
		return nil, errors.New("Error cutting the received path")
	}

	host, ok := enpointList[path]
	if !ok {
		log.Println("Error Getting the host from the endpoints list")
		return nil, errors.New("Error Getting the host from the endpoints list")
	}

	target := host + path
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Println("Error Parsing the target url", err)
		return nil, err
	}
	return targetURL, nil
}

func AddHeaders(r *http.Request) {
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
}
