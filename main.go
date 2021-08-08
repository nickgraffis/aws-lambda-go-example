package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
	"github.com/buckket/go-blurhash"
	"image/png"
	"os"
	"net/http"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	imageFile, _ := http.Get("test.png")
	loadedImage, err := png.Decode(imageFile)
	str, _ := blurhash.Encode(4, 3, loadedImage)
	if err != nil {
		// Handle errors
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: str,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
