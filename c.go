package java

type VaList interface {
	Extract(...any) error
}

type Ptr interface {
	Address() uintptr
}

type AnyPtr uintptr

type TypePtr[V any] interface {
	Ptr
	Get(index int) (V, error)
	Set(index int, value V) error
	ReadAt(b []V, off int64) (n int, err error)
	WriteAt(b []V, off int64) (n int, err error)
}

func (u AnyPtr) Address() uintptr {
	return uintptr(u)
}
