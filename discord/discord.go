package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordOutput struct {
	discordSession *discordgo.Session
	discordInit    bool
}

type DiscordMessage struct {
	Species            string
	FinderName         string
	FinderTown         string
	FinderPhone        string
	NumberOfAnimals    int
	Description        string
	AnimalContained    bool
	WillingToTransport bool // If true, the finder is willing to transport the animal
	Files              []*discordgo.File
}

func (dm DiscordMessage) Contained() string {
	if dm.AnimalContained {
		return "contained"
	}
	return "not contained"
}

func (dm DiscordMessage) Transport() string {
	if dm.WillingToTransport {
		return "willing to transport"
	}
	return "not willing to transport"
}

func (d *DiscordOutput) Init(token string) error {
	var err error
	if d.discordInit {
		return fmt.Errorf("ERR: already intiialized: %w", err)
	}

	log.Printf("INFO: Initializing Discord output with token %s", token)
	d.discordSession, err = discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("ERR: New(): %w", err)
	}

	log.Printf("INFO: Opening Discord session")
	err = d.discordSession.Open()
	if err != nil {
		return fmt.Errorf("ERR: Open(): %w", err)
	}

	d.discordInit = true
	return nil
}

func (d *DiscordOutput) SendMessage(channel string, msg DiscordMessage) (string, error) {
	content := fmt.Sprintf(
		"**%s** (%d)\n"+
			"Finder: %s (%s)\n"+
			"Contact: %s\n"+
			"Status: %s / %s\n"+
			"Description: %s\n",
		msg.Species, msg.NumberOfAnimals,
		msg.FinderName, msg.FinderTown,
		msg.FinderPhone,
		msg.Contained(), msg.Transport(),
		msg.Description)

	threadName := fmt.Sprintf("%s (%d) - %s (%s)", msg.Species, msg.NumberOfAnimals, msg.FinderName, msg.FinderTown)

	// Initial message ID
	m := discordgo.MessageSend{
		Content:         content,
		AllowedMentions: &discordgo.MessageAllowedMentions{},
	}
	if len(msg.Files) > 0 {
		log.Printf("INFO: Adding %d files to message", len(msg.Files))
		m.Files = msg.Files
	}

	/*
		log.Printf("INFO: Sending message to channel %s", channel)
		res, err := d.discordSession.ChannelMessageSend(channel, fmt.Sprintf("**%s**\n%s\n%s\n%s", msg.Title, msg.FinderName, msg.FinderTown, msg.FinderPhone))
		if err != nil {
			log.Printf("ERR: ChannelMessageSend: %s", err.Error())
			return "", fmt.Errorf("ERR: ChannelMessageSend(): %w", err)
		}
	*/

	// Initial message, from which to build the "channel"
	log.Printf("INFO: ChannelMessageSendComplex")
	res, err := d.discordSession.ChannelMessageSendComplex(channel, &m)
	if err != nil {
		log.Printf("ERR: ChannelMessageSendComplex: %s", err)
		return "", fmt.Errorf("ERR: ChannelMessageSendComplex(): %w", err)
	}

	// Create complex message thread
	{
		// Threads cannot have titles over 100 characters, otherwise we segfault
		x := threadName
		if len(x) > 100 {
			x = x[:95] + "..."
		}

		log.Printf("INFO: MessageThreadStartComplex")
		t, err := d.discordSession.MessageThreadStartComplex(channel, res.ID, &discordgo.ThreadStart{
			Name: x,
			//AutoArchiveDuration: 300,
			Invitable:        false,
			RateLimitPerUser: 30,
		})
		if err != nil {
			//log.Printf("ERR[%s]: MessageThreadStartComplex(): %s", agentMap[c.CallID], err.Error())
			log.Printf("ERR: MessageThreadStartComplex: %s", err)
			return "", fmt.Errorf("ERR: ChannelMessageSendComplex(): %w", err)
		}

		// Assign the discord channel to the thread
		log.Printf("INFO: Assigning thread %s for msg %s", t.ID, threadName)
		//callMap[c.CallID] = t

		// Set last updated to be long enough ago that this works
		//lastUpdatedMap[c.CallID] = time.Now().Local().Add(time.Hour * -24)
	}

	return res.ID, nil
}
