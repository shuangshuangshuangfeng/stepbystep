package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
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
func (self *DCONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushDouble(0)
}


type DCONST_1 struct {
	base.NoOperandsInstruction
}
func (self *DCONST_1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushDouble(1)
}

// float
type FCONST_0 struct {
	base.NoOperandsInstruction
}
func (self *FCONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(0)
}



type FCONST_1 struct {
	base.NoOperandsInstruction
}
func (self *FCONST_1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(1)
}



type FCONST_2 struct {
	base.NoOperandsInstruction
}
func (self *FCONST_2) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(2)
}


// int
// iconst_m1指令把int型-1推入操作数栈顶
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}
func (self *ICONST_0) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}
func (self *ICONST_1) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(1)
}


type ICONST_2 struct {
	base.NoOperandsInstruction
}
func (self *ICONST_2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(2)
}


type ICONST_3 struct {
	base.NoOperandsInstruction
}
func (self *ICONST_3) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(3)
}


type ICONST_4 struct {
	base.NoOperandsInstruction
}
func (self *ICONST_4) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(4)
}


type ICONST_5 struct {
	base.NoOperandsInstruction
}
func (self *ICONST_5) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(5)
}

// long
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}
func (self *LCONST_1) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushLong(1)
}