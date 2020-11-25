package main

import (
	"fmt"
	"log"
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
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})
	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		fmt.Println(m.Payload) // <PAYLOAD>
	})

	poll := &tb.Poll{
		Question:    "Friday Dota2 voting started!!!",
		Explanation: "Explanation",
	}

	poll.AddOptions("Yes", "No")

	b.Handle("/poll", func(m *tb.Message) {
		b.Send(m.Sender, poll)

	})

	b.Start()
}
