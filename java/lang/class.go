package lang

import (
	"iter"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
	lang_reflect "github.com/wnxd/microdbg-java/java/lang/reflect"
)

type StaticClass interface {
	Class() java.IClass
}

type Class struct {
	lang_reflect.Type
	lang_reflect.AnnotatedElement
	typ        reflect.Type
	static     reflect.Value
	ift        reflect.Type
	name       java.IString
	simpleName java.IString
	superClass java.IClass
	interfaces []java.IClass
}

type staticClass struct {
}

var (
	SClass     staticClass
	classClass = ClassResolve[*Class](&SClass)

	classMap sync.Map
)

func (cls *Class) GetClass() java.IClass {
	return classClass
}

func (cls *Class) HashCode() java.JInt {
	return java.JInt(uintptr(unsafe.Pointer(cls)))
}

func (cls *Class) Equals(obj java.IObject) java.JBoolean {
	if typeof, ok := obj.(interface{ typeof() reflect.Type }); ok {
		return cls.typeof() == typeof.typeof()
	} else if other, ok := obj.(interface {
		IsArray() java.JBoolean
		ComponentType() java.IClass
	}); ok && cls.IsArray() && other.IsArray() {
		return cls.ComponentType().Equals(other.ComponentType())
	}
	return cls == obj
}

func (cls *Class) ToString() java.IString {
	var prefix string
	if cls.IsInterface() {
		prefix = "interface "
	} else if cls.IsPrimitive() {
		prefix = ""
	} else {
		prefix = "class "
	}
	prefix += cls.GetName().String()
	return String(prefix)
}

func (cls *Class) NewInstance() java.IObject {
	if cls.IsInterface() || cls.IsPrimitive() || cls.IsArray() {
		return nil
	}
	return newInstance(cls, cls.typ)
}

func (cls *Class) GetName() java.IString {
	if cls.name == nil {
		if cls.IsPrimitive() {
			cls.name = String(cls.typ.Name())
		} else if cls.IsArray() {
			cls.name = String(strings.ReplaceAll(cls.DescriptorString().String(), "/", "."))
		} else {
			typ := typeIndirect(cls.typ)
			arr := strings.SplitAfterN(typ.PkgPath(), "/", 4)
			packageName := strings.ReplaceAll(arr[len(arr)-1], "/", ".")
			name, _, _ := strings.Cut(typ.Name(), "[")
			cls.name = String(packageName + "." + name)
		}
	}
	return cls.name
}

func (cls *Class) GetSimpleName() java.IString {
	if cls.simpleName == nil {
		if cls.IsArray() {
			cls.simpleName = String(getClass(cls.typ.Elem()).GetSimpleName().String() + "[]")
		} else {
			name := cls.GetName().String()
			cls.simpleName = String(name[strings.LastIndex(name, ".")+1:])
		}
	}
	return cls.simpleName
}

func (cls *Class) GetTypeName() java.IString {
	if cls.IsArray() {
		typ := cls.typ
		dimensions := 0
		for ; typ.Kind() == reflect.Slice; dimensions++ {
			typ = typ.Elem()
		}
		return String(getClass(typ).GetName().String() + strings.Repeat("[]", dimensions))
	}
	return cls.GetName()
}

func (cls *Class) DescriptorString() java.IString {
	if cls.IsPrimitive() {
		switch cls.typ.Kind() {
		case reflect.Bool:
			return String("Z")
		case reflect.Uint16:
			return String("C")
		case reflect.Int8:
			return String("B")
		case reflect.Int16:
			return String("S")
		case reflect.Int32:
			return String("I")
		case reflect.Int64:
			return String("J")
		case reflect.Float32:
			return String("F")
		case reflect.Float64:
			return String("D")
		}
		return nil
	} else if cls.IsArray() {
		return String("[" + getClass(cls.typ.Elem()).DescriptorString().String())
	}
	return String("L" + strings.ReplaceAll(cls.GetName().String(), ".", "/") + ";")
}

func (cls *Class) GetSuperclass() java.IClass {
	if cls.IsInterface() || cls.IsPrimitive() {
		return nil
	}
	if cls.superClass == nil && cls.ift == nil {
		cls.superClass = getSuperclass(cls.typ)
	}
	return cls.superClass
}

func (cls *Class) GetInterfaces() []java.IClass {
	if cls.interfaces == nil {
		cls.interfaces = getInterfaces(cls.typ)
	}
	return slices.Clone(cls.interfaces)
}

func (cls *Class) GetDeclaredConstructor(parameterTypes ...java.IClass) *lang_reflect.Constructor {
	if !cls.static.IsValid() {
		return nil
	}
	name, _, _ := strings.Cut(typeIndirect(cls.typ).Name(), "[")
	publicName := publicName(name)
	privateName := privateName(name)
	count := len(parameterTypes) + 1
	method, ok := matchDeclaredMethod(cls.static.Type(), func(method *reflect.Method) bool {
		if !method.Func.IsValid() {
			return false
		} else if count != method.Type.NumIn() {
			return false
		} else if !isOverload(method.Name, publicName) && !isOverload(method.Name, privateName) {
			return false
		}
		for i, in := range rangeMethodIn(method.Type, 1) {
			if !(parameterTypes[i].Equals(getClass(in))) {
				return false
			}
		}
		return true
	})
	if !ok {
		return nil
	}
	types := getMethodArgTypes(method)[1:]
	in := make([]reflect.Value, len(parameterTypes)+1)
	in[0] = cls.static
	return lang_reflect.NewConstructor(cls, parameterTypes, make([]java.IClass, 0), getMethodModifier(method), func(_ java.IObject, args ...any) any {
		for i := range args {
			if args[i] == nil {
				in[i+1] = reflect.Zero(types[i])
			} else {
				in[i+1] = reflect.ValueOf(args[i])
			}
		}
		out := method.Func.Call(in)
		if len(out) == 0 {
			return nil
		}
		return out[0].Interface()
	})
}

func (cls *Class) GetDeclaredField(name java.IString) *lang_reflect.Field {
	str := name.String()
	names := []string{publicName(str), privateName(str)}
	types := []reflect.Type{cls.typ, nil}
	if cls.static.IsValid() {
		types[1] = cls.static.Type()
	} else {
		types = types[:1]
	}
	typ, field := matchField(types, names)
	if typ == nil {
		return nil
	}
	mod := getFieldModifier(&field)
	var staticObj any
	if slices.Index(types, typ) == 1 {
		mod |= lang_reflect.Modifier_STATIC
		staticObj = cls.static.Interface()
	}
	return lang_reflect.NewField(cls, name, getClass(field.Type), mod, func(obj java.IObject) any {
		if staticObj != nil {
			return getFieldValue(staticObj, &field)
		} else if typ := reflect.TypeOf(obj); typ.Kind() != reflect.Pointer {
			// panic
			return nil
		} else if field, ok := typ.Elem().FieldByName(field.Name); ok {
			return getFieldValue(obj, &field)
		} else {
			// panic
			return nil
		}
	}, func(obj java.IObject, value any) {
		var o any
		var f *reflect.StructField
		if staticObj != nil {
			o = staticObj
			f = &field
		} else if typ := reflect.TypeOf(obj); typ.Kind() != reflect.Pointer {
			// panic
			return
		} else if field, ok := typ.Elem().FieldByName(field.Name); ok {
			o = obj
			f = &field
		} else {
			// panic
			return
		}
		if !reflect.TypeOf(value).AssignableTo(f.Type) {
			// panic
			return
		}
		setFieldValue(o, f, value)
	})
}

func (cls *Class) GetDeclaredMethod(name java.IString, parameterTypes ...java.IClass) *lang_reflect.Method {
	str := name.String()
	publicName := publicName(str)
	privateName := privateName(str)
	match := func(method *reflect.Method, offset int) bool {
		if offset+len(parameterTypes) != method.Type.NumIn() {
			return false
		} else if !isOverload(method.Name, publicName) && !isOverload(method.Name, privateName) {
			return false
		}
		for i, in := range rangeMethodIn(method.Type, offset) {
			if !(parameterTypes[i].Equals(getClass(in))) {
				return false
			}
		}
		return true
	}
	var mod java.JInt
	method, ok := matchDeclaredMethod(cls.typ, func(method *reflect.Method) bool {
		if method.Func.IsValid() {
			return match(method, 1)
		}
		return match(method, 0)
	})
	if !ok {
		if !cls.static.IsValid() {
			return nil
		}
		method, ok = matchDeclaredMethod(cls.static.Type(), func(method *reflect.Method) bool {
			if !method.Func.IsValid() {
				return false
			}
			return match(method, 1)
		})
		if !ok {
			return nil
		}
		mod |= lang_reflect.Modifier_STATIC
	}
	mod |= getMethodModifier(method)
	types := getMethodArgTypes(method)
	if method.Func.IsValid() {
		types = types[1:]
	}
	in := make([]reflect.Value, 1+len(parameterTypes))
	isStatic := mod&lang_reflect.Modifier_STATIC != 0
	if isStatic {
		in[0] = cls.static
	}
	return lang_reflect.NewMethod(cls, name, getMethodReturnClass(method), parameterTypes, make([]java.IClass, 0), mod, func(obj java.IObject, args ...any) any {
		f := &method.Func
		if isStatic {
		} else {
			typ := reflect.TypeOf(obj)
			funcType := method.Type
			if f.IsValid() && cls.Equals(obj.GetClass()) {
			} else {
				method, ok := matchMethod(typ, func(method *reflect.Method) bool {
					if !method.Func.IsValid() {
						return false
					}
					return match(method, 1)
				})
				if !ok {
					// panic
					return nil
				}
				f = &method.Func
				funcType = method.Type
			}
			if rcvrType := funcType.In(0); typ != rcvrType {
				in[0] = reflect.ValueOf(castPointer(obj, rcvrType))
			} else {
				in[0] = reflect.ValueOf(obj)
			}
		}
		for i := range args {
			if args[i] == nil {
				in[i+1] = reflect.Zero(types[i])
			} else {
				in[i+1] = reflect.ValueOf(args[i])
			}
		}
		out := f.Call(in)
		if len(out) == 0 {
			return nil
		}
		return out[0].Interface()
	})
}

func (cls *Class) GetConstructor(parameterTypes ...java.IClass) *lang_reflect.Constructor {
	if !cls.static.IsValid() {
		return nil
	}
	name, _, _ := strings.Cut(typeIndirect(cls.typ).Name(), "[")
	name = publicName(name)
	count := len(parameterTypes) + 1
	method, ok := matchDeclaredMethod(cls.static.Type(), func(method *reflect.Method) bool {
		if !method.IsExported() {
			return false
		} else if !method.Func.IsValid() {
			return false
		} else if count != method.Type.NumIn() {
			return false
		} else if !isOverload(method.Name, name) {
			return false
		}
		for i, in := range rangeMethodIn(method.Type, 1) {
			if !(parameterTypes[i].Equals(getClass(in))) {
				return false
			}
		}
		return true
	})
	if !ok {
		return nil
	}
	if !ok {
		return nil
	}
	types := getMethodArgTypes(method)[1:]
	in := make([]reflect.Value, 1+len(parameterTypes))
	in[0] = cls.static
	return lang_reflect.NewConstructor(cls, parameterTypes, make([]java.IClass, 0), lang_reflect.Modifier_PUBLIC, func(_ java.IObject, args ...any) any {
		for i := range args {
			if args[i] == nil {
				in[i+1] = reflect.Zero(types[i])
			} else {
				in[i+1] = reflect.ValueOf(args[i])
			}
		}
		out := method.Func.Call(in)
		if len(out) == 0 {
			return nil
		}
		return out[0].Interface()
	})
}

func (cls *Class) GetField(name java.IString) *lang_reflect.Field {
	names := []string{publicName(name.String())}
	types := []reflect.Type{cls.typ, nil}
	if cls.static.IsValid() {
		types[1] = cls.static.Type()
	} else {
		types = types[:1]
	}
	typ, field := matchField(types, names)
	if typ == nil {
		return nil
	}
	mod := lang_reflect.Modifier_PUBLIC
	var staticObj any
	if slices.Index(types, typ) == 1 {
		mod |= lang_reflect.Modifier_STATIC
		staticObj = cls.static.Interface()
	}
	return lang_reflect.NewField(cls, name, getClass(field.Type), mod, func(obj java.IObject) any {
		if staticObj != nil {
			return getFieldValue(staticObj, &field)
		} else if typ := reflect.TypeOf(obj); typ.Kind() != reflect.Pointer {
			// panic
			return nil
		} else if field, ok := typ.Elem().FieldByName(field.Name); ok {
			return getFieldValue(obj, &field)
		} else {
			// panic
			return nil
		}
	}, func(obj java.IObject, value any) {
		var o any
		var f *reflect.StructField
		if staticObj != nil {
			o = staticObj
			f = &field
		} else if typ := reflect.TypeOf(obj); typ.Kind() != reflect.Pointer {
			// panic
			return
		} else if field, ok := typ.Elem().FieldByName(field.Name); ok {
			o = obj
			f = &field
		} else {
			// panic
			return
		}
		if !reflect.TypeOf(value).AssignableTo(f.Type) {
			// panic
			return
		}
		setFieldValue(o, f, value)
	})
}

func (cls *Class) GetMethod(name java.IString, parameterTypes ...java.IClass) *lang_reflect.Method {
	publicName := publicName(name.String())
	match := func(method *reflect.Method, offset int) bool {
		if offset+len(parameterTypes) != method.Type.NumIn() {
			return false
		} else if !isOverload(method.Name, publicName) {
			return false
		}
		for i, in := range rangeMethodIn(method.Type, offset) {
			if !(parameterTypes[i].Equals(getClass(in))) {
				return false
			}
		}
		return true
	}
	mod := lang_reflect.Modifier_PUBLIC
	method, ok := matchExportedMethod(cls.typ, func(method *reflect.Method) bool {
		if method.Func.IsValid() {
			return match(method, 1)
		}
		return match(method, 0)
	})
	if !ok {
		if !cls.static.IsValid() {
			return nil
		}
		method, ok = matchExportedMethod(cls.static.Type(), func(method *reflect.Method) bool {
			if !method.Func.IsValid() {
				return false
			}
			return match(method, 1)
		})
		if !ok {
			return nil
		}
		mod |= lang_reflect.Modifier_STATIC
	}
	types := getMethodArgTypes(method)
	if method.Func.IsValid() {
		types = types[1:]
	}
	in := make([]reflect.Value, 1+len(parameterTypes))
	isStatic := mod&lang_reflect.Modifier_STATIC != 0
	if isStatic {
		in[0] = cls.static
	}
	return lang_reflect.NewMethod(cls, name, getMethodReturnClass(method), parameterTypes, make([]java.IClass, 0), mod, func(obj java.IObject, args ...any) any {
		f := &method.Func
		if isStatic {
		} else {
			typ := reflect.TypeOf(obj)
			funcType := method.Type
			if f.IsValid() && cls.Equals(obj.GetClass()) {
			} else {
				method, ok := matchMethod(typ, func(method *reflect.Method) bool {
					if !method.Func.IsValid() {
						return false
					}
					return match(method, 1)
				})
				if !ok {
					// panic
					return nil
				}
				f = &method.Func
				funcType = method.Type
			}
			if rcvrType := funcType.In(0); typ != rcvrType {
				in[0] = reflect.ValueOf(castPointer(obj, rcvrType))
			} else {
				in[0] = reflect.ValueOf(obj)
			}
		}
		for i := range args {
			if args[i] == nil {
				in[i+1] = reflect.Zero(types[i])
			} else {
				in[i+1] = reflect.ValueOf(args[i])
			}
		}
		out := f.Call(in)
		if len(out) == 0 {
			return nil
		}
		return out[0].Interface()
	})
}

func (cls *Class) IsInterface() java.JBoolean {
	return cls.typ.Kind() == reflect.Interface
}

func (cls *Class) IsPrimitive() java.JBoolean {
	switch cls.typ.Kind() {
	case reflect.Interface, reflect.Pointer, reflect.Slice:
		return false
	}
	return cls.typ.PkgPath() == ""
}

func (cls *Class) IsArray() java.JBoolean {
	return cls.typ.Kind() == reflect.Slice
}

func (cls *Class) IsInstance(obj java.IObject) java.JBoolean {
	if obj == nil {
		return false
	}
	return cls.IsAssignableFrom(obj.GetClass())
}

func (cls *Class) IsAssignableFrom(clazz java.IClass) java.JBoolean {
	if typeof, ok := clazz.(interface{ typeof() reflect.Type }); ok {
		typ1 := typeof.typeof()
		typ2 := cls.typeof()
		if typ1.Kind() == reflect.Pointer || typ2.Kind() == reflect.Interface {
			return assignableTo(typ2, typ1)
		}
	}
	for ; clazz != nil; clazz = clazz.GetSuperclass() {
		if cls.Equals(clazz) {
			return true
		}
	}
	return false
}

func (cls *Class) Cast(obj java.IObject) java.IObject {
	if !cls.IsInstance(obj) {
		return nil
	} else if cls.IsInterface() || cls.ift != nil || cls.Equals(obj.GetClass()) {
		return obj
	}
	return castPointer(obj, cls.typ).(java.IObject)
}

func (cls *Class) ComponentType() java.IClass {
	if cls.IsArray() {
		return getClass(cls.typ.Elem())
	}
	return nil
}

func (cls *Class) typeof() reflect.Type {
	if cls.ift != nil {
		return cls.ift
	}
	return cls.typ
}

func (staticClass) Class() java.IClass {
	return classClass
}

func (staticClass) ForName(className java.IString) java.IClass {
	if className == nil {
		return nil
	}
	return findClass(className.String())
}

func InstanceOf[T any](obj java.IObject) java.JBoolean {
	if obj == nil {
		return false
	}
	typ := reflect.TypeFor[T]()
	if typ.Kind() == reflect.Interface {
		_, ok := obj.(T)
		return ok
	}
	return getClass(typ).IsInstance(obj)
}

func ClassFor[T any]() java.IClass {
	return getClass(reflect.TypeFor[T]())
}

func ClassResolve[T any](static StaticClass) java.IClass {
	typ := reflect.TypeFor[T]()
	return resolveClass(typ, static)
}

func getObjectType(typ reflect.Type) reflect.Type {
	switch typ.Kind() {
	case reflect.Invalid, reflect.Bool, reflect.Int8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
	case reflect.Interface:
	case reflect.Slice:
		if typ.PkgPath() != "" {
			typ = reflect.SliceOf(typ.Elem())
		}
	case reflect.Struct:
		typ = reflect.PointerTo(typ)
		fallthrough
	default:
		if !typ.Implements(internal.BaseType) {
			return nil
		}
	}
	return typ
}

func getRType(typ reflect.Type) uintptr {
	return (*struct{ rtype, data uintptr })(unsafe.Pointer(&typ)).data
}

func getType(rtype uintptr) reflect.Type {
	typ := reflect.TypeOf(struct{}{})
	(*struct{ rtype, data uintptr })(unsafe.Pointer(&typ)).data = rtype
	return typ
}

func resolveClass(typ reflect.Type, static StaticClass) java.IClass {
	if typ = getObjectType(typ); typ == nil {
		return nil
	}
	rtype := getRType(typ)
	if cls, ok := classMap.Load(rtype); ok {
		return cls.(*Class)
	}
	return createClass(rtype, typ, static)
}

func createClass(rtype uintptr, typ reflect.Type, static StaticClass) *Class {
	cls := &Class{typ: typ}
	if static != nil {
		cls.static = reflect.ValueOf(static)
	}
	classMap.Store(rtype, cls)
	if typ.Kind() == reflect.Pointer {
		exportMethods(rtype)
	}
	return cls
}

func getClass(typ reflect.Type) *Class {
	if typ = getObjectType(typ); typ == nil {
		return nil
	}
	rtype := getRType(typ)
	if cls, ok := classMap.Load(rtype); ok {
		return cls.(*Class)
	}
	return &Class{typ: typ}
}

func findClass(name string) (cls *Class) {
	if strings.HasPrefix(name, "[") {
		return findArrayClass(name)
	}
	classMap.Range(func(key, value any) bool {
		v := value.(*Class)
		if v.GetName().String() == name {
			cls = v
			return false
		}
		return true
	})
	return
}

func findArrayClass(name string) *Class {
	for dimensions := 1; ; dimensions++ {
		switch name[1] {
		case 'Z':
			return booleanArrayClass
		case 'C':
			return charArrayClass
		case 'B':
			return byteArrayClass
		case 'S':
			return shortArrayClass
		case 'I':
			return intArrayClass
		case 'J':
			return longArrayClass
		case 'F':
			return floatArrayClass
		case 'D':
			return doubleArrayClass
		case 'L':
			cls := findClass(name[2 : len(name)-1])
			if cls == nil {
				break
			}
			typ := cls.typeof()
			for i := 0; i < dimensions; i++ {
				typ = reflect.SliceOf(typ)
			}
			return getClass(typ)
		case '[':
			name = name[1:]
		default:
			return nil
		}
	}
}

func getSuperclass(typ reflect.Type) java.IClass {
	switch typ.Kind() {
	case reflect.Interface:
		return nil
	case reflect.Pointer:
		return getSuperclass(typ.Elem())
	case reflect.Struct:
		if typ.NumField() == 0 {
		} else if field := typ.Field(0); !field.Anonymous {
		} else {
			return getClass(field.Type)
		}
	}
	return objectClass
}

func getInterfaces(typ reflect.Type) []java.IClass {
	switch typ.Kind() {
	case reflect.Pointer:
		typ := typ.Elem()
		ifs := make([]java.IClass, 0, typ.NumField())
		for _, field := range rangeField(typ, 0) {
			if !field.Anonymous {
				continue
			}
			if field.Type.Kind() == reflect.Interface {
				ifs = append(ifs, getClass(field.Type))
			}
		}
		return slices.Clip(ifs)
	case reflect.Slice:
		return arrayInterfaces
	}
	return make([]java.IClass, 0)
}

func aliasClass(dst, src reflect.Type) {
	if dst = getObjectType(dst); dst == nil {
		return
	}
	cls := getClass(src)
	if cls == nil {
		return
	}
	rtype := getRType(dst)
	classMap.Store(rtype, cls)
}

func newInstance(cls java.IClass, typ reflect.Type) java.IObject {
	var val reflect.Value
	if typ.Kind() == reflect.Pointer {
		val = reflect.New(typ.Elem())
		superInstance(cls, val.Elem())
	} else {
		val = reflect.Zero(typ)
		superInstance(cls, val)
	}
	obj := val.Interface()
	if o, ok := obj.(interface{ instance(cls java.IClass) }); ok {
		o.instance(cls)
	}
	return obj.(java.IObject)
}

func newArrayInstance(cls java.IClass, length int) java.IObject {
	clazz, ok := cls.(interface{ typeof() reflect.Type })
	if !ok {
		return nil
	}
	typ := reflect.SliceOf(clazz.typeof())
	return arrayPack(reflect.MakeSlice(typ, length, length).Interface())
}

func superInstance(cls java.IClass, value reflect.Value) {
	if value.Kind() != reflect.Struct {
		return
	} else if value.NumField() == 0 {
		return
	}
	field := value.Type().Field(0)
	if !field.Anonymous {
		return
	}
	switch field.Type.Kind() {
	case reflect.Interface:
		if field.Type != internal.BaseType {
			return
		}
		value.Field(0).Set(reflect.ValueOf(&Object{cls}))
	case reflect.Pointer:
		super := reflect.New(field.Type.Elem())
		value.Field(0).Set(super)
		superInstance(cls, super.Elem())
	case reflect.Struct:
		superInstance(cls, value.Field(0))
	default:
	}
}

func typeIndirect(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	return typ
}

func privateName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func publicName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func isOverload(fn, name string) bool {
	if !strings.HasPrefix(fn, name) {
		return false
	}
	suffix := fn[len(name):]
	if suffix == "" {
		return true
	} else if suffix[0] != '_' {
		return false
	}
	_, err := strconv.ParseInt(suffix[1:], 10, 0)
	return err == nil
}

func packObject(rtype, data uintptr) any {
	var i any
	e := (*struct{ rtype, data uintptr })(unsafe.Pointer(&i))
	e.rtype = rtype
	e.data = data
	return i
}

func castObject(obj java.IObject, cls java.IClass) any {
	clazz, ok := cls.(interface{ typeof() reflect.Type })
	if !ok {
		return obj
	}
	typ := clazz.typeof()
	ptr := (*struct{ rtype, data uintptr })(unsafe.Pointer(&obj)).data
	return packObject(getRType(typ), ptr)
}

func castPointerObject(obj java.IObject, cls java.IClass) any {
	clazz, ok := cls.(interface{ typeof() reflect.Type })
	if !ok {
		return obj
	}
	typ := clazz.typeof()
	ptr := *(*uintptr)((*struct{ rtype, data unsafe.Pointer })(unsafe.Pointer(&obj)).data)
	return packObject(getRType(typ), ptr)
}

func castPointer(obj any, typ reflect.Type) any {
	field, ok := reflect.TypeOf(obj).Elem().FieldByName(typeIndirect(typ).Name())
	if !ok {
		return nil
	}
	ptr := getFieldPointer(obj, &field)
	switch field.Type.Kind() {
	case reflect.Interface:
		if field.Type.NumMethod() == 0 {
			return *(*any)(ptr)
		}
		return *(*interface{ M() })(ptr)
	case reflect.Pointer:
		return packObject(getRType(field.Type), uintptr(ptr))
	}
	return packObject(getRType(reflect.PointerTo(field.Type)), uintptr(ptr))
}

func tryCast(val any, cls java.IClass) (any, bool) {
	if val == nil {
		return nil, true
	}
	switch v := val.(type) {
	case java.JBoolean:
		if cls == booleanClass {
			return Boolean(v), true
		} else if cls == SBoolean.TYPE {
			return v, true
		}
	case java.JByte:
		if cls == byteClass {
			return Byte(v), true
		} else if cls == SByte.TYPE {
			return v, true
		}
	case java.JChar:
		if cls == characterClass {
			return Character(v), true
		} else if cls == SCharacter.TYPE {
			return v, true
		}
	case java.JShort:
		if cls == shortClass {
			return Short(v), true
		} else if cls == SShort.TYPE {
			return v, true
		}
	case java.JInt:
		if cls == integerClass {
			return Integer(v), true
		} else if cls == SInteger.TYPE {
			return v, true
		}
	case java.JLong:
		if cls == longClass {
			return Long(v), true
		} else if cls == SLong.TYPE {
			return v, true
		}
	case java.JFloat:
		if cls == floatClass {
			return Float(v), true
		} else if cls == SFloat.TYPE {
			return v, true
		}
	case java.JDouble:
		if cls == doubleClass {
			return Double(v), true
		} else if cls == SDouble.TYPE {
			return v, true
		}
	case Boolean:
		if cls == SBoolean.TYPE {
			return java.JBoolean(v), true
		}
	case Byte:
		if cls == SByte.TYPE {
			return java.JByte(v), true
		}
	case Character:
		if cls == SCharacter.TYPE {
			return java.JChar(v), true
		}
	case Short:
		if cls == SShort.TYPE {
			return java.JShort(v), true
		}
	case Integer:
		if cls == SInteger.TYPE {
			return java.JInt(v), true
		}
	case Long:
		if cls == SLong.TYPE {
			return java.JLong(v), true
		}
	case Float:
		if cls == SFloat.TYPE {
			return java.JFloat(v), true
		}
	case Double:
		if cls == SDouble.TYPE {
			return java.JDouble(v), true
		}
	}
	if v, ok := val.(java.IObject); !ok {
	} else if v = cls.Cast(v); v != nil {
		return v, true
	}
	return nil, false
}

func assignableTo(dst, src reflect.Type) bool {
	if dst == src {
		return true
	} else if dKind := dst.Kind(); dKind == reflect.Interface {
		return src.AssignableTo(dst)
	} else if sKind := src.Kind(); sKind == reflect.Interface {
		return false
	} else if dKind != reflect.Pointer || sKind != reflect.Pointer {
		return false
	}
	dst = dst.Elem()
	src = src.Elem()
	field, ok := src.FieldByName(dst.Name())
	if !ok || !field.Anonymous {
		return false
	}
	typ := typeIndirect(field.Type)
	return typ == dst
}

func primitivePack(val any) java.IObject {
	switch v := val.(type) {
	case java.JBoolean:
		return Boolean(v)
	case java.JByte:
		return Byte(v)
	case java.JChar:
		return Character(v)
	case java.JShort:
		return Short(v)
	case java.JInt:
		return Integer(v)
	case java.JLong:
		return Long(v)
	case java.JFloat:
		return Float(v)
	case java.JDouble:
		return Double(v)
	}
	return nil
}

func primitiveUnpack(obj java.IObject) any {
	switch v := obj.(type) {
	case Boolean:
		return java.JBoolean(v)
	case Byte:
		return java.JByte(v)
	case Character:
		return java.JChar(v)
	case Short:
		return java.JShort(v)
	case Integer:
		return java.JInt(v)
	case Long:
		return java.JLong(v)
	case Float:
		return java.JFloat(v)
	case Double:
		return java.JDouble(v)
	}
	return nil
}

func matchField(types []reflect.Type, names []string) (reflect.Type, reflect.StructField) {
	for _, typ := range types {
		if typ.Kind() != reflect.Pointer {
			continue
		}
		for _, field := range rangeField(typ.Elem(), 0) {
			if !field.Anonymous && slices.Contains(names, field.Name) {
				return typ, field
			}
		}
	}
	return nil, reflect.StructField{}
}

func matchMethod(typ reflect.Type, f func(*reflect.Method) bool) (*reflect.Method, bool) {
	for method := range rangeMethod(typ) {
		if f(&method) {
			return &method, true
		}
	}
	return nil, false
}

func matchExportedMethod(typ reflect.Type, f func(*reflect.Method) bool) (*reflect.Method, bool) {
	return matchMethod(typ, func(method *reflect.Method) bool {
		if !method.IsExported() {
			return false
		}
		return f(method)
	})
}

func matchDeclaredMethod(typ reflect.Type, f func(*reflect.Method) bool) (*reflect.Method, bool) {
	return matchMethod(typ, func(method *reflect.Method) bool {
		if method.Func.IsValid() && method.Type.In(0) != typ {
			return false
		}
		return f(method)
	})
}

func getMethodArgTypes(method *reflect.Method) []reflect.Type {
	types := make([]reflect.Type, method.Type.NumIn())
	for i, typ := range rangeMethodIn(method.Type, 1) {
		types[i] = typ
	}
	return types
}

func getMethodReturnClass(method *reflect.Method) java.IClass {
	if method.Type.NumOut() == 0 {
		return SVoid.TYPE
	}
	return getClass(method.Type.Out(0))
}

func getMethodModifier(method *reflect.Method) java.JInt {
	first, _ := utf8.DecodeRuneInString(method.Name)
	if unicode.IsUpper(first) {
		return lang_reflect.Modifier_PUBLIC
	}
	return lang_reflect.Modifier_PRIVATE
}

func getFieldModifier(field *reflect.StructField) java.JInt {
	first, _ := utf8.DecodeRuneInString(field.Name)
	if unicode.IsUpper(first) {
		return lang_reflect.Modifier_PUBLIC
	}
	return lang_reflect.Modifier_PRIVATE
}

func getFieldPointer(obj any, field *reflect.StructField) unsafe.Pointer {
	typ := typeIndirect(reflect.TypeOf(obj))
	ptr := (*struct{ _, data unsafe.Pointer })(unsafe.Pointer(&obj)).data
	index := field.Index
	for ; len(index) != 1; index = index[1:] {
		f := typ.Field(index[0])
		typ = f.Type
		ptr = unsafe.Add(ptr, f.Offset)
		if typ.Kind() == reflect.Pointer {
			typ = typ.Elem()
			ptr = *(*unsafe.Pointer)(ptr)
		}
	}
	return unsafe.Add(ptr, field.Offset)
}

func getFieldValue(obj any, field *reflect.StructField) any {
	ptr := getFieldPointer(obj, field)
	switch field.Type.Kind() {
	case reflect.Interface:
		if field.Type.NumMethod() == 0 {
			return *(*any)(ptr)
		}
		return *(*interface{ M() })(ptr)
	case reflect.Pointer:
		return packObject(getRType(field.Type), *(*uintptr)(ptr))
	}
	return packObject(getRType(field.Type), uintptr(ptr))
}

func setFieldValue(obj any, field *reflect.StructField, value any) {
	ptr := getFieldPointer(obj, field)
	switch field.Type.Kind() {
	case reflect.Interface:
		if field.Type.NumMethod() == 0 {
			*(*any)(ptr) = value
		} else {
			ifaceE2I(getRType(field.Type), value, ptr)
		}
	case reflect.Pointer:
		*(*uintptr)(ptr) = (*struct{ _, data uintptr })(unsafe.Pointer(&value)).data
	default:
		size := field.Type.Size()
		copy(unsafe.Slice((*byte)(ptr), size), unsafe.Slice((*byte)((*struct{ _, data unsafe.Pointer })(unsafe.Pointer(&value)).data), size))
	}
}

func rangeField(typ reflect.Type, offset int) iter.Seq2[int, reflect.StructField] {
	return func(yield func(int, reflect.StructField) bool) {
		count := typ.NumField()
		for i := offset; i < count; i++ {
			if !yield(i, typ.Field(i)) {
				break
			}
		}
	}
}

func rangeMethod(typ reflect.Type) iter.Seq[reflect.Method] {
	return func(yield func(reflect.Method) bool) {
		count := typ.NumMethod()
		for i := 0; i < count; i++ {
			if !yield(typ.Method(i)) {
				break
			}
		}
	}
}

func rangeMethodIn(typ reflect.Type, offset int) iter.Seq2[int, reflect.Type] {
	return func(yield func(int, reflect.Type) bool) {
		count := typ.NumIn() - offset
		for i := 0; i < count; i++ {
			if !yield(i, typ.In(offset+i)) {
				break
			}
		}
	}
}

func init() {
	cls := classClass.(*Class)
	cls.ift = reflect.TypeFor[java.IClass]()
	aliasClass(cls.ift, cls.typ)
}
