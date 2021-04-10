package hetzner

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type Scraper func() ([]Server, error)

type Server struct {
	ID           int     `json:"key"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	CPU          string  `json:"cpu"`
	CPUBenchmark int     `json:"cpu_benchmark"`
	Traffic      string  `json:"traffic"`
	Bandwidth    int     `json:"bandwidth"`
	RAM          int     `json:"ram"`
	RAMDescr     string  `json:"ram_hr"`
	Price        string  `json:"price"`
	ECC          ECCFlag `json:"is_ecc"`
	DiskDescr    string  `json:"hdd_hr"`
	DiskSizeGB   int     `json:"hdd_size"`
	DiskCount    int     `json:"hdd_count"`
	Link         string  `json:"link"`
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

type ECCFlag bool

func (f *ECCFlag) UnmarshalJSON(raw []byte) error {
	err := json.Unmarshal(raw, (*bool)(f))
	if err == nil {
		return nil
	}

	var i int
	err = json.Unmarshal(raw, &i)
	if err != nil {
		return err
	}

	*f = i > 0
	return nil
}
