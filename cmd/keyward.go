package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

var (

	// KeySleep set the key default millisecond sleep time
	KeySleep = 5
)

func main() {
	fmt.Println("Press CTRL-C to stop the program.")
	for {
		robotgo.TypeStr("Hello, world!")                  // 模拟键盘输入
		robotgo.MilliSleep(500)                           // 暂停 500 毫秒
		robotgo.KeyTap("enter")                           // 模拟按下回车键
		time.Sleep(time.Duration(KeySleep) * time.Second) // 暂停 1 秒
	}
}
