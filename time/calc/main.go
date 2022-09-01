package main

import (
	"fmt"
	"time"
)

func main() {
	//某个时间
	//oneTimeStr := "2018-01-08 00:00:00"
	oneTimeStr := "2021-04-05 00:00:00"
	//时间转换成time.Time格式
	t, err := time.ParseInLocation("2006-01-02 15:04:05", oneTimeStr, time.Local)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取这个时间的基于这一年有多少天了
	yearDay := t.YearDay()
	//获取上一年的最后一天
	yesterdayYearEndDay := t.AddDate(0, 0, -yearDay)
	//获取上一年最后一天是星期几
	dayInWeek := int(yesterdayYearEndDay.Weekday())
	//第一周的总天数,默认是7天
	firstWeekDays := 7
	//如果上一年最后一天不是星期天，则第一周总天数是7-dayInWeek
	if dayInWeek != 0 {
		firstWeekDays = 7 - dayInWeek
	}
	week := 0
	//如果这一年的总天数小于第一周总天数，则是第一周，否则按照这一年多少天减去第一周的天数除以7+1 但是要考虑这一天减去第一周天数除以7会取整型，
	//所以需要处理两个数取余之后是否大于0，如果大于0 则多加一天，这样自然周就算出来了。
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		plusDay := 0
		if (yearDay-firstWeekDays)%7 > 0 {
			plusDay = 1
		}
		week = (yearDay-firstWeekDays)/7 + 1 + plusDay
	}
	fmt.Printf("%d%d", t.Year(), week)
}
