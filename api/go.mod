module github.com/jbuchbinder/rehabber-discord-integration/api

go 1.24.2

replace github.com/jbuchbinder/rehabber-discord-integration/discord => ../discord

require (
	github.com/bwmarrin/discordgo v0.29.0
	github.com/jbuchbinder/rehabber-discord-integration/discord v0.0.0-00010101000000-000000000000
	github.com/jbuchbinder/shims v0.0.0-20250315180801-ea13cafaf717
	github.com/labstack/echo/v4 v4.13.4
)

require (
	github.com/dpapathanasiou/go-recaptcha v0.0.0-20190121160230-be5090b17804 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)
