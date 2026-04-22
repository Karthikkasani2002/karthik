package postgres

import (

	"database/sql"

	"fmt"

	_ "github.com/lib/pq"

	"onboarding/internal/config"
)

type Client struct {

	DB *sql.DB
}

func New(cfg config.Config) *Client {

	conn := fmt.Sprintf(

		"host=%s port=%s user=%s password=%s dbname=%s sslmode=verify-full sslrootcert=%s",

		cfg.PostgresHost,

		cfg.PostgresPort,

		cfg.PostgresUser,

		cfg.PostgresPass,

		cfg.PostgresDB,

		cfg.PostgresCA,
	)

	db, err := sql.Open("postgres", conn)

	if err != nil {

		panic(err)
	}

	return &Client{DB: db}
}

func (c *Client) Close() {

	c.DB.Close()
}
