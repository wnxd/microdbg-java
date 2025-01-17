package lang

import java "github.com/wnxd/microdbg-java"

type Comparable[T java.IObject] interface {
	CompareTo(T) java.JInt
}
