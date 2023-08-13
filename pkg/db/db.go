package db

import (
	"context"
	"irg1008/next-go/ent"
	"irg1008/next-go/ent/migrate"
	"irg1008/next-go/pkg/config"
	"log"

	"github.com/labstack/echo/v4"
)

type DB struct {
	*ent.Client
	Ctx context.Context
}

type CustomContext struct {
	echo.Context
	DB *DB
}

func getDBConnection() *ent.Client {
	db, err := ent.Open("sqlite3", config.Env.DBUrl)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return db
}

func connectDB() (*ent.Client, context.Context) {
	client := getDBConnection()

	ctx := context.Background()
	err := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, ctx
}

func New() *DB {
	client, ctx := connectDB()
	return &DB{client, ctx}
}
