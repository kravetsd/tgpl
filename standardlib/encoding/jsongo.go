package encoding

import (
	"encoding/json"
	"fmt"
)

type StatementEntry struct {
	Effect   string
	Action   []string
	Resource string
}

type PolicyDocument struct {
	Version   string
	Statement []StatementEntry
}

func MarshalMe() {
	fmt.Println("This is about json marshalin")

	policy := PolicyDocument{
		Version: "2012-10-17",
		Statement: []StatementEntry{
			StatementEntry{
				Effect: "Allow",
				Action: []string{
					"logs:CreateLogGroup", // Allow for creating log groups
				},
				Resource: "RESOURCE ARN FOR logs:*",
			},
			StatementEntry{
				Effect: "Allow",
				// Allows for DeleteItem, GetItem, PutItem, Scan, and UpdateItem
				Action: []string{
					"dynamodb:DeleteItem",
					"dynamodb:GetItem",
					"dynamodb:PutItem",
					"dynamodb:Scan",
					"dynamodb:UpdateItem",
				},
				Resource: "RESOURCE ARN FOR dynamodb:*",
			},
		},
	}

	b, err := json.Marshal(&policy)
	if err != nil {
		fmt.Println("Error marshaling policy", err)
		return
	}

	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

	fmt.Println(string(b))

}
