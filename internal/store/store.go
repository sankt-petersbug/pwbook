package store

import (
    "errors"
    "time"

    "github.com/asdine/storm"
)

type Entry struct {
    Key string `storm:"id"`
    Value string
    CreatedAt time.Time
    ModifiedAt time.Time
}

type Store struct {
    db *storm.DB
}

func (store Store) Create(key string, value string) (Entry, error) {
    entry := Entry{}

    if err := store.db.One("Key", key, &entry); err == nil {
        return entry, errors.New("Already exists")
    }

    now := time.Now()
    entry.Key = key
    entry.Value = value
    entry.CreatedAt = now
    entry.ModifiedAt = now

    err := store.db.Save(&entry)

    return entry, err
}

func (store Store) Update(key string, value string) (Entry, error) {
    entry := Entry{Key: key, Value: value, ModifiedAt: time.Now()}
    err := store.db.Update(&entry)

    return entry, err
}

func (store Store) List() ([]Entry, error) {
    var entries []Entry

    err := store.db.All(&entries)

    return entries, err
}

func (store Store) Close() error {
    return store.db.Close()
}

func NewStore(path string) (Store, error) {
    store := Store{}

    db, err := storm.Open(path)
    if err != nil {
        return store, err
    }

    store.db = db

    return store, nil
}