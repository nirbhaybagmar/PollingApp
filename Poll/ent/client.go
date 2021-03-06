// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"poll/ent/migrate"

	"poll/ent/answer"
	"poll/ent/question"
	"poll/ent/user"
	"poll/ent/vote"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Answer is the client for interacting with the Answer builders.
	Answer *AnswerClient
	// Question is the client for interacting with the Question builders.
	Question *QuestionClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Vote is the client for interacting with the Vote builders.
	Vote *VoteClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Answer = NewAnswerClient(c.config)
	c.Question = NewQuestionClient(c.config)
	c.User = NewUserClient(c.config)
	c.Vote = NewVoteClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Answer:   NewAnswerClient(cfg),
		Question: NewQuestionClient(cfg),
		User:     NewUserClient(cfg),
		Vote:     NewVoteClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Answer:   NewAnswerClient(cfg),
		Question: NewQuestionClient(cfg),
		User:     NewUserClient(cfg),
		Vote:     NewVoteClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Answer.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Answer.Use(hooks...)
	c.Question.Use(hooks...)
	c.User.Use(hooks...)
	c.Vote.Use(hooks...)
}

// AnswerClient is a client for the Answer schema.
type AnswerClient struct {
	config
}

// NewAnswerClient returns a client for the Answer from the given config.
func NewAnswerClient(c config) *AnswerClient {
	return &AnswerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `answer.Hooks(f(g(h())))`.
func (c *AnswerClient) Use(hooks ...Hook) {
	c.hooks.Answer = append(c.hooks.Answer, hooks...)
}

// Create returns a create builder for Answer.
func (c *AnswerClient) Create() *AnswerCreate {
	mutation := newAnswerMutation(c.config, OpCreate)
	return &AnswerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Answer entities.
func (c *AnswerClient) CreateBulk(builders ...*AnswerCreate) *AnswerCreateBulk {
	return &AnswerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Answer.
func (c *AnswerClient) Update() *AnswerUpdate {
	mutation := newAnswerMutation(c.config, OpUpdate)
	return &AnswerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnswerClient) UpdateOne(a *Answer) *AnswerUpdateOne {
	mutation := newAnswerMutation(c.config, OpUpdateOne, withAnswer(a))
	return &AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnswerClient) UpdateOneID(id int) *AnswerUpdateOne {
	mutation := newAnswerMutation(c.config, OpUpdateOne, withAnswerID(id))
	return &AnswerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Answer.
func (c *AnswerClient) Delete() *AnswerDelete {
	mutation := newAnswerMutation(c.config, OpDelete)
	return &AnswerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AnswerClient) DeleteOne(a *Answer) *AnswerDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AnswerClient) DeleteOneID(id int) *AnswerDeleteOne {
	builder := c.Delete().Where(answer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnswerDeleteOne{builder}
}

// Query returns a query builder for Answer.
func (c *AnswerClient) Query() *AnswerQuery {
	return &AnswerQuery{
		config: c.config,
	}
}

// Get returns a Answer entity by its id.
func (c *AnswerClient) Get(ctx context.Context, id int) (*Answer, error) {
	return c.Query().Where(answer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnswerClient) GetX(ctx context.Context, id int) *Answer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuestionID queries the question_id edge of a Answer.
func (c *AnswerClient) QueryQuestionID(a *Answer) *QuestionQuery {
	query := &QuestionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(answer.Table, answer.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, answer.QuestionIDTable, answer.QuestionIDPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AnswerClient) Hooks() []Hook {
	return c.hooks.Answer
}

// QuestionClient is a client for the Question schema.
type QuestionClient struct {
	config
}

// NewQuestionClient returns a client for the Question from the given config.
func NewQuestionClient(c config) *QuestionClient {
	return &QuestionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `question.Hooks(f(g(h())))`.
func (c *QuestionClient) Use(hooks ...Hook) {
	c.hooks.Question = append(c.hooks.Question, hooks...)
}

// Create returns a create builder for Question.
func (c *QuestionClient) Create() *QuestionCreate {
	mutation := newQuestionMutation(c.config, OpCreate)
	return &QuestionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Question entities.
func (c *QuestionClient) CreateBulk(builders ...*QuestionCreate) *QuestionCreateBulk {
	return &QuestionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Question.
func (c *QuestionClient) Update() *QuestionUpdate {
	mutation := newQuestionMutation(c.config, OpUpdate)
	return &QuestionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *QuestionClient) UpdateOne(q *Question) *QuestionUpdateOne {
	mutation := newQuestionMutation(c.config, OpUpdateOne, withQuestion(q))
	return &QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *QuestionClient) UpdateOneID(id int) *QuestionUpdateOne {
	mutation := newQuestionMutation(c.config, OpUpdateOne, withQuestionID(id))
	return &QuestionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Question.
func (c *QuestionClient) Delete() *QuestionDelete {
	mutation := newQuestionMutation(c.config, OpDelete)
	return &QuestionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *QuestionClient) DeleteOne(q *Question) *QuestionDeleteOne {
	return c.DeleteOneID(q.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *QuestionClient) DeleteOneID(id int) *QuestionDeleteOne {
	builder := c.Delete().Where(question.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &QuestionDeleteOne{builder}
}

// Query returns a query builder for Question.
func (c *QuestionClient) Query() *QuestionQuery {
	return &QuestionQuery{
		config: c.config,
	}
}

// Get returns a Question entity by its id.
func (c *QuestionClient) Get(ctx context.Context, id int) (*Question, error) {
	return c.Query().Where(question.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *QuestionClient) GetX(ctx context.Context, id int) *Question {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAnswer queries the answer edge of a Question.
func (c *QuestionClient) QueryAnswer(q *Question) *AnswerQuery {
	query := &AnswerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, question.AnswerTable, question.AnswerPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Question.
func (c *QuestionClient) QueryUser(q *Question) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := q.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, question.UserTable, question.UserColumn),
		)
		fromV = sqlgraph.Neighbors(q.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *QuestionClient) Hooks() []Hook {
	return c.hooks.Question
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VoteClient is a client for the Vote schema.
type VoteClient struct {
	config
}

// NewVoteClient returns a client for the Vote from the given config.
func NewVoteClient(c config) *VoteClient {
	return &VoteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vote.Hooks(f(g(h())))`.
func (c *VoteClient) Use(hooks ...Hook) {
	c.hooks.Vote = append(c.hooks.Vote, hooks...)
}

// Create returns a create builder for Vote.
func (c *VoteClient) Create() *VoteCreate {
	mutation := newVoteMutation(c.config, OpCreate)
	return &VoteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vote entities.
func (c *VoteClient) CreateBulk(builders ...*VoteCreate) *VoteCreateBulk {
	return &VoteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vote.
func (c *VoteClient) Update() *VoteUpdate {
	mutation := newVoteMutation(c.config, OpUpdate)
	return &VoteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VoteClient) UpdateOne(v *Vote) *VoteUpdateOne {
	mutation := newVoteMutation(c.config, OpUpdateOne, withVote(v))
	return &VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VoteClient) UpdateOneID(id int) *VoteUpdateOne {
	mutation := newVoteMutation(c.config, OpUpdateOne, withVoteID(id))
	return &VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vote.
func (c *VoteClient) Delete() *VoteDelete {
	mutation := newVoteMutation(c.config, OpDelete)
	return &VoteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VoteClient) DeleteOne(v *Vote) *VoteDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VoteClient) DeleteOneID(id int) *VoteDeleteOne {
	builder := c.Delete().Where(vote.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VoteDeleteOne{builder}
}

// Query returns a query builder for Vote.
func (c *VoteClient) Query() *VoteQuery {
	return &VoteQuery{
		config: c.config,
	}
}

// Get returns a Vote entity by its id.
func (c *VoteClient) Get(ctx context.Context, id int) (*Vote, error) {
	return c.Query().Where(vote.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VoteClient) GetX(ctx context.Context, id int) *Vote {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuestion queries the question edge of a Vote.
func (c *VoteClient) QueryQuestion(v *Vote) *QuestionQuery {
	query := &QuestionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, id),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vote.QuestionTable, vote.QuestionColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAnswer queries the answer edge of a Vote.
func (c *VoteClient) QueryAnswer(v *Vote) *AnswerQuery {
	query := &AnswerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, id),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vote.AnswerTable, vote.AnswerColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Vote.
func (c *VoteClient) QueryUser(v *Vote) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vote.UserTable, vote.UserColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VoteClient) Hooks() []Hook {
	return c.hooks.Vote
}
