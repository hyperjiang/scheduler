package scheduler

import (
	"log"
	"time"
)

const defaultInterval = time.Second

// Runnable is a task that can be run
type Runnable func() error

// Scheduler runs jobs periodically
type Scheduler struct {
	Name     string
	quit     chan struct{}
	task     Runnable
	interval time.Duration
}

func (s *Scheduler) logPrefix() string {
	return "[scheduler][" + s.Name + "] "
}

// New creates a new Scheduler.
// Valid time units for interval are "ns", "us", "ms", "s", "m", "h"
func New(name string, task Runnable, interval string) *Scheduler {
	s := &Scheduler{
		Name: name,
	}
	s.quit = make(chan struct{})
	s.task = task

	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Printf(s.logPrefix()+"invalid interval string, set to %s by default\n", defaultInterval)
		d = defaultInterval
	}

	s.interval = d

	return s
}

// Start starts the scheduler
func (s *Scheduler) Start() {
	ticker := time.NewTicker(s.interval)

	log.Printf(s.logPrefix()+"started with interval: %s\n", s.interval)

	go func() {
		for {
			select {
			case <-s.quit:
				log.Println(s.logPrefix() + "quit")
				ticker.Stop()
				return
			case <-ticker.C:
				if err := s.task(); err != nil {
					log.Printf(s.logPrefix()+"run task error: %s\n", err.Error())
				}
			}
		}
	}()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
	s.quit <- struct{}{}
}
