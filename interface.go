package java

type IObject interface {
	GetClass() IClass
	HashCode() JInt
	Equals(IObject) JBoolean
	ToString() IString
}

type IClass interface {
	IObject
	NewInstance() IObject
	GetName() IString
	GetSimpleName() IString
	GetTypeName() IString
	DescriptorString() IString
	GetSuperclass() IClass
	GetInterfaces() []IClass
	IsInterface() JBoolean
	IsAssignableFrom(IClass) JBoolean
	IsPrimitive() JBoolean
	IsArray() JBoolean
	IsInstance(IObject) JBoolean
	Cast(IObject) IObject
}

type IString interface {
	IObject
	Length() JInt
	String() string
}

type IArray interface {
	IObject
	Length() JInt
}

type IGenericArray[V comparable] interface {
	IArray
	Get(JInt) V
	Set(JInt, V)
	Elements() []V
}

type IThrowable interface {
	IObject
	GetMessage() IString
}

type IMethod interface {
	IObject
	GetName() IString
	GetModifiers() JInt
	GetParameterTypes() []IClass
	GetParameterCount() JInt
	Call(IObject, ...any) IObject
	CallPrimitive(IObject, ...any) any
}

type IField interface {
	IObject
	GetName() IString
	GetModifiers() JInt
	GetType() IClass
	Get(IObject) IObject
	GetPrimitive(IObject) any
	Set(IObject, IObject)
	SetPrimitive(IObject, any)
}

type IBooleanArray = IGenericArray[JBoolean]
type IByteArray = IGenericArray[JByte]
type ICharArray = IGenericArray[JChar]
type IShortArray = IGenericArray[JShort]
type IIntArray = IGenericArray[JInt]
type ILongArray = IGenericArray[JLong]
type IFloatArray = IGenericArray[JFloat]
type IDoubleArray = IGenericArray[JDouble]
type IObjectArray = IGenericArray[IObject]
