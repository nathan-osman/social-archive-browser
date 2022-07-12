package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn maintains a connection to the database.
type Conn struct {
	*gorm.DB
}

// New creates a connection to the database.
func New(cfg *Config) (*Conn, error) {
	d, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
				cfg.Host,
				cfg.Port,
				cfg.Name,
				cfg.User,
				cfg.Password,
			),
		),
	)
	if err != nil {
		return nil, err
	}
	return &Conn{
		DB: d,
	}, nil
}

// Close shuts down the database connection.
func (c *Conn) Close() {
	db, _ := c.DB.DB()
	db.Close()
}
