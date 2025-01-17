package lang

import (
	"reflect"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/java/io"
)

type Character java.JChar

type staticCharacter struct {
	TYPE java.IClass
}

var (
	SCharacter     staticCharacter
	characterClass = ClassResolve[Character](&SCharacter)
)

func (c Character) GetClass() java.IClass {
	return characterClass
}

func (c Character) HashCode() java.JInt {
	return SCharacter.HashCode(java.JChar(c))
}

func (c Character) Equals(obj java.IObject) java.JBoolean {
	return c == obj
}

func (c Character) ToString() java.IString {
	return SString.ValueOf_4(java.JChar(c))
}

func (c Character) CompareTo(another Character) java.JInt {
	return SCharacter.Compare(java.JChar(c), java.JChar(another))
}

func (staticCharacter) Class() java.IClass {
	return characterClass
}

func (staticCharacter) HashCode(value java.JChar) java.JInt {
	return java.JInt(value)
}

func (staticCharacter) Compare(x, y java.JChar) java.JInt {
	return java.JInt(x - y)
}

func init() {
	typ := reflect.TypeFor[java.JChar]()
	SCharacter.TYPE = createClass(getRType(typ), charType{typ}, nil)
	cls := characterClass.(*Class)
	cls.interfaces = []java.IClass{io.SSerializable.Class(), ClassResolve[Comparable[Character]](nil)}
}
