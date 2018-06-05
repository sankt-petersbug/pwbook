/*
 Copyright © 2018 Sankt Petersbug <sankt.petersbug@gmail.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http:www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package store

import (
    "errors"
    "time"

    "github.com/asdine/storm"
)

// Entry holds information about name, password pair
type Entry struct {
    Key string `storm:"id"`
    Value string
    CreatedAt time.Time
    ModifiedAt time.Time
}

// ModifiedSince returns days elapsed since modifiedAt
func (e *Entry) ModifiedSince() int {
    d := time.Since(e.ModifiedAt)
    return int(d.Hours() / 24)
}

// Store provides APIs to interact with database
type Store struct {
    db *storm.DB
}

// Create create an entry and save it to store
func (s *Store) Create(key string, value string) (Entry, error) {
    entry := Entry{}

    if err := s.db.One("Key", key, &entry); err == nil {
        return entry, errors.New("Already exists")
    }

    now := time.Now()
    entry.Key = key
    entry.Value = value
    entry.CreatedAt = now
    entry.ModifiedAt = now

    err := s.db.Save(&entry)

    return entry, err
}

// Update updates an entry's value
func (s *Store) Update(key string, value string) (Entry, error) {
    entry := Entry{Key: key, Value: value, ModifiedAt: time.Now()}
    err := s.db.Update(&entry)

    return entry, err
}

// List returns all stored entries
func (s *Store) List() ([]Entry, error) {
    var entries []Entry

    err := s.db.All(&entries)

    return entries, err
}

// Delete deltes an entry from the store
func (s *Store) Delete(key string) error {
    entry := Entry{Key: key}

    err := s.db.DeleteStruct(&entry)

    return err
}

// Close closes internalDB
func (s *Store) Close() error {
    return s.db.Close()
}

// NewStore creates a new store with given filepath
func NewStore(path string) (Store, error) {
    s := Store{}

    db, err := storm.Open(path)
    if err != nil {
        return s, err
    }

    s.db = db

    return s, nil
}