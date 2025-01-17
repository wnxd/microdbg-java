package lang

import (
	java "github.com/wnxd/microdbg-java"
)

type Void struct {
	Object
}

type staticVoid struct {
	TYPE java.IClass
}

var (
	SVoid     staticVoid
	voidClass = ClassResolve[*Void](&SVoid)
)

func (staticVoid) Class() java.IClass {
	return voidClass
}

func init() {
	SVoid.TYPE = resolveClass(voidType{}, nil)
}
