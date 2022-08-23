package classfile

type ConstantClassInfo struct { // 代表类或者接口的符号引用
	cp ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader * ClassReader){
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string{
	return self.cp.getUtf8(self.nameIndex)
}

