package internal

import (
	_ "unsafe"

	java "github.com/wnxd/microdbg-java"
)

func StringPack(str string) java.IString {
	return stringPack(str)
}

//go:linkname stringPack github.com/wnxd/microdbg-java/java/lang.stringPack
func stringPack(str string) java.IString
