package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"github.com/KylerWilson01/fish-fact-bot/internal/commands"
)

var s *discordgo.Session

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Could not load env variables")
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic("Could not create bot")
	}
}

func init() {
	commands.InitHandler(s)
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	commands.RegisterHandlers(s)
}
