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
	port      = os.Getenv("PORT")       // sets automatically
	publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
	token     = os.Getenv("TOKEN")      // you must add it to your config vars
)

func main() {
	//https://api.telegram.org1476386207:AAEG5kTR8KL2A3xJN5xmGF5ODVwTxxnZWT4
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}
	pref := tb.Settings{
		Token: token,
		// URL:    "https://api.telegram.org/bot1476386207:AAEG5kTR8KL2A3xJN5xmGF5ODVwTxxnZWT4",
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle("/hello", func(m *tb.Message) {
		b.Reply(m, "Hello dotadrochers😈\n\nYou can control me by sending these commands:\n\n/hello - info about me.\n/start  - this command run process which sends voting poll in random time during the next 5 days from the moment it started.\n(use only one time its recursive process)")
	})

	poll := &tb.Poll{
		Question:    "Friday Dota2 voting started!!!",
		Explanation: "Explanation",
	}

	poll.AddOptions("Yes", "No")
	// max := 432000
	// min := 86400
	max := 50
	min := 10

	b.Handle("/start", func(m *tb.Message) {
		for {
			rand.Seed(time.Now().UnixNano())
			randomNum := random(min, max)
			fmt.Println(randomNum)
			time.Sleep(time.Duration(randomNum) * time.Second)
			b.Reply(m, poll)
			fmt.Println(max - randomNum)
			time.Sleep(time.Duration(max-randomNum+min+min) * time.Second)
		}
	})

	b.Start()

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
