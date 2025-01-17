package reflect

import (
	"slices"
	"strings"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type executable interface {
	java.IObject
	Member
	GetSharedParameterTypes() []java.IClass
	GetSharedExceptionTypes() []java.IClass
	GetParameterTypes() []java.IClass
	GetParameterCount() java.JInt
	GetExceptionTypes() []java.IClass
	specificToStringHeader(*strings.Builder)
}

type Executable struct {
	AccessibleObject
	Member
	method func(java.IObject, ...any) any
}

func (exec *Executable) GetDeclaringClass() java.IClass {
	panic("Not implemented")
}

func (exec *Executable) GetName() java.IString {
	panic("Not implemented")
}

func (exec *Executable) GetModifiers() java.JInt {
	panic("Not implemented")
}

func (exec *Executable) IsSynthetic() java.JBoolean {
	return SModifier.isSynthetic(internal.This[Member](exec).GetModifiers())
}

func (exec *Executable) GetSharedParameterTypes() []java.IClass {
	panic("Not implemented")
}

func (exec *Executable) GetSharedExceptionTypes() []java.IClass {
	panic("Not implemented")
}

func (exec *Executable) GetParameterTypes() []java.IClass {
	panic("Not implemented")
}

func (exec *Executable) GetParameterCount() java.JInt {
	panic("Not implemented")
}

func (exec *Executable) GetExceptionTypes() []java.IClass {
	panic("Not implemented")
}

func (exec *Executable) Call(obj java.IObject, args ...any) java.IObject {
	return internal.ObjectPack(exec.method(obj, args...))
}

func (exec *Executable) CallPrimitive(obj java.IObject, args ...any) any {
	return exec.method(obj, args...)
}

func (exec *Executable) equalParamTypes(params1, params2 []java.IClass) java.JBoolean {
	return slices.Equal(params1, params2)
}

func (exec *Executable) sharedToString(modifierMask java.JInt, isDefault java.JBoolean, parameterTypes, exceptionTypes []java.IClass) java.IString {
	this := internal.This[executable](exec)
	var sb strings.Builder
	mod := this.GetModifiers() & modifierMask
	if mod != 0 && !isDefault {
		sb.WriteString(SModifier.ToString(mod).String())
		sb.WriteByte(' ')
	} else {
		if access := mod & Modifier_ACCESS_MODIFIERS; access != 0 {
			sb.WriteString(SModifier.ToString(access).String())
			sb.WriteByte(' ')
		}
		if isDefault {
			sb.WriteString("default ")
		}
		mod = (mod & ^Modifier_ACCESS_MODIFIERS)
		if mod != 0 {
			sb.WriteString(SModifier.ToString(mod).String())
			sb.WriteByte(' ')
		}
	}
	this.specificToStringHeader(&sb)
	sb.WriteByte('(')
	for i, v := range parameterTypes {
		if i != 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(v.GetTypeName().String())
	}
	sb.WriteByte(')')
	if len(exceptionTypes) > 0 {
		sb.WriteString(" throws ")
		for i, v := range exceptionTypes {
			if i != 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(v.GetTypeName().String())
		}
	}
	return internal.StringPack(sb.String())
}

func (exec *Executable) specificToStringHeader(sb *strings.Builder) {
	panic("Not implemented")
}
