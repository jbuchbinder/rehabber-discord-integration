module github.com/jbuchbinder/rehabber-discord-integration/api

go 1.24.2

replace github.com/jbuchbinder/rehabber-discord-integration/discord => ../discord

require (
	github.com/bwmarrin/discordgo v0.29.0
	github.com/gabriel-vasile/mimetype v1.4.12
	github.com/jbuchbinder/rehabber-discord-integration/discord v0.0.0-20251029164615-efb600e47af6
	github.com/jbuchbinder/shims v0.0.0-20251029164657-6c80f5d6bc01
	github.com/labstack/echo/v4 v4.14.0
)

require (
	github.com/alexsasharegan/dotenv v0.0.0-20171113213728-090a4d1b5d42 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
)
