package main

import (
	"github.com/opinionated/scraper-core/net/server"
	"github.com/opinionated/scraper-core/scraper"
	"github.com/opinionated/utils/log"
	"net/http"
	"os"
)

func main() {
	infoFile, err := os.OpenFile("scrapeInfoLog.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	errFile, err := os.OpenFile("scrapeErrorLog.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	warnFile, err := os.OpenFile("scrapeWarnLog.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer infoFile.Close()
	defer errFile.Close()
	defer warnFile.Close()

	// create the new server
	s := server.NewScrapeServer()
	j := s.GetJefe()

	j.SetCycleTime(1)
	j.Start()

	log.Info("started jefe")

	// make server scrape WSJ
	rss := server.CreateSchedulableRSS(&scraper.NYTRSS{}, 10, j)
	j.AddSchedulable(rss)

	log.Info("going to start server")

	// start up the server
	http.HandleFunc("/", s.Handle())
	http.ListenAndServe(":8080", nil)
}
