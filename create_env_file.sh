#!/bin/bash

echo "Enter your AWS account ID: "
read aws_account_id

echo "Enter your AWS region: "
read aws_region

echo "Enter your DynamoDB table name: "
read table_name

dynamodb_table_arn="arn:aws:dynamodb:${aws_region}:${aws_account_id}:table/${table_name}"

echo "AWS_ACCOUNT_ID=${aws_account_id}" > .env.dev
echo "AWS_REGION=${aws_region}" >> .env.dev
echo "TABLE_NAME=${table_name}" >> .env.dev
echo "DYNAMODB_TABLE_ARN=${dynamodb_table_arn}" >> .env.dev

echo "The .env file has been created."
