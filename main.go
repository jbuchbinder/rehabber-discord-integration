package main

import (
	"embed"
	"flag"
	"fmt"
	"log"

	"github.com/jbuchbinder/rehabber-discord-integration/api"
	"github.com/labstack/echo/v4"
)

//go:embed ui
var uiFS embed.FS

var (
	embedded = flag.Bool("embedded", false, "Use embedded UI files instead of filesystem")
	port     = flag.Int("port", 1323, "Port to run the server on")
)

func main() {
	flag.Parse()
	log.Printf("INFO: Starting server")
	e := echo.New()
	api.InitApi(e, uiFS, *embedded)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
