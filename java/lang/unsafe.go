package lang

import "unsafe"

//go:linkname unsafe_New reflect.unsafe_New
func unsafe_New(uintptr) unsafe.Pointer

//go:linkname unsafe_NewArray reflect.unsafe_NewArray
func unsafe_NewArray(uintptr, int) unsafe.Pointer

//go:linkname ifaceE2I reflect.ifaceE2I
func ifaceE2I(uintptr, any, unsafe.Pointer)
