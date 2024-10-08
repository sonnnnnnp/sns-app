package main

import (
	"context"
	"log"

	"github.com/sonnnnnnp/sns-app/internal/adapter/gateway/db"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	db, err := db.Open(&db.ConnectionOptions{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Name:     cfg.DBName,
		Password: cfg.DBPassword,
		SSLMode:  cfg.DBSSLMode,
	})
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer db.Close()

	if err := seed(context.Background(), db); err != nil {
		log.Fatalf("failed seeding the database: %v", err)
	}
}

func seed(ctx context.Context, db *ent.Client) error {
	users := []ent.User{
		{
			Name:           "tekitou342",
			Nickname:       "テキトウ342",
			Biography:      "テキトウ342です。",
			AvatarImageURL: "https://i.pinimg.com/564x/b5/af/9e/b5af9e4469bfbe8c23b1e43b8dfa0062.jpg",
			BannerImageURL: "https://i.pinimg.com/564x/0d/14/6b/0d146b3876e366e49297a316af513fee.jpg",
		},
	}

	for _, u := range users {
		user, err := db.User.Create().
			SetName(u.Name).
			SetNickname(u.Nickname).
			SetBiography(u.Biography).
			SetAvatarImageURL(u.AvatarImageURL).
			SetBannerImageURL(u.BannerImageURL).
			Save(ctx)
		if err != nil {
			return err
		}

		if _, err := db.Post.Create().SetAuthor(user).SetText("シードされましたよ？").Save(ctx); err != nil {
			return err
		}
	}

	return nil
}
