package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Cheep2Workshop/proj-web/controller"
	"github.com/Cheep2Workshop/proj-web/grpc/dashboardserver"
	"github.com/Cheep2Workshop/proj-web/models/repo"
)

func main() {
	// initial context and waitgroup
	c, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// run grpc server
	go dashboardserver.Run(c, wg)
	// run http server
	go controller.RunGin(c, wg)

	// wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")
	cancel()
	wg.Wait()
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
