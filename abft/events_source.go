package abft

import (
	"github.com/ncog-earth/forest-base/hash"
	"github.com/ncog-earth/forest-base/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}
