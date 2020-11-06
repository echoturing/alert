// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/echoturing/alert/ent/alerthistory"
	"github.com/echoturing/alert/ent/schema/sub"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AlertHistoryCreate is the builder for creating a AlertHistory entity.
type AlertHistoryCreate struct {
	config
	mutation *AlertHistoryMutation
	hooks    []Hook
}

// SetAlertID sets the alert_id field.
func (ahc *AlertHistoryCreate) SetAlertID(i int64) *AlertHistoryCreate {
	ahc.mutation.SetAlertID(i)
	return ahc
}

// SetAlertName sets the alert_name field.
func (ahc *AlertHistoryCreate) SetAlertName(s string) *AlertHistoryCreate {
	ahc.mutation.SetAlertName(s)
	return ahc
}

// SetDetail sets the detail field.
func (ahc *AlertHistoryCreate) SetDetail(shd sub.AlertHistoryDetail) *AlertHistoryCreate {
	ahc.mutation.SetDetail(shd)
	return ahc
}

// SetCreatedAt sets the createdAt field.
func (ahc *AlertHistoryCreate) SetCreatedAt(t time.Time) *AlertHistoryCreate {
	ahc.mutation.SetCreatedAt(t)
	return ahc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (ahc *AlertHistoryCreate) SetNillableCreatedAt(t *time.Time) *AlertHistoryCreate {
	if t != nil {
		ahc.SetCreatedAt(*t)
	}
	return ahc
}

// SetUpdatedAt sets the updatedAt field.
func (ahc *AlertHistoryCreate) SetUpdatedAt(t time.Time) *AlertHistoryCreate {
	ahc.mutation.SetUpdatedAt(t)
	return ahc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (ahc *AlertHistoryCreate) SetNillableUpdatedAt(t *time.Time) *AlertHistoryCreate {
	if t != nil {
		ahc.SetUpdatedAt(*t)
	}
	return ahc
}

// SetID sets the id field.
func (ahc *AlertHistoryCreate) SetID(i int64) *AlertHistoryCreate {
	ahc.mutation.SetID(i)
	return ahc
}

// Mutation returns the AlertHistoryMutation object of the builder.
func (ahc *AlertHistoryCreate) Mutation() *AlertHistoryMutation {
	return ahc.mutation
}

// Save creates the AlertHistory in the database.
func (ahc *AlertHistoryCreate) Save(ctx context.Context) (*AlertHistory, error) {
	var (
		err  error
		node *AlertHistory
	)
	ahc.defaults()
	if len(ahc.hooks) == 0 {
		if err = ahc.check(); err != nil {
			return nil, err
		}
		node, err = ahc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ahc.check(); err != nil {
				return nil, err
			}
			ahc.mutation = mutation
			node, err = ahc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ahc.hooks) - 1; i >= 0; i-- {
			mut = ahc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ahc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ahc *AlertHistoryCreate) SaveX(ctx context.Context) *AlertHistory {
	v, err := ahc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ahc *AlertHistoryCreate) defaults() {
	if _, ok := ahc.mutation.CreatedAt(); !ok {
		v := alerthistory.DefaultCreatedAt()
		ahc.mutation.SetCreatedAt(v)
	}
	if _, ok := ahc.mutation.UpdatedAt(); !ok {
		v := alerthistory.DefaultUpdatedAt()
		ahc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ahc *AlertHistoryCreate) check() error {
	if _, ok := ahc.mutation.AlertID(); !ok {
		return &ValidationError{Name: "alert_id", err: errors.New("ent: missing required field \"alert_id\"")}
	}
	if _, ok := ahc.mutation.AlertName(); !ok {
		return &ValidationError{Name: "alert_name", err: errors.New("ent: missing required field \"alert_name\"")}
	}
	if _, ok := ahc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New("ent: missing required field \"detail\"")}
	}
	if _, ok := ahc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	return nil
}

func (ahc *AlertHistoryCreate) sqlSave(ctx context.Context) (*AlertHistory, error) {
	_node, _spec := ahc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ahc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (ahc *AlertHistoryCreate) createSpec() (*AlertHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &AlertHistory{config: ahc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: alerthistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: alerthistory.FieldID,
			},
		}
	)
	if id, ok := ahc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ahc.mutation.AlertID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alerthistory.FieldAlertID,
		})
		_node.AlertID = value
	}
	if value, ok := ahc.mutation.AlertName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alerthistory.FieldAlertName,
		})
		_node.AlertName = value
	}
	if value, ok := ahc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alerthistory.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := ahc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: alerthistory.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ahc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: alerthistory.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AlertHistoryCreateBulk is the builder for creating a bulk of AlertHistory entities.
type AlertHistoryCreateBulk struct {
	config
	builders []*AlertHistoryCreate
}

// Save creates the AlertHistory entities in the database.
func (ahcb *AlertHistoryCreateBulk) Save(ctx context.Context) ([]*AlertHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ahcb.builders))
	nodes := make([]*AlertHistory, len(ahcb.builders))
	mutators := make([]Mutator, len(ahcb.builders))
	for i := range ahcb.builders {
		func(i int, root context.Context) {
			builder := ahcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlertHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ahcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ahcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, ahcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ahcb *AlertHistoryCreateBulk) SaveX(ctx context.Context) []*AlertHistory {
	v, err := ahcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}