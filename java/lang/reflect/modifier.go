package reflect

import (
	"strings"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

const (
	Modifier_PUBLIC java.JInt = 1 << iota
	Modifier_PRIVATE
	Modifier_PROTECTED
	Modifier_STATIC
	Modifier_FINAL
	Modifier_SYNCHRONIZED
	Modifier_VOLATILE
	Modifier_TRANSIENT
	Modifier_NATIVE
	Modifier_INTERFACE
	Modifier_ABSTRACT
	Modifier_STRICT

	Modifier_ACCESS_MODIFIERS = Modifier_PUBLIC | Modifier_PROTECTED | Modifier_PRIVATE
)

type staticModifier struct {
}

var SModifier staticModifier

func (staticModifier) IsStatic(mod java.JInt) java.JBoolean {
	return (mod & Modifier_STATIC) != 0
}

func (staticModifier) isSynthetic(mod java.JInt) java.JBoolean {
	const SYNTHETIC = 0x00001000

	return mod&SYNTHETIC != 0
}

func (staticModifier) ConstructorModifiers() java.JInt {
	const CONSTRUCTOR_MODIFIERS = Modifier_PUBLIC | Modifier_PROTECTED | Modifier_PRIVATE

	return CONSTRUCTOR_MODIFIERS
}

func (staticModifier) MethodModifiers() java.JInt {
	const METHOD_MODIFIERS = Modifier_PUBLIC | Modifier_PROTECTED | Modifier_PRIVATE |
		Modifier_ABSTRACT | Modifier_STATIC | Modifier_FINAL |
		Modifier_SYNCHRONIZED | Modifier_NATIVE | Modifier_STRICT

	return METHOD_MODIFIERS
}

func (staticModifier) ToString(mod java.JInt) java.IString {
	var arr []string
	if (mod & Modifier_PUBLIC) != 0 {
		arr = append(arr, "public")
	}
	if (mod & Modifier_PROTECTED) != 0 {
		arr = append(arr, "protected")
	}
	if (mod & Modifier_PRIVATE) != 0 {
		arr = append(arr, "private")
	}
	if (mod & Modifier_ABSTRACT) != 0 {
		arr = append(arr, "abstract")
	}
	if (mod & Modifier_STATIC) != 0 {
		arr = append(arr, "static")
	}
	if (mod & Modifier_FINAL) != 0 {
		arr = append(arr, "final")
	}
	if (mod & Modifier_TRANSIENT) != 0 {
		arr = append(arr, "transient")
	}
	if (mod & Modifier_VOLATILE) != 0 {
		arr = append(arr, "volatile")
	}
	if (mod & Modifier_SYNCHRONIZED) != 0 {
		arr = append(arr, "synchronized")
	}
	if (mod & Modifier_NATIVE) != 0 {
		arr = append(arr, "native")
	}
	if (mod & Modifier_STRICT) != 0 {
		arr = append(arr, "strictfp")
	}
	if (mod & Modifier_INTERFACE) != 0 {
		arr = append(arr, "interface")
	}
	return internal.StringPack(strings.Join(arr, " "))
}
