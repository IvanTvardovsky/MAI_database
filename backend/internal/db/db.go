package db

import (
	"backend/internal/config"
	"backend/internal/schemas"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func ConnectToDB(cfg schemas.DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

var singleton sync.Once
var DB *sql.DB

func GetDB() (*sql.DB, error) {
	var Err error
	singleton.Do(func() {
		DB, Err = ConnectToDB(config.ProjectConfig.DB)
	})
	return DB, Err

}
