package lang

import (
	"reflect"

	java "github.com/wnxd/microdbg-java"
)

type Short java.JShort

type staticShort struct {
	TYPE java.IClass
}

var (
	SShort     staticShort
	shortClass = ClassResolve[Short](&SShort)
)

func (s Short) GetClass() java.IClass {
	return shortClass
}

func (s Short) HashCode() java.JInt {
	return SShort.HashCode(java.JShort(s))
}

func (s Short) Equals(obj java.IObject) java.JBoolean {
	return s == obj
}

func (s Short) ToString() java.IString {
	return SInteger.ToString(java.JInt(s))
}

func (s Short) ByteValue() java.JByte {
	return java.JByte(s)
}

func (s Short) ShortValue() java.JShort {
	return java.JShort(s)
}

func (s Short) IntValue() java.JInt {
	return java.JInt(s)
}

func (s Short) LongValue() java.JLong {
	return java.JLong(s)
}

func (s Short) FloatValue() java.JFloat {
	return java.JFloat(s)
}

func (s Short) DoubleValue() java.JDouble {
	return java.JDouble(s)
}

func (s Short) CompareTo(another Short) java.JInt {
	return SShort.Compare(java.JShort(s), java.JShort(another))
}

func (staticShort) Class() java.IClass {
	return shortClass
}

func (staticShort) HashCode(value java.JShort) java.JInt {
	return java.JInt(value)
}

func (staticShort) Compare(x, y java.JShort) java.JInt {
	return java.JInt(x - y)
}

func init() {
	typ := reflect.TypeFor[java.JShort]()
	SShort.TYPE = createClass(getRType(typ), shortType{typ}, nil)
	cls := shortClass.(*Class)
	cls.superClass = numberClass
	cls.interfaces = []java.IClass{ClassResolve[Comparable[Short]](nil)}
}
