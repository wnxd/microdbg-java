package lang

import (
	"strconv"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type Object struct {
	cls java.IClass
}

type staticObject struct {
}

var (
	SObject     staticObject
	objectClass = ClassResolve[*Object](&SObject)
)

func (obj *Object) GetClass() java.IClass {
	return obj.cls
}

func (obj *Object) HashCode() java.JInt {
	return java.JInt(uintptr(unsafe.Pointer(obj)))
}

func (obj *Object) Equals(o java.IObject) java.JBoolean {
	return obj == o
}

func (obj *Object) Clone() java.JObject {
	panic("Not implemented")
}

func (obj *Object) ToString() java.IString {
	return String(obj.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(obj.HashCode()), 16))
}

func (obj *Object) instance(cls java.IClass) {
	obj.cls = cls
}

func (staticObject) Class() java.IClass {
	return objectClass
}

func objectPack(obj any) java.IObject {
	if obj == nil {
		return nil
	} else if v, ok := obj.(java.IObject); ok {
		return v
	} else if v = primitivePack(obj); v != nil {
		return v
	}
	return arrayPack(obj)
}

func objectUnpack(obj java.IObject) any {
	if obj == nil {
		return nil
	} else if v := primitiveUnpack(obj); v != nil {
		return v
	}
	return arrayUnpack(obj)
}

func init() {
	cls := objectClass.(*Class)
	cls.ift = internal.BaseType
	aliasClass(cls.ift, cls.typ)
}
