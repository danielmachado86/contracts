package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var ErrNotFound = errors.New("resource not found")

type DynamoCreateUserParams struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
	User
}

func (s *DynamoDBStore) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	user := User{
		FirstName:      arg.FirstName,
		LastName:       arg.LastName,
		Username:       arg.Username,
		Email:          arg.Email,
		PasswordHashed: arg.PasswordHashed,
		ChangedAt:      time.Time{},
		CreatedAt:      time.Now(),
	}
	dynamoArg := DynamoCreateUserParams{
		Pk:   fmt.Sprintf("user#%s", arg.Username),
		Sk:   "profile",
		User: user,
	}

	item, err := attributevalue.MarshalMap(dynamoArg)
	if err != nil {
		return User{}, err
	}

	conditions := aws.String("attribute_not_exists(username)")
	_, err = s.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(s.TableName),
		Item:                item,
		ConditionExpression: conditions,
	})

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *DynamoDBStore) GetUser(ctx context.Context, username string) (User, error) {
	user := User{Username: username}
	response, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: user.GetKey(), TableName: aws.String(s.TableName),
	})

	if err != nil {
		return user, err
	}
	if len(response.Item) == 0 {
		return user, ErrNotFound
	}
	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		return user, err
	}

	return user, err

}

func (s *DynamoDBStore) DeleteUser(ctx context.Context, username string) error {
	user := User{Username: username}
	_, err := s.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: user.GetKey(), TableName: aws.String(s.TableName),
	})
	if err != nil {
		return err
	}
	return err
}

func (s *DynamoDBStore) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	user := User{
		Username: arg.Username,
	}
	var err error
	var response *dynamodb.UpdateItemOutput
	update := expression.Set(expression.Name("firstName"), expression.Value(arg.FirstName))
	update.Set(expression.Name("lastName"), expression.Value(arg.LastName))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return User{}, err
	} else {
		response, err = s.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(s.TableName),
			Key:                       user.GetKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			return user, err
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &user)
			if err != nil {
				return user, err
			}
		}
	}
	return user, err
}

func (user User) GetKey() map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(fmt.Sprintf("user#%s", user.Username))
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal("profile")
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{
		"pk": pk,
		"sk": sk,
	}
}

type DynamoCreateContractParams struct {
	Pk string `json:"pk"`
	Sk string `json:"sk"`
	Contract
}

func (s *DynamoDBStore) CreateContract(ctx context.Context, arg CreateContractParams) (Contract, error) {
	t := time.Now()
	contract := Contract{
		Template:  arg.Template,
		CreatedAt: t,
	}
	dynamoArg := DynamoCreateContractParams{
		Pk: fmt.Sprintf(
			"contract#%s#%s#%s",
			arg.Owner,
			contract.Template,
			t.Format(time.RFC3339),
		),
		Sk:       "info",
		Contract: contract,
	}

	item, err := attributevalue.MarshalMap(dynamoArg)
	if err != nil {
		return Contract{}, err
	}
	_, err = s.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.TableName), Item: item,
	})
	if err != nil {
		return Contract{}, err
	}
	return contract, nil
}

func (p GetContractParams) GetKey(rangeKey string) map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(fmt.Sprintf(
		"%s#%s#%s",
		p.Owner,
		p.Template,
		p.CreatedAt.Format(time.RFC3339),
	))
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal(rangeKey)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"pk": pk, "sk": sk}
}

func (s *DynamoDBStore) GetContract(ctx context.Context, arg GetContractParams) (Contract, error) {
	contract := Contract{
		Template: arg.Template,
	}
	response, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: arg.GetKey("contract"), TableName: aws.String(s.TableName),
	})
	if err != nil {
		return contract, err
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &contract)
		if err != nil {
			return contract, err
		}
	}
	return contract, err
}

func (s *DynamoDBStore) GetContractOwner(ctx context.Context, arg GetContractParams) (Party, error) {
	party := Party{
		ContractID: 0,
	}
	response, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: arg.GetKey("role#owner"), TableName: aws.String(s.TableName),
	})
	if err != nil {
		return party, err
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &party)
		if err != nil {
			return party, err
		}
	}
	return party, err
}

type DynamoCreateSessionParams struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
	Session
}

func (s *DynamoDBStore) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	session := Session{
		ID:           arg.ID.String(),
		Username:     arg.Username,
		RefreshToken: arg.RefreshToken,
		UserAgent:    arg.UserAgent,
		ClientIp:     arg.ClientIp,
		IsBlocked:    arg.IsBlocked,
		ExpiresAt:    arg.ExpiresAt,
		CreatedAt:    arg.CreatedAt,
	}
	dynamoArg := DynamoCreateSessionParams{
		Pk: fmt.Sprintf(
			"session#%s",
			arg.ID,
		),
		Sk: fmt.Sprintf(
			"user#%s",
			arg.Username,
		),
		Session: session,
	}

	item, err := attributevalue.MarshalMap(dynamoArg)
	if err != nil {
		return Session{}, err
	}
	_, err = s.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.TableName), Item: item,
	})
	if err != nil {
		return Session{}, err
	}
	return session, nil
}
func CreateLocalClient() (*dynamodb.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:       "aws",
				SigningRegion:     "us-east-1",
				Source:            aws.EndpointSourceCustom,
				URL:               "http://localhost:8000",
				HostnameImmutable: true,
			}, nil
		},
	)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "test")),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
