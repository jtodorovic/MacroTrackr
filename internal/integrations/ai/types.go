package ai

type RecipeRequest struct {
	Ingredients []string `json:"ingredients"`
	Goal        string   `json:"goal"`
}

type RecipeResponse struct {
	Title        string   `json:"title"`
	Servings     int      `json:"servings"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}
