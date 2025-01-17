package lang

import java "github.com/wnxd/microdbg-java"

type CharSequence interface {
	Length() java.JInt
	CharAt(java.JInt) java.JChar
	SubSequence(start, end java.JInt) CharSequence
}

var charSequenceClass = ClassResolve[CharSequence](nil)
