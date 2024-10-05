package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("author_id", uuid.UUID{}),
		field.String("content").Optional(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now().Truncate(time.Millisecond)
		}),
		field.Time("updated_at").Default(func() time.Time {
			return time.Now().Truncate(time.Millisecond)
		}),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Field("author_id").Unique().Required(),
	}
}
