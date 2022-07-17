package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}
func (self *ACONST_NULL) Execute(frame *rtda.Frame){
	frame.OperandStack().PushRef(nil)
}

// double
// dconst_0指令把double型0推入操作数栈顶
type DCONST_0 struct {
	base.NoOperandsInstruction
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

// float
type FCONST_0 struct {
	base.NoOperandsInstruction
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

// int
// iconst_m1指令把int型-1推入操作数栈顶
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}

// long
type LCONST_0 struct {
	base.NoOperandsInstruction
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}
