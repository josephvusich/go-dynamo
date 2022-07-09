package dynamo

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/require"
)

func TestMarshaling(t *testing.T) {
	assert := require.New(t)

	expectJson := `{"B":true,"I":-10,"S":"Hello World"}`

	type inner struct {
		B bool
		I int
		S string
	}

	type outer struct {
		String string      `dynamodbav:"String,omitempty"`
		Nested JSON[inner] `dynamodbav:"Nested"`
	}

	jAttr := JSON[inner]{
		Value: inner{
			B: true,
			I: -10,
			S: "Hello World",
		},
	}

	var outerStruct outer
	outerStruct.Nested = jAttr

	fields, err := attributevalue.MarshalMap(outerStruct)
	assert.NoError(err)
	assert.Len(fields, 1)
	assert.IsType(&types.AttributeValueMemberS{}, fields["Nested"])
	avs, ok := fields["Nested"].(*types.AttributeValueMemberS)
	assert.True(ok)
	assert.Equal(expectJson, avs.Value)

	av, err := attributevalue.Marshal(&jAttr)
	assert.NoError(err)
	assert.IsType(&types.AttributeValueMemberS{}, av)
	avs, ok = fields["Nested"].(*types.AttributeValueMemberS)
	assert.True(ok)
	assert.Equal(expectJson, avs.Value)

	var outerRecv outer
	err = attributevalue.UnmarshalMap(map[string]types.AttributeValue{
		"String": &types.AttributeValueMemberS{Value: "foobar"},
		"Nested": &types.AttributeValueMemberS{Value: expectJson},
	}, &outerRecv)
	assert.NoError(err)
	assert.Equal(outerRecv.String, "foobar")
	assert.Equal(jAttr, outerRecv.Nested)
}
