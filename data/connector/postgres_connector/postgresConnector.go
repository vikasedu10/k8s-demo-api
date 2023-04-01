package postgres_connector

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func PostgresConnector () (*pgxpool.Pool, bool) {
	var databaseURL string
	postgresDBName := os.Getenv("POSTGRES_DATABASE_NAME_CAC_PLATFORM")
	postgresDBUsername := os.Getenv("POSTGRES_DB_USERNAME")
	postgresDBPassword := os.Getenv("POSTGRES_DB_PASSWORD")
	dnsUrl := os.Getenv("DNS_URL")
	fmt.Println("postgresDBName: ", postgresDBName)
	fmt.Println("dnsUrl: ", dnsUrl)
	fmt.Println("postgresDBUsername: ", postgresDBUsername)
	fmt.Println("postgresDBPassword: ", postgresDBPassword)

	databaseURL = "postgres://" + postgresDBUsername + ":" + postgresDBPassword + "@"+dnsUrl+":5432/" + postgresDBName
	conn, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Printf("ERROR: Unable to connect to postgres: %v\n", err)
		return nil, false
	}
	return conn, true
}
