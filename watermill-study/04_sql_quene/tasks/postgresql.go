package tasks

import (
  "database/sql"
  _ "github.com/jackc/pgx/v5/stdlib" // required for pgx to work
)

func NewPostgresConnForQueue(dataSourceName string) (*sql.DB, error) {
  db, err := sql.Open("pgx", dataSourceName)
  if err != nil {
    return nil, err
  }

  return db, nil
}
