package stats

import "appliedConcurrency/models"

const workerCount = 3

type IStatsService interface {
	GetStats() models.Statistics
}

type statsService struct {
	result    IResult
	processed <-chan models.Order
	done      <-chan struct{}
	pStats    chan models.Statistics
}

func NewStats(processed <-chan models.Order, done <-chan struct{}) IStatsService {
	s := statsService{
		result:    &result{},
		processed: processed,
		done:      done,
		pStats:    make(chan models.Statistics, workerCount),
	}

	for i := 0; i < workerCount; i++ {
		go s.processStats()
	}

	go s.reconcile()

	return &s
}

func (s *statsService) GetStats() models.Statistics {
	return s.result.Get()
}

func (s *statsService) processStats() {
	for {
		select {
		case newOrder := <-s.processed:
			newStat := s.processOrder(newOrder)
			s.pStats <- newStat
		case <-s.done:
			return
		}
	}
}

func (s *statsService) processOrder(order models.Order) models.Statistics {
	if order.Status == models.OrderStatus_Completed {
		return models.Statistics{
			CompletedOrders: 1,
			Revenue:         *order.Total,
		}
	}

	return models.Statistics{
		RejectedOrders: 1,
	}
}

func (s *statsService) reconcile() {
	for {
		select {
		case stat := <-s.pStats:
			s.result.Combine(stat)
		case <-s.done:
			return
		}
	}
}
