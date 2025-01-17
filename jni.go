package java

const (
	JNI_OK        = 0
	JNI_ERR       = -1
	JNI_EDETACHED = -2
	JNI_EVERSION  = -3
	JNI_ENOMEM    = -4
	JNI_EEXIST    = -5
	JNI_EINVAL    = -6

	JNI_VERSION_1_1 = 0x00010001
	JNI_VERSION_1_2 = 0x00010002
	JNI_VERSION_1_4 = 0x00010004
	JNI_VERSION_1_6 = 0x00010006
	JNI_VERSION_1_8 = 0x00010008

	JNI_COMMIT = 1
	JNI_ABORT  = 2
)

type JBoolean = bool
type JByte = int8
type JChar = uint16
type JShort = int16
type JInt = int32
type JLong = int64
type JFloat = float32
type JDouble = float64
type JSize = JInt

type JObject interface{}
type JClass interface{ JObject }
type JString interface{ JObject }
type JArray interface{ JObject }
type JGenericArray[E any] interface{ JArray }
type JThrowable interface{ JObject }
type JWeak JObject

type JFieldID interface{}
type JMethodID interface{}

type JValue interface {
	JBoolean() JBoolean
	JByte() JByte
	JChar() JChar
	JShort() JShort
	JInt() JInt
	JLong() JLong
	JFloat() JFloat
	JDouble() JDouble
	JObject() JObject
}

type JObjectRefType int32

const (
	JNIInvalidRefType JObjectRefType = iota
	JNILocalRefType
	JNIGlobalRefType
	JNIWeakGlobalRefType
)

type JNINativeMethod struct {
	Name      string
	Signature string
	FnPtr     AnyPtr
}

type JNIEnv interface {
	GetVersion() JInt
	DefineClass(string, JObject, []JByte) JClass
	FindClass(string) JClass
	FromReflectedMethod(JObject) JMethodID
	FromReflectedField(JObject) JFieldID
	ToReflectedMethod(JClass, JMethodID, JBoolean) JObject
	GetSuperclass(JClass) JClass
	IsAssignableFrom(JClass, JClass) JBoolean
	ToReflectedField(JClass, JFieldID, JBoolean) JObject
	Throw(JThrowable) JInt
	ThrowNew(JClass, string) JInt
	ExceptionOccurred() JThrowable
	ExceptionDescribe()
	ExceptionClear()
	FatalError(string)
	PushLocalFrame(JInt) JInt
	PopLocalFrame(JObject) JObject
	NewGlobalRef(JObject) JObject
	DeleteGlobalRef(JObject)
	DeleteLocalRef(JObject)
	IsSameObject(JObject, JObject) JBoolean
	NewLocalRef(JObject) JObject
	EnsureLocalCapacity(JInt) JInt
	AllocObject(JClass) JObject
	NewObject(JClass, JMethodID, ...any) JObject
	NewObjectV(JClass, JMethodID, VaList) JObject
	NewObjectA(JClass, JMethodID, TypePtr[JValue]) JObject
	GetObjectClass(JObject) JClass
	IsInstanceOf(JObject, JClass) JBoolean
	GetMethodID(JClass, string, string) JMethodID
	CallObjectMethod(JObject, JMethodID, ...any) JObject
	CallObjectMethodV(JObject, JMethodID, VaList) JObject
	CallObjectMethodA(JObject, JMethodID, TypePtr[JValue]) JObject
	CallBooleanMethod(JObject, JMethodID, ...any) JBoolean
	CallBooleanMethodV(JObject, JMethodID, VaList) JBoolean
	CallBooleanMethodA(JObject, JMethodID, TypePtr[JValue]) JBoolean
	CallByteMethod(JObject, JMethodID, ...any) JByte
	CallByteMethodV(JObject, JMethodID, VaList) JByte
	CallByteMethodA(JObject, JMethodID, TypePtr[JValue]) JByte
	CallCharMethod(JObject, JMethodID, ...any) JChar
	CallCharMethodV(JObject, JMethodID, VaList) JChar
	CallCharMethodA(JObject, JMethodID, TypePtr[JValue]) JChar
	CallShortMethod(JObject, JMethodID, ...any) JShort
	CallShortMethodV(JObject, JMethodID, VaList) JShort
	CallShortMethodA(JObject, JMethodID, TypePtr[JValue]) JShort
	CallIntMethod(JObject, JMethodID, ...any) JInt
	CallIntMethodV(JObject, JMethodID, VaList) JInt
	CallIntMethodA(JObject, JMethodID, TypePtr[JValue]) JInt
	CallLongMethod(JObject, JMethodID, ...any) JLong
	CallLongMethodV(JObject, JMethodID, VaList) JLong
	CallLongMethodA(JObject, JMethodID, TypePtr[JValue]) JLong
	CallFloatMethod(JObject, JMethodID, ...any) JFloat
	CallFloatMethodV(JObject, JMethodID, VaList) JFloat
	CallFloatMethodA(JObject, JMethodID, TypePtr[JValue]) JFloat
	CallDoubleMethod(JObject, JMethodID, ...any) JDouble
	CallDoubleMethodV(JObject, JMethodID, VaList) JDouble
	CallDoubleMethodA(JObject, JMethodID, TypePtr[JValue]) JDouble
	CallVoidMethod(JObject, JMethodID, ...any)
	CallVoidMethodV(JObject, JMethodID, VaList)
	CallVoidMethodA(JObject, JMethodID, TypePtr[JValue])
	CallNonvirtualObjectMethod(JObject, JClass, JMethodID, ...any) JObject
	CallNonvirtualObjectMethodV(JObject, JClass, JMethodID, VaList) JObject
	CallNonvirtualObjectMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JObject
	CallNonvirtualBooleanMethod(JObject, JClass, JMethodID, ...any) JBoolean
	CallNonvirtualBooleanMethodV(JObject, JClass, JMethodID, VaList) JBoolean
	CallNonvirtualBooleanMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JBoolean
	CallNonvirtualByteMethod(JObject, JClass, JMethodID, ...any) JByte
	CallNonvirtualByteMethodV(JObject, JClass, JMethodID, VaList) JByte
	CallNonvirtualByteMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JByte
	CallNonvirtualCharMethod(JObject, JClass, JMethodID, ...any) JChar
	CallNonvirtualCharMethodV(JObject, JClass, JMethodID, VaList) JChar
	CallNonvirtualCharMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JChar
	CallNonvirtualShortMethod(JObject, JClass, JMethodID, ...any) JShort
	CallNonvirtualShortMethodV(JObject, JClass, JMethodID, VaList) JShort
	CallNonvirtualShortMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JShort
	CallNonvirtualIntMethod(JObject, JClass, JMethodID, ...any) JInt
	CallNonvirtualIntMethodV(JObject, JClass, JMethodID, VaList) JInt
	CallNonvirtualIntMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JInt
	CallNonvirtualLongMethod(JObject, JClass, JMethodID, ...any) JLong
	CallNonvirtualLongMethodV(JObject, JClass, JMethodID, VaList) JLong
	CallNonvirtualLongMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JLong
	CallNonvirtualFloatMethod(JObject, JClass, JMethodID, ...any) JFloat
	CallNonvirtualFloatMethodV(JObject, JClass, JMethodID, VaList) JFloat
	CallNonvirtualFloatMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JFloat
	CallNonvirtualDoubleMethod(JObject, JClass, JMethodID, ...any) JDouble
	CallNonvirtualDoubleMethodV(JObject, JClass, JMethodID, VaList) JDouble
	CallNonvirtualDoubleMethodA(JObject, JClass, JMethodID, TypePtr[JValue]) JDouble
	CallNonvirtualVoidMethod(JObject, JClass, JMethodID, ...any)
	CallNonvirtualVoidMethodV(JObject, JClass, JMethodID, VaList)
	CallNonvirtualVoidMethodA(JObject, JClass, JMethodID, TypePtr[JValue])
	GetFieldID(JClass, string, string) JFieldID
	GetObjectField(JObject, JFieldID) JObject
	GetBooleanField(JObject, JFieldID) JBoolean
	GetByteField(JObject, JFieldID) JByte
	GetCharField(JObject, JFieldID) JChar
	GetShortField(JObject, JFieldID) JShort
	GetIntField(JObject, JFieldID) JInt
	GetLongField(JObject, JFieldID) JLong
	GetFloatField(JObject, JFieldID) JFloat
	GetDoubleField(JObject, JFieldID) JDouble
	SetObjectField(JObject, JFieldID, JObject)
	SetBooleanField(JObject, JFieldID, JBoolean)
	SetByteField(JObject, JFieldID, JByte)
	SetCharField(JObject, JFieldID, JChar)
	SetShortField(JObject, JFieldID, JShort)
	SetIntField(JObject, JFieldID, JInt)
	SetLongField(JObject, JFieldID, JLong)
	SetFloatField(JObject, JFieldID, JFloat)
	SetDoubleField(JObject, JFieldID, JDouble)
	GetStaticMethodID(JClass, string, string) JMethodID
	CallStaticObjectMethod(JClass, JMethodID, ...any) JObject
	CallStaticObjectMethodV(JClass, JMethodID, VaList) JObject
	CallStaticObjectMethodA(JClass, JMethodID, TypePtr[JValue]) JObject
	CallStaticBooleanMethod(JClass, JMethodID, ...any) JBoolean
	CallStaticBooleanMethodV(JClass, JMethodID, VaList) JBoolean
	CallStaticBooleanMethodA(JClass, JMethodID, TypePtr[JValue]) JBoolean
	CallStaticByteMethod(JClass, JMethodID, ...any) JByte
	CallStaticByteMethodV(JClass, JMethodID, VaList) JByte
	CallStaticByteMethodA(JClass, JMethodID, TypePtr[JValue]) JByte
	CallStaticCharMethod(JClass, JMethodID, ...any) JChar
	CallStaticCharMethodV(JClass, JMethodID, VaList) JChar
	CallStaticCharMethodA(JClass, JMethodID, TypePtr[JValue]) JChar
	CallStaticShortMethod(JClass, JMethodID, ...any) JShort
	CallStaticShortMethodV(JClass, JMethodID, VaList) JShort
	CallStaticShortMethodA(JClass, JMethodID, TypePtr[JValue]) JShort
	CallStaticIntMethod(JClass, JMethodID, ...any) JInt
	CallStaticIntMethodV(JClass, JMethodID, VaList) JInt
	CallStaticIntMethodA(JClass, JMethodID, TypePtr[JValue]) JInt
	CallStaticLongMethod(JClass, JMethodID, ...any) JLong
	CallStaticLongMethodV(JClass, JMethodID, VaList) JLong
	CallStaticLongMethodA(JClass, JMethodID, TypePtr[JValue]) JLong
	CallStaticFloatMethod(JClass, JMethodID, ...any) JFloat
	CallStaticFloatMethodV(JClass, JMethodID, VaList) JFloat
	CallStaticFloatMethodA(JClass, JMethodID, TypePtr[JValue]) JFloat
	CallStaticDoubleMethod(JClass, JMethodID, ...any) JDouble
	CallStaticDoubleMethodV(JClass, JMethodID, VaList) JDouble
	CallStaticDoubleMethodA(JClass, JMethodID, TypePtr[JValue]) JDouble
	CallStaticVoidMethod(JClass, JMethodID, ...any)
	CallStaticVoidMethodV(JClass, JMethodID, VaList)
	CallStaticVoidMethodA(JClass, JMethodID, TypePtr[JValue])
	GetStaticFieldID(JClass, string, string) JFieldID
	GetStaticObjectField(JClass, JFieldID) JObject
	GetStaticBooleanField(JClass, JFieldID) JBoolean
	GetStaticByteField(JClass, JFieldID) JByte
	GetStaticCharField(JClass, JFieldID) JChar
	GetStaticShortField(JClass, JFieldID) JShort
	GetStaticIntField(JClass, JFieldID) JInt
	GetStaticLongField(JClass, JFieldID) JLong
	GetStaticFloatField(JClass, JFieldID) JFloat
	GetStaticDoubleField(JClass, JFieldID) JDouble
	SetStaticObjectField(JClass, JFieldID, JObject)
	SetStaticBooleanField(JClass, JFieldID, JBoolean)
	SetStaticByteField(JClass, JFieldID, JByte)
	SetStaticCharField(JClass, JFieldID, JChar)
	SetStaticShortField(JClass, JFieldID, JShort)
	SetStaticIntField(JClass, JFieldID, JInt)
	SetStaticLongField(JClass, JFieldID, JLong)
	SetStaticFloatField(JClass, JFieldID, JFloat)
	SetStaticDoubleField(JClass, JFieldID, JDouble)
	NewString([]JChar) JString
	GetStringLength(JString) JSize
	GetStringChars(JString) []JChar
	ReleaseStringChars(JString, []JChar)
	NewStringUTF(string) JString
	GetStringUTFLength(JString) JSize
	GetStringUTFChars(JString) []byte
	ReleaseStringUTFChars(JString, []byte)
	GetArrayLength(JArray) JSize
	NewObjectArray(JSize, JClass, JObject) JGenericArray[JObject]
	GetObjectArrayElement(JGenericArray[JObject], JSize) JObject
	SetObjectArrayElement(JGenericArray[JObject], JSize, JObject)
	NewBooleanArray(JSize) JGenericArray[JBoolean]
	NewByteArray(JSize) JGenericArray[JByte]
	NewCharArray(JSize) JGenericArray[JChar]
	NewShortArray(JSize) JGenericArray[JShort]
	NewIntArray(JSize) JGenericArray[JInt]
	NewLongArray(JSize) JGenericArray[JLong]
	NewFloatArray(JSize) JGenericArray[JFloat]
	NewDoubleArray(JSize) JGenericArray[JDouble]
	GetBooleanArrayElements(JGenericArray[JBoolean]) []JBoolean
	GetByteArrayElements(JGenericArray[JByte]) []JByte
	GetCharArrayElements(JGenericArray[JChar]) []JChar
	GetShortArrayElements(JGenericArray[JShort]) []JShort
	GetIntArrayElements(JGenericArray[JInt]) []JInt
	GetLongArrayElements(JGenericArray[JLong]) []JLong
	GetFloatArrayElements(JGenericArray[JFloat]) []JFloat
	GetDoubleArrayElements(JGenericArray[JDouble]) []JDouble
	ReleaseBooleanArrayElements(JGenericArray[JBoolean], []JBoolean, JInt)
	ReleaseByteArrayElements(JGenericArray[JByte], []JByte, JInt)
	ReleaseCharArrayElements(JGenericArray[JChar], []JChar, JInt)
	ReleaseShortArrayElements(JGenericArray[JShort], []JShort, JInt)
	ReleaseIntArrayElements(JGenericArray[JInt], []JInt, JInt)
	ReleaseLongArrayElements(JGenericArray[JLong], []JLong, JInt)
	ReleaseFloatArrayElements(JGenericArray[JFloat], []JFloat, JInt)
	ReleaseDoubleArrayElements(JGenericArray[JDouble], []JDouble, JInt)
	GetBooleanArrayRegion(JGenericArray[JBoolean], JSize, []JBoolean)
	GetByteArrayRegion(JGenericArray[JByte], JSize, []JByte)
	GetCharArrayRegion(JGenericArray[JChar], JSize, []JChar)
	GetShortArrayRegion(JGenericArray[JShort], JSize, []JShort)
	GetIntArrayRegion(JGenericArray[JInt], JSize, []JInt)
	GetLongArrayRegion(JGenericArray[JLong], JSize, []JLong)
	GetFloatArrayRegion(JGenericArray[JFloat], JSize, []JFloat)
	GetDoubleArrayRegion(JGenericArray[JDouble], JSize, []JDouble)
	SetBooleanArrayRegion(JGenericArray[JBoolean], JSize, []JBoolean)
	SetByteArrayRegion(JGenericArray[JByte], JSize, []JByte)
	SetCharArrayRegion(JGenericArray[JChar], JSize, []JChar)
	SetShortArrayRegion(JGenericArray[JShort], JSize, []JShort)
	SetIntArrayRegion(JGenericArray[JInt], JSize, []JInt)
	SetLongArrayRegion(JGenericArray[JLong], JSize, []JLong)
	SetFloatArrayRegion(JGenericArray[JFloat], JSize, []JFloat)
	SetDoubleArrayRegion(JGenericArray[JDouble], JSize, []JDouble)
	RegisterNatives(JClass, []JNINativeMethod) JInt
	UnregisterNatives(JClass) JInt
	MonitorEnter(JObject) JInt
	MonitorExit(JObject) JInt
	GetJavaVM(*JavaVM) JInt
	GetStringRegion(JString, JSize, []JChar)
	GetStringUTFRegion(JString, JSize, []byte)
	GetPrimitiveArrayCritical(JArray) []byte
	ReleasePrimitiveArrayCritical(JArray, []byte, JInt)
	GetStringCritical(JString) []JChar
	ReleaseStringCritical(JString, []JChar)
	NewWeakGlobalRef(JObject) JWeak
	DeleteWeakGlobalRef(JWeak)
	ExceptionCheck() JBoolean
	NewDirectByteBuffer(AnyPtr, JLong) JObject
	GetDirectBufferAddress(JObject) AnyPtr
	GetDirectBufferCapacity(JObject) JLong
	GetObjectRefType(JObject) JObjectRefType
}

type JavaVM interface {
	DestroyJavaVM() JInt
	AttachCurrentThread(*JNIEnv, any) JInt
	DetachCurrentThread() JInt
	GetEnv(*JNIEnv, JInt) JInt
	AttachCurrentThreadAsDaemon(*JNIEnv, any) JInt
}
