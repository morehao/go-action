package utils

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"unicode"
)

func GetFirstLetter(str string) string {
	var res = ""
	for _, v := range str {
		fmt.Println("isLetter:", unicode.IsLetter(v))
		isChinese := unicode.Is(unicode.Han, v)
		if isChinese {
			p := pinyin.NewArgs()
			p.Style = pinyin.FirstLetter
			word := pinyin.LazyPinyin(str, p)
			res = word[0]
			break
		}
		// 不要调整代码顺序
		if unicode.IsLetter(v) {
			res = string(v)
		} else {
			res = "#"
		}
		break
	}
	return strings.ToUpper(res)
}
