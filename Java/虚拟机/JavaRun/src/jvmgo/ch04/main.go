package main

import (
	"fmt"
)

import "jvmgo/ch04/classpath"
import "jvmgo/ch04/classfile"
import "jvmgo/ch04/rtda"


func startJVM(cmd  *Cmd){
	fmt.Printf("start JVM ...\n")
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

// 测试局部变量表
func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

// 测试操作数栈
func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile{
	classData, _, err := cp.ReadClass(className)
	if err != nil{
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil{
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile){

	fmt.Printf("version : %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constant count: %v \n", len(cf.ConstantPool()))
	fmt.Printf("access flags : 0x%x \n", cf.AccessFlags())
	fmt.Printf("this class : %v \n", cf.ClassName())
	fmt.Printf("super class : %v \n", cf.SuperClassName())
	fmt.Printf("interfaces : %v \n", cf.InterfaceNames())
	fmt.Printf("field counts : %v \n", len(cf.Fields()))

	for _, f := range cf.Fields(){
		fmt.Printf("\t %s \n", f.Name())
	}

	fmt.Printf("mathods count : %v \n", len(cf.Methods()))
	for _, m := range cf.Methods(){
		fmt.Printf("\t %s \n", m.Name())
	}
}


func main() {
	cmd := parseCmd()
	if cmd.versionFlag{
		fmt.Println("version 0.0.1")
	}else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	}else{
		startJVM(cmd)
	}
}

// 1.13 后默认使用go mod 1.11
// 1.16 embed
// 1.18 后支持泛型