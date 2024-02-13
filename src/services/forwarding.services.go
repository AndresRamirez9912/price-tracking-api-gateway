package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"price-tracking-api-gateway/src/constants"
	"price-tracking-api-gateway/src/utils"
	"strconv"
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

	target := os.Getenv(constants.SCHEME) + "://" + os.Getenv(host) + "/api" + path
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

func AddBodyElement(r *http.Request, elementName string, elementValue interface{}) error {
	// Decode the current Body
	body := make(map[string]interface{})
	err := utils.GetBody(r.Body, &body)
	if err != nil {
		return err
	}
	r.Body.Close()

	// Add the desired element
	body[elementName] = elementValue
	modifiedBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error Serializing the modified body")
		return err
	}

	// Update the Closer with the new Body request
	r.Body = io.NopCloser(bytes.NewBuffer(modifiedBody))
	r.ContentLength = int64(len(modifiedBody))                              // Update parametter
	r.Header.Set(constants.CONTENT_LENGTH, strconv.Itoa(len(modifiedBody))) // Update header
	return nil
}
