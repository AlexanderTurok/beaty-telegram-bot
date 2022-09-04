package repository

import (
	"database/sql"
)

func NewPostgresDB(config Config) (*sql.DB, error) {
	// db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))
	db, err := sql.Open("postgres", "dbname=tgbeaty_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
