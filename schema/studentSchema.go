package schema

import (
	"github.com/graphql-go/graphql"
)

var TestingResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TestingResponse",
		Fields: graphql.Fields{
			"Status_Code": &graphql.Field{
				Type: graphql.Int,
			},
			"Message": &graphql.Field{
				Type: graphql.String,
			},
			"Data": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var StudentResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StudentResponse",
		Fields: graphql.Fields{
			"Status_Code": &graphql.Field{
				Type: graphql.Int,
			},
			"Message": &graphql.Field{
				Type: graphql.String,
			},
			"Data": &graphql.Field{
				Type: graphql.NewList(StudentDetails),
			},
		},
	},
)

var StudentDetails = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StudentDetails",
		Fields: graphql.Fields{
			"Student_Name": &graphql.Field{
				Type: graphql.String,
			},
			"Student_Roll": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var StudentArg = graphql.FieldConfigArgument {
	"Student_Roll": &graphql.ArgumentConfig {
		Type: graphql.Int,
	},
	"Student_Name": &graphql.ArgumentConfig {
		Type: graphql.String,
	},
}