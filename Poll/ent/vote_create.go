// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll/ent/answer"
	"poll/ent/question"
	"poll/ent/user"
	"poll/ent/vote"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoteCreate is the builder for creating a Vote entity.
type VoteCreate struct {
	config
	mutation *VoteMutation
	hooks    []Hook
}

// SetQuestionID sets the "question_id" field.
func (vc *VoteCreate) SetQuestionID(i int) *VoteCreate {
	vc.mutation.SetQuestionID(i)
	return vc
}

// SetAnswerID sets the "answer_id" field.
func (vc *VoteCreate) SetAnswerID(i int) *VoteCreate {
	vc.mutation.SetAnswerID(i)
	return vc
}

// SetUserID sets the "user_id" field.
func (vc *VoteCreate) SetUserID(i int) *VoteCreate {
	vc.mutation.SetUserID(i)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VoteCreate) SetCreatedAt(t time.Time) *VoteCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VoteCreate) SetNillableCreatedAt(t *time.Time) *VoteCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VoteCreate) SetUpdatedAt(t time.Time) *VoteCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VoteCreate) SetNillableUpdatedAt(t *time.Time) *VoteCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VoteCreate) SetID(i int) *VoteCreate {
	vc.mutation.SetID(i)
	return vc
}

// SetQuestion sets the "question" edge to the Question entity.
func (vc *VoteCreate) SetQuestion(q *Question) *VoteCreate {
	return vc.SetQuestionID(q.ID)
}

// SetAnswer sets the "answer" edge to the Answer entity.
func (vc *VoteCreate) SetAnswer(a *Answer) *VoteCreate {
	return vc.SetAnswerID(a.ID)
}

// SetUser sets the "user" edge to the User entity.
func (vc *VoteCreate) SetUser(u *User) *VoteCreate {
	return vc.SetUserID(u.ID)
}

// Mutation returns the VoteMutation object of the builder.
func (vc *VoteCreate) Mutation() *VoteMutation {
	return vc.mutation
}

// Save creates the Vote in the database.
func (vc *VoteCreate) Save(ctx context.Context) (*Vote, error) {
	var (
		err  error
		node *Vote
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VoteCreate) SaveX(ctx context.Context) *Vote {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VoteCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VoteCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VoteCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := vote.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := vote.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VoteCreate) check() error {
	if _, ok := vc.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question_id", err: errors.New(`ent: missing required field "question_id"`)}
	}
	if _, ok := vc.mutation.AnswerID(); !ok {
		return &ValidationError{Name: "answer_id", err: errors.New(`ent: missing required field "answer_id"`)}
	}
	if _, ok := vc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := vc.mutation.QuestionID(); !ok {
		return &ValidationError{Name: "question", err: errors.New("ent: missing required edge \"question\"")}
	}
	if _, ok := vc.mutation.AnswerID(); !ok {
		return &ValidationError{Name: "answer", err: errors.New("ent: missing required edge \"answer\"")}
	}
	if _, ok := vc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (vc *VoteCreate) sqlSave(ctx context.Context) (*Vote, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (vc *VoteCreate) createSpec() (*Vote, *sqlgraph.CreateSpec) {
	var (
		_node = &Vote{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vote.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vote.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vote.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := vc.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vote.QuestionTable,
			Columns: []string{vote.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.QuestionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.AnswerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vote.AnswerTable,
			Columns: []string{vote.AnswerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AnswerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VoteCreateBulk is the builder for creating many Vote entities in bulk.
type VoteCreateBulk struct {
	config
	builders []*VoteCreate
}

// Save creates the Vote entities in the database.
func (vcb *VoteCreateBulk) Save(ctx context.Context) ([]*Vote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vote, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VoteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VoteCreateBulk) SaveX(ctx context.Context) []*Vote {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VoteCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VoteCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
