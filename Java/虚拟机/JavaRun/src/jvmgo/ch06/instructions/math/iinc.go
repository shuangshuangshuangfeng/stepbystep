package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type INNC struct {
	Index uint
	Const int32
}

// 从字节码中读取操作数
func (self *INNC) FetchOperands(reader *base.BytecodeReader){
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *INNC) Execute(frame *rtda.Frame){
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}




