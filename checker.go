package hetzner

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

type Checker struct {
	Scraper         Scraper
	MinRequirements []Server
	Interval        time.Duration
}

func (a *Checker) Run(srvchan chan Server) {
	for {
		err := a.Check(srvchan)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Check round finished, waiting %s...", a.Interval.String())
		time.Sleep(a.Interval)
	}
}

func (a *Checker) Check(srvchan chan Server) error {
	servers, err := a.Scraper()
	if err != nil {
		return err
	}

	for _, srv := range servers {
		for _, req := range a.MinRequirements {
			if srv.Bandwidth < req.Bandwidth {
				continue
			}

			if srv.CPUBenchmark < req.CPUBenchmark {
				continue
			}

			if srv.RAM < req.RAM {
				continue
			}
			if !srv.ECC && req.ECC {
				continue
			}

			if srv.DiskCount < req.DiskCount {
				continue
			}
			if srv.DiskSizeGB < req.DiskSizeGB {
				continue
			}

			if !regexp.MustCompile(req.Name).MatchString(srv.Name) {
				continue
			}

			if !regexp.MustCompile(req.Category).MatchString(srv.Category) {
				continue
			}

			if !regexp.MustCompile(req.Traffic).MatchString(srv.Traffic) {
				continue
			}

			if !regexp.MustCompile(req.RAMDescr).MatchString(srv.RAMDescr) {
				continue
			}

			if !regexp.MustCompile(req.DiskDescr).MatchString(srv.DiskDescr) {
				continue
			}

			if parseFloat(srv.Price) > parseFloat(req.Price) {
				continue
			}

			srvchan <- srv
		}
	}

	return nil
}

func parseFloat(s string) float64 {
	r, _ := strconv.ParseFloat(s, 64)
	return r
}
