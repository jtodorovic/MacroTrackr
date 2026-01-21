package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jtodorovic/macrotrackr/internal/models"
)

var DynamoClient *dynamodb.Client

func InitDynamoDB() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("Unable to load AWS config: %v", err)
	}

	DynamoClient = dynamodb.NewFromConfig(cfg)
	return nil
}

func init() {
	err := InitDynamoDB()
	if err != nil {
		panic("Failed to initialize DynamoDB client: " + err.Error())
	}
}

func FindFoodLogsByUserAndDateDynamoDB(userID int64, date string) ([]models.FoodLogDynamo, error) {
	if DynamoClient == nil {
		return nil, fmt.Errorf("DynamoDB client not initialized")
	}

	input := &dynamodb.QueryInput{
		TableName:              aws.String("FoodLogs"),
		KeyConditionExpression: aws.String("userID = :uid AND logDate = :date"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", userID)},
			":date": &types.AttributeValueMemberS{Value: date},
		},
	}

	res, err := DynamoClient.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("Failed to query FoodLogs from DynamoDB: %w", err)
	}

	var logs []models.FoodLogDynamo
	err = attributevalue.UnmarshalListOfMaps(res.Items, &logs)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal FoodLogs: %w", err)
	}

	return logs, nil
}

func GetLatestGoalForUserDynamoDB(userID int64) (*models.DailyGoalDynamo, error) {
	if DynamoClient == nil {
		return nil, fmt.Errorf("DynamoDB client not initialized")
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String("DailyGoals"),
		Key: map[string]types.AttributeValue{
			"userID": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", userID)},
		},
	}

	res, err := DynamoClient.GetItem(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch daily goal from DynamoDB: %w", err)
	}

	if res.Item == nil {
		return nil, fmt.Errorf("no daily goal found for user %d", userID)
	}

	var goal models.DailyGoalDynamo
	err = attributevalue.UnmarshalMap(res.Item, &goal)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal daily goal: %w", err)
	}

	return &goal, nil
}
