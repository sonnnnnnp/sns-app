package seeder

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (s *Seeder) seedUsers() error {
	users := []db.UpdateUserParams{
		{
			UserID:         uuid.Nil,
			Name:           ptr("osaka_oba"),
			Nickname:       ptr("大阪のおばちゃん"),
			Biography:      ptr("大阪のおばちゃんです。電話好きです"),
			AvatarImageUrl: ptr("https://i.pinimg.com/474x/78/84/6b/78846bd3997e4a15c93fb99325ff4665.jpg"),
			BannerImageUrl: ptr("https://i.pinimg.com/474x/66/6e/a3/666ea384c69d89b9d291c7afb4dc2d40.jpg"),
		},
		{
			UserID:         uuid.Nil,
			Name:           ptr("noahhhhhh"),
			Nickname:       ptr("のあ"),
			Biography:      ptr("三重/社会人/20の代\n仲良くなったらフォロー返す💁🏼‍♀️\nだる絡みはすぐ消すかも"),
			AvatarImageUrl: ptr("https://i.pinimg.com/474x/85/e6/47/85e64767b129d2cae2d1c47b1ed0aece.jpg"),
			BannerImageUrl: ptr("https://i.pinimg.com/474x/6c/22/d4/6c22d47f9c2f2f3e3830d115e1e23d9e.jpg"),
		},
		{
			UserID:         uuid.Nil,
			Name:           ptr("kkeapii_377642"),
			Nickname:       ptr("KK"),
			Biography:      ptr("監視中"),
			AvatarImageUrl: ptr("https://i.pinimg.com/474x/9a/4c/52/9a4c52067795bb1e6bedb3a640d018fb.jpg"),
			BannerImageUrl: ptr("https://i.pinimg.com/474x/7f/24/e8/7f24e8bda2436b3545370801fb996556.jpg"),
		},
		{
			UserID:         uuid.Nil,
			Name:           ptr("sssakaski8017"),
			Nickname:       ptr("佐々木"),
			Biography:      ptr("てきとうに絡んでください"),
			AvatarImageUrl: ptr("https://i.pinimg.com/474x/fe/6b/11/fe6b1199b1644c1552fa13f08a02ea50.jpg"),
			BannerImageUrl: ptr("https://i.pinimg.com/474x/3b/12/c1/3b12c1963c568493393e9c56053769dd.jpg"),
		},
		{
			UserID:         uuid.Nil,
			Name:           ptr("no_depl"),
			Nickname:       ptr("ひまちゃん"),
			Biography:      ptr("22"),
			AvatarImageUrl: ptr("https://i.pinimg.com/474x/13/a8/70/13a87082505ddbe8004517892b203cdb.jpg"),
			BannerImageUrl: ptr("https://i.pinimg.com/474x/e5/a6/32/e5a632ed8c4d9071aea99290f8ff3f8c.jpg"),
		},
	}

	for i := range users {
		uID, err := s.queries.CreateUser(context.Background(), nil)
		if err != nil {
			return err
		}

		if err := s.queries.UpdateUser(context.Background(), db.UpdateUserParams{
			UserID:         uID,
			Name:           users[i].Name,
			Nickname:       users[i].Nickname,
			Biography:      users[i].Biography,
			AvatarImageUrl: users[i].AvatarImageUrl,
			BannerImageUrl: users[i].BannerImageUrl,
		}); err != nil {
			return err
		}

		text := "Hello, everyone! I'm using SNS-App to share my thoughts and experiences."
		s.queries.CreatePost(context.Background(), db.CreatePostParams{
			AuthorID: uID,
			Text:     &text,
		})
	}

	return nil
}
