package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// `User`型の "owner "という逆エッジを作成し
		// `Ref`メソッドを使って明示的に
		// (Userスキーマの)"cars"エッジを参照します
		edge.From("owner", User.Type).
			Ref("cars").
			// エッジをuniqueに設定することで、
			// 1台の車は1人のオーナーのみが所有することを保証する
			Unique(),
	}
}
