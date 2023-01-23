package store

import (
	"fmt"
	"sync"

	"github.com/igor-baiborodine/distributed-services-with-go-workshop/SecureYourServices/AuthenticateServerWithTLS/internal/model"
	"golang.org/x/exp/slices"
)

type BookingStore struct {
	mu       sync.Mutex
	bookings []model.Booking
}

func NewBookingStore() (*BookingStore, error) {
	return &BookingStore{}, nil
}

func (c *BookingStore) GetByUUID(uuid string) (model.Booking, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	idx := slices.IndexFunc(c.bookings, func(b model.Booking) bool { return b.UUID == uuid })
	if idx == -1 {
		return model.Booking{}, fmt.Errorf("booking not found")
	}
	return c.bookings[idx], nil
}

func (c *BookingStore) Create(b model.Booking) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bookings = append(c.bookings, b)
	return nil
}
