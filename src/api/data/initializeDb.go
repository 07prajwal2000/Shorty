package data

import (
	"context"
	"log"
	"os"
	data "shorty/data/postgresData"

	"github.com/jackc/pgx/v5"
)

var dbContext = context.Background()
var postgresDb *data.Queries
var conn *pgx.Conn

func InitializeDb() {
	connectionString := os.Getenv("POSTGRES_CONN")
	// InitializeUrlData()
	conn, err := pgx.Connect(dbContext, connectionString)
	if err != nil {
		log.Fatalf("error connecting db.\nError: %s", err.Error())
		return
	}
	postgresDb = data.New(conn)
}

func UnInitializeDb() {
	conn.Close(dbContext)
}

func CreateUrlRepository() IUrlRepository {
	return &UrlPostgresRepository{}
}
