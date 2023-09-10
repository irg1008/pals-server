package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password").Sensitive(),
		field.Enum("role").Values("admin", "user").Default("user"),
		field.Bool("is_confirmed").Default(false).StructTag(`json:"-"`),
		field.Time("created_at").Default(time.Now).Immutable().StructTag(`json:"-"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requests", AuthRequest.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
