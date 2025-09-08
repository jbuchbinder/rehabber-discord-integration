package api

import (
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/jbuchbinder/rehabber-discord-integration/discord"
	"github.com/jbuchbinder/shims"
	"github.com/labstack/echo/v4"
)

const (
	// uploadPath is the directory where uploaded files will be stored
	uploadPath = "/tmp"
)

func InitApi(e *echo.Echo, fs embed.FS, embedded bool) {
	// Register the POST route for form submission
	e.POST("/api/post", PostForm)

	// Recaptcha
	e.GET("/api/recaptcha", func(c echo.Context) error {
		log.Printf("INFO: Serving recaptcha")
		return c.JSON(http.StatusOK, echo.Map{
			"siteKey": os.Getenv("RECAPTCHA_SITE_KEY"),
		})
	})
	// Register the GET route for serving the UI

	if embedded {
		log.Printf("INFO: Serving embedded UI files")
		e.StaticFS("/", fs)
		return
	}

	log.Printf("INFO: Serving static UI files from filesystem")
	fsx := echo.MustSubFS(e.Filesystem, "ui")
	e.StaticFS("/", fsx)
}

func PostForm(c echo.Context) error {
	// Get name
	finderName := c.FormValue("finderName")
	if finderName == "" {
		log.Printf("ERROR: Finder name is required")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Finder name is required",
		})
	}
	finderTown := c.FormValue("finderTown")
	if finderTown == "" {
		log.Printf("ERROR: Finder town is required")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Finder town is required",
		})
	}
	finderEmail := c.FormValue("finderEmail")
	finderPhone := c.FormValue("finderPhone")
	if finderPhone == "" {
		log.Printf("ERROR: Finder phone is required")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Finder phone is required",
		})
	}
	species := c.FormValue("species")
	if species == "" {
		log.Printf("ERROR: Species is required")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Species is required",
		})
	}
	numberOfAnimals := int(shims.SingleValueDiscardError(strconv.ParseInt(c.FormValue("numberOfAnimals"), 32, 10)))
	if numberOfAnimals == 0 {
		log.Printf("ERROR: Number of animals is required")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Number of animals is required",
		})
	}

	description := c.FormValue("description")
	if description == "" {
		log.Printf("ERROR: Description is required")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Description is required",
		})
	}
	animalContained, _ := strconv.ParseBool(c.FormValue("animalContained"))
	willingToTransport, _ := strconv.ParseBool(c.FormValue("willingToTransport"))

	log.Printf("INFO: Received form submission: species=%s, numberOfAnimals=%d, finderName=%s, finderTown=%s, finderPhone=%s, finderEmail=%s, description=%s, animalContained=%t, willingToTransport=%t",
		species, numberOfAnimals, finderName, finderTown, finderPhone, finderEmail, description, animalContained, willingToTransport)

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	m := discord.DiscordMessage{
		Species:            species,
		FinderName:         finderName,
		FinderTown:         finderTown,
		FinderPhone:        finderPhone,
		NumberOfAnimals:    numberOfAnimals,
		Description:        description,
		AnimalContained:    animalContained,
		WillingToTransport: willingToTransport,
		Files:              []*discordgo.File{},
	}

	log.Printf("DEBUG: %#v", form.File)
	log.Printf("INFO: Received %d files in form submission", len(form.File["file[]"]))
	files := form.File["file[]"]

	removeFiles := []string{}

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		defer src.Close()

		// Destination
		//uploadedFileName := file.Filename
		uploadedFilePath := path.Join(uploadPath, file.Filename)

		log.Printf("INFO: Saving file %s to %s", file.Filename, uploadedFilePath)
		dst, err := os.Create(uploadedFilePath)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		m.Files = append(m.Files, &discordgo.File{
			Name:        file.Filename,
			ContentType: "image/jpeg",
			Reader:      shims.SingleValueDiscardError(os.Open(uploadedFilePath)),
		},
		)
		removeFiles = append(removeFiles, uploadedFilePath)

		log.Printf("INFO: Uploaded file %s to %s", file.Filename, uploadedFilePath)
	}

	d := &discord.DiscordOutput{}
	err = d.Init(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Printf("ERROR: Failed to initialize Discord: %v", err)
	}

	channelID := discord.GetDiscordChannelID(species)
	if channelID == "" {
		log.Printf("ERROR: Invalid species '%s' provided", species)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid species provided",
		})
	}

	// TODO:FIXME: routing logic
	d.SendMessage(channelID, m) // Initialize with a dummy token

	for _, rf := range removeFiles {
		os.Remove(rf)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"name":   finderName,
		"status": "uploaded",
	})
}
