package main

import (
	"log"

	"github.com/Cheep2Workshop/proj-web/controller"
	"github.com/Cheep2Workshop/proj-web/grpc/dashboardserver"
	"github.com/Cheep2Workshop/proj-web/orm"
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
	origin := orm.DbConfig{
		Account:  "root",
		Password: "QMKAJNNjNK9vBO88",
		Ip:       "127.0.0.1",
		Port:     "3306",
		DbName:   "mysql",
	}
	dest := orm.DbConfig{
		Account:  "root",
		Password: "QMKAJNNjNK9vBO88",
		Ip:       "127.0.0.1",
		Port:     "3306",
		DbName:   "dashboard",
	}
	orm.InitRemoteDB(&origin, &dest)
}
