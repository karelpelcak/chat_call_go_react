package main

import (
	"fmt"

	"github.com/karelpelcak/chat_call/internal/db"
)

func main() {
	db.InitDB()
	fmt.Println("Hello, world!")
}
