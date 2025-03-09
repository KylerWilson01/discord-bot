package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Says hello world",
		},
		{
			Name:        "list",
			Description: "lists the saved comics",
		},
		{
			Name:        "save",
			Description: "saves a comic with the given id",
		},
		{
			Name:        "random",
			Description: "gets a random xkcd comic",
		},
		{
			Name:        "latest",
			Description: "gets the latest xkcd comic",
		},
	}

	handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Hello %s", i.Member.User.Mention()),
				},
			})
		},
		"latest": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			c := http.Client{}
			resp, err := c.Get("https://xkcd.com/info.0.json")
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.ErrCodeGeneralError,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("Could not retrieve latest comic: %s", err.Error()),
					},
				})
				return
			}

			decoder := json.NewDecoder(resp.Body)

			var comic Comic
			err = decoder.Decode(&comic)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.ErrCodeGeneralError,
					Data: &discordgo.InteractionResponseData{
						Content: "Could not decode comic",
					},
				})
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: comic.Img,
				},
			})
		},
		"get": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			c := http.Client{}
			resp, err := c.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", 24))
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.ErrCodeGeneralError,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("Could not retrieve latest comic: %s", err.Error()),
					},
				})
				return
			}

			decoder := json.NewDecoder(resp.Body)

			var comic Comic
			err = decoder.Decode(&comic)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.ErrCodeGeneralError,
					Data: &discordgo.InteractionResponseData{
						Content: "Could not decode comic",
					},
				})
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Hello %s", i.Member.User.Mention()),
				},
			})
		},
		"save": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Hello %s", i.Member.User.Mention()),
				},
			})
		},
	}
)

// InitHandler creates the handlers
func InitHandler(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

// RegisterHandlers registers all the handlers
func RegisterHandlers(s *discordgo.Session) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
