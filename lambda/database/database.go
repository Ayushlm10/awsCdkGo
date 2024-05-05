package database

import (
	"lambda-go/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "users"
)

type DynamoDBClient struct {
	dbStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() *DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)
	return &DynamoDBClient{
		dbStore: db,
	}
}

// check if a user with the given username exists in the database
func (db *DynamoDBClient) DoesUserExist(username string) (bool, error) {
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
func (db *DynamoDBClient) InsertUser(user models.RegisterUser) error {
	userInput := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
		},
	}

	_, err := db.dbStore.PutItem(userInput)
	if err != nil {
		return err
	}
	return nil
}
