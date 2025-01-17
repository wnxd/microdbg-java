package main

import (
	"fmt"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/java/lang"
	lang_reflect "github.com/wnxd/microdbg-java/java/lang/reflect"
)

type N interface {
	CC()
}

type I interface {
	N
	Test()
}

type B struct {
	java.IObject
	I
}

type C struct {
	B
	test java.IObject
}

func (b *B) call() I {
	fmt.Printf("%p %T: B.Call\n", b, b)
	return b
}

// func (c *C) call() *C {
// 	return c
// }

var (
	_, _ = lang.Export.MethodByName("call")
)

func main() {
	cls := lang.SString.Class().(*lang.Class)
	ctor := cls.GetDeclaredConstructor(lang.SArray[java.JChar]().Class())
	fmt.Println(ctor.ToString())
	method := cls.GetDeclaredMethod(lang.String("substring"), lang.SInteger.TYPE, lang.SInteger.TYPE)
	fmt.Println(method.ToString())

	str := ctor.NewInstance(lang.ArrayOf([]java.JChar{'a', 'b', 'c', 'd', 'e', 'f', 'g'}))
	fmt.Println(str)

	sub := method.Invoke(str, lang.Integer(1), lang.Integer(3))
	fmt.Println(sub)

	var obj java.IObject = lang.ClassResolve[C](nil).NewInstance()

	// mm.Func.Call([]reflect.Value{reflect.ValueOf(obj)})

	cls = obj.GetClass().(*lang.Class)

	// field := cls.GetDeclaredField(lang.String("test"))
	// fmt.Println(field.ToString())

	method = cls.GetDeclaredMethod(lang.String("call"))
	fmt.Println(method.ToString())

	fmt.Println(method.Invoke(obj))

	method = lang.ClassResolve[B](nil).(*lang.Class).GetDeclaredMethod(lang.String("call"))
	fmt.Println(method.ToString())

	fmt.Println(method.Invoke(obj))

	field := cls.GetDeclaredField(lang.String("test"))
	fmt.Println(field.ToString())

	fmt.Println(field.Get(obj))

	field.Set(obj, lang.Boolean(true))

	fmt.Println(field.Get(obj))

	fmt.Println(obj.GetClass().GetSuperclass().ToString())

	fmt.Println(lang.InstanceOf[lang.Object](obj))

	arr := lang.ArrayOf(make([]lang.Boolean, 5))
	fmt.Println(arr.GetClass().GetName())

	fmt.Println(lang.SClass.ForName(lang.String("[Ljava.lang.Object;")).ToString())

	method = lang.ClassFor[lang.Number]().(*lang.Class).GetMethod(lang.String("intValue"))
	fmt.Println(method.ToString())

	obj = method.Invoke(lang.Float(5))
	fmt.Println(obj.ToString())

	obj = lang_reflect.SArray.NewInstance(lang.SClass.ForName(lang.String("java.lang.Object")), 5)
	fmt.Println(obj.ToString())

	fmt.Println(obj.GetClass().Equals(lang.ClassFor[[]lang.String]()))

	lang_reflect.SArray.Set(obj, 0, lang.String("w"))

	fmt.Println(lang_reflect.SArray.Get(obj, 0))

	arr = lang.ArrayOf(make([]lang.Boolean, 5))

	lang_reflect.SArray.Set(arr, 0, lang.Boolean(true))

	fmt.Println(lang_reflect.SArray.Get(arr, 0))
}
