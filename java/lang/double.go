package lang

import (
	"math"
	"reflect"
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type Double java.JDouble

type staticDouble struct {
	TYPE java.IClass
}

var (
	SDouble     staticDouble
	doubleClass = ClassResolve[Double](&SDouble)
)

func (d Double) GetClass() java.IClass {
	return doubleClass
}

func (d Double) HashCode() java.JInt {
	return SDouble.HashCode(java.JDouble(d))
}

func (d Double) Equals(obj java.IObject) java.JBoolean {
	return d == obj
}

func (d Double) ToString() java.IString {
	return SDouble.ToString(java.JDouble(d))
}

func (d Double) ByteValue() java.JByte {
	return java.JByte(d)
}

func (d Double) ShortValue() java.JShort {
	return java.JShort(d)
}

func (d Double) IntValue() java.JInt {
	return java.JInt(d)
}

func (d Double) LongValue() java.JLong {
	return java.JLong(d)
}

func (d Double) FloatValue() java.JFloat {
	return java.JFloat(d)
}

func (d Double) DoubleValue() java.JDouble {
	return java.JDouble(d)
}

func (d Double) CompareTo(another Double) java.JInt {
	return SDouble.Compare(java.JDouble(d), java.JDouble(another))
}

func (staticDouble) Class() java.IClass {
	return doubleClass
}

func (staticDouble) HashCode(value java.JDouble) java.JInt {
	bits := math.Float64bits(value)
	return (java.JInt)(bits ^ (bits >> 32))
}

func (staticDouble) ToString(value java.JDouble) String {
	return String(strconv.FormatFloat(float64(value), 'f', -1, 64))
}

func (staticDouble) Compare(d1, d2 java.JDouble) java.JInt {
	if d1 < d2 {
		return -1
	} else if d1 > d2 {
		return 1
	}
	b1 := math.Float64bits(d1)
	b2 := math.Float64bits(d2)
	if b1 == b2 {
		return 0
	} else if b1 < b2 {
		return -1
	}
	return 1
}

func init() {
	typ := reflect.TypeFor[java.JDouble]()
	SDouble.TYPE = createClass(getRType(typ), doubleType{typ}, nil)
	cls := doubleClass.(*Class)
	cls.superClass = numberClass
	cls.interfaces = []java.IClass{ClassResolve[Comparable[Double]](nil)}
}
