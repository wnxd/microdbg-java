package reflect

import java "github.com/wnxd/microdbg-java"

type Member interface {
	GetDeclaringClass() java.IClass
	GetName() java.IString
	GetModifiers() java.JInt
	IsSynthetic() java.JBoolean
}
