package query_resolver

import (
	"platform-api/data/models/response"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func GetTestingResponse(param graphql.ResolveParams) (interface{}, error) {
	fmt.Println("Testing GetTestingResponse function executed successfully!")
	return response.TestingResponse{
		Status_Code: http.StatusOK,
		Message:     "Connection to API is working.",
		Data:        "Status Executed.",
	}, nil
}
