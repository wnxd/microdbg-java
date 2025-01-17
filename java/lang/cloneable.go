package lang

type Cloneable interface {
}

var cloneableClass = ClassResolve[Cloneable](nil)
