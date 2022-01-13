package uuid

import (
	"github.com/rogpeppe/fastuuid"
)

var fuuid *fastuuid.Generator = fastuuid.MustNewGenerator()

// New generates an UUID using fastuuid
func New() string {
	return fuuid.Hex128()
}
