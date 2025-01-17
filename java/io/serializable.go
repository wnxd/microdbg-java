package io

import (
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-java/internal"
)

type Serializable interface {
}

type staticSerializable struct {
}

var (
	SSerializable     staticSerializable
	serializableClass = internal.ClassResolve[Serializable](&SSerializable)
)

func (staticSerializable) Class() java.IClass {
	return serializableClass
}
