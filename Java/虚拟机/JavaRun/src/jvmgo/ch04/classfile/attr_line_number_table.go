package classfile

// LineNumberTable属性表存放方法的行号信息
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute)readInfo(reader *ClassReader){
	LineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, LineNumberTableLength)
	for i := range self.lineNumberTable{
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
