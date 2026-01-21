package models

type FoodLogDynamo struct {
	ID        string `dynamodbav:"id"` // string for UUID
	UserID    int64  `dynamodbav:"userID"`
	Name      string `dynamodbav:"name"`
	Calories  int32  `dynamodbav:"calories"`
	Protein   int32  `dynamodbav:"protein"`
	Carbs     int32  `dynamodbav:"carbs"`
	Fats      int32  `dynamodbav:"fats"`
	MealType  string `dynamodbav:"mealType"`
	CreatedAt string `dynamodbav:"createdAt"` // store as ISO string
	LogDate   string `dynamodbav:"logDate"`   // new field for YYYY-MM-DD
}

type DailyGoalDynamo struct {
	UserID         int64 `dynamodbav:"userID"`
	CaloriesTarget int32 `dynamodbav:"caloriesTarget"`
	ProteinTarget  int32 `dynamodbav:"proteinTarget"`
	CarbsTarget    int32 `dynamodbav:"carbsTarget"`
	FatsTarget     int32 `dynamodbav:"fatsTarget"`
}
