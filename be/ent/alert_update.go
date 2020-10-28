// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/echoturing/alert/ent/alert"
	"github.com/echoturing/alert/ent/predicate"
	"github.com/echoturing/alert/ent/schema"
	"github.com/echoturing/alert/ent/schema/sub"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AlertUpdate is the builder for updating Alert entities.
type AlertUpdate struct {
	config
	hooks      []Hook
	mutation   *AlertMutation
	predicates []predicate.Alert
}

// Where adds a new predicate for the builder.
func (au *AlertUpdate) Where(ps ...predicate.Alert) *AlertUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetName sets the name field.
func (au *AlertUpdate) SetName(s string) *AlertUpdate {
	au.mutation.SetName(s)
	return au
}

// SetChannels sets the channels field.
func (au *AlertUpdate) SetChannels(si schema.ChannelIDS) *AlertUpdate {
	au.mutation.SetChannels(si)
	return au
}

// SetRule sets the rule field.
func (au *AlertUpdate) SetRule(s sub.Rule) *AlertUpdate {
	au.mutation.SetRule(s)
	return au
}

// SetStatus sets the status field.
func (au *AlertUpdate) SetStatus(ss schema.AlertStatus) *AlertUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(ss)
	return au
}

// AddStatus adds ss to status.
func (au *AlertUpdate) AddStatus(ss schema.AlertStatus) *AlertUpdate {
	au.mutation.AddStatus(ss)
	return au
}

// SetState sets the state field.
func (au *AlertUpdate) SetState(ss schema.AlertState) *AlertUpdate {
	au.mutation.ResetState()
	au.mutation.SetState(ss)
	return au
}

// AddState adds ss to state.
func (au *AlertUpdate) AddState(ss schema.AlertState) *AlertUpdate {
	au.mutation.AddState(ss)
	return au
}

// SetUpdatedAt sets the updatedAt field.
func (au *AlertUpdate) SetUpdatedAt(t time.Time) *AlertUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the AlertMutation object of the builder.
func (au *AlertUpdate) Mutation() *AlertMutation {
	return au.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AlertUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AlertUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AlertUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AlertUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AlertUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := alert.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AlertUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   alert.Table,
			Columns: alert.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: alert.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldName,
		})
	}
	if value, ok := au.mutation.Channels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldChannels,
		})
	}
	if value, ok := au.mutation.Rule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldRule,
		})
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldStatus,
		})
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldStatus,
		})
	}
	if value, ok := au.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldState,
		})
	}
	if value, ok := au.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldState,
		})
	}
	if au.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: alert.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: alert.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alert.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AlertUpdateOne is the builder for updating a single Alert entity.
type AlertUpdateOne struct {
	config
	hooks    []Hook
	mutation *AlertMutation
}

// SetName sets the name field.
func (auo *AlertUpdateOne) SetName(s string) *AlertUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetChannels sets the channels field.
func (auo *AlertUpdateOne) SetChannels(si schema.ChannelIDS) *AlertUpdateOne {
	auo.mutation.SetChannels(si)
	return auo
}

// SetRule sets the rule field.
func (auo *AlertUpdateOne) SetRule(s sub.Rule) *AlertUpdateOne {
	auo.mutation.SetRule(s)
	return auo
}

// SetStatus sets the status field.
func (auo *AlertUpdateOne) SetStatus(ss schema.AlertStatus) *AlertUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(ss)
	return auo
}

// AddStatus adds ss to status.
func (auo *AlertUpdateOne) AddStatus(ss schema.AlertStatus) *AlertUpdateOne {
	auo.mutation.AddStatus(ss)
	return auo
}

// SetState sets the state field.
func (auo *AlertUpdateOne) SetState(ss schema.AlertState) *AlertUpdateOne {
	auo.mutation.ResetState()
	auo.mutation.SetState(ss)
	return auo
}

// AddState adds ss to state.
func (auo *AlertUpdateOne) AddState(ss schema.AlertState) *AlertUpdateOne {
	auo.mutation.AddState(ss)
	return auo
}

// SetUpdatedAt sets the updatedAt field.
func (auo *AlertUpdateOne) SetUpdatedAt(t time.Time) *AlertUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the AlertMutation object of the builder.
func (auo *AlertUpdateOne) Mutation() *AlertMutation {
	return auo.mutation
}

// Save executes the query and returns the updated entity.
func (auo *AlertUpdateOne) Save(ctx context.Context) (*Alert, error) {
	var (
		err  error
		node *Alert
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AlertUpdateOne) SaveX(ctx context.Context) *Alert {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AlertUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AlertUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AlertUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := alert.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AlertUpdateOne) sqlSave(ctx context.Context) (_node *Alert, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   alert.Table,
			Columns: alert.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: alert.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Alert.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldName,
		})
	}
	if value, ok := auo.mutation.Channels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldChannels,
		})
	}
	if value, ok := auo.mutation.Rule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldRule,
		})
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldStatus,
		})
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldStatus,
		})
	}
	if value, ok := auo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldState,
		})
	}
	if value, ok := auo.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: alert.FieldState,
		})
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: alert.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: alert.FieldUpdatedAt,
		})
	}
	_node = &Alert{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alert.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
