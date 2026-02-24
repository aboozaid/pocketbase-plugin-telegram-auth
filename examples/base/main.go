package main

import (
	"log"

	tgAuthPlugin "github.com/aboozaid/pocketbase-plugin-telegram-auth"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	// Setup tg auth for users collection
	tgAuthPlugin.MustRegister(app, &tgAuthPlugin.Options{
		BotToken:      "YOUR_SUPER_SECRET_BOT_TOKEN", // Better to use ENV variable for that
		CollectionKey: "users",
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
