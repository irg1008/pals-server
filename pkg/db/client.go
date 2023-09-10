package db

import (
	"context"
	"irg1008/pals/ent"
	"irg1008/pals/ent/migrate"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type DB struct {
	*ent.Client
	Ctx context.Context
}

type CustomContext struct {
	echo.Context
	DB *DB
}

func getDBConnection(url string) *ent.Client {
	db, err := ent.Open("postgres", url)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return db
}

func connectDB(url string) (*ent.Client, context.Context) {
	client := getDBConnection(url)

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

func NewDB(url string) *DB {
	client, ctx := connectDB(url)
	return &DB{client, ctx}
}
