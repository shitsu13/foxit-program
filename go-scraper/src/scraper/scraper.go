package scraper

import (
	"encoding/json"
	"fmt"

	"go-scraper/src/dao/model"
	"go-scraper/src/dcard"

	"github.com/gocolly/colly/v2"

	log "github.com/sirupsen/logrus"
)

var scraper *Scraper

type Scraper struct {
	collector        *colly.Collector
	detail_collector *colly.Collector
	content          map[int64]*model.Post
	keys             []int64
}

func init() {
	scraper = new(Scraper)
	scraper.collector = colly.NewCollector(
		colly.AllowedDomains("www.dcard.tw"),
		colly.AllowURLRevisit(),
	)
	scraper.detail_collector = scraper.collector.Clone()
	scraper.content = make(map[int64]*model.Post, 0)
	scraper.keys = make([]int64, 0)
}

func GetScraper() *Scraper {
	return scraper
}

func (s *Scraper) GetKeys() []int64 {
	return s.keys
}

// on listener for posts
func (s *Scraper) OnPostsListener() *Scraper {
	s.collector.OnResponse(func(r *colly.Response) {
		posts := make([]dcard.Posts, 0)
		if err := json.Unmarshal(r.Body, &posts); err != nil {
			log.Panicf("response posts body unmarshal error: %s\n", err)
		}

		for _, v := range posts {
			pid, title, categories, topics, media, created_at := v.Id, v.Title, v.Categories, v.Topics, v.Media, v.CreatedAt

			post := model.NewPost(pid, title, created_at)
			if len(categories) > 0 {
				post.Categories = categories
			}
			if len(topics) > 0 {
				post.Topics = topics
			}

			urls := make([]string, 0)
			for _, vv := range media {
				urls = append(urls, vv.URL)
			}
			post.MediaURL = urls

			s.keys = append(s.keys, pid)
			s.content[pid] = post

			if err := scrapeDetail(pid); err != nil {
				log.Debugf("scrape detail error: %s\n", err)
				continue
			}
		}
	})

	s.collector.OnError(func(r *colly.Response, e error) {
		log.Errorf("request url: %s\nfailed with response: %v\nerror: %s", r.Request.URL, r, e)
	})

	return s
}

// on listener for post detail
func (s *Scraper) OnPostDetailListener() *Scraper {
	s.detail_collector.OnResponse(func(r *colly.Response) {
		post := dcard.Post{}
		if err := json.Unmarshal(r.Body, &post); err != nil {
			log.Panicf("response post detail body unmarshal error: %s\n", err)
		}

		s.content[post.Id].Content = post.Content
	})

	s.detail_collector.OnError(func(r *colly.Response, e error) {
		log.Errorf("request url: %s\nfailed with response: %v\nerror: %s", r.Request.URL, r, e)
	})

	return s
}

// run to start scraping
func (s *Scraper) Run() (err error) {
	url := fmt.Sprintf("https://%s/%s?popular=false", dcard.DCARD_URL, fmt.Sprintf(dcard.FORUM_API, "apple"))
	if err = s.collector.Visit(url); err != nil {
		return
	}
	return
}

// save data to storage
func (s *Scraper) Save() {
	posts := make([]interface{}, 0)
	content, keys := s.content, s.keys

	for _, v := range keys {
		posts = append(posts, content[v])
	}

	model.SavePosts(posts)
}

// scrape detail for each post
func scrapeDetail(id int64) (err error) {
	url := fmt.Sprintf("https://%s/%s", dcard.DCARD_URL, fmt.Sprintf(dcard.POST_API, id))
	if err = scraper.detail_collector.Visit(url); err != nil {
		return
	}
	return
}
