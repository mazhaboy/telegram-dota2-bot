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

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	poll := &tb.Poll{
		Question:    "Friday Dota2 voting started!!!",
		Explanation: "Explanation",
	}

	poll.AddOptions("Yes", "No")

	b.Handle("/start", func(m *tb.Message) {
		rand.Seed(time.Now().UnixNano())
		randomNum := random(86400, 518400)
		fmt.Println(randomNum)
		time.Sleep(time.Duration(randomNum) * time.Second)
		b.Reply(m, poll)

	})

	b.Start()
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
