package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}