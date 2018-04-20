package db

import (
	"log"
	"errors"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamoDB struct {
	*dynamodb.DynamoDB
}

func NewDynamoDB() Store {
	settings := getCredentials()
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(settings.DynamoDb.Region),
		Credentials: credentials.NewSharedCredentials(settings.DynamoDb.FileName, settings.DynamoDb.Profile),
	})
	if err != nil {
		log.Fatal(err)
	}
	svc := dynamodb.New(sess)
	return &dynamoDB{
		svc,
	}
}

func (db *dynamoDB) Add(id string, user User) error {
	av, _ := dynamodbattribute.Marshal(user)
	params := &dynamodb.PutItemInput{
		Item:      map[string]*dynamodb.AttributeValue{"object": av},
		TableName: aws.String("Asset_Data"),
	}
	_, err := db.PutItem(params)
	if err != nil {

	}
	return nil
}

func (db *dynamoDB) Get(id string) (User, error) {
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("person"),
	}
	_, err := db.DescribeTable(req)
	if err != nil {
		panic(err)
	}
	return User{}, errors.New("not found")
}

func (db *dynamoDB) Update(user User) error {
	return nil
}

func (db *dynamoDB) Delete(id string) error {
	return nil
}
