package counterrepo

import (
	tracking "interview1-assessment/internal/tracking"

	"sync"

	hyper "github.com/jmatom/hyper"
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

	hllInitial := hyper.New(uint32(18), true)
	hll, _ := c.visits.LoadOrStore(trackingEvent.Url().Hash(), hllInitial)
	hll.(*hyper.HyperLogLog).Add([]byte(trackingEvent.Uid().String()))

	return nil
}

func (c *CounterRepository) GetVisits(url tracking.Url) uint64 {
	hll, ok := c.visits.Load(url.Hash())
	if ok {
		return hll.(*hyper.HyperLogLog).Count()
	}

	return 0
}

func NewCounterRepository() *CounterRepository {
	return &CounterRepository{}
}
