module github.com/jbuchbinder/rehabber-discord-integration

go 1.24.2

replace (
	github.com/jbuchbinder/rehabber-discord-integration/api => ./api
	github.com/jbuchbinder/rehabber-discord-integration/discord => ./discord
)

require (
	github.com/alexsasharegan/dotenv v0.0.0-20171113213728-090a4d1b5d42
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dpapathanasiou/go-recaptcha v0.0.0-20190121160230-be5090b17804
	github.com/jbuchbinder/rehabber-discord-integration/api v0.0.0-00010101000000-000000000000
	github.com/jbuchbinder/shims v0.0.0-20250818154854-22c0ac83b788
	github.com/labstack/echo/v4 v4.13.4
)

require (
	github.com/gabriel-vasile/mimetype v1.4.10 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jbuchbinder/rehabber-discord-integration/discord v0.0.0-00010101000000-000000000000 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.29.0 // indirect
)
