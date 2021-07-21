package base

import (
	"bp_service/lib/structs"
	"context"
	"log"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

//creates all dependency connections
func createConnections(dependencies *structs.Depend) {
	createPostgresconnection(dependencies)
}

func createPostgresconnection(dependencies *structs.Depend) {
	log.Println((*dependencies).Config)
	if ((*dependencies).Config.Postgres == structs.Database{}) {
		return
	}

	if (*dependencies).Connection.Postgres != nil {
		return
	}
	rawPgxConfig := (*dependencies).Config.Postgres
	pgxConfig, err := pgxpool.ParseConfig(rawPgxConfig.ConnectionString)
	conn, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal("Error connecting to postgres", err)
	}
	dependencies.Connection.Postgres = conn
}
