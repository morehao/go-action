## 对于其他类型，只要实现了`String`方法，就可以定义该类型的字符串表示形式。

```go
package main

import "fmt"

func main() {
	studentInfo := student{
		Name: "张三",
		Age:  1,
	}
	fmt.Println(studentInfo)
}

type student struct {
	Name string
	Age  int64
}

func (s student) String() string {
	return fmt.Sprintf("我是%s, 今年%d岁", s.Name, s.Age)
}
```