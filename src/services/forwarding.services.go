package services

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"price-tracking-api-gateway/src/constants"
	"price-tracking-api-gateway/src/utils"
)

func GetTargetRoute(r *http.Request) (*url.URL, error) {
	// Read the Endpoint register
	enpointList, err := utils.ReadEndpointList()
	if err != nil {
		return nil, err
	}

	// Determine the desired host based on the received path
	path, err := utils.GetEndpoint(r)
	if err != nil {
		return nil, err
	}

	host, ok := enpointList[path]
	if !ok {
		log.Println("Error Getting the host from the endpoints list")
		return nil, errors.New("Error Getting the host from the endpoints list")
	}

	target := os.Getenv(constants.SCHEME) + "://" + os.Getenv(host) + path
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
