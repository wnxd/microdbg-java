package reflect

import (
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type Array struct {
	java.IObject
}

type staticArray struct {
}

var (
	SArray     staticArray
	arrayClass = internal.ClassResolve[Array](&SArray)
)

func (staticArray) Class() java.IClass {
	return arrayClass
}

func (staticArray) NewInstance(componentType java.IClass, length java.JInt) java.IObject {
	return internal.NewArrayInstance(componentType, int(length))
}

func (staticArray) Get(array java.IObject, index java.JInt) java.IObject {
	if arr, ok := array.(java.IGenericArray[java.IObject]); ok {
		return arr.Get(index)
	}
	return nil
}

func (staticArray) GetBoolean(array java.IObject, index java.JInt) java.JBoolean {
	if arr, ok := array.(java.IGenericArray[java.JBoolean]); ok {
		return arr.Get(index)
	}
	return false
}

func (staticArray) GetByte(array java.IObject, index java.JInt) java.JByte {
	if arr, ok := array.(java.IGenericArray[java.JByte]); ok {
		return arr.Get(index)
	}
	return 0
}

func (staticArray) GetChar(array java.IObject, index java.JInt) java.JChar {
	if arr, ok := array.(java.IGenericArray[java.JChar]); ok {
		return arr.Get(index)
	}
	return 0
}

func (staticArray) GetShort(array java.IObject, index java.JInt) java.JShort {
	switch arr := array.(type) {
	case java.IGenericArray[java.JByte]:
		return java.JShort(arr.Get(index))
	case java.IGenericArray[java.JShort]:
		return arr.Get(index)
	}
	return 0
}

func (staticArray) GetInt(array java.IObject, index java.JInt) java.JInt {
	switch arr := array.(type) {
	case java.IGenericArray[java.JByte]:
		return java.JInt(arr.Get(index))
	case java.IGenericArray[java.JChar]:
		return java.JInt(arr.Get(index))
	case java.IGenericArray[java.JShort]:
		return java.JInt(arr.Get(index))
	case java.IGenericArray[java.JInt]:
		return arr.Get(index)
	}
	return 0
}

func (staticArray) GetLong(array java.IObject, index java.JInt) java.JLong {
	switch arr := array.(type) {
	case java.IGenericArray[java.JByte]:
		return java.JLong(arr.Get(index))
	case java.IGenericArray[java.JChar]:
		return java.JLong(arr.Get(index))
	case java.IGenericArray[java.JShort]:
		return java.JLong(arr.Get(index))
	case java.IGenericArray[java.JInt]:
		return java.JLong(arr.Get(index))
	case java.IGenericArray[java.JLong]:
		return arr.Get(index)
	}
	return 0
}

func (staticArray) GetFloat(array java.IObject, index java.JInt) java.JFloat {
	switch arr := array.(type) {
	case java.IGenericArray[java.JByte]:
		return java.JFloat(arr.Get(index))
	case java.IGenericArray[java.JChar]:
		return java.JFloat(arr.Get(index))
	case java.IGenericArray[java.JShort]:
		return java.JFloat(arr.Get(index))
	case java.IGenericArray[java.JInt]:
		return java.JFloat(arr.Get(index))
	case java.IGenericArray[java.JLong]:
		return java.JFloat(arr.Get(index))
	case java.IGenericArray[java.JFloat]:
		return arr.Get(index)
	}
	return 0
}

func (staticArray) GetDouble(array java.IObject, index java.JInt) java.JDouble {
	switch arr := array.(type) {
	case java.IGenericArray[java.JByte]:
		return java.JDouble(arr.Get(index))
	case java.IGenericArray[java.JChar]:
		return java.JDouble(arr.Get(index))
	case java.IGenericArray[java.JShort]:
		return java.JDouble(arr.Get(index))
	case java.IGenericArray[java.JInt]:
		return java.JDouble(arr.Get(index))
	case java.IGenericArray[java.JLong]:
		return java.JDouble(arr.Get(index))
	case java.IGenericArray[java.JFloat]:
		return java.JDouble(arr.Get(index))
	case java.IGenericArray[java.JDouble]:
		return arr.Get(index)
	}
	return 0
}

func (staticArray) Set(array java.IObject, index java.JInt, value java.IObject) {
	if arr, ok := array.(java.IGenericArray[java.IObject]); ok {
		arr.Set(index, value)
	}
}

func (staticArray) SetBoolean(array java.IObject, index java.JInt, z java.JBoolean) {
	if arr, ok := array.(java.IGenericArray[java.JBoolean]); ok {
		arr.Set(index, z)
	}
}

func (staticArray) SetByte(array java.IObject, index java.JInt, b java.JByte) {
	switch arr := array.(type) {
	case java.IGenericArray[java.JByte]:
		arr.Set(index, b)
	case java.IGenericArray[java.JShort]:
		arr.Set(index, java.JShort(b))
	case java.IGenericArray[java.JInt]:
		arr.Set(index, java.JInt(b))
	case java.IGenericArray[java.JLong]:
		arr.Set(index, java.JLong(b))
	case java.IGenericArray[java.JFloat]:
		arr.Set(index, java.JFloat(b))
	case java.IGenericArray[java.JDouble]:
		arr.Set(index, java.JDouble(b))
	}
}

func (staticArray) SetChar(array java.IObject, index java.JInt, c java.JChar) {
	switch arr := array.(type) {
	case java.IGenericArray[java.JChar]:
		arr.Set(index, c)
	case java.IGenericArray[java.JInt]:
		arr.Set(index, java.JInt(c))
	case java.IGenericArray[java.JLong]:
		arr.Set(index, java.JLong(c))
	case java.IGenericArray[java.JFloat]:
		arr.Set(index, java.JFloat(c))
	case java.IGenericArray[java.JDouble]:
		arr.Set(index, java.JDouble(c))
	}
}

func (staticArray) SetShort(array java.IObject, index java.JInt, s java.JShort) {
	switch arr := array.(type) {
	case java.IGenericArray[java.JShort]:
		arr.Set(index, s)
	case java.IGenericArray[java.JInt]:
		arr.Set(index, java.JInt(s))
	case java.IGenericArray[java.JLong]:
		arr.Set(index, java.JLong(s))
	case java.IGenericArray[java.JFloat]:
		arr.Set(index, java.JFloat(s))
	case java.IGenericArray[java.JDouble]:
		arr.Set(index, java.JDouble(s))
	}
}

func (staticArray) SetInt(array java.IObject, index java.JInt, i java.JInt) {
	switch arr := array.(type) {
	case java.IGenericArray[java.JInt]:
		arr.Set(index, i)
	case java.IGenericArray[java.JLong]:
		arr.Set(index, java.JLong(i))
	case java.IGenericArray[java.JFloat]:
		arr.Set(index, java.JFloat(i))
	case java.IGenericArray[java.JDouble]:
		arr.Set(index, java.JDouble(i))
	}
}

func (staticArray) SetLong(array java.IObject, index java.JInt, l java.JLong) {
	switch arr := array.(type) {
	case java.IGenericArray[java.JLong]:
		arr.Set(index, l)
	case java.IGenericArray[java.JFloat]:
		arr.Set(index, java.JFloat(l))
	case java.IGenericArray[java.JDouble]:
		arr.Set(index, java.JDouble(l))
	}
}

func (staticArray) SetFloat(array java.IObject, index java.JInt, f java.JFloat) {
	switch arr := array.(type) {
	case java.IGenericArray[java.JFloat]:
		arr.Set(index, f)
	case java.IGenericArray[java.JDouble]:
		arr.Set(index, java.JDouble(f))
	}
}

func (staticArray) SetDouble(array java.IObject, index java.JInt, d java.JDouble) {
	if arr, ok := array.(java.IGenericArray[java.JDouble]); ok {
		arr.Set(index, d)
	}
}
