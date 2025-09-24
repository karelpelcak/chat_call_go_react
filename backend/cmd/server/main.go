package main

import (
	"log"

	"github.com/karelpelcak/chat_call/internal/api/router"
	"github.com/karelpelcak/chat_call/internal/db"
)

func main() {
	db.InitDB()
	db.Migration()
	defer db.DB.Close();

	r := router.SetupRouter()
	log.Print("Server running on port :3333")
	r.Run(":3333")
}
