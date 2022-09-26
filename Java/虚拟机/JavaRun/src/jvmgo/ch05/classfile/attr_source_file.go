package classfile

type SourceFileAttribute struct {
	cp ConstantPool
	SourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader){
	self.SourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string{
	return self.cp.getUtf8(self.SourceFileIndex)
}

