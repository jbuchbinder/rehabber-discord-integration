package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/jbuchbinder/rehabber-discord-integration/api"
	"github.com/labstack/echo/v4"
)

//go:embed ui
var uiFS embed.FS

var (
	embedded            = flag.Bool("embedded", false, "Use embedded UI files instead of filesystem")
	port                = flag.Int("port", 1323, "Port to run the server on")
	recaptchaPrivateKey = flag.String("recaptcha-private-key", "", "Recaptcha private key (RECAPTCHA_PRIVATE_KEY environment variable)")
)

func main() {
	flag.Parse()
	err := dotenv.Load()
	if err != nil {
		panic(err)
	}

	if *recaptchaPrivateKey == "" {
		*recaptchaPrivateKey = os.Getenv("RECAPTCHA_PRIVATE_KEY")
	}

	recaptcha.Init(*recaptchaPrivateKey)

	log.Printf("INFO: Starting server")
	e := echo.New()
	api.InitApi(e, uiFS, *embedded)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
