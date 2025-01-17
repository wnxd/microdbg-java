package reflect

import java "github.com/wnxd/microdbg-java"

type AccessibleObject struct {
	java.IObject
}

func (obj *AccessibleObject) SetAccessible(flag java.JBoolean) {
}
