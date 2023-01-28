package classfile


// 属性表中存放的属性名是常量池索引
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// 获取单个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo{
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

// 创建属性实例
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool)AttributeInfo{
	// Java 虚拟机规范中预定义了23种属性，先解析其中的8种
	switch attrName {
	case "Code": return &CodeAttribute{cp:cp}
	case "ConstantValue": return &ConstantValueAttribute{}
	//case "Deprecated": return &DeprecatedAttribute{}
	case "Exceptions": return &ExceptionsAttribute{}
	case "LineNumberTable": return &LineNumberTableAttribute{}
	//case "LocalVariableTable": return &LocalVariableTableAttribute{}
	case "SourceFile": return &SourceFileAttribute{cp: cp}
	//case "Synthetic": return &SyntheticAttribute{}
	default:
		return &UnparseAttribute{attrName, attrLen, nil}
	}
}

func readAttributes(reader *ClassReader, cp ConstantPool)[]AttributeInfo{
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i:= range attributes{
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}