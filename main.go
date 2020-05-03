package main

import (
	"github.com/s1250040/go-yumemori/db"
	"github.com/s1250040/go-yumemori/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
