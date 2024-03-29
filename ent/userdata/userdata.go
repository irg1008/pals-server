// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package userdata

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the userdata type in the database.
	Label = "user_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthID holds the string denoting the auth_id field in the database.
	FieldAuthID = "auth_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the userdata in the database.
	Table = "user_data"
)

// Columns holds all SQL columns for userdata fields.
var Columns = []string{
	FieldID,
	FieldAuthID,
	FieldName,
	FieldEmail,
	FieldPicture,
	FieldRole,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Role defines the type for the "role" enum field.
type Role string

// RoleUser is the default value of the Role enum.
const DefaultRole = RoleUser

// Role values.
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleUser:
		return nil
	default:
		return fmt.Errorf("userdata: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the UserData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAuthID orders the results by the auth_id field.
func ByAuthID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPicture orders the results by the picture field.
func ByPicture(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPicture, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
