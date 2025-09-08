package discord

import "strings"

var (
	DiscordChannelMap = map[string]string{
		"bats":                    "1372949251750887575",
		"birds: ducks and geese":  "1372949132435787988",
		"birds: hawks and owls":   "1372949051942764585",
		"birds: raptors":          "1372948018869047506",
		"birds: small":            "1372948964797845545",
		"birds: other":            "1372949428012453968",
		"deer":                    "1372948696580489330",
		"foxes":                   "1372948653617971240",
		"opossums":                "1372947919757643807",
		"rabbits":                 "1380561508190126301",
		"raccoons":                "1373025872415686786",
		"reptiles and amphibians": "1372948802008387654",
		"skunks":                  "1372948609284444363",
	}
)

func GetDiscordChannelID(name string) string {
	if id, ok := DiscordChannelMap[strings.ToLower(name)]; ok {
		return id
	}
	return ""
}
