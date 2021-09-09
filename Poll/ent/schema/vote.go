package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.Int("question_id"),
		field.Int("answer_id"),
		field.Int("user_id"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("question", Question.Type).Field("question_id").Unique().Required(),
		edge.To("answer", Answer.Type).Field("answer_id").Unique().Required(),
		edge.To("user", User.Type).Field("user_id").Unique().Required(),
	}
}
