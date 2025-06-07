package service

import (
	"database/sql"
	_"github.com/lib/pq"
)

func OpenDB()(*sql.DB, error) {
	connetionString := "user=diego host=localhost port=5432 dbname=equipamento sslmode=disable"
		db, err := sql.Open("postgres",connetionString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
