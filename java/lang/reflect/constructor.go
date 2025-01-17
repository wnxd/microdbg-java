package reflect

import (
	"slices"
	"strings"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type Constructor struct {
	Executable
	clazz          java.IClass
	parameterTypes []java.IClass
	exceptionTypes []java.IClass
	modifiers      java.JInt
}

type staticConstructor struct {
}

var (
	SConstructor     staticConstructor
	constructorClass = internal.ClassResolve[Constructor](&SConstructor)
)

func (ctor *Constructor) GetDeclaringClass() java.IClass {
	return ctor.clazz
}

func (ctor *Constructor) GetName() java.IString {
	return ctor.GetDeclaringClass().GetName()
}

func (ctor *Constructor) GetModifiers() java.JInt {
	return ctor.modifiers
}

func (ctor *Constructor) GetSharedParameterTypes() []java.IClass {
	return ctor.parameterTypes
}

func (ctor *Constructor) GetSharedExceptionTypes() []java.IClass {
	return ctor.exceptionTypes
}

func (ctor *Constructor) GetParameterTypes() []java.IClass {
	return slices.Clone(ctor.parameterTypes)
}

func (ctor *Constructor) GetParameterCount() java.JInt {
	return java.JInt(len(ctor.parameterTypes))
}

func (ctor *Constructor) GetExceptionTypes() []java.IClass {
	return slices.Clone(ctor.exceptionTypes)
}

func (ctor *Constructor) NewInstance(initargs ...java.IObject) java.IObject {
	if len(initargs) != len(ctor.parameterTypes) {
		// panic
		return nil
	}
	arr := make([]any, len(initargs))
	for i := range arr {
		v, ok := internal.TryCast(initargs[i], ctor.parameterTypes[i])
		if !ok {
			// panic
			return nil
		}
		arr[i] = v
	}
	return internal.ObjectPack(ctor.method(nil, arr...))
}

func (ctor *Constructor) HashCode() java.JInt {
	return ctor.GetDeclaringClass().GetName().HashCode()
}

func (ctor *Constructor) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(*Constructor); ok {
		if !ctor.GetDeclaringClass().Equals(other.GetDeclaringClass()) {
			return false
		}
		return ctor.equalParamTypes(ctor.parameterTypes, other.parameterTypes)
	}
	return false
}

func (ctor *Constructor) ToString() java.IString {
	return ctor.sharedToString(SModifier.ConstructorModifiers(), false, ctor.parameterTypes, ctor.exceptionTypes)
}

func (ctor *Constructor) specificToStringHeader(sb *strings.Builder) {
	sb.WriteString(ctor.GetDeclaringClass().GetTypeName().String())
}

func (staticConstructor) Class() java.IClass {
	return constructorClass
}

func NewConstructor(cls java.IClass, parameterTypes, exceptionTypes []java.IClass, modifiers java.JInt, method func(java.IObject, ...any) any) *Constructor {
	ctor := constructorClass.NewInstance().(*Constructor)
	ctor.clazz = cls
	ctor.parameterTypes = parameterTypes
	ctor.exceptionTypes = exceptionTypes
	ctor.modifiers = modifiers
	ctor.method = method
	return ctor
}
