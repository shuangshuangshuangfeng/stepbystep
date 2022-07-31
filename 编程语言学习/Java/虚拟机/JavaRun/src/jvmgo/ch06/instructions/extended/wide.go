package extended

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader){
	cpcode := reader.ReadUint8()
	switch cpcode {
		// todo 这里没写， p283
	}
}

func (self *WIDE) Execute(frame *rtda.Frame){
	self.modifiedInstruction.Execute(frame)
}


