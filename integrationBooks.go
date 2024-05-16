package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetIntegrationBooksBySubject(subject string) (*GetBooksRespDTO, error) {
	var response GetBooksRespDTO

	url := "http://openlibrary.org/subjects/%s.json?"

	url = fmt.Sprintf(url, subject)

	getReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request to Open Library : %v", err)
	}

	getReq.Header["Accept"] = []string{"application/json"}
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(getReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request to Open Library: %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println(err)
		bodyBytes, err := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		return nil, fmt.Errorf("failed to decode response: %v [RESP BODY: %s]", err, bodyString)
	}

	return &response, nil
}
