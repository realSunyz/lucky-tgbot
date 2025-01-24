package main

import (
	"github.com/realSunyz/lucky-tgbot/plugin/reborn"
	"github.com/realSunyz/lucky-tgbot/plugin/slash"
	"github.com/realSunyz/lucky-tgbot/plugin/torf"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Error creating bot: ", err)
		return
	}

	rebornData, err := reborn.InitRebornList("plugin/reborn/countries.json")
	if err != nil {
		log.Fatal("Error initializing rebornData: ", err)
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	b.Handle("/reborn", func(c tele.Context) error {
		return reborn.Execute(c, r, rebornData)
	})

	// b.Handle("/info", info.Execute)

	b.Handle(tele.OnText, func(c tele.Context) error {
		inputText := c.Text()

		if strings.HasPrefix(inputText, "/") {
			return slash.Execute(c)
		} else {
			return torf.Execute(c, r)
		}
	})

	b.Start()
}
