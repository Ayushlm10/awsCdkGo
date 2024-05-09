package database

import (
	"fmt"
	"lambda-go/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	TABLE_NAME = "users"
)

type DatabaseClient interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user models.User) error
	GetUser(username string) (models.User, error)
}

type DynamoDBClient struct {
	dbStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)
	return DynamoDBClient{
		dbStore: db,
	}
}

// check if a user with the given username exists in the database
func (db DynamoDBClient) DoesUserExist(username string) (bool, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	result, err := db.dbStore.GetItem(input)
	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

// insert a user into the the database
func (db DynamoDBClient) InsertUser(user models.User) error {
	userInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.PasswordHash),
			},
		},
	}

	_, err := db.dbStore.PutItem(userInput)
	if err != nil {
		return err
	}
	return nil
}

func (db DynamoDBClient) GetUser(username string) (models.User, error) {
	var user models.User
	userInput := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	userResult, err := db.dbStore.GetItem(userInput)
	if err != nil {
		return user, err
	}

	if userResult.Item == nil {
		return user, fmt.Errorf("no user exists")
	}

	err = dynamodbattribute.UnmarshalMap(userResult.Item, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
