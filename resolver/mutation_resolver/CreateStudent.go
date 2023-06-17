package mutation_resolver

import (
	"platform-api/data/connector/postgres_connector"
	"context"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type studentResponse struct {
	StatusCode int
	Message string	
	Data []studentInfo
}

type studentInfo struct {
	Student_Roll int
    Student_Name string
}

func AddClusterDetail(params graphql.ResolveParams) (interface{}, error) {
	fmt.Println("Enters student Add")
	studentName, okStudentName := params.Args["Student_Name"].(string)
	if !okStudentName {
		return studentResponse{StatusCode: http.StatusBadRequest, Message: "Provide parametre - Student_Name"}, nil
	}
	studentRoll, okStudentRoll := params.Args["Student_Roll"].(int)
	if !okStudentRoll {
		return studentResponse{StatusCode: http.StatusBadRequest, Message: "Provide parametre - Student_Roll"}, nil
	}
	dbPool, isDbPool := postgres_connector.PostgresConnector()
	if !isDbPool {
		return studentResponse{StatusCode: http.StatusBadRequest, Message: "Unable to connect to Postgres DB"}, nil
	}
	defer dbPool.Close()
	_, err := dbPool.Exec(context.Background(), "select cac_master.add_new_student_detail($1,$2)", studentName, studentRoll)
	if err != nil {
		responses := studentResponse{
			StatusCode: http.StatusInternalServerError, 
			Message: "ERROR: Unable to add student details: " + err.Error(),
		}
		return responses, nil
	}
	var studentsObj []studentInfo
	studentOjb := studentInfo {
		Student_Name: studentName,
		Student_Roll: studentRoll,
	}	
	studentsObj = append(studentsObj, studentOjb)
	return studentResponse{StatusCode: http.StatusOK, Message: "Success: Student details added to DB", Data: studentsObj}, nil
}