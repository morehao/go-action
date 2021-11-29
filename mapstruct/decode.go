package mapstruct

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
}

type Friend1 struct {
	Person
}

type Friend2 struct {
	Person `mapstructure:",squash"` // 解析内嵌的结构体需要增加tag： `mapstructure:",squash"`
}

func Decode() {
	datas := []string{`
    {
      "type": "friend1",
      "person": {
        "name":"dj"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Printf("friend1:%+v", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Printf("friend2:%+v", f2)
		}
	}
}
