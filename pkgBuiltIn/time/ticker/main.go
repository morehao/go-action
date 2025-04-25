package main

import (
	"fmt"
	"time"
)

func main() {
	tickerDemo()
	tickerLaunch()
}

func tickerDemo() {
	// 周期性地打印日志
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		fmt.Println("tickerDemo")
	}
}

func tickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 30 // 每车最大装载人数
	passengers := make([]string, 0, maxPassenger)
	for {
		passenger := getNewPassenger() // 获取一个新乘客
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}
		select {
		case <-ticker.C:
			launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger {
				launch(passengers)
				passengers = []string{}
			}

		}
	}
}

func getNewPassenger() string {
	return "passenger"
}

func launch(passengers []string) {
	fmt.Println("The bus launch!")
}
