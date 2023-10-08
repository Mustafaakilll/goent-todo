// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFullname holds the string denoting the fullname field in the database.
	FieldFullname = "fullname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeTodos holds the string denoting the todos edge name in mutations.
	EdgeTodos = "todos"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TodosTable is the table that holds the todos relation/edge.
	TodosTable = "todos"
	// TodosInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodosInverseTable = "todos"
	// TodosColumn is the table column denoting the todos relation/edge.
	TodosColumn = "user_todos"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFullname,
	FieldEmail,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// FullnameValidator is a validator for the "fullname" field. It is called by the builders before save.
	FullnameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFullname orders the results by the fullname field.
func ByFullname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullname, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByTodosCount orders the results by todos count.
func ByTodosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodosStep(), opts...)
	}
}

// ByTodos orders the results by todos terms.
func ByTodos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TodosTable, TodosColumn),
	)
}
