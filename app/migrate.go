package app

import (
 "log"
 "github.com/golang-migrate/migrate/v4"
 "github.com/golang-migrate/migrate/v4/database/postgres"
 _ "github.com/golang-migrate/migrate/v4/source/file"
 _ "github.com/lib/pq"
)

func (a *App) Migrate() {
 driver, err := postgres.WithInstance(a.DB, &postgres.Config{})
 if err != nil {
  log.Println("migrate1",err)
 }
 m, err := migrate.NewWithDatabaseInstance(
  "file://./migrations/","stori", driver)
 if err != nil {
  log.Println("migrate2",err)
 }
 if err := m.Steps(2); err != nil {
  log.Println("migrate3",err)
 }
}