package utils

import (
	"fmt"

	"github.com/ncog-earth/forest-base/hash"
	"github.com/ncog-earth/forest-base/inter/idx"
)

// NameOf returns human readable string representation.
func NameOf(p idx.ValidatorID) string {
	if name := hash.GetNodeName(p); len(name) > 0 {
		return name
	}

	return fmt.Sprintf("%d", p)
}
