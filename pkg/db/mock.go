package db

import (
	"context"
	"irg1008/pals/ent"
	"irg1008/pals/ent/enttest"
	"irg1008/pals/ent/migrate"
	"testing"
)

func GetMockedDB(t *testing.T) *DB {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	ctx := context.Background()

	return &DB{Client: client, Ctx: ctx}
}
