// Copyright (c) 2025 SoursopID
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cache

import (
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Cache[K comparable, V any] interface {
	Get(K) (V, error)
	Set(K, V) error
	Delete(K) error
	Clear() error
}

var DefaultMessageCache Cache[string, *events.Message]
var DefaultSentCache Cache[string, *waE2E.Message]
var DefaultGroupInfoCache Cache[string, *types.GroupInfo]
