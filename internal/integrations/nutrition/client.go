package nutrition

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CalorieNinjasClient struct {
	apiKey string
}

type CalorieNinjasResponse struct {
	Items []struct {
		Name     string  `json:"name"`
		Calories float64 `json:"calories"`
		Protein  float64 `json:"protein_g"`
		Carbs    float64 `json:"carbohydrates_total_g"` // ?
		Fats     float64 `json:"fat_total_g"`
	} `json:"items"`
}

func NewCalorieNinjasClient(apiKey string) *CalorieNinjasClient {
	return &CalorieNinjasClient{
		apiKey: apiKey,
	}
}

func (c *CalorieNinjasClient) GetNutrition(foodQuery string) (*CalorieNinjasResponse, error) {
	query := url.QueryEscape(foodQuery)
	url := fmt.Sprintf(
		"https://api.api-ninjas.com/v1/nutrition?query=1lb+brisket+and+fries",
		query,
	)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", c.apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	fmt.Print(res.Body)

	var result CalorieNinjasResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
