package service

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestUser(t *testing.T) {
	user := &User{
		Username: "zhangsan",
		Age:      20,
	}
	// 转换为protobuf
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
