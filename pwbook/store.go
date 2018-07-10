package pwbook

import (
	"errors"
	"time"

	"github.com/asdine/storm"
)

// Store represents the DataStore
type Store interface {
	Create(key string, value string) (Entry, error)
	List() ([]Entry, error)
	Remove(key string) error
	Update(key string, value string) (Entry, error)
}

// PWBookStore is an instance of Store
type PWBookStore struct {
	db *storm.DB
}

// Create an entry and save it to store
func (s *PWBookStore) Create(key string, value string) (Entry, error) {
	entry := Entry{}

	if err := s.db.One("Key", key, &entry); err == nil {
		return entry, errors.New("Already exists")
	}

	now := time.Now().UTC()
	entry.Key = key
	entry.Value = value
	entry.CreatedAt = now
	entry.ModifiedAt = now

	err := s.db.Save(&entry)

	return entry, err
}

// Update an entry's value
func (s *PWBookStore) Update(key string, value string) (Entry, error) {
	entry := Entry{Key: key, Value: value, ModifiedAt: time.Now()}
	err := s.db.Update(&entry)

	return entry, err
}

// List all stored entries
func (s *PWBookStore) List() ([]Entry, error) {
	var entries []Entry

	err := s.db.All(&entries)

	return entries, err
}

// Remove an entry from the store
func (s *PWBookStore) Remove(key string) error {
	entry := Entry{Key: key}

	err := s.db.DeleteStruct(&entry)

	return err
}

// Close internalDB
func (s *PWBookStore) Close() error {
	return s.db.Close()
}

// NewPWBookStore initialize a store with given filepath
func NewPWBookStore(path string) (PWBookStore, error) {
	s := PWBookStore{}

	db, err := storm.Open(path)
	if err != nil {
		return s, err
	}

	s.db = db

	return s, nil
}
