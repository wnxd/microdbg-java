package internal

import (
	"reflect"
	_ "unsafe"

	java "github.com/wnxd/microdbg-java"
)

var BaseType = reflect.TypeFor[java.IObject]()

func ObjectPack(val any) java.IObject {
	return objectPack(val)
}

func ObjectUnpack(obj java.IObject) any {
	return objectUnpack(obj)
}

//go:linkname objectPack github.com/wnxd/microdbg-java/java/lang.objectPack
func objectPack(obj any) java.IObject

//go:linkname objectUnpack github.com/wnxd/microdbg-java/java/lang.objectUnpack
func objectUnpack(obj java.IObject) any
