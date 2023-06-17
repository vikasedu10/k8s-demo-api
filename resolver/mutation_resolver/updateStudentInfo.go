package mutation_resolver

import (
	"fmt"
	"platform-api/data/models/response"
	"net/http"
	"github.com/graphql-go/graphql"
)

func UpdateStudentInfo(params graphql.ResolveParams) (interface{}, error) {
	fmt.Println("Student Records updated.")	
	return response.StudentResponse{Status_Code: http.StatusOK, Message: "Executed", Data: "Student records updated"}, nil
}