package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// test connection
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return &Adapter{db: db}, nil
}


func (da Adapter) CloseDBConnection() {
	if err := da.db.Close(); err != nil {
		log.Fatalf("failed to close db: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").
		Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).
		ToSql()

	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}