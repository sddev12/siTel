package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SetSession(url, sessionId string) error {
	data := map[string]string{
		"sessionId": sessionId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to marshall sessionId to json: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("unable to create http request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request to set-session failed: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	return nil
}
