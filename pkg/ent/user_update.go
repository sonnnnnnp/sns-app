// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/favorite"
	"github.com/sonnnnnnp/sns-app/pkg/ent/post"
	"github.com/sonnnnnnp/sns-app/pkg/ent/predicate"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of the "nickname" field.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.mutation.ClearNickname()
	return uu
}

// SetAvatarImageURL sets the "avatar_image_url" field.
func (uu *UserUpdate) SetAvatarImageURL(s string) *UserUpdate {
	uu.mutation.SetAvatarImageURL(s)
	return uu
}

// SetNillableAvatarImageURL sets the "avatar_image_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatarImageURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatarImageURL(*s)
	}
	return uu
}

// ClearAvatarImageURL clears the value of the "avatar_image_url" field.
func (uu *UserUpdate) ClearAvatarImageURL() *UserUpdate {
	uu.mutation.ClearAvatarImageURL()
	return uu
}

// SetBannerImageURL sets the "banner_image_url" field.
func (uu *UserUpdate) SetBannerImageURL(s string) *UserUpdate {
	uu.mutation.SetBannerImageURL(s)
	return uu
}

// SetNillableBannerImageURL sets the "banner_image_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBannerImageURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetBannerImageURL(*s)
	}
	return uu
}

// ClearBannerImageURL clears the value of the "banner_image_url" field.
func (uu *UserUpdate) ClearBannerImageURL() *UserUpdate {
	uu.mutation.ClearBannerImageURL()
	return uu
}

// SetBiography sets the "biography" field.
func (uu *UserUpdate) SetBiography(s string) *UserUpdate {
	uu.mutation.SetBiography(s)
	return uu
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBiography(s *string) *UserUpdate {
	if s != nil {
		uu.SetBiography(*s)
	}
	return uu
}

// ClearBiography clears the value of the "biography" field.
func (uu *UserUpdate) ClearBiography() *UserUpdate {
	uu.mutation.ClearBiography()
	return uu
}

// SetBirthdate sets the "birthdate" field.
func (uu *UserUpdate) SetBirthdate(t time.Time) *UserUpdate {
	uu.mutation.SetBirthdate(t)
	return uu
}

// SetNillableBirthdate sets the "birthdate" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBirthdate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetBirthdate(*t)
	}
	return uu
}

// ClearBirthdate clears the value of the "birthdate" field.
func (uu *UserUpdate) ClearBirthdate() *UserUpdate {
	uu.mutation.ClearBirthdate()
	return uu
}

// SetLineID sets the "line_id" field.
func (uu *UserUpdate) SetLineID(s string) *UserUpdate {
	uu.mutation.SetLineID(s)
	return uu
}

// SetNillableLineID sets the "line_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLineID(s *string) *UserUpdate {
	if s != nil {
		uu.SetLineID(*s)
	}
	return uu
}

// ClearLineID clears the value of the "line_id" field.
func (uu *UserUpdate) ClearLineID() *UserUpdate {
	uu.mutation.ClearLineID()
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uu *UserUpdate) AddPostIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddPostIDs(ids...)
	return uu
}

// AddPosts adds the "posts" edges to the Post entity.
func (uu *UserUpdate) AddPosts(p ...*Post) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPostIDs(ids...)
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowerIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddFollowerIDs(ids...)
	return uu
}

// AddFollowers adds the "followers" edges to the User entity.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowingIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddFollowingIDs(ids...)
	return uu
}

// AddFollowing adds the "following" edges to the User entity.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// AddFavoriteIDs adds the "favorites" edge to the Favorite entity by IDs.
func (uu *UserUpdate) AddFavoriteIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFavoriteIDs(ids...)
	return uu
}

// AddFavorites adds the "favorites" edges to the Favorite entity.
func (uu *UserUpdate) AddFavorites(f ...*Favorite) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddFavoriteIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (uu *UserUpdate) ClearPosts() *UserUpdate {
	uu.mutation.ClearPosts()
	return uu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (uu *UserUpdate) RemovePostIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemovePostIDs(ids...)
	return uu
}

// RemovePosts removes "posts" edges to Post entities.
func (uu *UserUpdate) RemovePosts(p ...*Post) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePostIDs(ids...)
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uu *UserUpdate) ClearFollowers() *UserUpdate {
	uu.mutation.ClearFollowers()
	return uu
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveFollowerIDs(ids...)
	return uu
}

// RemoveFollowers removes "followers" edges to User entities.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uu *UserUpdate) ClearFollowing() *UserUpdate {
	uu.mutation.ClearFollowing()
	return uu
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveFollowingIDs(ids...)
	return uu
}

// RemoveFollowing removes "following" edges to User entities.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
}

// ClearFavorites clears all "favorites" edges to the Favorite entity.
func (uu *UserUpdate) ClearFavorites() *UserUpdate {
	uu.mutation.ClearFavorites()
	return uu
}

// RemoveFavoriteIDs removes the "favorites" edge to Favorite entities by IDs.
func (uu *UserUpdate) RemoveFavoriteIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFavoriteIDs(ids...)
	return uu
}

// RemoveFavorites removes "favorites" edges to Favorite entities.
func (uu *UserUpdate) RemoveFavorites(f ...*Favorite) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveFavoriteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uu.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uu.mutation.AvatarImageURL(); ok {
		_spec.SetField(user.FieldAvatarImageURL, field.TypeString, value)
	}
	if uu.mutation.AvatarImageURLCleared() {
		_spec.ClearField(user.FieldAvatarImageURL, field.TypeString)
	}
	if value, ok := uu.mutation.BannerImageURL(); ok {
		_spec.SetField(user.FieldBannerImageURL, field.TypeString, value)
	}
	if uu.mutation.BannerImageURLCleared() {
		_spec.ClearField(user.FieldBannerImageURL, field.TypeString)
	}
	if value, ok := uu.mutation.Biography(); ok {
		_spec.SetField(user.FieldBiography, field.TypeString, value)
	}
	if uu.mutation.BiographyCleared() {
		_spec.ClearField(user.FieldBiography, field.TypeString)
	}
	if value, ok := uu.mutation.Birthdate(); ok {
		_spec.SetField(user.FieldBirthdate, field.TypeTime, value)
	}
	if uu.mutation.BirthdateCleared() {
		_spec.ClearField(user.FieldBirthdate, field.TypeTime)
	}
	if value, ok := uu.mutation.LineID(); ok {
		_spec.SetField(user.FieldLineID, field.TypeString, value)
	}
	if uu.mutation.LineIDCleared() {
		_spec.ClearField(user.FieldLineID, field.TypeString)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !uu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoritesTable,
			Columns: []string{user.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFavoritesIDs(); len(nodes) > 0 && !uu.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoritesTable,
			Columns: []string{user.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoritesTable,
			Columns: []string{user.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of the "nickname" field.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.mutation.ClearNickname()
	return uuo
}

// SetAvatarImageURL sets the "avatar_image_url" field.
func (uuo *UserUpdateOne) SetAvatarImageURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarImageURL(s)
	return uuo
}

// SetNillableAvatarImageURL sets the "avatar_image_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatarImageURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatarImageURL(*s)
	}
	return uuo
}

// ClearAvatarImageURL clears the value of the "avatar_image_url" field.
func (uuo *UserUpdateOne) ClearAvatarImageURL() *UserUpdateOne {
	uuo.mutation.ClearAvatarImageURL()
	return uuo
}

// SetBannerImageURL sets the "banner_image_url" field.
func (uuo *UserUpdateOne) SetBannerImageURL(s string) *UserUpdateOne {
	uuo.mutation.SetBannerImageURL(s)
	return uuo
}

// SetNillableBannerImageURL sets the "banner_image_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBannerImageURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBannerImageURL(*s)
	}
	return uuo
}

// ClearBannerImageURL clears the value of the "banner_image_url" field.
func (uuo *UserUpdateOne) ClearBannerImageURL() *UserUpdateOne {
	uuo.mutation.ClearBannerImageURL()
	return uuo
}

// SetBiography sets the "biography" field.
func (uuo *UserUpdateOne) SetBiography(s string) *UserUpdateOne {
	uuo.mutation.SetBiography(s)
	return uuo
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBiography(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBiography(*s)
	}
	return uuo
}

// ClearBiography clears the value of the "biography" field.
func (uuo *UserUpdateOne) ClearBiography() *UserUpdateOne {
	uuo.mutation.ClearBiography()
	return uuo
}

// SetBirthdate sets the "birthdate" field.
func (uuo *UserUpdateOne) SetBirthdate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirthdate(t)
	return uuo
}

// SetNillableBirthdate sets the "birthdate" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBirthdate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetBirthdate(*t)
	}
	return uuo
}

// ClearBirthdate clears the value of the "birthdate" field.
func (uuo *UserUpdateOne) ClearBirthdate() *UserUpdateOne {
	uuo.mutation.ClearBirthdate()
	return uuo
}

// SetLineID sets the "line_id" field.
func (uuo *UserUpdateOne) SetLineID(s string) *UserUpdateOne {
	uuo.mutation.SetLineID(s)
	return uuo
}

// SetNillableLineID sets the "line_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLineID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLineID(*s)
	}
	return uuo
}

// ClearLineID clears the value of the "line_id" field.
func (uuo *UserUpdateOne) ClearLineID() *UserUpdateOne {
	uuo.mutation.ClearLineID()
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uuo *UserUpdateOne) AddPostIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddPostIDs(ids...)
	return uuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (uuo *UserUpdateOne) AddPosts(p ...*Post) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPostIDs(ids...)
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddFollowerIDs(ids...)
	return uuo
}

// AddFollowers adds the "followers" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddFollowingIDs(ids...)
	return uuo
}

// AddFollowing adds the "following" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// AddFavoriteIDs adds the "favorites" edge to the Favorite entity by IDs.
func (uuo *UserUpdateOne) AddFavoriteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFavoriteIDs(ids...)
	return uuo
}

// AddFavorites adds the "favorites" edges to the Favorite entity.
func (uuo *UserUpdateOne) AddFavorites(f ...*Favorite) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddFavoriteIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (uuo *UserUpdateOne) ClearPosts() *UserUpdateOne {
	uuo.mutation.ClearPosts()
	return uuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (uuo *UserUpdateOne) RemovePostIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemovePostIDs(ids...)
	return uuo
}

// RemovePosts removes "posts" edges to Post entities.
func (uuo *UserUpdateOne) RemovePosts(p ...*Post) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePostIDs(ids...)
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowers() *UserUpdateOne {
	uuo.mutation.ClearFollowers()
	return uuo
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveFollowerIDs(ids...)
	return uuo
}

// RemoveFollowers removes "followers" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowing() *UserUpdateOne {
	uuo.mutation.ClearFollowing()
	return uuo
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveFollowingIDs(ids...)
	return uuo
}

// RemoveFollowing removes "following" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
}

// ClearFavorites clears all "favorites" edges to the Favorite entity.
func (uuo *UserUpdateOne) ClearFavorites() *UserUpdateOne {
	uuo.mutation.ClearFavorites()
	return uuo
}

// RemoveFavoriteIDs removes the "favorites" edge to Favorite entities by IDs.
func (uuo *UserUpdateOne) RemoveFavoriteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFavoriteIDs(ids...)
	return uuo
}

// RemoveFavorites removes "favorites" edges to Favorite entities.
func (uuo *UserUpdateOne) RemoveFavorites(f ...*Favorite) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveFavoriteIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uuo.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uuo.mutation.AvatarImageURL(); ok {
		_spec.SetField(user.FieldAvatarImageURL, field.TypeString, value)
	}
	if uuo.mutation.AvatarImageURLCleared() {
		_spec.ClearField(user.FieldAvatarImageURL, field.TypeString)
	}
	if value, ok := uuo.mutation.BannerImageURL(); ok {
		_spec.SetField(user.FieldBannerImageURL, field.TypeString, value)
	}
	if uuo.mutation.BannerImageURLCleared() {
		_spec.ClearField(user.FieldBannerImageURL, field.TypeString)
	}
	if value, ok := uuo.mutation.Biography(); ok {
		_spec.SetField(user.FieldBiography, field.TypeString, value)
	}
	if uuo.mutation.BiographyCleared() {
		_spec.ClearField(user.FieldBiography, field.TypeString)
	}
	if value, ok := uuo.mutation.Birthdate(); ok {
		_spec.SetField(user.FieldBirthdate, field.TypeTime, value)
	}
	if uuo.mutation.BirthdateCleared() {
		_spec.ClearField(user.FieldBirthdate, field.TypeTime)
	}
	if value, ok := uuo.mutation.LineID(); ok {
		_spec.SetField(user.FieldLineID, field.TypeString, value)
	}
	if uuo.mutation.LineIDCleared() {
		_spec.ClearField(user.FieldLineID, field.TypeString)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !uuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoritesTable,
			Columns: []string{user.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFavoritesIDs(); len(nodes) > 0 && !uuo.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoritesTable,
			Columns: []string{user.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoritesTable,
			Columns: []string{user.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
