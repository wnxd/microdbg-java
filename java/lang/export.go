package lang

import (
	"reflect"
	_ "unsafe"
)

//go:linkname Export github.com/wnxd/microdbg-java/internal.Export
var Export reflect.Type
