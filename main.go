package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "https://api.telegram.org",

		Token:  "1476386207:AAEG5kTR8KL2A3xJN5xmGF5ODVwTxxnZWT4",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		// https://api.telegram.org1476386207:AAEG5kTR8KL2A3xJN5xmGF5ODVwTxxnZWT4/getUpdates
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Reply(m, "Hello dotadrochers😈\nLet's start our game.")
	})

	poll := &tb.Poll{
		Question:    "Friday Dota2 voting started!!!",
		Explanation: "Explanation",
	}

	poll.AddOptions("Yes", "No")
	max := 432000
	min := 86400

	b.Handle("/kettik", func(m *tb.Message) {
		for {
			rand.Seed(time.Now().UnixNano())
			randomNum := random(min, max)
			fmt.Println(randomNum)
			time.Sleep(time.Duration(randomNum) * time.Second)
			b.Reply(m, poll)
			fmt.Println(max - randomNum)
			time.Sleep(time.Duration(max-randomNum) * time.Second)
		}
	})

	b.Start()
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
