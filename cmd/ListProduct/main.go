package main

import (
	"encoding/json"
	"strconv"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/lambda"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func ListProduct(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {	
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		TableName: aws.String("Products"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500},
	}

	var products []Product
	for _, item := range result.Items {
		price, err := srtconv.Atoi(*item["price"].N)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500},
		}

		products = append(products, Product{
			ID: *item["id"].S,
			Name: *item["name"].S,
			Price, price,
		})
	}

	body, err := json.Marshal(products)
	if err != nil {if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500},
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers: map[string][string]{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func main(){
	lambda.Start(ListProduct)
}