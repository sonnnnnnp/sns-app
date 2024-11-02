package db

import (
	"database/sql"
	"fmt"

)

type ConnectionOptions struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
	SSLMode  string
}

func Open(opt *ConnectionOptions) (*sql.DB, error) {
	return sql.Open("pgx", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", opt.Host, opt.Port, opt.User, opt.Name, opt.Password, opt.SSLMode))
}
