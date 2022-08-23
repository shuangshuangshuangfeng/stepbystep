package main

import (
	"flag"
)

import "fmt"
import "os"

type Cmd struct {
	helpFlag	 bool
	versionFlag  bool
	XjreOption   string
	cpOption     string
	class        string
	args 		 []string
}

func printUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}


func parseCmd() *Cmd{
	cmd :=&Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version")

	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0{
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}
