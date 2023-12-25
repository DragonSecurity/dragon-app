package main

import (
	"github.com/dragonsecurity/dragon"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	getDragon := &dragon.Dragon{}
	err = getDragon.New(path)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		App: getDragon,
	}
	return app
}
