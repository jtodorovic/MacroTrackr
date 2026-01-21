package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jtodorovic/macrotrackr/internal/handlers"
)

func main() {
	lambda.Start(handlers.GetTodaysSummaryLambda)
}
