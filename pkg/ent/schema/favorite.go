package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Favorite holds the schema definition for the Favorite entity.
type Favorite struct {
	ent.Schema
}

// Fields of the Favorite.
func (Favorite) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("post_id", uuid.UUID{}),
		field.Time("created_at").Default(func() time.Time {
			return time.Now().Truncate(time.Millisecond)
		}),
	}
}

// Edges of the Favorite.
func (Favorite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("favorites").Field("user_id").Unique().Required(),
		edge.From("post", Post.Type).Ref("favorites").Field("post_id").Unique().Required(),
	}
}

// Indexes of the Favorite.
func (Favorite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "post_id").Unique(),
	}
}
