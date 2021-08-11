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
		return
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../sql/",
		"dashboard",
		driver,
	)
	if err != nil {
		log.Println(err.Error())
		return
	}

	step := -1
	version, dirty, err := m.Version()
	log.Printf("Ver:%v -> %v,Step:%v, Dirty:%v", version, int(version)+step, step, dirty)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = m.Steps(step)
	if err != nil {
		log.Println(err.Error())
	}
}
