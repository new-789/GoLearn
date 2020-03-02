package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now() // 获取当前系统时间
	fmt.Printf("当前时间为:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	fmt.Printf("时间戳:%d\n", timestamp)
}

// 时间戳转换为 time 时间
func testTimetamp(timestamp int64) {
	// 时间戳转换为 time 方法，返回一个时间对象
	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("转换后结果：%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func processtask() {
	fmt.Printf("定时任务自动执行该函数\n")
}

// 定时器的简单使用
func testTicker() {
	// 创建一个定时器，参数为指定的时间，返回一个(类似数组)的管道对象
	// 我们可以对这个管道进行遍历,在遍历中我们可以执行我们需要执行的任务(任务需提前编写到函数或其它方法中)
	tickerObj := time.Tick(time.Second * 2)
	for i := range tickerObj { // 遍历管道
		fmt.Printf("%v\n", i)
		processtask()
	}
}

func testTimeConst() {
	fmt.Printf("NanoSecond:%d\n", time.Nanosecond)
	fmt.Printf("MicroSecond:%d\n", time.Microsecond)
	fmt.Printf("MilliSecond:%d\n", time.Millisecond)
	fmt.Printf("Second:%d\n", time.Second)
	fmt.Printf("Minute:%d\n", time.Minute)
	fmt.Printf("Hour:%d\n", time.Hour)
}

func testFormatTime() {
	now := time.Now()
	timeStr1 := now.Format("2006/1/02 15:04:05")
	timeStr2 := now.Format("2006/1/02 15:04")
	timeStr3 := now.Format("02/1/2006 15:04")
	timeStr4 := now.Format("2006-1-02 15:04")
	timeStr5 := now.Format("2006年1月02日 15:04")
	fmt.Printf("time1:%s\n", timeStr1)
	fmt.Printf("time2:%s\n", timeStr2)
	fmt.Printf("time3:%s\n", timeStr3)
	fmt.Printf("time4:%s\n", timeStr4)
	fmt.Printf("time5:%s\n", timeStr5)
}

func timeFormat(NowTime time.Time) (formatTime string) {
	// 函数接收时间类型为 time.Time,时间戳则为 int64，格式化后的时间则为 string 类型
	formatTime = NowTime.Format("2006/01/02 15:04:05")
	return formatTime
}

func testFormat() {
	now := time.Now()
	timeStr := fmt.Sprintf("方式二：%02d/%02d/%02d %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("方式二：%s", timeStr)
}

func costTime() {
	start := time.Now().UnixNano() // 程序执行前获取当前纳秒时间戳
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()     // 程序结束后获取当前纳秒时间戳
	costtime := (end - start) / 1000 // 除以一千将纳秒转成微秒
	fmt.Printf("该程序的执行时间为: %d 微妙\n", costtime)
}

func main() {
	// testTime()
	// testTimetamp(time.Now().Unix())
	// testTicker()
	// testTimeConst()
	// testFormatTime()
	timeformat := timeFormat(time.Now())
	fmt.Printf("方式一：format time is: %s\n", timeformat)
	testFormat()
	costTime()
}
