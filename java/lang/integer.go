package lang

import (
	"reflect"
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type Integer java.JInt

type staticInteger struct {
	TYPE java.IClass
}

var (
	SInteger     staticInteger
	integerClass = ClassResolve[Integer](&SInteger)
)

func (i Integer) GetClass() java.IClass {
	return integerClass
}

func (i Integer) HashCode() java.JInt {
	return SInteger.HashCode(java.JInt(i))
}

func (i Integer) Equals(obj java.IObject) java.JBoolean {
	return i == obj
}

func (i Integer) ToString() java.IString {
	return SInteger.ToString(java.JInt(i))
}

func (i Integer) ByteValue() java.JByte {
	return java.JByte(i)
}

func (i Integer) ShortValue() java.JShort {
	return java.JShort(i)
}

func (i Integer) IntValue() java.JInt {
	return java.JInt(i)
}

func (i Integer) LongValue() java.JLong {
	return java.JLong(i)
}

func (i Integer) FloatValue() java.JFloat {
	return java.JFloat(i)
}

func (i Integer) DoubleValue() java.JDouble {
	return java.JDouble(i)
}

func (i Integer) CompareTo(another Integer) java.JInt {
	return SInteger.Compare(java.JInt(i), java.JInt(another))
}

func (staticInteger) Class() java.IClass {
	return integerClass
}

func (staticInteger) HashCode(value java.JInt) java.JInt {
	return value
}

func (staticInteger) ToString(i java.JInt) java.IString {
	return SInteger.ToString_1(i, 10)
}

func (staticInteger) ToString_1(i, radix java.JInt) java.IString {
	return String(strconv.FormatInt(int64(i), int(radix)))
}

func (staticInteger) Compare(x, y java.JInt) java.JInt {
	if x < y {
		return -1
	} else if x == y {
		return 0
	}
	return 1
}

func init() {
	typ := reflect.TypeFor[java.JInt]()
	SInteger.TYPE = createClass(getRType(typ), intType{typ}, nil)
	cls := integerClass.(*Class)
	cls.superClass = numberClass
	cls.interfaces = []java.IClass{ClassResolve[Comparable[Integer]](nil)}
}
