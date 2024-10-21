package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/vanvanni/sudjoo/internal/config"
	"github.com/vanvanni/sudjoo/internal/logger"
)

func main() {
	config := config.GetConfig()

	logger.Info("Getting sudo permisions.")
	d, _ := discordgo.New("Bot " + config.Token)
	d.ChannelMessageSend(config.Channels["NT_PLAY"], "Boo!")
}
