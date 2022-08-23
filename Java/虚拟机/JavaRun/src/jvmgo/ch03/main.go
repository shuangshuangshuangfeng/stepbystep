package main


import (
	"fmt"
	"strings"
)

import "jvmgo/ch03/classpath"
import "jvmgo/ch03/classfile"


func startJVM(cmd  *Cmd){
	fmt.Printf("start JVM ...\n")
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath:%v\nclass:%v\nargs:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Printf(cmd.class)
	printClassInfo(cf)
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