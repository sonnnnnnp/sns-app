package db

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

type ConnectionOptions struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
	SSLMode  string
}

func Open(opt *ConnectionOptions) (*ent.Client, error) {
	return ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", opt.Host, opt.Port, opt.User, opt.Name, opt.Password, opt.SSLMode))
}
