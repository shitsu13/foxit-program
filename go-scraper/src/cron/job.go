package cron

import (
	"go-scraper/src/dao/model"
	"go-scraper/src/scraper"

	log "github.com/sirupsen/logrus"
)

// run scraper
func scrape() {
	s := scraper.GetScraper()
	s.OnPostsListener().OnPostDetailListener()
	if err := s.Run(); err != nil {
		log.Panicf("scrape error: %s\n", err)
	}

	keys := s.GetKeys()
	latest_pid := keys[0]

	// check if it's latest, then needed to update
	if exist, _, _ := model.IsPostExistedByPId(latest_pid); !exist {
		// remove all and save data
		model.DeletePosts()
		s.Save()
	}
}
