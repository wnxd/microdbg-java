package lang

import (
	"reflect"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/java/io"
)

type String string

type staticString struct {
}

var (
	SString     staticString
	stringClass = ClassResolve[String](&SString)
)

func (str String) GetClass() java.IClass {
	return stringClass
}

func (str String) HashCode() java.JInt {
	ptr := (*struct{ rtype, data unsafe.Pointer })(unsafe.Pointer((&str))).data
	return int32(uintptr(ptr))
}

func (str String) Equals(obj java.IObject) java.JBoolean {
	return str == obj
}

func (str String) ToString() java.IString {
	return str
}

func (str String) CompareTo(another java.IString) java.JInt {
	return java.JInt(strings.Compare(str.String(), another.String()))
}

func (str String) Length() java.JInt {
	return java.JInt(len(str))
}

func (str String) CharAt(index java.JInt) java.JChar {
	var runeIndex java.JInt
	for i, w := 0, len(str); i < w; i += utf8.RuneLen(rune(str[i])) {
		if runeIndex == index {
			r, _ := utf8.DecodeRuneInString(string(str[i:]))
			return java.JChar(r)
		}
		runeIndex++
	}
	return 0
}

func (str String) SubSequence(start, end java.JInt) CharSequence {
	return str.Substring(start, end).(CharSequence)
}

func (str String) Substring(beginIndex, endIndex java.JInt) java.IString {
	return String([]rune(str)[beginIndex:endIndex])
}

func (str String) GetBytes() []java.JByte {
	return unsafe.Slice((*java.JByte)(unsafe.Pointer(unsafe.StringData(string(str)))), len(str))
}

func (str String) ToCharArray() []java.JChar {
	return utf16.Encode([]rune(str))
}

func (str String) String() string {
	return string(str)
}

func (staticString) Class() java.IClass {
	return stringClass
}

func (staticString) String(original java.IString) java.IString {
	return String(original.String())
}

func (staticString) String_1(value []java.JChar) java.IString {
	return SString.String_2(value, 0, java.JInt(len(value)))
}

func (staticString) String_2(value []java.JChar, offset, count java.JInt) java.IString {
	if len := java.JInt(len(value)); offset < 0 || offset >= len || count < 0 {
		return nil
	} else if end := offset + count; end > len {
		return nil
	} else {
		value = value[offset:end]
	}
	return String(utf16.Decode(value))
}

func (staticString) ValueOf(obj java.IObject) java.IString {
	if obj == nil {
		return String("null")
	}
	return obj.ToString()
}

func (staticString) ValueOf_1(data []java.JChar) java.IString {
	return SString.String_1(data)
}

func (staticString) ValueOf_2(data []java.JChar, offset, count java.JInt) java.IString {
	return SString.String_2(data, offset, count)
}

func (staticString) ValueOf_3(b java.JBoolean) java.IString {
	if b {
		return String("true")
	}
	return String("false")
}

func (staticString) ValueOf_4(c java.JChar) java.IString {
	return String(rune(c))
}

func (staticString) ValueOf_5(i java.JInt) java.IString {
	return SInteger.ToString(i)
}

func (staticString) ValueOf_6(l java.JLong) java.IString {
	return SLong.ToString(l)
}

func (staticString) ValueOf_7(f java.JFloat) java.IString {
	return SFloat.ToString(f)
}

func (staticString) ValueOf_8(d java.JDouble) java.IString {
	return SDouble.ToString(d)
}

func stringPack(str string) java.IString {
	return String(str)
}

func init() {
	cls := stringClass.(*Class)
	cls.ift = reflect.TypeFor[java.IString]()
	cls.interfaces = []java.IClass{io.SSerializable.Class(), ClassResolve[Comparable[java.IString]](nil), charSequenceClass}
	aliasClass(cls.ift, cls.typ)
}
