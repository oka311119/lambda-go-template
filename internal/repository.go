package internal

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ItemRepository interface {
	FindAll() ([]*Item, error)
	Save(ctx context.Context, item Item) error
	Delete(ctx context.Context, id string) error
}

// An implementation of ItemRepository using DynamoDB.
type ddbItemRepository struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewDdbItemRepository() ItemRepository {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("MY_AWS_REGION")),
	}))
	svc := dynamodb.New(mySession)

	return &ddbItemRepository{
		svc:       svc,
		tableName: os.Getenv("DYNAMODB_TABLE"),
	}
}

func (r *ddbItemRepository) FindAll() ([]*Item, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	output, err := r.svc.Scan(input)
	if err != nil {
		return nil, err
	}

	var items []*Item
	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ddbItemRepository) Save(ctx context.Context, item Item) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.svc.PutItemWithContext(ctx, input)
	return err
}

func (r *ddbItemRepository) Delete(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	_, err := r.svc.DeleteItemWithContext(ctx, input)
	return err
}
