package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jtodorovic/macrotrackr/internal/integrations/nutrition"
	"github.com/jtodorovic/macrotrackr/internal/models"
)

type NutritionService struct {
	client *nutrition.CalorieNinjasClient
}

func NewNutritionService(client *nutrition.CalorieNinjasClient) *NutritionService {
	return &NutritionService{
		client: client,
	}
}

type USDASearchResponse struct {
	Foods []models.USDAFood `json:"foods"`
}

func FetchNutrition(food string) (*models.USDAFood, error) {
	apiKey := os.Getenv("USDA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("USDA_API_KEY not set")
	}

	params := url.Values{}
	params.Set("query", food)
	params.Set("pageSize", "1")
	params.Set("api_key", apiKey)

	req, err := http.NewRequest(
		"GET",
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
		return nil, fmt.Errorf(
			"USDA error (%d): %s",
			resp.StatusCode,
			string(body),
		)
	}

	var result USDASearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Foods) == 0 {
		return nil, fmt.Errorf("no food found")
	}

	return &result.Foods[0], nil
}
