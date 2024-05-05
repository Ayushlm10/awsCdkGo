package api

import (
	"fmt"
	"lambda-go/database"
	"lambda-go/models"
)

type ApiHandler struct {
	dbstore database.DynamoDBClient
}

func NewApiHandler(dbstore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbstore: dbstore,
	}
}

func (api *ApiHandler) RegisterUserHandler(event models.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("username and password must be provided")
	}

	doesUserExist, err := api.dbstore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}

	if doesUserExist {
		return fmt.Errorf("a user with that username already exists")
	}
	err = api.dbstore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error registering the user")
	}
	return nil
}
