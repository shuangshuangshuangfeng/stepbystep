package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type BIPUSH struct {
	val int8 // Push byte
}
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16 // push short
}



