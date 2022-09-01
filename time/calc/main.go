package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(WeekIntervalTime(0))
	// fmt.Println(MonthIntervalTime(0))
	// fmt.Println(SubWeek(time.Unix(1661184000, 10), time.Now()))
	fmt.Println(SubMonth(time.Unix(1661184000, 10), time.Now(), true))
}

// WeekIntervalTime 获取某周的开始和结束时间,week为0本周,-1上周，1下周以此类推
func WeekIntervalTime(week int) (startTime, endTime string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	// 周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}
	year, month, day := now.Date()
	thisWeek := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	startTime = thisWeek.AddDate(0, 0, offset+7*week).Format("2006-01-02") + " 00:00:00"
	endTime = thisWeek.AddDate(0, 0, offset+6+7*week).Format("2006-01-02") + " 23:59:59"
	return startTime, endTime
}

// MonthIntervalTime 获取某月的开始和结束时间mon为0本月,-1上月，1下月以此类推
func MonthIntervalTime(mon int) (startTime, endTime string) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	startTime = thisMonth.AddDate(0, mon, 0).Format("2006-01-02") + " 00:00:00"
	endTime = thisMonth.AddDate(0, mon+1, -1).Format("2006-01-02") + " 23:59:59"
	return startTime, endTime
}

func SubWeek(startTime, endTime time.Time) int {
	subHours := endTime.Sub(startTime).Hours()
	weeks := subHours / (7 * 24)
	return int(weeks)
}

// SubMonth 计算日期相差多少月, isNatureMonth:true-只关心自然月的差值，false-需要关心两个时间的日期大小
func SubMonth(startTime, endTime time.Time, isNatureMonth bool) (month int) {
	startY := startTime.Year()
	endY := endTime.Year()
	startM := int(startTime.Month())
	endM := int(endTime.Month())
	startD := startTime.Day()
	endD := endTime.Day()

	yearInterval := endY - startY
	// 如果 endDay的 月-日 小于 startDay的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if endM < startM || endM == startM && endD < startD {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (endM + 12) - startM
	if !isNatureMonth && endD < startD {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}
