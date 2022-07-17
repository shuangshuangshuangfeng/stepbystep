package classpath

import (
	"os"
	"path/filepath"
)

// 用于存放三种类型的路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// jreOption 启动路径
// cpOption 扩展路径
func Parse(jreOption, cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 判断目录是否存在
func exists(path string) bool{
	if _, err := os.Stat(path); err != nil{
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}

func getJreDir(jreOption string) string{
	if jreOption != "" && exists(jreOption){
		return jreOption
	}
	if exists("./jre"){
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != ""{
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}


func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption == ""{
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error){
	className = className+".class"
	// 启动类路径搜索class文件
	if data, entry, err := self.bootClasspath.readClass(className); err == nil{
		return data, entry, err
	}
	// 扩展类路径搜索class文件
	if data, entry, err := self.extClasspath.readClass(className); err == nil{
		return data, entry, err
	}
	// 用户类路径搜索class文件
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string{
	return self.userClasspath.String()
}

