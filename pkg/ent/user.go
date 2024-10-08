// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// AvatarImageURL holds the value of the "avatar_image_url" field.
	AvatarImageURL string `json:"avatar_image_url,omitempty"`
	// BannerImageURL holds the value of the "banner_image_url" field.
	BannerImageURL string `json:"banner_image_url,omitempty"`
	// Biography holds the value of the "biography" field.
	Biography string `json:"biography,omitempty"`
	// Birthdate holds the value of the "birthdate" field.
	Birthdate time.Time `json:"birthdate,omitempty"`
	// LineID holds the value of the "line_id" field.
	LineID string `json:"line_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// Favorites holds the value of the favorites edge.
	Favorites []*Favorite `json:"favorites,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// FavoritesOrErr returns the Favorites value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FavoritesOrErr() ([]*Favorite, error) {
	if e.loadedTypes[3] {
		return e.Favorites, nil
	}
	return nil, &NotLoadedError{edge: "favorites"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldName, user.FieldNickname, user.FieldAvatarImageURL, user.FieldBannerImageURL, user.FieldBiography, user.FieldLineID:
			values[i] = new(sql.NullString)
		case user.FieldBirthdate, user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldAvatarImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_image_url", values[i])
			} else if value.Valid {
				u.AvatarImageURL = value.String
			}
		case user.FieldBannerImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_image_url", values[i])
			} else if value.Valid {
				u.BannerImageURL = value.String
			}
		case user.FieldBiography:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biography", values[i])
			} else if value.Valid {
				u.Biography = value.String
			}
		case user.FieldBirthdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthdate", values[i])
			} else if value.Valid {
				u.Birthdate = value.Time
			}
		case user.FieldLineID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field line_id", values[i])
			} else if value.Valid {
				u.LineID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return NewUserClient(u.config).QueryPosts(u)
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return NewUserClient(u.config).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return NewUserClient(u.config).QueryFollowing(u)
}

// QueryFavorites queries the "favorites" edge of the User entity.
func (u *User) QueryFavorites() *FavoriteQuery {
	return NewUserClient(u.config).QueryFavorites(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("avatar_image_url=")
	builder.WriteString(u.AvatarImageURL)
	builder.WriteString(", ")
	builder.WriteString("banner_image_url=")
	builder.WriteString(u.BannerImageURL)
	builder.WriteString(", ")
	builder.WriteString("biography=")
	builder.WriteString(u.Biography)
	builder.WriteString(", ")
	builder.WriteString("birthdate=")
	builder.WriteString(u.Birthdate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("line_id=")
	builder.WriteString(u.LineID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
