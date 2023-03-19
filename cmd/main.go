package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	handler "github.com/oka311119/lambda-go-template/internal"
)

// The main function serves as the entry point for the Lambda function.
func main() {
	// Load environment variables
	godotenv.Load()

	h := handler.NewHandler()
	lambda.Start(h.Invoke)
}
