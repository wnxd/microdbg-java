package reflect

import (
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type Field struct {
	AccessibleObject
	Member
	clazz     java.IClass
	name      java.IString
	typ       java.IClass
	modifiers java.JInt
	get       func(java.IObject) any
	set       func(java.IObject, any)
}

type staticField struct {
}

var (
	SField     staticField
	fieldClass = internal.ClassResolve[Field](&SField)
)

func (f *Field) GetDeclaringClass() java.IClass {
	return f.clazz
}

func (f *Field) GetName() java.IString {
	return f.name
}

func (f *Field) GetModifiers() java.JInt {
	return f.modifiers
}

func (f *Field) IsSynthetic() java.JBoolean {
	return SModifier.isSynthetic(f.GetModifiers())
}

func (f *Field) GetType() java.IClass {
	return f.typ
}

func (f *Field) Get(obj java.IObject) java.IObject {
	return internal.ObjectPack(f.get(obj))
}

func (f *Field) GetBoolean(obj java.IObject) java.JBoolean {
	return f.get(obj).(java.JBoolean)
}

func (f *Field) GetByte(obj java.IObject) java.JByte {
	return f.get(obj).(java.JByte)
}

func (f *Field) GetChar(obj java.IObject) java.JChar {
	return f.get(obj).(java.JChar)
}

func (f *Field) GetShort(obj java.IObject) java.JShort {
	return f.get(obj).(java.JShort)
}

func (f *Field) GetInt(obj java.IObject) java.JInt {
	return f.get(obj).(java.JInt)
}

func (f *Field) GetLong(obj java.IObject) java.JLong {
	return f.get(obj).(java.JLong)
}

func (f *Field) GetFloat(obj java.IObject) java.JFloat {
	return f.get(obj).(java.JFloat)
}

func (f *Field) GetDouble(obj java.IObject) java.JDouble {
	return f.get(obj).(java.JDouble)
}

func (f *Field) Set(obj java.IObject, value java.IObject) {
	v, ok := internal.TryCast(value, f.typ)
	if !ok {
		// panic
		return
	}
	f.set(obj, v)
}

func (f *Field) SetBoolean(obj java.IObject, value java.JBoolean) {
	f.set(obj, value)
}

func (f *Field) SetByte(obj java.IObject, value java.JByte) {
	f.set(obj, value)
}

func (f *Field) SetChar(obj java.IObject, value java.JChar) {
	f.set(obj, value)
}

func (f *Field) SetShort(obj java.IObject, value java.JShort) {
	f.set(obj, value)
}

func (f *Field) SetInt(obj java.IObject, value java.JInt) {
	f.set(obj, value)
}

func (f *Field) SetLong(obj java.IObject, value java.JLong) {
	f.set(obj, value)
}

func (f *Field) SetFloat(obj java.IObject, value java.JFloat) {
	f.set(obj, value)
}

func (f *Field) SetDouble(obj java.IObject, value java.JDouble) {
	f.set(obj, value)
}

func (f *Field) GetPrimitive(obj java.IObject) any {
	return f.get(obj)
}

func (f *Field) SetPrimitive(obj java.IObject, value any) {
	f.set(obj, value)
}

func (f *Field) HashCode() java.JInt {
	return f.GetDeclaringClass().GetName().HashCode() ^ f.GetName().HashCode()
}

func (f *Field) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(*Field); ok {
		return (f.GetDeclaringClass() == other.GetDeclaringClass()) &&
			(f.GetName() == other.GetName()) &&
			(f.GetType() == other.GetType())
	}
	return false
}

func (f *Field) ToString() java.IString {
	var modifier string
	if mod := f.GetModifiers(); mod != 0 {
		modifier = SModifier.ToString(mod).String() + " "
	}
	return internal.StringPack(modifier +
		f.GetType().GetTypeName().String() + " " +
		f.GetDeclaringClass().GetTypeName().String() + "." +
		f.GetName().String())
}

func (staticField) Class() java.IClass {
	return fieldClass
}

func NewField(cls java.IClass, name java.IString, typ java.IClass, modifiers java.JInt, get func(java.IObject) any, set func(java.IObject, any)) *Field {
	f := fieldClass.NewInstance().(*Field)
	f.clazz = cls
	f.name = name
	f.typ = typ
	f.modifiers = modifiers
	f.get = get
	f.set = set
	return f
}
