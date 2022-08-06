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
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("resource not found")

type AV struct {
	Pk string `json:"pk" dynamodbav:"pk"`
	Sk string `json:"sk" dynamodbav:"sk"`
}

type DynamoCreateUserParams struct {
	AV
	User
}

func (s *DynamoDBStore) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	av := AV{
		Pk: fmt.Sprintf("USER#%s", arg.Username),
		Sk: "PROFILE",
	}
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
		AV:   av,
		User: user,
	}

	eAv := AV{
		Pk: arg.Email,
		Sk: "field#email",
	}

	emailAv, err := attributevalue.MarshalMap(eAv)
	if err != nil {
		return User{}, err
	}
	item, err := attributevalue.MarshalMap(dynamoArg)
	if err != nil {
		return User{}, err
	}

	_, err = s.db.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName:                           &s.TableName,
					Item:                                item,
					ConditionExpression:                 aws.String("attribute_not_exists(pk)"),
					ReturnValuesOnConditionCheckFailure: types.ReturnValuesOnConditionCheckFailureAllOld,
				},
			},
			{
				Put: &types.Put{
					TableName:                           &s.TableName,
					Item:                                emailAv,
					ConditionExpression:                 aws.String("attribute_not_exists(pk)"),
					ReturnValuesOnConditionCheckFailure: types.ReturnValuesOnConditionCheckFailureAllOld,
				},
			},
		},
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

type ContractView struct {
	AV
	Contract
}

type PartyView struct {
	AV
	Username  string `dynamodbav:"username"`
	FirstName string `dynamodbav:"firstName"`
	LastName  string `dynamodbav:"lastName"`
	Role      string `dynamodbav:"role"`
}

func (s *DynamoDBStore) CreateContract(ctx context.Context, arg CreateContractParams) (Contract, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Contract{}, err
	}
	party := PartyView{
		AV: AV{
			Pk: fmt.Sprintf("contract#%s", id),
			Sk: "role#owner",
		},
		Username:  arg.Owner.Username,
		FirstName: arg.Owner.FirstName,
		LastName:  arg.Owner.LastName,
		Role:      "owner",
	}
	contract := ContractView{
		AV: AV{
			Pk: fmt.Sprintf("contract#%s", id),
			Sk: "info",
		},
		Contract: Contract{
			ID:        id.String(),
			Template:  arg.Template,
			CreatedAt: time.Now(),
		},
	}

	mContract, err := attributevalue.MarshalMap(contract)
	if err != nil {
		return Contract{}, err
	}
	mParty, err := attributevalue.MarshalMap(party)
	if err != nil {
		return Contract{}, err
	}
	_, err = s.db.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName:                           &s.TableName,
					Item:                                mContract,
					ConditionExpression:                 aws.String("attribute_not_exists(pk)"),
					ReturnValuesOnConditionCheckFailure: types.ReturnValuesOnConditionCheckFailureAllOld,
				},
			},
			{
				Put: &types.Put{
					TableName:                           &s.TableName,
					Item:                                mParty,
					ConditionExpression:                 aws.String("attribute_not_exists(sk)"),
					ReturnValuesOnConditionCheckFailure: types.ReturnValuesOnConditionCheckFailureAllOld,
				},
			},
		},
	})
	if err != nil {
		return Contract{}, err
	}
	return contract.Contract, nil
}

func (c Contract) GetKey(rangeKey string) map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(fmt.Sprintf("contract#%s", c.ID))
	if err != nil {
		panic(err)
	}
	sk, err := attributevalue.Marshal(rangeKey)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"pk": pk, "sk": sk}
}

func (s *DynamoDBStore) GetContract(ctx context.Context, id string) (Contract, error) {
	contract := Contract{
		ID: id,
	}
	response, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: contract.GetKey("info"), TableName: aws.String(s.TableName),
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

func (s *DynamoDBStore) GetContractOwner(ctx context.Context, id string) (Party, error) {
	contract := Contract{
		ID: id,
	}
	party := Party{
		ContractID: 0,
	}
	response, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: contract.GetKey("role#owner"), TableName: aws.String(s.TableName),
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
