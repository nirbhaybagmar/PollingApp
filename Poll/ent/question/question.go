// Code generated by entc, DO NOT EDIT.

package question

import (
	"time"
)

const (
	// Label holds the string label denoting the question type in the database.
	Label = "question"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPollExpiry holds the string denoting the pollexpiry field in the database.
	FieldPollExpiry = "poll_expiry"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAnswer holds the string denoting the answer edge name in mutations.
	EdgeAnswer = "answer"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the question in the database.
	Table = "questions"
	// AnswerTable is the table that holds the answer relation/edge. The primary key declared below.
	AnswerTable = "question_answer"
	// AnswerInverseTable is the table name for the Answer entity.
	// It exists in this package in order to avoid circular dependency with the "answer" package.
	AnswerInverseTable = "answers"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "questions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for question fields.
var Columns = []string{
	FieldID,
	FieldBody,
	FieldUserID,
	FieldPollExpiry,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// AnswerPrimaryKey and AnswerColumn2 are the table columns denoting the
	// primary key for the answer relation (M2M).
	AnswerPrimaryKey = []string{"question_id", "answer_id"}
)

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)