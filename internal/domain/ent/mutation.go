// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent/predicate"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	username      *string
	display_name  *string
	avatar_url    *string
	cover_url     *string
	biography     *string
	birthdate     *time.Time
	line_id       *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetDisplayName sets the "display_name" field.
func (m *UserMutation) SetDisplayName(s string) {
	m.display_name = &s
}

// DisplayName returns the value of the "display_name" field in the mutation.
func (m *UserMutation) DisplayName() (r string, exists bool) {
	v := m.display_name
	if v == nil {
		return
	}
	return *v, true
}

// OldDisplayName returns the old "display_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDisplayName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDisplayName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDisplayName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDisplayName: %w", err)
	}
	return oldValue.DisplayName, nil
}

// ClearDisplayName clears the value of the "display_name" field.
func (m *UserMutation) ClearDisplayName() {
	m.display_name = nil
	m.clearedFields[user.FieldDisplayName] = struct{}{}
}

// DisplayNameCleared returns if the "display_name" field was cleared in this mutation.
func (m *UserMutation) DisplayNameCleared() bool {
	_, ok := m.clearedFields[user.FieldDisplayName]
	return ok
}

// ResetDisplayName resets all changes to the "display_name" field.
func (m *UserMutation) ResetDisplayName() {
	m.display_name = nil
	delete(m.clearedFields, user.FieldDisplayName)
}

// SetAvatarURL sets the "avatar_url" field.
func (m *UserMutation) SetAvatarURL(s string) {
	m.avatar_url = &s
}

// AvatarURL returns the value of the "avatar_url" field in the mutation.
func (m *UserMutation) AvatarURL() (r string, exists bool) {
	v := m.avatar_url
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatarURL returns the old "avatar_url" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAvatarURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatarURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatarURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatarURL: %w", err)
	}
	return oldValue.AvatarURL, nil
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (m *UserMutation) ClearAvatarURL() {
	m.avatar_url = nil
	m.clearedFields[user.FieldAvatarURL] = struct{}{}
}

// AvatarURLCleared returns if the "avatar_url" field was cleared in this mutation.
func (m *UserMutation) AvatarURLCleared() bool {
	_, ok := m.clearedFields[user.FieldAvatarURL]
	return ok
}

// ResetAvatarURL resets all changes to the "avatar_url" field.
func (m *UserMutation) ResetAvatarURL() {
	m.avatar_url = nil
	delete(m.clearedFields, user.FieldAvatarURL)
}

// SetCoverURL sets the "cover_url" field.
func (m *UserMutation) SetCoverURL(s string) {
	m.cover_url = &s
}

// CoverURL returns the value of the "cover_url" field in the mutation.
func (m *UserMutation) CoverURL() (r string, exists bool) {
	v := m.cover_url
	if v == nil {
		return
	}
	return *v, true
}

// OldCoverURL returns the old "cover_url" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCoverURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCoverURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCoverURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCoverURL: %w", err)
	}
	return oldValue.CoverURL, nil
}

// ClearCoverURL clears the value of the "cover_url" field.
func (m *UserMutation) ClearCoverURL() {
	m.cover_url = nil
	m.clearedFields[user.FieldCoverURL] = struct{}{}
}

// CoverURLCleared returns if the "cover_url" field was cleared in this mutation.
func (m *UserMutation) CoverURLCleared() bool {
	_, ok := m.clearedFields[user.FieldCoverURL]
	return ok
}

// ResetCoverURL resets all changes to the "cover_url" field.
func (m *UserMutation) ResetCoverURL() {
	m.cover_url = nil
	delete(m.clearedFields, user.FieldCoverURL)
}

// SetBiography sets the "biography" field.
func (m *UserMutation) SetBiography(s string) {
	m.biography = &s
}

// Biography returns the value of the "biography" field in the mutation.
func (m *UserMutation) Biography() (r string, exists bool) {
	v := m.biography
	if v == nil {
		return
	}
	return *v, true
}

// OldBiography returns the old "biography" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldBiography(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBiography is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBiography requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBiography: %w", err)
	}
	return oldValue.Biography, nil
}

// ClearBiography clears the value of the "biography" field.
func (m *UserMutation) ClearBiography() {
	m.biography = nil
	m.clearedFields[user.FieldBiography] = struct{}{}
}

// BiographyCleared returns if the "biography" field was cleared in this mutation.
func (m *UserMutation) BiographyCleared() bool {
	_, ok := m.clearedFields[user.FieldBiography]
	return ok
}

// ResetBiography resets all changes to the "biography" field.
func (m *UserMutation) ResetBiography() {
	m.biography = nil
	delete(m.clearedFields, user.FieldBiography)
}

// SetBirthdate sets the "birthdate" field.
func (m *UserMutation) SetBirthdate(t time.Time) {
	m.birthdate = &t
}

// Birthdate returns the value of the "birthdate" field in the mutation.
func (m *UserMutation) Birthdate() (r time.Time, exists bool) {
	v := m.birthdate
	if v == nil {
		return
	}
	return *v, true
}

// OldBirthdate returns the old "birthdate" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldBirthdate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBirthdate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBirthdate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBirthdate: %w", err)
	}
	return oldValue.Birthdate, nil
}

// ClearBirthdate clears the value of the "birthdate" field.
func (m *UserMutation) ClearBirthdate() {
	m.birthdate = nil
	m.clearedFields[user.FieldBirthdate] = struct{}{}
}

// BirthdateCleared returns if the "birthdate" field was cleared in this mutation.
func (m *UserMutation) BirthdateCleared() bool {
	_, ok := m.clearedFields[user.FieldBirthdate]
	return ok
}

// ResetBirthdate resets all changes to the "birthdate" field.
func (m *UserMutation) ResetBirthdate() {
	m.birthdate = nil
	delete(m.clearedFields, user.FieldBirthdate)
}

// SetLineID sets the "line_id" field.
func (m *UserMutation) SetLineID(s string) {
	m.line_id = &s
}

// LineID returns the value of the "line_id" field in the mutation.
func (m *UserMutation) LineID() (r string, exists bool) {
	v := m.line_id
	if v == nil {
		return
	}
	return *v, true
}

// OldLineID returns the old "line_id" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldLineID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLineID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLineID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLineID: %w", err)
	}
	return oldValue.LineID, nil
}

// ClearLineID clears the value of the "line_id" field.
func (m *UserMutation) ClearLineID() {
	m.line_id = nil
	m.clearedFields[user.FieldLineID] = struct{}{}
}

// LineIDCleared returns if the "line_id" field was cleared in this mutation.
func (m *UserMutation) LineIDCleared() bool {
	_, ok := m.clearedFields[user.FieldLineID]
	return ok
}

// ResetLineID resets all changes to the "line_id" field.
func (m *UserMutation) ResetLineID() {
	m.line_id = nil
	delete(m.clearedFields, user.FieldLineID)
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.display_name != nil {
		fields = append(fields, user.FieldDisplayName)
	}
	if m.avatar_url != nil {
		fields = append(fields, user.FieldAvatarURL)
	}
	if m.cover_url != nil {
		fields = append(fields, user.FieldCoverURL)
	}
	if m.biography != nil {
		fields = append(fields, user.FieldBiography)
	}
	if m.birthdate != nil {
		fields = append(fields, user.FieldBirthdate)
	}
	if m.line_id != nil {
		fields = append(fields, user.FieldLineID)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	case user.FieldDisplayName:
		return m.DisplayName()
	case user.FieldAvatarURL:
		return m.AvatarURL()
	case user.FieldCoverURL:
		return m.CoverURL()
	case user.FieldBiography:
		return m.Biography()
	case user.FieldBirthdate:
		return m.Birthdate()
	case user.FieldLineID:
		return m.LineID()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldDisplayName:
		return m.OldDisplayName(ctx)
	case user.FieldAvatarURL:
		return m.OldAvatarURL(ctx)
	case user.FieldCoverURL:
		return m.OldCoverURL(ctx)
	case user.FieldBiography:
		return m.OldBiography(ctx)
	case user.FieldBirthdate:
		return m.OldBirthdate(ctx)
	case user.FieldLineID:
		return m.OldLineID(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldDisplayName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDisplayName(v)
		return nil
	case user.FieldAvatarURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatarURL(v)
		return nil
	case user.FieldCoverURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCoverURL(v)
		return nil
	case user.FieldBiography:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBiography(v)
		return nil
	case user.FieldBirthdate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBirthdate(v)
		return nil
	case user.FieldLineID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLineID(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldDisplayName) {
		fields = append(fields, user.FieldDisplayName)
	}
	if m.FieldCleared(user.FieldAvatarURL) {
		fields = append(fields, user.FieldAvatarURL)
	}
	if m.FieldCleared(user.FieldCoverURL) {
		fields = append(fields, user.FieldCoverURL)
	}
	if m.FieldCleared(user.FieldBiography) {
		fields = append(fields, user.FieldBiography)
	}
	if m.FieldCleared(user.FieldBirthdate) {
		fields = append(fields, user.FieldBirthdate)
	}
	if m.FieldCleared(user.FieldLineID) {
		fields = append(fields, user.FieldLineID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldDisplayName:
		m.ClearDisplayName()
		return nil
	case user.FieldAvatarURL:
		m.ClearAvatarURL()
		return nil
	case user.FieldCoverURL:
		m.ClearCoverURL()
		return nil
	case user.FieldBiography:
		m.ClearBiography()
		return nil
	case user.FieldBirthdate:
		m.ClearBirthdate()
		return nil
	case user.FieldLineID:
		m.ClearLineID()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldDisplayName:
		m.ResetDisplayName()
		return nil
	case user.FieldAvatarURL:
		m.ResetAvatarURL()
		return nil
	case user.FieldCoverURL:
		m.ResetCoverURL()
		return nil
	case user.FieldBiography:
		m.ResetBiography()
		return nil
	case user.FieldBirthdate:
		m.ResetBirthdate()
		return nil
	case user.FieldLineID:
		m.ResetLineID()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
