package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low int32
	high int32
	jumpOffsets []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	index := stack.PopInt()
	var offset int
	if index >= self.low && index <= self.high{
		offset = int(self.jumpOffsets[index-self.low])
	}else{
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}


