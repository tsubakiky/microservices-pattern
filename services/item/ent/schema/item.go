package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("customer_id"),
		field.String("title"),
		field.Int64("price").
			Positive(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return nil
}
