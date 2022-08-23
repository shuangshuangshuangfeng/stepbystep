package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
)

type ClassLoader struct { // 类加载器
	cp       *classpath.Classpath
	classMap map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class // 类已经加载
	}
	return self.loadNonArrayClass(name)
}

// 数组类加载
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 1. 找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
	// 2. 解析class文件，并生成虚拟机可以使用的类数据，并放入方法区
	class := self.defineClass(data)
	// 3. 连接
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

// 类的链接
func link(class *Class) {
	// 1. 验证
	verify(class)
	// 2. 准备
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// 给类变量分配空间并给与初始值
func prepare(class *Class) {
	// 1. 计算实例字段的个数
	calcInstanceFieldSlotIds(class)
	// 2. 计算静态字段的个数，同时给他们编号
	calcStaticFieldSlotIds(class)
	// 3. 给类变量分配空间，并给他们赋初值
	allocAndInitStaticVars(class)
}

// 1. 计算实例字段的个数
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 2. 计算静态字段的个数，同时给他们编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 3. 给类变量分配空间，并给他们赋初值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo") // 在第8章实现
		}
	}
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	// 1. 将class文件数据转换为Class结构体
	class := parseClass(data)

	//2. superClass和interfaces字段存放超类名和直接接口表
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf) // 见6.1.1小节
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}
