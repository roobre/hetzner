package hetzner

import (
	"testing"
)

func TestScrape(t *testing.T) {
	servers, err := Scrape()
	if err != nil {
		t.Fatal(err)
	}

	if len(servers) <= 0 {
		t.Fatal("Returned empty list of servers")
	}
}

func TestServers(t *testing.T) {
	serverList, err := servers(serversUrl())
	if err != nil {
		t.Fatal(err)
	}

	if len(serverList) <= 0 {
		t.Fatal("Returned empty list of servers")
	}
}

func TestAuctions(t *testing.T) {
	auctions, err := servers(auctionUrl())
	if err != nil {
		t.Fatal(err)
	}

	if len(auctions) <= 0 {
		t.Fatal("Returned empty list of auctions")
	}
}
