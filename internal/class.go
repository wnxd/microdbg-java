package internal

import (
	"reflect"
	_ "unsafe"

	java "github.com/wnxd/microdbg-java"
)

type StaticClass interface {
	Class() java.IClass
}

func ClassResolve[T any](static StaticClass) java.IClass {
	typ := reflect.TypeFor[T]()
	return resolveClass(typ, static)
}

func This[T any](obj java.IObject) (v T) {
	v, _ = castObject(obj, obj.GetClass()).(T)
	return
}

//go:linkname resolveClass github.com/wnxd/microdbg-java/java/lang.resolveClass
func resolveClass(typ reflect.Type, static StaticClass) java.IClass

//go:linkname castObject github.com/wnxd/microdbg-java/java/lang.castObject
func castObject(obj java.IObject, cls java.IClass) any

//go:linkname NewArrayInstance github.com/wnxd/microdbg-java/java/lang.newArrayInstance
func NewArrayInstance(cls java.IClass, length int) java.IObject

//go:linkname TryCast github.com/wnxd/microdbg-java/java/lang.tryCast
func TryCast(val any, cls java.IClass) (any, bool)
