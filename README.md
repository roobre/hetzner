# hetzner

Hetzner provides a mechanism to receive alerts when Hetzner servers matching requirements become available.

## Design

This package features a modular, object-oriented design which should make integration easy:

- The `Scrapper` function, which returns a list of servers
  + `HetznerScrape()` Implements this for Hetzner
- The `Checker` object, which compares the list of servers returned by `Scrapper` with a reference one. Additionally, it handles running the whole thing periodically. 
- The `Alerter` object, which tracks matching servers and invokes the user defined `Send` function. Alerter keeps track of recently notified servers for you.

In the provided `cmd/telegram.go` file, a very simple integration sending alerts to a Telegram chat can be seen.
