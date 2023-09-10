// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"irg1008/pals/ent/authrequest"
	"irg1008/pals/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AuthRequest is the model entity for the AuthRequest schema.
type AuthRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expiresAt"`
	// Token holds the value of the "token" field.
	Token uuid.UUID `json:"token,omitempty"`
	// Type holds the value of the "type" field.
	Type authrequest.Type `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"-"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthRequestQuery when eager-loading is set.
	Edges        AuthRequestEdges `json:"-"`
	selectValues sql.SelectValues
}

// AuthRequestEdges holds the relations/edges for other nodes in the graph.
type AuthRequestEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthRequestEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authrequest.FieldActive:
			values[i] = new(sql.NullBool)
		case authrequest.FieldID, authrequest.FieldUserID:
			values[i] = new(sql.NullInt64)
		case authrequest.FieldType:
			values[i] = new(sql.NullString)
		case authrequest.FieldExpiresAt, authrequest.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case authrequest.FieldToken:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthRequest fields.
func (ar *AuthRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int(value.Int64)
		case authrequest.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				ar.Active = value.Bool
			}
		case authrequest.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				ar.ExpiresAt = value.Time
			}
		case authrequest.FieldToken:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value != nil {
				ar.Token = *value
			}
		case authrequest.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ar.Type = authrequest.Type(value.String)
			}
		case authrequest.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ar.CreatedAt = value.Time
			}
		case authrequest.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ar.UserID = int(value.Int64)
			}
		default:
			ar.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AuthRequest.
// This includes values selected through modifiers, order, etc.
func (ar *AuthRequest) Value(name string) (ent.Value, error) {
	return ar.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the AuthRequest entity.
func (ar *AuthRequest) QueryUser() *UserQuery {
	return NewAuthRequestClient(ar.config).QueryUser(ar)
}

// Update returns a builder for updating this AuthRequest.
// Note that you need to call AuthRequest.Unwrap() before calling this method if this AuthRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AuthRequest) Update() *AuthRequestUpdateOne {
	return NewAuthRequestClient(ar.config).UpdateOne(ar)
}

// Unwrap unwraps the AuthRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AuthRequest) Unwrap() *AuthRequest {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthRequest is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AuthRequest) String() string {
	var builder strings.Builder
	builder.WriteString("AuthRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", ar.Active))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(ar.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(fmt.Sprintf("%v", ar.Token))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ar.Type))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ar.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ar *AuthRequest) MarshalJSON() ([]byte, error) {
	type Alias AuthRequest
	return json.Marshal(&struct {
		*Alias
		AuthRequestEdges
	}{
		Alias:            (*Alias)(ar),
		AuthRequestEdges: ar.Edges,
	})
}

// AuthRequests is a parsable slice of AuthRequest.
type AuthRequests []*AuthRequest
