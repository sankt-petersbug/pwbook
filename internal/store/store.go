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

func (store Store) Delete(key string) error {
    entry := Entry{Key: key}

    err := store.db.DeleteStruct(&entry)

    return err
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