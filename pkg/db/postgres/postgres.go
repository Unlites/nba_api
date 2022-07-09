package postgres

import (
	"fmt"
	"os"

	"github.com/Unlites/nba_api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.PgHost,
		cfg.Postgres.PgPort,
		cfg.Postgres.PgUsername,
		os.Getenv("PG_PASSWORD"),
		cfg.Postgres.PgDBName,
		cfg.Postgres.PgSSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
