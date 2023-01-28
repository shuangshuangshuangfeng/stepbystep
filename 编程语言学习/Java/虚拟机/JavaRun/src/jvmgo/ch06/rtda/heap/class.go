package heap

import (
	"jvmgo/ch06/classfile"
	"os"
)

// 可以放进方法区的类
type Class struct {
	accessFlags        uint16 // 类的访问标志
	name               string // 类名
	superClassName     string // 父类名
	interfaceNames     []string
	constantPool       *classfile.ConstantPool // 运行时常量指针池
	fields             []*Field                // 字段表
	methods            []*Method               // 方法表
	loader             *ClassLoader            // 类加载器指针
	superClass         *Class
	interfaces         []*Class
	instanceSlotCount uint   // 实例变量占据的空间大小
	staticSlotCount    uint   // 类变量占据的空间大小
	staticVars         Slots // 静态变量
}

// 根据class文件创建class对象
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}


