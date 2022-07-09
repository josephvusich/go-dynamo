# go-dynamo

A utility library for interacting with DynamoDB in [AWS SDK v2](https://github.com/aws/aws-sdk-go-v2).

Supports marshaling of JSON-serializable types to/from DynamoDB String attributes.

## Example
```go
package main

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/josephvusich/go-dynamo"
)

type Nested struct {
	String string
}

type JSONRecord struct {
	Number int
	Blob   dynamo.JSON[Nested]
}

type MapRecord struct {
	Number int
	Blob   Nested
}

func main() {
	jsonRecord := &JSONRecord{Number: -1}
	jsonRecord.Blob.Value = Nested{"Hello World"}
	
	attributevalue.MarshalMap(jsonRecord)
	// Output record represented as DynamoDB JSON:
	// {
	//   "Number": { "N": "-1" },
	//   "Blob":   { "S": "{\"String\":\"Hello World\"}" }
	// }
	
	mapRecord := &MapRecord{Number: -1}
	mapRecord.Blob = Nested{"Hello World"}
	
	attributevalue.MarshalMap(mapRecord)
	// Output record represented as DynamoDB JSON:
	// {
	//   "Number": { "N": "-1" },
	//   "Blob":   { "M": {
	//     "String": { "S": "Hello World" }
	//   }}
	// }
}
```
