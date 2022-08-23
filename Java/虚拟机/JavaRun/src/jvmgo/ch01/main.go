package ch01


import (
	"fmt"
)


func startJVM(cmd  *Cmd){
	fmt.Printf("start JVM ...")
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
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