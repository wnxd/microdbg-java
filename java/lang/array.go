package lang

import (
	"reflect"
	"slices"
	"strconv"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/java/io"
)

type primitiveArray[V comparable] []V

type objectArray struct {
	cls   java.IClass
	ecls  java.IClass
	elem  []java.IObject
	rtype uintptr
}

type staticArray[V comparable] struct {
}

var (
	arrayInterfaces   []java.IClass
	booleanArrayClass *Class
	charArrayClass    *Class
	byteArrayClass    *Class
	shortArrayClass   *Class
	intArrayClass     *Class
	longArrayClass    *Class
	floatArrayClass   *Class
	doubleArrayClass  *Class
)

func SArray[V comparable]() staticArray[V] {
	return staticArray[V]{}
}

func (arr primitiveArray[V]) GetClass() java.IClass {
	return SArray[V]().Class()
}

func (arr primitiveArray[V]) HashCode() java.JInt {
	ptr := (*struct{ data unsafe.Pointer })(unsafe.Pointer((&arr))).data
	return int32(uintptr(ptr))
}

func (arr primitiveArray[V]) Equals(obj java.IObject) java.JBoolean {
	if arr, ok := obj.(primitiveArray[V]); ok {
		return java.JBoolean(slices.Equal(arr, arr))
	}
	return false
}

func (arr primitiveArray[V]) ToString() java.IString {
	return String(arr.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(arr.HashCode()), 16))
}

func (arr primitiveArray[V]) Length() java.JInt {
	return java.JInt(len(arr))
}

func (arr primitiveArray[V]) Get(index java.JInt) V {
	return arr[index]
}

func (arr primitiveArray[V]) Set(index java.JInt, value V) {
	arr[index] = value
}

func (arr primitiveArray[V]) Elements() []V {
	return arr
}

func (arr *objectArray) GetClass() java.IClass {
	return arr.cls
}

func (arr *objectArray) HashCode() java.JInt {
	return java.JInt(uintptr(unsafe.Pointer(arr)))
}

func (arr *objectArray) Equals(obj java.IObject) java.JBoolean {
	if arr, ok := obj.(*objectArray); ok {
		return java.JBoolean(slices.Equal(arr.elem, arr.elem))
	}
	return false
}

func (arr *objectArray) ToString() java.IString {
	return String(arr.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(arr.HashCode()), 16))
}

func (arr *objectArray) Length() java.JInt {
	return java.JInt(len(arr.elem))
}

func (arr *objectArray) Get(index java.JInt) java.IObject {
	return arr.elem[index]
}

func (arr *objectArray) Set(index java.JInt, value java.IObject) {
	if value == nil || arr.ecls.IsInstance(value) {
		arr.elem[index] = value
	}
}

func (arr *objectArray) Elements() []java.IObject {
	return arr.elem
}

func (staticArray[V]) Class() java.IClass {
	return ClassResolve[[]V](nil)
}

func ArrayOf(val any) java.IObject {
	return arrayPack(val)
}

func arrayPack(obj any) java.IObject {
	switch v := obj.(type) {
	case []java.JBoolean:
		return primitiveArray[java.JBoolean](v)
	case []java.JByte:
		return primitiveArray[java.JByte](v)
	case []java.JChar:
		return primitiveArray[java.JChar](v)
	case []java.JShort:
		return primitiveArray[java.JShort](v)
	case []java.JInt:
		return primitiveArray[java.JInt](v)
	case []java.JLong:
		return primitiveArray[java.JLong](v)
	case []java.JFloat:
		return primitiveArray[java.JFloat](v)
	case []java.JDouble:
		return primitiveArray[java.JDouble](v)
	}
	typ := reflect.TypeOf(obj)
	if typ.Kind() != reflect.Slice {
		return nil
	}
	elemType := typ.Elem()
	ecls := getClass(elemType)
	if ecls == nil {
		return nil
	}
	ptr := (*struct {
		rtype uintptr
		slice *struct {
			data unsafe.Pointer
			len  int
		}
	})(unsafe.Pointer(&obj))
	arr := &objectArray{cls: getClass(typ), ecls: ecls, rtype: ptr.rtype, elem: make([]java.IObject, ptr.slice.len)}
	elemPtr := ptr.slice.data
	elemSize := elemType.Size()
	var cast func(unsafe.Pointer) java.IObject
	switch elemType.Kind() {
	case reflect.Interface:
		if elemType.NumMethod() == 0 {
			cast = func(p unsafe.Pointer) java.IObject {
				ifi := *(*any)(p)
				if ifi == nil {
					return nil
				}
				return ifi.(java.IObject)
			}
		} else {
			cast = func(p unsafe.Pointer) java.IObject {
				ifi := *(*interface{ M() })(p)
				if ifi == nil {
					return nil
				}
				return ifi.(java.IObject)
			}
		}
	case reflect.Pointer:
		rtype := getRType(elemType)
		cast = func(p unsafe.Pointer) java.IObject {
			p = *(*unsafe.Pointer)(p)
			if p == nil {
				return nil
			}
			return packObject(rtype, uintptr(p)).(java.IObject)
		}
	default:
		rtype := getRType(elemType)
		cast = func(p unsafe.Pointer) java.IObject {
			return packObject(rtype, uintptr(p)).(java.IObject)
		}
	}
	for i := 0; i < ptr.slice.len; i++ {
		arr.elem[i] = cast(elemPtr)
		elemPtr = unsafe.Add(elemPtr, elemSize)
	}
	return arr
}

func arrayUnpack(obj java.IObject) any {
	switch v := obj.(type) {
	case primitiveArray[java.JBoolean]:
		return []java.JBoolean(v)
	case primitiveArray[java.JByte]:
		return []java.JByte(v)
	case primitiveArray[java.JChar]:
		return []java.JChar(v)
	case primitiveArray[java.JShort]:
		return []java.JShort(v)
	case primitiveArray[java.JInt]:
		return []java.JInt(v)
	case primitiveArray[java.JLong]:
		return []java.JLong(v)
	case primitiveArray[java.JFloat]:
		return []java.JFloat(v)
	case primitiveArray[java.JDouble]:
		return []java.JDouble(v)
	case *objectArray:
		var i any
		var typ reflect.Type
		write := func(p unsafe.Pointer, o java.IObject, n uintptr) {
			data := (*struct{ _, data unsafe.Pointer })(unsafe.Pointer(&o)).data
			copy(unsafe.Slice((*byte)(p), n), unsafe.Slice((*byte)(data), n))
		}
		switch v.ecls {
		case stringClass:
			i = make([]String, len(v.elem))
		case booleanClass:
			i = make([]Boolean, len(v.elem))
		case byteClass:
			i = make([]Byte, len(v.elem))
		case characterClass:
			i = make([]Character, len(v.elem))
		case shortClass:
			i = make([]Short, len(v.elem))
		case integerClass:
			i = make([]Integer, len(v.elem))
		case longClass:
			i = make([]Long, len(v.elem))
		case floatClass:
			i = make([]Float, len(v.elem))
		case doubleClass:
			i = make([]Double, len(v.elem))
		default:
			typ = getType(v.rtype).Elem()
			rtype := getRType(typ)
			slice := new(struct {
				data unsafe.Pointer
				len  int
				cap  int
			})
			slice.len = len(v.elem)
			slice.cap = slice.len
			slice.data = unsafe_NewArray(rtype, slice.len)
			ptr := (*struct {
				rtype uintptr
				slice unsafe.Pointer
			})(unsafe.Pointer(&i))
			ptr.rtype = v.rtype
			ptr.slice = unsafe.Pointer(slice)
			switch typ.Kind() {
			case reflect.Interface:
				if typ.NumMethod() == 0 {
					write = func(p unsafe.Pointer, o java.IObject, _ uintptr) {
						*(*any)(p) = o
					}
				} else {
					write = func(p unsafe.Pointer, o java.IObject, _ uintptr) {
						ifaceE2I(rtype, o, p)
					}
				}
			case reflect.Pointer:
				write = func(p unsafe.Pointer, o java.IObject, _ uintptr) {
					*(*unsafe.Pointer)(p) = (*struct{ _, data unsafe.Pointer })(unsafe.Pointer(&o)).data
				}
			default:
				write = func(p unsafe.Pointer, o java.IObject, n uintptr) {
					data := (*struct{ _, data unsafe.Pointer })(unsafe.Pointer(&o)).data
					copy(unsafe.Slice((*byte)(p), n), unsafe.Slice((*byte)(data), n))
				}
			}
		}
		if typ == nil {
			typ = reflect.TypeOf(i).Elem()
		}
		ptr := (*struct {
			_     uintptr
			slice *struct{ data unsafe.Pointer }
		})(unsafe.Pointer(&i)).slice.data
		size := typ.Size()
		for _, o := range v.elem {
			if o != nil {
				write(ptr, o, size)
			}
			ptr = unsafe.Add(ptr, size)
		}
		return i
	}
	return obj
}

func init() {
	arrayInterfaces = []java.IClass{cloneableClass, io.SSerializable.Class()}
	typ := reflect.TypeFor[[]java.JBoolean]()
	booleanArrayClass = createClass(getRType(typ), typ, nil)
	booleanArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JChar]()
	charArrayClass = createClass(getRType(typ), typ, nil)
	charArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JByte]()
	byteArrayClass = createClass(getRType(typ), typ, nil)
	byteArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JShort]()
	shortArrayClass = createClass(getRType(typ), typ, nil)
	shortArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JInt]()
	intArrayClass = createClass(getRType(typ), typ, nil)
	intArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JLong]()
	longArrayClass = createClass(getRType(typ), typ, nil)
	longArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JFloat]()
	floatArrayClass = createClass(getRType(typ), typ, nil)
	floatArrayClass.interfaces = arrayInterfaces
	typ = reflect.TypeFor[[]java.JDouble]()
	doubleArrayClass = createClass(getRType(typ), typ, nil)
	doubleArrayClass.interfaces = arrayInterfaces
}
