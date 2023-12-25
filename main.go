package main

import (
	"github.com/dragonsecurity/dragon"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type application struct {
	App *dragon.Dragon
	wg  sync.WaitGroup
}

func main() {
	d := initApplication()
	go d.listenForShutdown()
	err := d.App.ListenAndServe()
	d.App.ErrorLog.Println(err)
}

func (a *application) shutdown() {
	a.wg.Wait()
}

func (a *application) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	a.App.InfoLog.Println("Received signal", s.String())
	a.shutdown()

	os.Exit(0)
}
