package main

import (
	"github.com/prashantkhandelwal/devbuzz/db"
	"github.com/prashantkhandelwal/devbuzz/server"
)

func main() {

	db.InitDB()

	server.Run()

}
