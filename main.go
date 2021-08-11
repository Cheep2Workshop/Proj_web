package main

import (
	"log"

	"github.com/Cheep2Workshop/proj-web/controller"
	"github.com/Cheep2Workshop/proj-web/grpc/dashboardserver"
	"github.com/Cheep2Workshop/proj-web/models/repo"
)

func main() {
	// run grpc server
	go dashboardserver.Run()

	i := make(chan int)
	go controller.RunGin(i)
	result := <-i
	log.Println(result)
}

func init() {
	origin := repo.DbConfig{
		Account:  "root",
		Password: "QMKAJNNjNK9vBO88",
		Ip:       "127.0.0.1",
		Port:     "3306",
		DbName:   "mysql",
	}
	dest := repo.DbConfig{
		Account:  "root",
		Password: "QMKAJNNjNK9vBO88",
		Ip:       "127.0.0.1",
		Port:     "3306",
		DbName:   "dashboard",
	}
	repo.InitRemoteDB(&origin, &dest)

	// .sql create table
	// v1.0 -> v1.1 true (success)
}
