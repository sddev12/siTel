package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// setSession makes a POST request to the session service '/set-session' endpoint to store the sessionId in redis
// Parameters:
//
//	url - string: url for the session service
//	sessionId - string: the sessionId to send
//
// Returns:
//
//	error
//	response - SessionResponse : the response of the /set-session call
func setSession(url, sessionId string) (SessionResponse, error) {
	data := map[string]string{
		"sessionId": sessionId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return SessionResponse{}, fmt.Errorf("unable to marshall sessionId to json: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return SessionResponse{}, fmt.Errorf("unable to create http request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return SessionResponse{}, fmt.Errorf("request to set-session failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SessionResponse{}, fmt.Errorf("failed to read response body from set-session request: %v", err)
	}

	var sessionResponse SessionResponse
	err = json.Unmarshal(body, &sessionResponse)
	if err != nil {
		return SessionResponse{}, fmt.Errorf("failed to parse set-session response body to json: %v", err)
	}

	return sessionResponse, nil
}
