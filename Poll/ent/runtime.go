// Code generated by entc, DO NOT EDIT.

package ent

import (
	"poll/ent/answer"
	"poll/ent/question"
	"poll/ent/schema"
	"poll/ent/user"
	"poll/ent/vote"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescNumOfVotes is the schema descriptor for numOfVotes field.
	answerDescNumOfVotes := answerFields[2].Descriptor()
	// answer.DefaultNumOfVotes holds the default value on creation for the numOfVotes field.
	answer.DefaultNumOfVotes = answerDescNumOfVotes.Default.(int)
	// answerDescCreatedAt is the schema descriptor for created_at field.
	answerDescCreatedAt := answerFields[3].Descriptor()
	// answer.DefaultCreatedAt holds the default value on creation for the created_at field.
	answer.DefaultCreatedAt = answerDescCreatedAt.Default.(func() time.Time)
	// answerDescUpdatedAt is the schema descriptor for updated_at field.
	answerDescUpdatedAt := answerFields[4].Descriptor()
	// answer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	answer.DefaultUpdatedAt = answerDescUpdatedAt.Default.(func() time.Time)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescCreatedAt is the schema descriptor for created_at field.
	questionDescCreatedAt := questionFields[4].Descriptor()
	// question.DefaultCreatedAt holds the default value on creation for the created_at field.
	question.DefaultCreatedAt = questionDescCreatedAt.Default.(func() time.Time)
	// questionDescUpdatedAt is the schema descriptor for updated_at field.
	questionDescUpdatedAt := questionFields[5].Descriptor()
	// question.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	question.DefaultUpdatedAt = questionDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	voteFields := schema.Vote{}.Fields()
	_ = voteFields
	// voteDescCreatedAt is the schema descriptor for created_at field.
	voteDescCreatedAt := voteFields[4].Descriptor()
	// vote.DefaultCreatedAt holds the default value on creation for the created_at field.
	vote.DefaultCreatedAt = voteDescCreatedAt.Default.(func() time.Time)
	// voteDescUpdatedAt is the schema descriptor for updated_at field.
	voteDescUpdatedAt := voteFields[5].Descriptor()
	// vote.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vote.DefaultUpdatedAt = voteDescUpdatedAt.Default.(func() time.Time)
}
