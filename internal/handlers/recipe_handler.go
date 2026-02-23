package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/internal/integrations/ai"
)

func RecipeHandler(context *gin.Context) {
	var req ai.RecipeRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if len(req.Ingredients) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Ingredients required",
		})
		return
	}

	recipe, err := ai.GenerateRecipe(req.Ingredients, req.Goal)

	fmt.Print(recipe)

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "Internal service unavailable",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Recipe generated.", "recipe": recipe})
}
