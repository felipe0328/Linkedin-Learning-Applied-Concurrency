package stats

import (
	"appliedConcurrency/models"
	"sync"
)

type IResult interface {
	Get() models.Statistics
	Combine(stats models.Statistics)
}

type result struct {
	latest models.Statistics
	lock   sync.Mutex
}

func (r *result) Get() models.Statistics {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.latest
}

func (r *result) Combine(stats models.Statistics) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.latest = models.Combine(r.latest, stats)
}
