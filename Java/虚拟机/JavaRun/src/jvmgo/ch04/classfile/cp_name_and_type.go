package classfile

// Java虚拟机定义了一种简单的语法描述字段和方法
// 然后按照规则生成描述符
// 详细描述 见P137

type ConstantNameAndTypeInfo struct {
	nameIndex		uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo)readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}



