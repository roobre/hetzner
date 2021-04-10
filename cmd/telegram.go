package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"roob.re/hetzner"
	hetznerparser "roob.re/hetzner/hetzner"
	"time"
)

func main() {
	token := flag.String("token", "", "Telegram bot token")
	chatid := flag.String("chat-id", "", "Telegram chat id")
	interval := flag.Duration("interval", 5*time.Minute, "Time to wait between checks")
	flag.Parse()

	checker := &hetzner.Checker{
		Scraper:  hetznerparser.Scrape,
		Interval: *interval,
		CheckerFunc: hetzner.BetterThan(
			hetzner.Server{
				CPUBenchmark: 10000,
				RAM:          32,
				DiskCount:    2,
				DiskSizeGB:   512,
				DiskDescr:    "(?i:ssd|nvme)",
				ECC:          true,
				Price:        "60",
			},
		),
	}

	var send func(server hetzner.Server) error // Sender function
	if len(*token) == 0 || len(*chatid) == 0 {
		// Just print to stderr if no bot token or chat-id is passed
		log.Println("Invalid token or chat-id, falling back to just printing")
		send = func(server hetzner.Server) error {
			log.Print(server)
			return nil
		}
	} else {
		// Report to telegram chat
		send = func(server hetzner.Server) error {
			resp, err := http.Get("https://api.telegram.org/bot" + *token + "/sendMessage?chat_id=" + *chatid + "&text=" + url.QueryEscape(server.String()))
			if err != nil {
				return err
			}
			_ = resp.Body.Close()
			return nil
		}
	}

	alerter := hetzner.Alerter{
		Realert: 2 * time.Hour * 24 * 30,
		Send:    send,
	}

	checker.Run(alerter.Start())
}
