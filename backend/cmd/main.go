package main

import (
	"os"
	"socialNetwork/api/handlers"
	"socialNetwork/pkg/db"

	_ "github.com/mattn/go-sqlite3"
)

var port = "8080"

func main() {
	db.OpenDB("database")
	os.Setenv("PORT", port)
	go handlers.HandlerMain(port)
	select {}
}
