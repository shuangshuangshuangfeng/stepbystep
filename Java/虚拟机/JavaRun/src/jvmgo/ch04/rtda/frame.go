package rtda

// Java虚拟机栈
// 其中保存着方法执行的状态，主要由局部变量表和操作数栈
type Frame struct {
	lower *Frame
	localVars LocalVars // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame{
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame)LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack{
	return self.operandStack
}

