package classfile

import (
	"fmt"
)

// 反应了Java虚拟机规范定义的class文件格式
type ClassFile struct { // Class文件定义
	// magic uint 32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}


// 将byte[]解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error){
	defer func(){
		if r:=recover(); r != nil{
			var ok bool
			err, ok = r.(error)
			if !ok{
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 依次调用其他方法解析class文件
func (self *ClassFile) read(reader *ClassReader){
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16() // 类名  存的是常量池索引
	self.superClass = reader.readUint16() // 超类名 存的是常量池索引
	self.interfaces = reader.readUint16s() // 接口索引表 ， 存的是常量池索引
	self.fields = readMembers(reader, self.constantPool) //字段表
	self.methods = readMembers(reader, self.constantPool) // 方法表
	self.attributes = readAttributes(reader, self.constantPool) // 属性表

}


// 魔数， 很多文件格式都会规定满足该格式的文件必须以几个固定字节开头
// 这几个字节叫魔数，起到标识作用
// class文件的魔数是： 0xCAFEBABE
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 魔数之后是class文件的次版本号和主版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader){
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	switch  self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52, 58:
		if self.minorVersion == 0{
			return
		}

	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16{
	return self.minorVersion
}


func (self *ClassFile) MajorVersion() uint16{
	return self.majorVersion
}

// 版本号之后是常量池
func (self *ClassFile) ConstantPool() ConstantPool{
	return self.constantPool
}

// 类的访问标志
// 指出class文件定义的是类还是接口，访问级别是public还是private
func (self *ClassFile) AccessFlags() uint16{
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo{
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo{
	return self.methods
}

func (self *ClassFile) ClassName() string{ // 从常量池种查找类名
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string{ // 从常量池种查找超类名
	if self.superClass > 0{
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}


func (self *ClassFile) InterfaceNames() []string{ // 从常量池种查找接口名
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

