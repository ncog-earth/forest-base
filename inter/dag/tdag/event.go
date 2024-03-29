package tdag

import (
	"github.com/ncog-earth/forest-base/hash"
	"github.com/ncog-earth/forest-base/inter/dag"
)

type TestEvent struct {
	dag.MutableBaseEvent
	Name string
}

func (e *TestEvent) AddParent(id hash.Event) {
	parents := e.Parents()
	parents.Add(id)
	e.SetParents(parents)
}
