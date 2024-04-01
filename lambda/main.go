package main

import (
	"fmt"
	"lambda-func/app"
	"lambda-func/middleware"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("heyy, %s!", event.Username), nil
}

func ProtectedHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Protected",
	}, nil
}

func main() {
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	myApp := app.NewApp()
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.Path {
		case "/register":
			return myApp.ApiHandler.RegisterUserHandler(req)
		case "/login":
			return myApp.ApiHandler.LoginUserHandler(req)
		case "/protected":
			return middleware.ValidateJWTMiddleware(ProtectedHandler)(req)
		default:
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       "Not Found",
			}, nil
		}
	})
}
