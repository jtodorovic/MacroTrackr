package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func GenerateRecipe(ingredients []string, goal string) (*RecipeResponse, error) {
	reqBody := RecipeRequest{
		Ingredients: ingredients,
		Goal:        goal,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 300 * time.Second,
	}

	service_url := os.Getenv("AI_SERVICE_URL")

	resp, err := client.Post(
		service_url,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer resp.Body.Close()

	var recipe RecipeResponse
	err = json.NewDecoder(resp.Body).Decode(&recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
