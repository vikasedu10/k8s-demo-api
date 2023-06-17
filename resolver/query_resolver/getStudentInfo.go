package query_resolver

import (
	"platform-api/data/connector/postgres_connector"
	"context"
	"log"
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

func GetStudentList (param graphql.ResolveParams) (interface{}, error) {
	dbPool, isDBPool := postgres_connector.PostgresConnector()
	defer dbPool.Close()
	if !isDBPool {
		return studentResponse{StatusCode: http.StatusInternalServerError, Message: "Unable to connect to postgres database"}, nil
	}
	rows, err := dbPool.Query(context.Background(), "SELECT * FROM cac_master.get_student_info();")
	if err!= nil {
		log.Println("ERROR: ", err)
		return studentResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}, nil
	}
	var studentsObj [] studentInfo
	for rows.Next() {
		var singleObj studentInfo
		err := rows.Scan(&singleObj.Student_Roll, &singleObj.Student_Name)
		if err != nil {
			log.Println("ERROR: ", err)
            return studentResponse{StatusCode: http.StatusBadRequest, Message: err.Error()}, nil
		}
		studentsObj = append(studentsObj, singleObj)
	}
	return studentResponse{StatusCode: 200, Message: "Query Executed", Data: studentsObj}, nil
}