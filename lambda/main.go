package main

import (
	"fmt"
	"lambda-go/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEventMessage struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEventMessage) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	return fmt.Sprintf("Succesfully called by - %s", event.Username), nil
}
func main() {
	myApp := app.NewApp()
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}
