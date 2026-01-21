package usda

type Nutrient struct {
	NutrientID int     `json:"nutrientId"`
	Value      float64 `json:"value"`
}

type Food struct {
	Description string     `json:"description"`
	Nutrients   []Nutrient `json:"foodNutrients"`
}

type SearchResponse struct {
	Foods []Food `json:"foods"`
}
