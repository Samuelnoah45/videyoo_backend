package db

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"server/config"
)

func Connect() (*sql.DB, error) {
	port, err := strconv.Atoi(config.DB_PORT)
    if err != nil {
		return nil, err
    }
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    config.HOST,port, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}
