package main

import (
	"github.com/opinionated/scraper-core/net/server"
	"github.com/opinionated/scraper-core/scraper"
	"github.com/opinionated/utils/log"
	"net/http"
)

func main() {
	// make logs write to standard out
	log.InitStd()

	// create the new server
	s := server.NewScrapeServer()
	j := s.GetJefe()

	j.SetCycleTime(1)
	j.Start()

	log.Info("started jefe")

	// make server scrape WSJ
	rss := server.CreateSchedulableRSS(&scraper.WSJRSS{}, 10, j)
	j.AddSchedulable(rss)

	log.Info("going to start server")

	// start up the server
	http.HandleFunc("/", s.Handle())
	http.ListenAndServe(":8080", nil)
}
