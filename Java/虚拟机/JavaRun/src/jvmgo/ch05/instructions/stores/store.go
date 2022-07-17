package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type LSTORE struct {
	base.Index8Instruction
}
func _lstore(frame *rtda.Frame, index uint){
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}
func (self *LSTORE_1) Execute(frame *rtda.Frame){
	_lstore(frame, 1)
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
func (self *LSTORE_2) Execute(frame *rtda.Frame){
	_lstore(frame, 2)
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}
func (self *LSTORE_3) Execute(frame *rtda.Frame){
	_lstore(frame, 3)
}