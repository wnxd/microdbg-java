package lang

import (
	"unsafe"
)

type UncommonType struct {
	PkgPath uint32
	Mcount  uint16
	Xcount  uint16
	Moff    uint32
	_       uint32
}

func exportMethods(rtype uintptr) {
	ut := uncommon(rtype)
	if ptr := (*UncommonType)(ut); ptr.Xcount != ptr.Mcount {
		patchExportMethods(ut)
	}
}

//go:linkname uncommon internal/abi.(*Type).Uncommon
func uncommon(uintptr) unsafe.Pointer
