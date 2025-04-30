// Copyright (c) 2025 SoursopID
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cache

import (
	"database/sql"
	"encoding/json"
)

type SQLCache[V any] struct {
	db *sql.DB
}

func NewSQLCache[V any](db *sql.DB) Cache[string, V] {
	sq := SQLCache[V]{
		db: db,
	}
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cache (key TEXT PRIMARY KEY, value BLOB)")
	if err != nil {
		panic(err)
	}
	return &sq
}

var _ Cache[string, any] = (*SQLCache[any])(nil)

func (s *SQLCache[V]) Get(key string) (V, error) {
	var value V
	var jb []byte
	err := s.db.QueryRow("SELECT value FROM cache WHERE key = ?", key).Scan(&jb)
	if err != nil {
		return value, err
	}
	err = json.Unmarshal(jb, &value)
	return value, err
}

func (s *SQLCache[V]) Set(key string, value V) error {
	jb, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("INSERT OR REPLACE INTO cache (key, value) VALUES (?, ?)", key, jb)
	return err
}

func (s *SQLCache[V]) Delete(key string) error {
	_, err := s.db.Exec("DELETE FROM cache WHERE key = ?", key)
	return err
}

func (s *SQLCache[V]) Clear() error {
	_, err := s.db.Exec("DELETE FROM cache")
	return err
}
