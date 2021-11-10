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
	// fmt.Printf("The animal: %s\n", animal.AnimalCategory)
	fmt.Println(animal.GetAnimalName())
	animal.SetAnimalName("newName")
	fmt.Println(animal.GetAnimalName())
}
