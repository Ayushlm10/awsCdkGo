package app

import (
	"lambda-go/api"
	"lambda-go/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	dbClient := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(dbClient)
	return App{
		ApiHandler: apiHandler,
	}
}
