package discord

import (
	"os"
	"testing"

	"github.com/alexsasharegan/dotenv"
	"github.com/bwmarrin/discordgo"
	"github.com/jbuchbinder/shims"
)

func Test_Discord(t *testing.T) {
	dotenv.Load()
	d := &DiscordOutput{}
	err := d.Init(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		t.Fatalf("Failed to initialize Discord: %v", err)
	}
	m := DiscordMessage{
		Species:            "Opossum",
		FinderName:         "Test Finder",
		FinderTown:         "Test Town",
		FinderPhone:        "1234567890",
		NumberOfAnimals:    1,
		Description:        "Test Description",
		AnimalContained:    true,
		WillingToTransport: true,
		Files: []*discordgo.File{
			{
				Name:        "opossum-walking.jpg",
				ContentType: "image/jpeg",
				Reader:      shims.SingleValueDiscardError(os.Open("testdata/opossum-walking.jpg")),
			},
		},
	}

	d.SendMessage("1372947919757643807", m) // Initialize with a dummy token0
}
