package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEventMessage struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEventMessage) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("Username cannot be empty")
	}
	return fmt.Sprintf("Succesfully called by - %s", event.Username), nil
}
func main() {
	lambda.Start(HandleRequest)
}
