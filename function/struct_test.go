package function

import (
	"fmt"
	"testing"
)

func Test_AnimalCategory(t *testing.T) {
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	/*如果Animal未实现String方法，则递归调用子字段的String方法；如果Animal实现了String方法，则子字段String方法会被屏蔽
	如果想调用子字段的String函数，需要通过链式调用的方式进行*/
	fmt.Printf("The animal: %s\n", animal)
	fmt.Println(animal.GetAnimalName())
	// 指针类型的接受者才会修改原值
	animal.SetAnimalName("newName")
	fmt.Println(animal.GetAnimalName())
}

type Food interface {
	setColor(int)
	setPrice(int)
}

type Apple struct {
	price int
	color int
}

func (a Apple) setColor(cc int) {
	a.color = cc
}
func (a *Apple) setPrice(pp int) {
	a.price = pp
}

func InterfaceTest3() {
	var apple1 Food
	apple1 = Apple{price: 15, color: 3} // 这里提示错误，值接收者没有setPrice方法
	apple1.setColor(1)
	apple1.setPrice(16)

	var apple2 Food
	apple2 = &Apple{price: 15, color: 3}
	apple2.setColor(1)
	apple2.setPrice(16)
}
