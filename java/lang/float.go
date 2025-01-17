package lang

import (
	"math"
	"reflect"
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type Float java.JFloat

type staticFloat struct {
	TYPE java.IClass
}

var (
	SFloat     staticFloat
	floatClass = ClassResolve[Float](&SFloat)
)

func (f Float) GetClass() java.IClass {
	return floatClass
}

func (f Float) HashCode() java.JInt {
	return SFloat.HashCode(java.JFloat(f))
}

func (f Float) Equals(obj java.IObject) java.JBoolean {
	return f == obj
}

func (f Float) ToString() java.IString {
	return SFloat.ToString(java.JFloat(f))
}

func (f Float) ByteValue() java.JByte {
	return java.JByte(f)
}

func (f Float) ShortValue() java.JShort {
	return java.JShort(f)
}

func (f Float) IntValue() java.JInt {
	return java.JInt(f)
}

func (f Float) LongValue() java.JLong {
	return java.JLong(f)
}

func (f Float) FloatValue() java.JFloat {
	return java.JFloat(f)
}

func (f Float) DoubleValue() java.JDouble {
	return java.JDouble(f)
}

func (f Float) CompareTo(another Float) java.JInt {
	return SFloat.Compare(java.JFloat(f), java.JFloat(another))
}

func (staticFloat) Class() java.IClass {
	return floatClass
}

func (staticFloat) HashCode(value java.JFloat) java.JInt {
	return java.JInt(math.Float32bits(value))
}

func (staticFloat) ToString(value java.JFloat) java.IString {
	return String(strconv.FormatFloat(float64(value), 'f', -1, 32))
}

func (staticFloat) Compare(f1, f2 java.JFloat) java.JInt {
	if f1 < f2 {
		return -1
	} else if f1 > f2 {
		return 1
	}
	b1 := math.Float32bits(f1)
	b2 := math.Float32bits(f2)
	if b1 == b2 {
		return 0
	} else if b1 < b2 {
		return -1
	}
	return 1
}

func init() {
	typ := reflect.TypeFor[java.JFloat]()
	SFloat.TYPE = createClass(getRType(typ), floatType{typ}, nil)
	cls := floatClass.(*Class)
	cls.superClass = numberClass
	cls.interfaces = []java.IClass{ClassResolve[Comparable[Float]](nil)}
}
