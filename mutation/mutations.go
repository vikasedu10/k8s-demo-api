package mutation

import (
	"github.com/graphql-go/graphql"
	"platform-api/schema"
	"platform-api/resolver/mutation_resolver"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
    Fields: graphql.Fields{
		// http://localhost:9091/goengine?query=mutation+_{updateStudentInfo(Student_Roll:2,Student_Name: "Vikas"){Status_Code,Message,Data}}
		"updateStudentInfo": &graphql.Field{
            Type: schema.TestingResponse,
			Description: "Update cluster details",
			Args: schema.StudentArg,
			Resolve: mutation_resolver.UpdateStudentInfo,
		},
		// http://localhost:9091/goengine?query=mutation+_{updateStudentInfo(Student_Roll:2,Student_Name: "Vikas"){Status_Code,Message,Data}}
		"createStudent": &graphql.Field{
            Type: schema.StudentResponse,
			Description: "Update cluster details",
			Args: schema.StudentArg,
			Resolve: mutation_resolver.UpdateStudentInfo,
		},
	},
})