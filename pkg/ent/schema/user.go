package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("username").NotEmpty().Unique(),
		field.String("display_name").Optional(),
		field.String("avatar_url").Optional(),
		field.String("cover_url").Optional(),
		field.String("biography").Optional(),
		field.Time("birthdate").Optional(),
		field.String("line_id").Optional(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now().Truncate(time.Millisecond)
		}),
		field.Time("updated_at").Default(func() time.Time {
			return time.Now().Truncate(time.Millisecond)
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("following", User.Type).From("followers"),
	}
}
