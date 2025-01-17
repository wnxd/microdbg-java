package lang

import (
	"reflect"
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type Long java.JLong

type staticLong struct {
	TYPE java.IClass
}

var (
	SLong     staticLong
	longClass = ClassResolve[Long](&SLong)
)

func (l Long) GetClass() java.IClass {
	return longClass
}

func (l Long) HashCode() java.JInt {
	return SLong.HashCode(java.JLong(l))
}

func (l Long) Equals(obj java.IObject) java.JBoolean {
	return l == obj
}

func (l Long) ToString() java.IString {
	return SLong.ToString(java.JLong(l))
}

func (l Long) ByteValue() java.JByte {
	return java.JByte(l)
}

func (l Long) ShortValue() java.JShort {
	return java.JShort(l)
}

func (l Long) IntValue() java.JInt {
	return java.JInt(l)
}

func (l Long) LongValue() java.JLong {
	return java.JLong(l)
}

func (l Long) FloatValue() java.JFloat {
	return java.JFloat(l)
}

func (l Long) DoubleValue() java.JDouble {
	return java.JDouble(l)
}

func (l Long) CompareTo(another Long) java.JInt {
	return SLong.Compare(java.JLong(l), java.JLong(another))
}

func (staticLong) Class() java.IClass {
	return longClass
}

func (staticLong) HashCode(value java.JLong) java.JInt {
	v := uint64(value)
	return java.JInt(v ^ (v >> 32))
}

func (staticLong) ToString(i java.JLong) java.IString {
	return SLong.ToString_1(i, 10)
}

func (staticLong) ToString_1(i java.JLong, radix java.JInt) java.IString {
	return String(strconv.FormatInt(i, int(radix)))
}

func (staticLong) Compare(x, y java.JLong) java.JInt {
	if x < y {
		return -1
	} else if x == y {
		return 0
	}
	return 1
}

func init() {
	typ := reflect.TypeFor[java.JLong]()
	SLong.TYPE = createClass(getRType(typ), longType{typ}, nil)
	cls := longClass.(*Class)
	cls.superClass = numberClass
	cls.interfaces = []java.IClass{ClassResolve[Comparable[Long]](nil)}
}
