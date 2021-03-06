package hetzner

import (
	"testing"
)

func TestScrape(t *testing.T) {
	servers, err := HetznerScrape()
	if err != nil {
		t.Fatal(err)
	}

	if len(servers) <= 0 {
		t.Fatal("Returned empty list of servers")
	}
}

