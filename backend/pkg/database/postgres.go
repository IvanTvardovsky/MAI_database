package database

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Init(config *config.Config) *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Storage.Host, config.Storage.Port, config.Storage.Username, config.Storage.Password, config.Storage.Database)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		_ = fmt.Errorf("can not connect to database")
	}

	err = db.Ping()

	return db
}
