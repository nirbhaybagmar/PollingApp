// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"poll/ent/question"
	"poll/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Question is the model entity for the Question schema.
type Question struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// PollExpiry holds the value of the "pollExpiry" field.
	PollExpiry time.Time `json:"pollExpiry,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QuestionQuery when eager-loading is set.
	Edges QuestionEdges `json:"edges"`
}

// QuestionEdges holds the relations/edges for other nodes in the graph.
type QuestionEdges struct {
	// Answer holds the value of the answer edge.
	Answer []*Answer `json:"answer,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AnswerOrErr returns the Answer value or an error if the edge
// was not loaded in eager-loading.
func (e QuestionEdges) AnswerOrErr() ([]*Answer, error) {
	if e.loadedTypes[0] {
		return e.Answer, nil
	}
	return nil, &NotLoadedError{edge: "answer"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QuestionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Question) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case question.FieldID, question.FieldUserID:
			values[i] = new(sql.NullInt64)
		case question.FieldBody:
			values[i] = new(sql.NullString)
		case question.FieldPollExpiry, question.FieldCreatedAt, question.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Question", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Question fields.
func (q *Question) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case question.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			q.ID = int(value.Int64)
		case question.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				q.Body = value.String
			}
		case question.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				q.UserID = int(value.Int64)
			}
		case question.FieldPollExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field pollExpiry", values[i])
			} else if value.Valid {
				q.PollExpiry = value.Time
			}
		case question.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				q.CreatedAt = value.Time
			}
		case question.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				q.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryAnswer queries the "answer" edge of the Question entity.
func (q *Question) QueryAnswer() *AnswerQuery {
	return (&QuestionClient{config: q.config}).QueryAnswer(q)
}

// QueryUser queries the "user" edge of the Question entity.
func (q *Question) QueryUser() *UserQuery {
	return (&QuestionClient{config: q.config}).QueryUser(q)
}

// Update returns a builder for updating this Question.
// Note that you need to call Question.Unwrap() before calling this method if this Question
// was returned from a transaction, and the transaction was committed or rolled back.
func (q *Question) Update() *QuestionUpdateOne {
	return (&QuestionClient{config: q.config}).UpdateOne(q)
}

// Unwrap unwraps the Question entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (q *Question) Unwrap() *Question {
	tx, ok := q.config.driver.(*txDriver)
	if !ok {
		panic("ent: Question is not a transactional entity")
	}
	q.config.driver = tx.drv
	return q
}

// String implements the fmt.Stringer.
func (q *Question) String() string {
	var builder strings.Builder
	builder.WriteString("Question(")
	builder.WriteString(fmt.Sprintf("id=%v", q.ID))
	builder.WriteString(", body=")
	builder.WriteString(q.Body)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", q.UserID))
	builder.WriteString(", pollExpiry=")
	builder.WriteString(q.PollExpiry.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(q.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(q.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Questions is a parsable slice of Question.
type Questions []*Question

func (q Questions) config(cfg config) {
	for _i := range q {
		q[_i].config = cfg
	}
}
