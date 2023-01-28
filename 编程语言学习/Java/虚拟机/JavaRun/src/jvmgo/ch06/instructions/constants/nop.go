package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self NOP) Execute(frame *rtda.Frame){
	// 什么都不做
}
