package hetzner

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

type Checker struct {
	Scraper     Scraper
	CheckerFunc CheckerFunc
	Interval    time.Duration
}

// CheckerFunc is a funciton that returns true if a Server matches requirements
type CheckerFunc func(*Server) bool

// BetterThan returns a CheckerFunc that returns true if the server being checked has better specs than the reference server
func BetterThan(req Server) CheckerFunc {
	return func(srv *Server) bool {
		if srv.Bandwidth < req.Bandwidth {
			return false
		}

		if srv.CPUBenchmark < req.CPUBenchmark {
			return false
		}

		if srv.RAM < req.RAM {
			return false
		}
		if !srv.ECC && req.ECC {
			return false
		}

		if srv.DiskCount < req.DiskCount {
			return false
		}
		if srv.DiskSizeGB < req.DiskSizeGB {
			return false
		}

		if !regexp.MustCompile(req.Name).MatchString(srv.Name) {
			return false
		}

		if !regexp.MustCompile(req.Category).MatchString(srv.Category) {
			return false
		}

		if !regexp.MustCompile(req.Traffic).MatchString(srv.Traffic) {
			return false
		}

		if !regexp.MustCompile(req.RAMDescr).MatchString(srv.RAMDescr) {
			return false
		}

		if !regexp.MustCompile(req.DiskDescr).MatchString(srv.DiskDescr) {
			return false
		}

		if parseFloat(srv.Price) > parseFloat(req.Price) {
			return false
		}

		return true
	}
}

// AnyOf returns a CheckerFunc that returns true if any of the passed CheckerFunc args returns true
func AnyOf(funcs ...CheckerFunc) CheckerFunc {
	return func(server *Server) bool {
		for _, f := range funcs {
			if f(server) {
				return true
			}
		}
		return false
	}
}

// AllOf returns a CheckerFunc that returns true if all of the passed CheckerFunc args returns true
func AllOf(funcs ...CheckerFunc) CheckerFunc {
	return func(server *Server) bool {
		for _, f := range funcs {
			if !f(server) {
				return false
			}
		}
		return true
	}
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
		if a.CheckerFunc(&srv) {
			srvchan <- srv
		}
	}

	return nil
}

func parseFloat(s string) float64 {
	r, _ := strconv.ParseFloat(s, 64)
	return r
}
