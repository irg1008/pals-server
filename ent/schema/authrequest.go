package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthRequest holds the schema definition for the AuthRequest entity.
type AuthRequest struct {
	ent.Schema
}

// Fields of the AuthRequest.
func (AuthRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active"),
		field.Time("expires_at"),
		field.UUID("token", uuid.UUID{}).Unique().Immutable().Default(uuid.New),
		field.Enum("type").Values("confirmEmail", "resetPassword"),
		field.Time("created_at").Default(time.Now).Immutable().StructTag(`json:"-"`),
		field.Int("user_id"),
	}
}

// Edges of the AuthRequest.
func (AuthRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("requests").
			Unique().
			Required().
			Field("user_id"),
	}
}
