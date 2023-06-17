package main

import (
	"fmt"
	"platform-api/config"
	"platform-api/mutation"
	"platform-api/query"
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"

	"platform-api/data/services/cors_policy"
)

var PlatformSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: query.Query,
		Mutation: mutation.MutationType,
	},
)

func init() {
    // loads values from .env into the system before main execution
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }
}

func main () {	
	config.ConfigurationServer("localhost")
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Testing home")
	})
	http.HandleFunc("/goengine", func(w http.ResponseWriter, r *http.Request) {
		cors_policy.SetupCorsResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		graphqlHttpHandler := handler.New(&handler.Config {
		    Schema: &PlatformSchema,
			Pretty: true,
			GraphiQL: true,
        })
		graphqlHttpHandler.ServeHTTP(w, r)
	})
	portNumber := 9091
	fmt.Printf("Server is up and running at port http://localhost:%d\n", portNumber)
	http.ListenAndServe(":9091", nil)
}
