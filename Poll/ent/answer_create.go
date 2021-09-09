// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll/ent/answer"
	"poll/ent/question"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AnswerCreate is the builder for creating a Answer entity.
type AnswerCreate struct {
	config
	mutation *AnswerMutation
	hooks    []Hook
}

// SetBody sets the "body" field.
func (ac *AnswerCreate) SetBody(s string) *AnswerCreate {
	ac.mutation.SetBody(s)
	return ac
}

// SetNumOfVotes sets the "numOfVotes" field.
func (ac *AnswerCreate) SetNumOfVotes(i int) *AnswerCreate {
	ac.mutation.SetNumOfVotes(i)
	return ac
}

// SetNillableNumOfVotes sets the "numOfVotes" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableNumOfVotes(i *int) *AnswerCreate {
	if i != nil {
		ac.SetNumOfVotes(*i)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AnswerCreate) SetCreatedAt(t time.Time) *AnswerCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableCreatedAt(t *time.Time) *AnswerCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AnswerCreate) SetUpdatedAt(t time.Time) *AnswerCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AnswerCreate) SetNillableUpdatedAt(t *time.Time) *AnswerCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AnswerCreate) SetID(i int) *AnswerCreate {
	ac.mutation.SetID(i)
	return ac
}

// AddQuestionIDIDs adds the "question_id" edge to the Question entity by IDs.
func (ac *AnswerCreate) AddQuestionIDIDs(ids ...int) *AnswerCreate {
	ac.mutation.AddQuestionIDIDs(ids...)
	return ac
}

// AddQuestionID adds the "question_id" edges to the Question entity.
func (ac *AnswerCreate) AddQuestionID(q ...*Question) *AnswerCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return ac.AddQuestionIDIDs(ids...)
}

// Mutation returns the AnswerMutation object of the builder.
func (ac *AnswerCreate) Mutation() *AnswerMutation {
	return ac.mutation
}

// Save creates the Answer in the database.
func (ac *AnswerCreate) Save(ctx context.Context) (*Answer, error) {
	var (
		err  error
		node *Answer
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnswerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnswerCreate) SaveX(ctx context.Context) *Answer {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnswerCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnswerCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnswerCreate) defaults() {
	if _, ok := ac.mutation.NumOfVotes(); !ok {
		v := answer.DefaultNumOfVotes
		ac.mutation.SetNumOfVotes(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := answer.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := answer.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnswerCreate) check() error {
	if _, ok := ac.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "body"`)}
	}
	if _, ok := ac.mutation.NumOfVotes(); !ok {
		return &ValidationError{Name: "numOfVotes", err: errors.New(`ent: missing required field "numOfVotes"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (ac *AnswerCreate) sqlSave(ctx context.Context) (*Answer, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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

func (ac *AnswerCreate) createSpec() (*Answer, *sqlgraph.CreateSpec) {
	var (
		_node = &Answer{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: answer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: answer.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: answer.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := ac.mutation.NumOfVotes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: answer.FieldNumOfVotes,
		})
		_node.NumOfVotes = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: answer.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: answer.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.QuestionIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   answer.QuestionIDTable,
			Columns: answer.QuestionIDPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AnswerCreateBulk is the builder for creating many Answer entities in bulk.
type AnswerCreateBulk struct {
	config
	builders []*AnswerCreate
}

// Save creates the Answer entities in the database.
func (acb *AnswerCreateBulk) Save(ctx context.Context) ([]*Answer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Answer, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnswerMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnswerCreateBulk) SaveX(ctx context.Context) []*Answer {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnswerCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnswerCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
