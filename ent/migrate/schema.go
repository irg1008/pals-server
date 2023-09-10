// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthRequestsColumns holds the columns for the "auth_requests" table.
	AuthRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "active", Type: field.TypeBool},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "token", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"confirmEmail", "resetPassword"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// AuthRequestsTable holds the schema information for the "auth_requests" table.
	AuthRequestsTable = &schema.Table{
		Name:       "auth_requests",
		Columns:    AuthRequestsColumns,
		PrimaryKey: []*schema.Column{AuthRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "auth_requests_users_requests",
				Columns:    []*schema.Column{AuthRequestsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "is_confirmed", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserDataColumns holds the columns for the "user_data" table.
	UserDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "user"}, Default: "user"},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UserDataTable holds the schema information for the "user_data" table.
	UserDataTable = &schema.Table{
		Name:       "user_data",
		Columns:    UserDataColumns,
		PrimaryKey: []*schema.Column{UserDataColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthRequestsTable,
		UsersTable,
		UserDataTable,
	}
)

func init() {
	AuthRequestsTable.ForeignKeys[0].RefTable = UsersTable
}
