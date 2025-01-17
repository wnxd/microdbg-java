package reflect

import (
	"slices"
	"strings"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type Method struct {
	Executable
	clazz          java.IClass
	name           java.IString
	returnType     java.IClass
	parameterTypes []java.IClass
	exceptionTypes []java.IClass
	modifiers      java.JInt
}

type staticMethod struct {
}

var (
	SMethod     staticMethod
	methodClass = internal.ClassResolve[Method](&SMethod)
)

func (m *Method) GetDeclaringClass() java.IClass {
	return m.clazz
}

func (m *Method) GetName() java.IString {
	return m.name
}

func (m *Method) GetModifiers() java.JInt {
	return m.modifiers
}

func (m *Method) GetReturnType() java.IClass {
	return m.returnType
}

func (m *Method) GetSharedParameterTypes() []java.IClass {
	return m.parameterTypes
}

func (m *Method) GetSharedExceptionTypes() []java.IClass {
	return m.exceptionTypes
}

func (m *Method) GetParameterTypes() []java.IClass {
	return slices.Clone(m.parameterTypes)
}

func (m *Method) GetParameterCount() java.JInt {
	return java.JInt(len(m.parameterTypes))
}

func (m *Method) GetExceptionTypes() []java.IClass {
	return slices.Clone(m.exceptionTypes)
}

func (m *Method) IsDefault() java.JBoolean {
	return (m.GetModifiers()&(Modifier_ABSTRACT|Modifier_PUBLIC|Modifier_STATIC) == Modifier_PUBLIC) && m.GetDeclaringClass().IsInterface()
}

func (m *Method) Invoke(obj java.IObject, args ...java.IObject) java.IObject {
	if SModifier.IsStatic(m.modifiers) {
	} else if !m.clazz.IsInstance(obj) {
		// panic
		return nil
	} else if len(args) != len(m.parameterTypes) {
		// panic
		return nil
	}
	arr := make([]any, len(args))
	for i := range arr {
		v, ok := internal.TryCast(args[i], m.parameterTypes[i])
		if !ok {
			// panic
			return nil
		}
		arr[i] = v
	}
	return internal.ObjectPack(m.method(obj, arr...))
}

func (m *Method) HashCode() java.JInt {
	return m.GetDeclaringClass().GetName().HashCode() ^ m.GetName().HashCode()
}

func (m *Method) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(*Method); ok {
		if !m.GetDeclaringClass().Equals(other.GetDeclaringClass()) || !m.GetName().Equals(other.GetName()) {
			return false
		} else if !m.returnType.Equals(other.GetReturnType()) {
			return false
		}
		return m.equalParamTypes(m.parameterTypes, other.parameterTypes)
	}
	return false
}

func (m *Method) ToString() java.IString {
	return m.sharedToString(SModifier.MethodModifiers(), m.IsDefault(), m.parameterTypes, m.exceptionTypes)
}

func (m *Method) specificToStringHeader(sb *strings.Builder) {
	sb.WriteString(m.GetReturnType().GetTypeName().String())
	sb.WriteByte(' ')
	sb.WriteString(m.GetDeclaringClass().GetTypeName().String())
	sb.WriteByte(' ')
	sb.WriteString(m.GetName().String())
}

func (staticMethod) Class() java.IClass {
	return methodClass
}

func NewMethod(cls java.IClass, name java.IString, returnType java.IClass, parameterTypes, exceptionTypes []java.IClass, modifiers java.JInt, method func(java.IObject, ...any) any) *Method {
	m := methodClass.NewInstance().(*Method)
	m.clazz = cls
	m.name = name
	m.returnType = returnType
	m.parameterTypes = parameterTypes
	m.exceptionTypes = exceptionTypes
	m.modifiers = modifiers
	m.method = method
	return m
}
