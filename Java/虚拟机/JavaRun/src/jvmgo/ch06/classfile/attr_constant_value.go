package classfile

type ConstantValueAttribute struct { // 定长属性，
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader){
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16{
	return self.constantValueIndex
}
