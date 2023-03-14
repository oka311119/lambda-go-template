package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

// Handler function for AWS Lambda
func Handler(ctx context.Context, request Request) (Response, error) {
	// Create a new AWS session with the specified region
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	// Create a new DynamoDB client
	svc := dynamodb.New(mySession)

	// Unmarshal the request body into an Item struct
	item := Item{}
	if err := json.Unmarshal([]byte(request.Body), &item); err != nil {
		// If there is an error unmarshalling the request body, panic
		panic(err)
	}

	// Create a new PutItemInput struct to put the item in DynamoDB
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String("my-table"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(item.Id),
			},
			"name": {
				S: aws.String(item.Name),
			},
		},
	}

	// Put the item in DynamoDB
	_, putErr := svc.PutItem(putItemInput)
	if putErr != nil {
		// If there is an error putting the item in DynamoDB, panic
		panic(putErr)
	}

	// Create a new Response object with a success message and headers
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "save success",
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
