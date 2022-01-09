package main

import "fmt"

// 结构体内嵌
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ // 注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() // 乐乐会汪汪汪~
	d1.move() // 乐乐会动！
}

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
