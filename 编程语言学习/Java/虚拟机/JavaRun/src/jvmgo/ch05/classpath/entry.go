package classpath

import (
	"os"
	"strings"
)

/*
类路径想象成一个大的整体，它由启动类路径、扩展类路径和用户类路径三个小路径构成。
三个小路径又分别由更小的路径构成。是不是很像组合模式（composite pattern）
 */


// 用于存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)


// 定义Entry接口
type Entry interface{
	// 负责寻找和加载类
	readClass(className string) ([]byte, Entry, error) //比如读取java.lang.Object类， 传入的参数应该是java/lang/Object.class
	// 作用相当于Java种得toString， 返回变量得字符串表示
	String() string
}

func newEntry(path string) Entry{
	if strings.Contains(path, pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, "zip"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}


