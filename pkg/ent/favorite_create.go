// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/favorite"
	"github.com/sonnnnnnp/sns-app/pkg/ent/post"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

// FavoriteCreate is the builder for creating a Favorite entity.
type FavoriteCreate struct {
	config
	mutation *FavoriteMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (fc *FavoriteCreate) SetUserID(u uuid.UUID) *FavoriteCreate {
	fc.mutation.SetUserID(u)
	return fc
}

// SetPostID sets the "post_id" field.
func (fc *FavoriteCreate) SetPostID(u uuid.UUID) *FavoriteCreate {
	fc.mutation.SetPostID(u)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FavoriteCreate) SetCreatedAt(t time.Time) *FavoriteCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FavoriteCreate) SetNillableCreatedAt(t *time.Time) *FavoriteCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FavoriteCreate) SetUser(u *User) *FavoriteCreate {
	return fc.SetUserID(u.ID)
}

// SetPost sets the "post" edge to the Post entity.
func (fc *FavoriteCreate) SetPost(p *Post) *FavoriteCreate {
	return fc.SetPostID(p.ID)
}

// Mutation returns the FavoriteMutation object of the builder.
func (fc *FavoriteCreate) Mutation() *FavoriteMutation {
	return fc.mutation
}

// Save creates the Favorite in the database.
func (fc *FavoriteCreate) Save(ctx context.Context) (*Favorite, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FavoriteCreate) SaveX(ctx context.Context) *Favorite {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FavoriteCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FavoriteCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FavoriteCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := favorite.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FavoriteCreate) check() error {
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Favorite.user_id"`)}
	}
	if _, ok := fc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Favorite.post_id"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Favorite.created_at"`)}
	}
	if len(fc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Favorite.user"`)}
	}
	if len(fc.mutation.PostIDs()) == 0 {
		return &ValidationError{Name: "post", err: errors.New(`ent: missing required edge "Favorite.post"`)}
	}
	return nil
}

func (fc *FavoriteCreate) sqlSave(ctx context.Context) (*Favorite, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FavoriteCreate) createSpec() (*Favorite, *sqlgraph.CreateSpec) {
	var (
		_node = &Favorite{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(favorite.Table, sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt))
	)
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(favorite.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.UserTable,
			Columns: []string{favorite.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.PostTable,
			Columns: []string{favorite.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FavoriteCreateBulk is the builder for creating many Favorite entities in bulk.
type FavoriteCreateBulk struct {
	config
	err      error
	builders []*FavoriteCreate
}

// Save creates the Favorite entities in the database.
func (fcb *FavoriteCreateBulk) Save(ctx context.Context) ([]*Favorite, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Favorite, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FavoriteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FavoriteCreateBulk) SaveX(ctx context.Context) []*Favorite {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FavoriteCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FavoriteCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
