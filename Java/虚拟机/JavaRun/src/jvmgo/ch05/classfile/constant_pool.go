package classfile


// 常量池是一个表， 但是有三点需要注意：
// 1. 表头给出的常量池比实际大1
// 2. 有效常量池的索引是1~n-1
// 3. CONSTANT_Long_info 和 CONSTANT_Double_info 各占两个位置
type ConstantPool []ConstantInfo // 常量池

func readConstantPool(reader *ClassReader) ConstantPool{
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i:=1; i<cpCount; i++{
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

// 按照索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	if cpInfo := self[index]; cpInfo != nil{
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 从常量池查找字段或者方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16)(string, string){
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string{
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)

	return self.getUtf8(classInfo.nameIndex)
}

// 从常量池种查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string{
	utf8info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8info.str
}




