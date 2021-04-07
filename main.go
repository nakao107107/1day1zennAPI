package main

import (
	"1day1zennAPI/db"
	"1day1zennAPI/server"
)

func main() {
	db.Init()
	server.Init()
}