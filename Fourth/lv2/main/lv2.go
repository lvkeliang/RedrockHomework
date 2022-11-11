package main

import (
	"fmt"
	"time"
)

func alarm(alarmName string, alarmTime string) {
	loc := time.Local
	alarmTimer, _ := time.ParseInLocation("15:04", alarmTime, loc)
	alarmTimer = alarmTimer.AddDate(int(time.Now().Year()), int(time.Now().Month())-1, int(time.Now().Day())-1)
	var wait time.Duration
	//获取第一次打印所需的等待时间
	if time.Now().Before(alarmTimer) {
		wait = alarmTimer.Sub(time.Now())
	} else {
		alarmTimer = alarmTimer.AddDate(0, 0, 1)
		wait = alarmTimer.Sub(time.Now())
	}
	//fmt.Printf("Year : %v\nmonth :%v\ndate :%v\n", int(time.Now().Year()), int(time.Now().Month()), int(time.Now().Day()))
	//fmt.Printf("alarmTimer: %v\n", alarmTimer)
	//fmt.Printf("time.Now(): %v\n", time.Now())
	//fmt.Printf("wait: %v\n", wait)
	timer := time.NewTimer(wait)
	<-timer.C
	fmt.Println(alarmName)

	//第一次打印后每间隔24h打印一次
	ticker := time.NewTicker(time.Hour * 24)
	for {
		select {
		case <-ticker.C:
			fmt.Println(alarmName)

		}
	}
}

func main() {
	fmt.Println("Press enter to stop.")
	go alarm("我还能再战4小时！", "2:00")

	go alarm("我要去图书馆开卷！", "6:00")

	go func() {
		var dur time.Duration = 30000000000
		ticker := time.NewTicker(dur)
		for {
			select {
			case <-ticker.C:
				fmt.Println("芜湖！起飞！")
			}

		}
	}()

	//fmt.Println(time.Parse("15:04", "20:48"))
	var a string
	fmt.Scanf("%v", &a)
	//fmt.Println("Finished")
}
