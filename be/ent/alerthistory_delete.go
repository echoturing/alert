// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/echoturing/alert/ent/alerthistory"
	"github.com/echoturing/alert/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AlertHistoryDelete is the builder for deleting a AlertHistory entity.
type AlertHistoryDelete struct {
	config
	hooks      []Hook
	mutation   *AlertHistoryMutation
	predicates []predicate.AlertHistory
}

// Where adds a new predicate to the delete builder.
func (ahd *AlertHistoryDelete) Where(ps ...predicate.AlertHistory) *AlertHistoryDelete {
	ahd.predicates = append(ahd.predicates, ps...)
	return ahd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ahd *AlertHistoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ahd.hooks) == 0 {
		affected, err = ahd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ahd.mutation = mutation
			affected, err = ahd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ahd.hooks) - 1; i >= 0; i-- {
			mut = ahd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ahd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahd *AlertHistoryDelete) ExecX(ctx context.Context) int {
	n, err := ahd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ahd *AlertHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: alerthistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: alerthistory.FieldID,
			},
		},
	}
	if ps := ahd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ahd.driver, _spec)
}

// AlertHistoryDeleteOne is the builder for deleting a single AlertHistory entity.
type AlertHistoryDeleteOne struct {
	ahd *AlertHistoryDelete
}

// Exec executes the deletion query.
func (ahdo *AlertHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ahdo.ahd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{alerthistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ahdo *AlertHistoryDeleteOne) ExecX(ctx context.Context) {
	ahdo.ahd.ExecX(ctx)
}
