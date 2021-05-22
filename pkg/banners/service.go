package banners

import (
	"github.com/Ulugbek999/http.git/pkg/types"
	"context"
	"errors"
	"sync"

)

var count int64

//Service - is a banner management service.
type Service struct {
	mu    sync.RWMutex
	items []*types.Banner
}

//newService - create a service.
func NewService() *Service {
	return &Service{items: make([]*types.Banner, 0)}
}

//All - returns all existing banners.
func (s *Service) All(ctx context.Context) ([]*types.Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

//ByID - returns banner by ID.
func (s *Service) ByID(ctx context.Context, id int64) (*types.Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}
	return nil, errors.New("item not found")
}

//Save - saves/updates banner.
func (s *Service) Save(ctx context.Context, item *types.Banner) (*types.Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// var count int64
	if item.ID == 0 {
		count++
		item.ID = count
		s.items = append(s.items, item)
		return item, nil
	}

	for i, banner := range s.items {
		if banner.ID == item.ID {
			s.items[i] = item
			return item, nil
		}
	}
	return nil, errors.New("item not found")
}

//RemoveByID - removes banner by ID.
func (s *Service) RemoveByID(ctx context.Context, id int64) (*types.Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, banner := range s.items {
		if banner.ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return banner, nil
		}
	}
	return nil, errors.New("item not found")
}
