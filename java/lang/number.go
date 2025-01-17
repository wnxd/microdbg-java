package lang

import (
	"reflect"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/java/io"
)

type INumber interface {
	ByteValue() java.JByte
	ShortValue() java.JShort
	IntValue() java.JInt
	LongValue() java.JLong
	FloatValue() java.JFloat
	DoubleValue() java.JDouble
}

type Number struct {
	Object
	io.Serializable
}

var numberClass = ClassResolve[Number](nil)

func (num *Number) ByteValue() java.JByte {
	panic("Not implemented")
}

func (num *Number) ShortValue() java.JShort {
	panic("Not implemented")
}

func (num *Number) IntValue() java.JInt {
	panic("Not implemented")
}

func (num *Number) LongValue() java.JLong {
	panic("Not implemented")
}

func (num *Number) FloatValue() java.JFloat {
	panic("Not implemented")
}

func (num *Number) DoubleValue() java.JDouble {
	panic("Not implemented")
}

func init() {
	cls := numberClass.(*Class)
	cls.ift = reflect.TypeFor[INumber]()
	aliasClass(cls.ift, cls.typ)
}
