package rtda



/*   Java运行时数据区示意图

 |------------Thread------------|
 |            (pc)              |
 | |-------(Java Stack)-------| |
 | | |--------Frame----------|| |
 | | |    Local Variable     || |
 | | |    Operand Stack      || |
 | | |-----------------------|| |
 | |--------------------------| |
 |------------------------------|


*/
type Thread struct {
	pc int
	stack *Stack
}

func NewThread() *Thread{
	return &Thread{
		stack: newStack(1024),
	}
	return nil
}

func (self *Thread) PC() int{
	return self.pc
}

func (self *Thread) SetPC(pc int){
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame){
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame{
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame{
	return NewFrame(self, maxLocals, maxStack)
}

