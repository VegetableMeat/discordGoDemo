package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/VegetableMeat/discordGoDemo/config"
	"github.com/VegetableMeat/discordGoDemo/handlers"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Discordへのセッションを開始する
	dg, err := discordgo.New("Bot " + config.GetConf().Discord.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(handlers.MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
