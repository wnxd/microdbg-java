package lang

import (
	"reflect"
)

type booleanType struct {
	reflect.Type
}

type charType struct {
	reflect.Type
}

type byteType struct {
	reflect.Type
}

type shortType struct {
	reflect.Type
}

type intType struct {
	reflect.Type
}

type longType struct {
	reflect.Type
}

type floatType struct {
	reflect.Type
}

type doubleType struct {
	reflect.Type
}

type voidType struct {
	reflect.Type
}

func (booleanType) Name() string {
	return "boolean"
}

func (charType) Name() string {
	return "char"
}

func (byteType) Name() string {
	return "byte"
}

func (shortType) Name() string {
	return "short"
}

func (intType) Name() string {
	return "int"
}

func (longType) Name() string {
	return "long"
}

func (floatType) Name() string {
	return "float"
}

func (doubleType) Name() string {
	return "double"
}

func (voidType) Kind() reflect.Kind {
	return reflect.Invalid
}

func (voidType) Name() string {
	return "void"
}

func (voidType) PkgPath() string {
	return ""
}
