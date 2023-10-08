package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.Int64("created_at").Default(time.Now().UnixMilli()).SchemaType(map[string]string{
			dialect.Postgres: "bigint",
		}).Immutable(),
		field.Int64("due_date").Optional().Nillable().SchemaType(map[string]string{
			dialect.Postgres: "bigint",
		}),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("owner", User.Type).
			Ref("todos"),
	}
}
