package lang

import (
	"reflect"
	"strconv"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/java/io"
)

type Boolean java.JBoolean

type staticBoolean struct {
	TYPE java.IClass
}

var (
	SBoolean     staticBoolean
	booleanClass = ClassResolve[Boolean](&SBoolean)
)

func (b Boolean) GetClass() java.IClass {
	return booleanClass
}

func (b Boolean) HashCode() java.JInt {
	return SBoolean.HashCode(java.JBoolean(b))
}

func (b Boolean) Equals(obj java.IObject) java.JBoolean {
	return b == obj
}

func (b Boolean) ToString() java.IString {
	return String(strconv.FormatBool(bool(b)))
}

func (b Boolean) BooleanValue() java.JBoolean {
	return java.JBoolean(b)
}

func (staticBoolean) Class() java.IClass {
	return booleanClass
}

func (staticBoolean) HashCode(value java.JBoolean) java.JInt {
	if value {
		return 1231
	}
	return 1237
}

func init() {
	typ := reflect.TypeFor[java.JBoolean]()
	SBoolean.TYPE = createClass(getRType(typ), booleanType{typ}, nil)
	cls := booleanClass.(*Class)
	cls.interfaces = []java.IClass{io.SSerializable.Class(), ClassResolve[Comparable[Boolean]](nil)}
}
