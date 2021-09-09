// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnswersColumns holds the columns for the "answers" table.
	AnswersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "body", Type: field.TypeString},
		{Name: "num_of_votes", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AnswersTable holds the schema information for the "answers" table.
	AnswersTable = &schema.Table{
		Name:       "answers",
		Columns:    AnswersColumns,
		PrimaryKey: []*schema.Column{AnswersColumns[0]},
	}
	// QuestionsColumns holds the columns for the "questions" table.
	QuestionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "body", Type: field.TypeString},
		{Name: "poll_expiry", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// QuestionsTable holds the schema information for the "questions" table.
	QuestionsTable = &schema.Table{
		Name:       "questions",
		Columns:    QuestionsColumns,
		PrimaryKey: []*schema.Column{QuestionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "questions_users_user",
				Columns:    []*schema.Column{QuestionsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VotesColumns holds the columns for the "votes" table.
	VotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "question_id", Type: field.TypeInt, Nullable: true},
		{Name: "answer_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// VotesTable holds the schema information for the "votes" table.
	VotesTable = &schema.Table{
		Name:       "votes",
		Columns:    VotesColumns,
		PrimaryKey: []*schema.Column{VotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "votes_questions_question",
				Columns:    []*schema.Column{VotesColumns[3]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "votes_answers_answer",
				Columns:    []*schema.Column{VotesColumns[4]},
				RefColumns: []*schema.Column{AnswersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "votes_users_user",
				Columns:    []*schema.Column{VotesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// QuestionAnswerColumns holds the columns for the "question_answer" table.
	QuestionAnswerColumns = []*schema.Column{
		{Name: "question_id", Type: field.TypeInt},
		{Name: "answer_id", Type: field.TypeInt},
	}
	// QuestionAnswerTable holds the schema information for the "question_answer" table.
	QuestionAnswerTable = &schema.Table{
		Name:       "question_answer",
		Columns:    QuestionAnswerColumns,
		PrimaryKey: []*schema.Column{QuestionAnswerColumns[0], QuestionAnswerColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "question_answer_question_id",
				Columns:    []*schema.Column{QuestionAnswerColumns[0]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "question_answer_answer_id",
				Columns:    []*schema.Column{QuestionAnswerColumns[1]},
				RefColumns: []*schema.Column{AnswersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnswersTable,
		QuestionsTable,
		UsersTable,
		VotesTable,
		QuestionAnswerTable,
	}
)

func init() {
	QuestionsTable.ForeignKeys[0].RefTable = UsersTable
	VotesTable.ForeignKeys[0].RefTable = QuestionsTable
	VotesTable.ForeignKeys[1].RefTable = AnswersTable
	VotesTable.ForeignKeys[2].RefTable = UsersTable
	QuestionAnswerTable.ForeignKeys[0].RefTable = QuestionsTable
	QuestionAnswerTable.ForeignKeys[1].RefTable = AnswersTable
}
