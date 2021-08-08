package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res, err = http.Get("https://i.natgeofe.com/n/622db4d3-6f73-462d-b08f-4fb528467f50/01-flamingo-friendships-nationalgeographic_1477297.jpg")
	if err != nil || res.StatusCode != 200 {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: err,
		}, nil
	}
	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: err,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: m,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
