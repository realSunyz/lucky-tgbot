package main

import (
	"github.com/realSunyz/lucky-tgbot/plugin/reborn"
	"github.com/realSunyz/lucky-tgbot/plugin/slash"
	"github.com/realSunyz/lucky-tgbot/plugin/torf"
	"log"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token: os.Getenv("TOKEN"),

		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	rebornData, err := reborn.InitRebornList("plugin/reborn/countries.json")
	if err != nil {
		log.Fatal(err)
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Error creating bot: ", err)
		return
	}

	b.Handle("/reborn", func(c tele.Context) error {
		return reborn.Execute(c, rebornData)
	})
	// b.Handle("/info", info.Execute)

	b.Handle(tele.OnText, func(c tele.Context) error {
		inputText := c.Text()

		if strings.HasPrefix(inputText, "/") {
			return slash.Execute(c)
		} else {
			return torf.Execute(c)
		}
	})

	b.Start()
}
