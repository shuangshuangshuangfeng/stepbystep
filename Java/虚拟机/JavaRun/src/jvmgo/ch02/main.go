package main


import (
	"fmt"
	"strings"
)

import "jvmgo/ch02/classpath"


func startJVM(cmd  *Cmd){
	fmt.Printf("start JVM ...\n")
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v\n class:%v\n args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil{
		fmt.Printf("Could not fund or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data :%v\n", classData)
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