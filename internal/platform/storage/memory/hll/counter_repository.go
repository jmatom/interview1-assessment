package counterrepo

import (
	"fmt"
	tracking "interview1-assessment/internal/tracking"

	"sync"

	"github.com/axiomhq/hyperloglog"
)

type CounterRepository struct {
	// visits map[string]*hyperloglog.Sketch
	visits sync.Map
}

// implement interface counter
// addVisit(trackingEvent tracking.TrackingEvent) error
// getVisits(url tracking.Url)
func (c *CounterRepository) AddVisit(trackingEvent tracking.TrackingEvent) error {
	// For the given url, get the hll
	// add user uid
	hllInitial := hyperloglog.New16()
	hll, loaded := c.visits.LoadOrStore(trackingEvent.Url().Hash(), hllInitial)
	if loaded {
		fmt.Printf("loaded is true! insert %s", trackingEvent.Uid().String())
		hll.(*hyperloglog.Sketch).Insert([]byte(trackingEvent.Uid().String()))
		c.visits.Store(trackingEvent.Url().Hash(), hll)
	} else {
		fmt.Printf("loaded is false! insert %s", trackingEvent.Uid().String())
		hllInitial.Insert([]byte(trackingEvent.Uid().String()))
		c.visits.Store(trackingEvent.Url().Hash(), hllInitial)
	}

	return nil
}

func (c *CounterRepository) GetVisits(url tracking.Url) uint64 {
	hll, ok := c.visits.Load(url.Hash())
	if ok {
		return hll.(*hyperloglog.Sketch).Estimate()
	}

	return 0
}

func NewCounterRepository() *CounterRepository {
	return &CounterRepository{}
}
