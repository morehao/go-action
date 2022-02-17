package main

type Student struct {
	Name string
	Age  int
}

// 函数StudentRegister()内部的s为局部变量，其值通过函数返回值返回，s本身为一个指针，其指向的内存地址不会是栈，而是堆
func StudentRegister(name string, age int) *Student {
	s := new(Student) // 局部变量s逃逸到堆中
	s.Name = name
	s.Age = age
	return s
}

func main() {

	StudentRegister("xiaoming", 18)
}
