// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Nulandmori/micorservices-pattern/services/item/ent/item"
	"github.com/google/uuid"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID string `json:"customer_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Price holds the value of the "price" field.
	Price int64 `json:"price,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldPrice:
			values[i] = new(sql.NullInt64)
		case item.FieldCustomerID, item.FieldTitle:
			values[i] = new(sql.NullString)
		case item.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Item", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case item.FieldCustomerID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[j])
			} else if value.Valid {
				i.CustomerID = value.String
			}
		case item.FieldTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[j])
			} else if value.Valid {
				i.Title = value.String
			}
		case item.FieldPrice:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[j])
			} else if value.Valid {
				i.Price = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", customer_id=")
	builder.WriteString(i.CustomerID)
	builder.WriteString(", title=")
	builder.WriteString(i.Title)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", i.Price))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}