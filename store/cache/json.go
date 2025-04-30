// Copyright (c) 2025 SoursopID
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cache

import (
	"encoding/json"
	"os"
	"path"
)

type JSONCache[V any] struct {
	folder string
}

func NewJSONCache[V any](folder string) Cache[string, V] {
	sq := JSONCache[V]{
		folder: folder,
	}

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0755)
	}

	return &sq
}

var _ Cache[string, any] = (*JSONCache[any])(nil)

func (s *JSONCache[V]) Get(key string) (V, error) {
	var value V
	filePath := path.Join(s.folder, key+".json")
	jb, err := os.ReadFile(filePath)
	if err != nil {
		return value, err
	}

	err = json.Unmarshal(jb, &value)
	return value, err
}

func (s *JSONCache[V]) Set(key string, value V) error {
	filePath := path.Join(s.folder, key+".json")
	jb, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jb, 0644)
	return err
}

func (s *JSONCache[V]) Delete(key string) error {
	return os.Remove(path.Join(s.folder, key+".json"))
}

func (s *JSONCache[V]) Clear() error {
	return os.RemoveAll(s.folder)
}
