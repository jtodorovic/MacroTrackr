package usda

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SearchFood(query string, pageSize int) (*SearchResponse, error) {
	apiKey := os.Getenv("USDA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("USDA_API_KEY not set")
	}

	params := url.Values{}
	params.Set("query", query)
	params.Set("pageSize", fmt.Sprintf("%d", pageSize))
	params.Set("api_key", apiKey)

	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.nal.usda.gov/fdc/v1/foods/search?"+params.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("USDA error (%d): %s", resp.StatusCode, string(body))
	}

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
