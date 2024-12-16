package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	// 汉字
	hanzi := "你好世界"
	result := getFirstLetter(hanzi)
	fmt.Printf("汉字：%s，首字母：%s\n", hanzi, result)

	// 英文
	english := "Hello"
	result = getFirstLetter(english)
	fmt.Printf("英文：%s，首字母：%s\n", english, result)
}

func getFirstLetter(s string) string {
	// 创建拼音转换器
	p := pinyin.NewArgs()

	// 获取拼音
	pinyinStr := pinyin.Pinyin(s, p)

	// 提取首字母
	var initials string
	for _, py := range pinyinStr {
		if len(py) > 0 {
			initials += string(py[0])
		}
	}

	return initials
}
