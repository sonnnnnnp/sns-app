package seeder

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/db"
)

type User struct {
	Name      string
	Nickname  string
	Biography string
}

func (s *Seeder) seedUsers() error {
	users := make([]User, 5)

	users[0] = User{
		Name:      "john_doe",
		Nickname:  "John Doe",
		Biography: "I'm a software engineer",
	}
	users[1] = User{
		Name:      "jane_smith",
		Nickname:  "Jane Smith",
		Biography: "I'm a graphic designer",
	}
	users[2] = User{
		Name:      "alice_j",
		Nickname:  "Alice Johnson",
		Biography: "I'm a product manager",
	}
	users[3] = User{
		Name:      "bob_b",
		Nickname:  "Bob Brown",
		Biography: "I'm a data scientist",
	}
	users[4] = User{
		Name:      "charlie_d",
		Nickname:  "Charlie Davis",
		Biography: "I'm a marketing specialist",
	}

	for i := range users {
		u, err := s.queries.CreateUser(context.Background(), nil)
		if err != nil {
			return err
		}

		if err := s.queries.UpdateUser(context.Background(), db.UpdateUserParams{
			UserID:    u.ID,
			Name:      &users[i].Name,
			Nickname:  &users[i].Nickname,
			Biography: &users[i].Biography,
		}); err != nil {
			return err
		}

		text := "Hello, everyone! I'm using SNS-App to share my thoughts and experiences."
		s.queries.CreatePost(context.Background(), db.CreatePostParams{
			AuthorID: u.ID,
			Text:     &text,
		})
	}

	return nil
}
