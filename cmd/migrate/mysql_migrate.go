package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:QMKAJNNjNK9vBO88@tcp(127.0.0.1:3306)/dashboard?multiStatements=true")
	if err != nil {
		log.Println(err.Error())
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../sql/",
		"dashboard",
		driver,
	)
	if err != nil {
		log.Println(err.Error())
	}

	err = m.Steps(1)
	if err != nil {
		log.Println(err.Error())
	}
}
