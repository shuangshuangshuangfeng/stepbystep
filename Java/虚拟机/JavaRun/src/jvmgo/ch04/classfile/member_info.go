package classfile



type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes [] AttributeInfo
}

// cp 常量池指针
// 读取字段表或者方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo{
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members{
		members[i] = readMember(reader, cp)
	}
	return members
}


// 读取字段或者方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo{
	return &MemberInfo{
		cp : cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16{
	return self.accessFlags
}

func (self *MemberInfo) Name() string{
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string{
	return self.cp.getUtf8(self.descriptorIndex)
}

