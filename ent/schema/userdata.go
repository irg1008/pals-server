package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserData holds the schema definition for the UserData entity.
type UserData struct {
	ent.Schema
}

// Fields of the UserData.
func (UserData) Fields() []ent.Field {
	return []ent.Field{
		field.String("auth_id").Unique().Immutable().Comment("User ID from auth service. Prefixed with email, google, etc").StructTag(`json:"authId"`),
		field.String("name"),
		field.String("email").Optional().Unique(),
		field.String("picture").Optional(),
		field.Enum("role").Values("admin", "user").Default("user"),
		field.Time("created_at").Default(time.Now).Immutable().StructTag(`json:"-"`),
	}
}

// Edges of the UserData.
func (UserData) Edges() []ent.Edge {
	return nil
}
