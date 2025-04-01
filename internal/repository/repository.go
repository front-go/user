package repository

import (
	"fmt"
	"log"

	"github.com/front-go/user/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	conn sqlx.DB
}

func NewRepository(cfg *config.Config) *Repository {

	connectCmd := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.Host, cfg.Postgres.Port)

	conn, err := sqlx.Connect("postgres", connectCmd)
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{conn: *conn}
}
