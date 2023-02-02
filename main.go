package main

import (
	"epc.com/api/db"
	"epc.com/api/server"
)

func main() {
	db.Init()
	server.Init()
}
