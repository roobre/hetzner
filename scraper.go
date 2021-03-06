package hetzner

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
)

type Scraper func() ([]Server, error)

type Server struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CPU          string `json:"cpu"`
	CPUBenchmark int    `json:"cpu_benchmark"`
	Traffic      string `json:"traffic"`
	Bandwidth    int    `json:"bandwidth"`
	RAM          int    `json:"ram"`
	RAMDescr     string `json:"ram_hr"`
	Price        string `json:"price"`
	ECC          int    `json:"is_ecc"`
	DiskDescr    string `json:"hdd_hr"`
	DiskSizeGB   int    `json:"hdd_size"`
	DiskCount    int    `json:"hdd_count"`
	Link         string `json:"link"`
}

const hetznerURL = "https://www.hetzner.com/dedicated-rootserver/getServer"

func HetznerScrape() ([]Server, error) {
	resp, err := http.Get(hetznerURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("hetzner returned status %d", resp.StatusCode)
	}

	var response struct {
		List []Server `json:"server"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.List, nil
}

func (s *Server) String() string {
	server := "Hetzner server:\n"
	out, err := yaml.Marshal(s)
	if err != nil {
		return ""
	}

	server += string(out)
	return server
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
