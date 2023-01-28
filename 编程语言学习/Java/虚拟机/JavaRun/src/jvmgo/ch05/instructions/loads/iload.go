package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ILOAD struct {
	base.Index8Instruction
}
func _iload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame){
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_0) Execute(frame *rtda.Frame){
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_1) Execute(frame *rtda.Frame){
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_2) Execute(frame *rtda.Frame){
	_iload(frame, 2)
}


type ILOAD_3 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_3) Execute(frame *rtda.Frame){
	_iload(frame, 3)
}



