package main

import (
	u "github.com/OldTyT/buy_elephant/bots"
	"github.com/OldTyT/buy_elephant/modules/message"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	u.Run(
		&message.Message{},
	)
}
