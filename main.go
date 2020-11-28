package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	port      = os.Getenv("PORT")
	publicURL = os.Getenv("PUBLIC_URL")
	token     = os.Getenv("TOKEN")
)

func main() {

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}
	pref := tb.Settings{
		Token: token,

		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Reply(m, "Hello dotadrochers😈\n\nYou can control me by sending these commands:\n\n/hello - info about me.\n/start  - this command run process which sends voting poll in random time during the next 5 days from the moment it started.\n(use only one time during the week)")
	})

	poll := &tb.Poll{
		Question:    "Friday Dota2 voting started!!!",
		Explanation: "Explanation",
	}

	poll.AddOptions("Yes", "No")
	max := 432000
	min := 0

	b.Handle("/start", func(m *tb.Message) {
		if m.Sender.Username == "majazzzzz" {
			rand.Seed(time.Now().UnixNano())
			randomNum := random(min, max)
			fmt.Println(randomNum)
			time.Sleep(time.Duration(randomNum) * time.Second)
			b.Reply(m, poll)
		} else {
			b.Reply(m, "You are not alowed to use this command(")
		}
	})

	b.Start()

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
