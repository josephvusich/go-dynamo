package dynamo

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type JSON[T any] struct{ Value T }

var _ attributevalue.Marshaler = (*JSON[struct{}])(nil)
var _ attributevalue.Unmarshaler = (*JSON[struct{}])(nil)

func (j JSON[T]) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	b, err := json.Marshal(&j.Value)
	if err != nil {
		return nil, err
	}
	return &types.AttributeValueMemberS{Value: string(b)}, nil
}

func (j *JSON[T]) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	s, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return errors.New("expected string attribute")
	}
	return json.Unmarshal([]byte(s.Value), &j.Value)
}
