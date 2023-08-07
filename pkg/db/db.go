package db

import (
	"context"
	"database/sql"
	"irg1008/next-go/ent"
	"irg1008/next-go/pkg/config"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "modernc.org/sqlite"
)

func getDBConnection() *sql.DB {
	db, err := sql.Open("sqlite", config.GetConfig().DB_URL)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return db
}

func ConnectDB() {
	db := getDBConnection()
	defer db.Close()

	driver := entsql.OpenDB(dialect.SQLite, db)
	client := ent.NewClient(ent.Driver(driver))
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
