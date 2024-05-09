package api

import (
	"encoding/json"
	"lambda-go/database"
	"lambda-go/models"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbstore database.DatabaseClient
}

func NewApiHandler(dbstore database.DatabaseClient) ApiHandler {
	return ApiHandler{
		dbstore: dbstore,
	}
}

func (api *ApiHandler) RegisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user models.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if user.Username == "" || user.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Username and password are required",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	doesUserExist, err := api.dbstore.DoesUserExist(user.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if doesUserExist {
		return events.APIGatewayProxyResponse{
			Body:       "username already exists",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	dbUser, err := models.NewUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	err = api.dbstore.InsertUser(dbUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error registering the user",
			StatusCode: http.StatusBadRequest,
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       "User registered successfully",
		StatusCode: http.StatusOK,
	}, nil
}

func (api *ApiHandler) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var user LoginRequest

	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	dbUser, err := api.dbstore.GetUser(user.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Couldn't find user with username " + user.Username,
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if !models.ValidatePassword(dbUser.PasswordHash, user.Password) {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid user credentials",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Succesfully logged in",
		StatusCode: http.StatusOK,
	}, nil
}
