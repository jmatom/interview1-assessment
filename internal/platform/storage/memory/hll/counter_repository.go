package counterrepo

import (
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
	hll, _ := c.visits.LoadOrStore(trackingEvent.Url().Hash(), hllInitial)
	hll.(*hyperloglog.Sketch).Insert([]byte(trackingEvent.Uid().String()))

	return nil
}

func (c *CounterRepository) GetVisits(url tracking.Url) uint64 {
	// sfmt.Printf("getVisits: %s\n", url.Hash())
	hll, ok := c.visits.Load(url.Hash())
	if ok {
		return hll.(*hyperloglog.Sketch).Estimate()
	}

	return 0
}

func NewCounterRepository() *CounterRepository {
	return &CounterRepository{}
}
