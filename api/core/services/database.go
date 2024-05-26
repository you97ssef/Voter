package services

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DBService struct {
    db *sql.DB
}

func NewDBService(dataSourceName string) *DBService {
    db, err := sql.Open("libsql", dataSourceName)
    if err != nil {
        panic(err)
    }

    if err := db.Ping(); err != nil {
        panic(err)
    }

    return &DBService{db: db}
}

func (s *DBService) Execute(query string, args ...interface{}) (sql.Result, error) {
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(args...)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (s *DBService) Select(query string, args ...interface{}) (*sql.Rows, error) {
    rows, err := s.db.Query(query, args...)
    if err != nil {
        return nil, err
    }

    return rows, nil
}