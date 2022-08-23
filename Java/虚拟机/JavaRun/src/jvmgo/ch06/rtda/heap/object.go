package heap

type Object struct { // 创建一个临时结构体，用于表示对象
	// todo
	class  *Class // 存放对象的Class指针
	fields Slots  // 存放实例变量
}
