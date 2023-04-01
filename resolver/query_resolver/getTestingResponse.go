package query_resolver

import (
	"cac-platform-api/data/models/response"
	"github.com/graphql-go/graphql"
	"net/http"
)

func GetTestingResponse (param graphql.ResolveParams) (interface{}, error) {
	return response.TestingResponse {
		Status_Code: http.StatusOK,
		Message: "Connection to API is working.",
		Data: "Status Executed.",
	}, nil
}