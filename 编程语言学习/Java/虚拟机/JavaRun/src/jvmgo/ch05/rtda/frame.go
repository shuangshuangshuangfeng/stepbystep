package rtda

type Frame struct {
	lower *Frame
	localVars LocalVars // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
	thread *Thread
	nextPC int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame{
	return &Frame{
		thread: thread,
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

func (self *Frame) Thread() *Thread{
	return self.thread
}

func (self *Frame) NextPC() int{
	return self.nextPC
}

func (self *Frame) SetNextPC(n int){
	self.nextPC = n
}