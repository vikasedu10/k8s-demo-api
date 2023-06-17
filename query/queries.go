package query

import (
	"platform-api/resolver/query_resolver"
	"platform-api/schema"
	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(
	graphql.ObjectConfig {
        Name: "Query",
        Fields: graphql.Fields{
			// http://localhost:9091/goengine?query={getTestingResponse{Message,Status_Code,Data}}
			"getTestingResponse": &graphql.Field {
				Type: schema.TestingResponse,
				Description: "Get Testing Response to check API execution",
				Resolve: query_resolver.GetTestingResponse,
			},
			// http://localhost:9091/goengine?query={getStudentList{Message,Status_Code,Data}}
			"getStudentList": &graphql.Field {
				Type: schema.StudentResponse,
				Description: "Get list of students",
				Resolve: query_resolver.GetStudentList,
			},
        },
	},
)
