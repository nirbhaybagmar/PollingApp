// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"poll/ent/answer"
	"poll/ent/predicate"
	"poll/ent/question"
	"poll/ent/user"
	"poll/ent/vote"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoteQuery is the builder for querying Vote entities.
type VoteQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Vote
	// eager-loading edges.
	withQuestion *QuestionQuery
	withAnswer   *AnswerQuery
	withUser     *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoteQuery builder.
func (vq *VoteQuery) Where(ps ...predicate.Vote) *VoteQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VoteQuery) Limit(limit int) *VoteQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VoteQuery) Offset(offset int) *VoteQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VoteQuery) Unique(unique bool) *VoteQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *VoteQuery) Order(o ...OrderFunc) *VoteQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryQuestion chains the current query on the "question" edge.
func (vq *VoteQuery) QueryQuestion() *QuestionQuery {
	query := &QuestionQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, selector),
			sqlgraph.To(question.Table, question.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vote.QuestionTable, vote.QuestionColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAnswer chains the current query on the "answer" edge.
func (vq *VoteQuery) QueryAnswer() *AnswerQuery {
	query := &AnswerQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, selector),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vote.AnswerTable, vote.AnswerColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (vq *VoteQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, vote.UserTable, vote.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vote entity from the query.
// Returns a *NotFoundError when no Vote was found.
func (vq *VoteQuery) First(ctx context.Context) (*Vote, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VoteQuery) FirstX(ctx context.Context) *Vote {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vote ID from the query.
// Returns a *NotFoundError when no Vote ID was found.
func (vq *VoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VoteQuery) FirstIDX(ctx context.Context) int {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Vote entity is not found.
// Returns a *NotFoundError when no Vote entities are found.
func (vq *VoteQuery) Only(ctx context.Context) (*Vote, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vote.Label}
	default:
		return nil, &NotSingularError{vote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VoteQuery) OnlyX(ctx context.Context) *Vote {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vote ID in the query.
// Returns a *NotSingularError when exactly one Vote ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vq *VoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = &NotSingularError{vote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Votes.
func (vq *VoteQuery) All(ctx context.Context) ([]*Vote, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VoteQuery) AllX(ctx context.Context) []*Vote {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vote IDs.
func (vq *VoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := vq.Select(vote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VoteQuery) IDsX(ctx context.Context) []int {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VoteQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VoteQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VoteQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VoteQuery) Clone() *VoteQuery {
	if vq == nil {
		return nil
	}
	return &VoteQuery{
		config:       vq.config,
		limit:        vq.limit,
		offset:       vq.offset,
		order:        append([]OrderFunc{}, vq.order...),
		predicates:   append([]predicate.Vote{}, vq.predicates...),
		withQuestion: vq.withQuestion.Clone(),
		withAnswer:   vq.withAnswer.Clone(),
		withUser:     vq.withUser.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithQuestion tells the query-builder to eager-load the nodes that are connected to
// the "question" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoteQuery) WithQuestion(opts ...func(*QuestionQuery)) *VoteQuery {
	query := &QuestionQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withQuestion = query
	return vq
}

// WithAnswer tells the query-builder to eager-load the nodes that are connected to
// the "answer" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoteQuery) WithAnswer(opts ...func(*AnswerQuery)) *VoteQuery {
	query := &AnswerQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withAnswer = query
	return vq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoteQuery) WithUser(opts ...func(*UserQuery)) *VoteQuery {
	query := &UserQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withUser = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		QuestionID int `json:"question_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vote.Query().
//		GroupBy(vote.FieldQuestionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VoteQuery) GroupBy(field string, fields ...string) *VoteGroupBy {
	group := &VoteGroupBy{config: vq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		QuestionID int `json:"question_id,omitempty"`
//	}
//
//	client.Vote.Query().
//		Select(vote.FieldQuestionID).
//		Scan(ctx, &v)
//
func (vq *VoteQuery) Select(fields ...string) *VoteSelect {
	vq.fields = append(vq.fields, fields...)
	return &VoteSelect{VoteQuery: vq}
}

func (vq *VoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !vote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VoteQuery) sqlAll(ctx context.Context) ([]*Vote, error) {
	var (
		nodes       = []*Vote{}
		_spec       = vq.querySpec()
		loadedTypes = [3]bool{
			vq.withQuestion != nil,
			vq.withAnswer != nil,
			vq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Vote{config: vq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vq.withQuestion; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Vote)
		for i := range nodes {
			fk := nodes[i].QuestionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(question.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "question_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Question = n
			}
		}
	}

	if query := vq.withAnswer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Vote)
		for i := range nodes {
			fk := nodes[i].AnswerID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(answer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "answer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Answer = n
			}
		}
	}

	if query := vq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Vote)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (vq *VoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VoteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vq *VoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vote.Table,
			Columns: vote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vote.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vote.FieldID)
		for i := range fields {
			if fields[i] != vote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vote.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = vote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VoteGroupBy is the group-by builder for Vote entities.
type VoteGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VoteGroupBy) Aggregate(fns ...AggregateFunc) *VoteGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VoteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vgb *VoteGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoteGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vgb *VoteGroupBy) StringsX(ctx context.Context) []string {
	v, err := vgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vgb *VoteGroupBy) StringX(ctx context.Context) string {
	v, err := vgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoteGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vgb *VoteGroupBy) IntsX(ctx context.Context) []int {
	v, err := vgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vgb *VoteGroupBy) IntX(ctx context.Context) int {
	v, err := vgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoteGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vgb *VoteGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vgb *VoteGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoteGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vgb *VoteGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoteGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vgb *VoteGroupBy) BoolX(ctx context.Context) bool {
	v, err := vgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vgb *VoteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vgb.fields {
		if !vote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VoteGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}

// VoteSelect is the builder for selecting fields of Vote entities.
type VoteSelect struct {
	*VoteQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VoteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VoteQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vs *VoteSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoteSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vs *VoteSelect) StringsX(ctx context.Context) []string {
	v, err := vs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vs *VoteSelect) StringX(ctx context.Context) string {
	v, err := vs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoteSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vs *VoteSelect) IntsX(ctx context.Context) []int {
	v, err := vs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vs *VoteSelect) IntX(ctx context.Context) int {
	v, err := vs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoteSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vs *VoteSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vs *VoteSelect) Float64X(ctx context.Context) float64 {
	v, err := vs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoteSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vs *VoteSelect) BoolsX(ctx context.Context) []bool {
	v, err := vs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vs *VoteSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = fmt.Errorf("ent: VoteSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vs *VoteSelect) BoolX(ctx context.Context) bool {
	v, err := vs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vs *VoteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}