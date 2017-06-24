package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type CatFactResponse struct {
	Success string   `json:"success"`
	Facts   []string `json:"facts"`
}

const endpoint = "http://catfacts-api.appspot.com/api/facts"

// get fact from "Cat Facts" API
func GetFact() (string, error) {
	data, err := makeRequest(endpoint)
	if err != nil {
		return "", err
	}

	fact, err := getFactFromJSON(data)
	if err != nil {
		return "", err
	}

	return fact, nil
}

func makeRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf("%s returned %d\n", endpoint, resp.StatusCode)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getFactFromJSON(data []byte) (string, error) {
	factResponse := CatFactResponse{}
	err := json.Unmarshal(data, factResponse)
	if err != nil {
		return "", err
	}
	if len(factResponse.Facts) == 0 {
		return "", errors.New("No go-cat-facts facts returned from API")
	}

	return factResponse.Facts[0], err
}