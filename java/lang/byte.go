package lang

import (
	"reflect"

	java "github.com/wnxd/microdbg-java"
)

type Byte java.JByte

type staticByte struct {
	TYPE java.IClass
}

var (
	SByte     staticByte
	byteClass = ClassResolve[Byte](&SByte)
)

func (b Byte) GetClass() java.IClass {
	return byteClass
}

func (b Byte) HashCode() java.JInt {
	return SByte.HashCode(java.JByte(b))
}

func (b Byte) Equals(obj java.IObject) java.JBoolean {
	return b == obj
}

func (b Byte) ToString() java.IString {
	return SInteger.ToString(java.JInt(b))
}

func (b Byte) ByteValue() java.JByte {
	return java.JByte(b)
}

func (b Byte) ShortValue() java.JShort {
	return java.JShort(b)
}

func (b Byte) IntValue() java.JInt {
	return java.JInt(b)
}

func (b Byte) LongValue() java.JLong {
	return java.JLong(b)
}

func (b Byte) FloatValue() java.JFloat {
	return java.JFloat(b)
}

func (b Byte) DoubleValue() java.JDouble {
	return java.JDouble(b)
}

func (b Byte) CompareTo(another Byte) java.JInt {
	return SByte.Compare(java.JByte(b), java.JByte(another))
}

func (staticByte) Class() java.IClass {
	return byteClass
}

func (staticByte) HashCode(value java.JByte) java.JInt {
	return java.JInt(value)
}

func (staticByte) Compare(x, y java.JByte) java.JInt {
	return java.JInt(x - y)
}

func init() {
	typ := reflect.TypeFor[java.JByte]()
	SByte.TYPE = createClass(getRType(typ), byteType{typ}, nil)
	cls := byteClass.(*Class)
	cls.superClass = numberClass
	cls.interfaces = []java.IClass{ClassResolve[Comparable[Byte]](nil)}
}
