package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"price-tracking-api-gateway/src/constants"
	"price-tracking-api-gateway/src/models"
	"strings"
)

func ReadEndpointList() (map[string]string, error) {
	jsonFile, err := os.Open("src/endpointList.json")
	if err != nil {
		log.Println("Error reading the endpoint list file", err)
		return nil, err
	}
	defer jsonFile.Close()

	var endpoints map[string]string
	err = json.NewDecoder(jsonFile).Decode(&endpoints)
	if err != nil {
		log.Println("Error reading the endpoint list into a map", err)
		return nil, err
	}
	return endpoints, nil
}

func GetEndpoint(r *http.Request) (string, error) {
	_, path, found := strings.Cut(r.URL.Path, "/api")
	if !found {
		log.Println("Error cutting the received path")
		return "", errors.New("Error cutting the received path")
	}
	return path, nil
}

func GetBody(body io.ReadCloser, receiver any) error {
	err := json.NewDecoder(body).Decode(receiver)
	if err != nil {
		log.Println("Error decoding the Body of the request", err)
		return err
	}
	return nil
}

func VerifyUserByJWT(accessToken string) (*models.GetUserResponse, error) {
	// Create the request
	bodyRequest := &models.GetUserRequest{AccessToken: accessToken}
	jsonBody, err := json.Marshal(bodyRequest)
	if err != nil {
		log.Println("Error marshalling the body request for user verifying", err)
		return nil, err
	}
	// authURL = http://price-tracking-auth:3001/api/getUser
	authURL := fmt.Sprintf("%s://%s/api/getUser", os.Getenv(constants.SCHEME), os.Getenv(constants.AUTH_HOST))
	req, err := http.NewRequest(http.MethodPost, authURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Error creating the HTTP request for user verifying", err)
		return nil, err
	}

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the HTTP request for user verifying", err)
		return nil, err
	}
	defer response.Body.Close()

	// Get the information
	userData := &models.GetUserResponse{}
	err = GetBody(response.Body, userData)
	if err != nil {
		log.Println("Error Getting the Body response for user verifying", err)
		return nil, err
	}

	return userData, nil
}

func GetUserByJWTResponse(response *models.GetUserResponse) *models.User {
	user := &models.User{}
	userAttributes := make(map[string]string)

	for _, v := range response.UserAttributes {
		userAttributes[v.Name] = v.Value
	}

	user.Email = userAttributes[constants.USER_EMAIL_ATTRIBUTE]
	user.Id = userAttributes[constants.USER_SUB_ATTRIBUTE]
	user.UserName = userAttributes[constants.USER_NAME_ATTRIBUTE]
	return user
}
