package cron

import (
	"time"

	"github.com/go-co-op/gocron"

	log "github.com/sirupsen/logrus"
)

var c *Cron

type Cron struct {
	scheduler *gocron.Scheduler
	startAt   time.Time
}

func init() {
	c = New()
}

func New() *Cron {
	c := new(Cron)
	c.scheduler = gocron.NewScheduler(time.Local)
	c.startAt = duration()

	return c
}

// add job
func (c *Cron) addJob(jobFun interface{}, params ...interface{}) *Cron {
	if _, err := c.scheduler.Every("6h").StartAt(c.startAt).SingletonMode().Do(jobFun, params...); err != nil {
		log.Panicf("cron, add job error: %s\n", err)
	}
	return c
}

func (c *Cron) start() {
	c.scheduler.StartAsync()
}

func (c *Cron) stop() {
	c.scheduler.Stop()
}

/*
	duration by start at hour
*/
func duration() (d time.Time) {
	now := time.Now()
	min := now.Minute()

	duration := 0
	if min > 0 {
		duration = 60 - min
	}

	d = now.Add(time.Duration(duration) * time.Minute).Truncate(time.Minute)

	return
}

/*
	run tasks
*/
func Start() {
	c.addJob(scrape)
	c.start()
}

/*
	stop task
*/
func Stop() {
	c.stop()
}
