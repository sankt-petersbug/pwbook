package commands

import (
	"github.com/sankt-petersbug/pwbook/pwbook"
)

// FakeStore fakes Store interface
type FakeStore struct {
	// CreateFunc is used to override Create function
	CreateFunc func(key string, value string) (pwbook.Entry, error)

	// ListFunc is used to override List function
	ListFunc func() ([]pwbook.Entry, error)

	// RemoveFunc is used to override Remove function
	RemoveFunc func(key string) error

	// UpdateFunc is used to override Update function
	UpdateFunc func(key string, value string) (pwbook.Entry, error)

	entries []pwbook.Entry
}

// Create an entry
func (s *FakeStore) Create(key string, value string) (pwbook.Entry, error) {
	if s.CreateFunc != nil {
		return s.CreateFunc(key, value)
	}

	entry := pwbook.Entry{
		Key:   key,
		Value: value,
	}

	return entry, nil
}

// List all entries
func (s *FakeStore) List() ([]pwbook.Entry, error) {
	if s.ListFunc != nil {
		return s.ListFunc()
	}

	return []pwbook.Entry{}, nil
}

// Remove an entry with matching key
func (s *FakeStore) Remove(key string) error {
	if s.RemoveFunc != nil {
		return s.RemoveFunc(key)
	}

	return nil
}

// Update an entry
func (s *FakeStore) Update(key string, value string) (pwbook.Entry, error) {
	if s.UpdateFunc != nil {
		return s.UpdateFunc(key, value)
	}

	entry := pwbook.Entry{
		Key:   key,
		Value: value,
	}

	return entry, nil
}
