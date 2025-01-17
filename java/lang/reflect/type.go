package reflect

import java "github.com/wnxd/microdbg-java"

type Type interface {
	GetTypeName() java.IString
}
