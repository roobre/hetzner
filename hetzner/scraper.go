package hetzner

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"roob.re/hetzner"
	"time"
)

// Parses hetzner server list from their JSON format
func servers(url string) ([]hetzner.Server, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("hetzner returned status %d", resp.StatusCode)
	}

	var response struct {
		List []hetzner.Server `json:"server"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.List, nil
}

// Return available auctions and servers from hetzner
func Scrape() ([]hetzner.Server, error) {
	serverList, err := servers(serversUrl())
	if err != nil {
		return nil, err
	}

	auctionList, err := servers(auctionUrl())
	if err != nil {
		return nil, err
	}

	log.Printf("Hetzner: Fetched %d servers and %d auctions", len(serverList), len(auctionList))

	return append(serverList, auctionList...), nil
}

func serversUrl() string {
	return "https://www.hetzner.com/dedicated-rootserver/getServer"
}
func auctionUrl() string {
	return fmt.Sprintf("https://www.hetzner.com/a_hz_serverboerse/live_data.json?m=%d", time.Now().Unix())
}

//    {
//      "id": 13758,
//      "key": 13758,
//      "name": "AX41",
//      "description": "<p>The AX41 Dedicated Root Server houses the powerful third generation Ryzen 5 3600 CPU from AMD, which is based on Zen 2 architecture. With 6 cores and 12 threads, this server's optimized for creative workloads and applications that have high multi-threading requirements. In addition, the AX41 comes with two 2 TB HDDs and 64 GB of DDR4 RAM.</p>",
//      "category": "Dedicated Root Server AX",
//      "cat_id": 2839,
//      "cores": 6,
//      "threads": 12,
//      "cpu": "AMD Ryzen 5 3600 Hexa-Core",
//      "cpu_fullname": "AMD Ryzen 5 3600 Hexa-Core Matisse (Zen2) Simultaneous Multithreading",
//      "cpu_benchmark": 19914,
//      "is_highio": false,
//      "traffic": "Unlimited",
//      "bandwith": 1000,
//      "ram": 64,
//      "ram_hr": "64 GB DDR4",
//      "price": "40.46",
//      "is_multiprice": true,
//      "is_disabled": false,
//      "is_pre_order": false,
//      "disabled_text": null,
//      "setup_price": "46.41",
//      "hdd_hr": "2x 2 TB SATA Enterprise HDD",
//      "hdd_size": 2000,
//      "hdd_count": 2,
//      "is_special": false,
//      "link": "dedicated-rootserver/ax41",
//      "countries": [
//        {
//          "name": "Germany",
//          "shortcode": "de"
//        },
//        {
//          "name": "Finland",
//          "shortcode": "fi"
//        }
//      ],
//      "is_ecc": 0,
//      "is_configureable": 1,
//      "canConfigureDisk": [],
//      "isLocationDependentMonthlyPrice": true,
//      "isLocationDependentSetupPrice": false,
//      "datacenter": [
//        {
//          "datacenter": "FSN1",
//          "name": "Falkenstein",
//          "shortname": "FSN",
//          "country": "Germany",
//          "country_shortcode": "de"
//        },
//        {
//          "datacenter": "HEL1",
//          "name": "Helsinki",
//          "shortname": "HEL",
//          "country": "Finland",
//          "country_shortcode": "fi"
//        }
//      ],
//      "specials": []
//    },
