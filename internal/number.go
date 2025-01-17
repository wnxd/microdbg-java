package internal

import java "github.com/wnxd/microdbg-java"

type Number interface {
	ByteValue() java.JByte
	ShortValue() java.JShort
	IntValue() java.JInt
	LongValue() java.JLong
	FloatValue() java.JFloat
	DoubleValue() java.JDouble
}
