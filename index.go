package main

import (
	"fmt"
	"jx3/jx"
	"time"

	"github.com/go-vgo/robotgo"
)

// Point 屏幕坐标点
type Point struct {
	x, y int
}

// Rectangle 截图矩形
type Rectangle struct {
	lt, rb Point
}

func main() {
	fmt.Println(time.Now())
	rect := Rectangle{Point{35, 366}, Point{154, 378}}
	bitmap := robotgo.CaptureScreen(rect.lt.x, rect.lt.y, rect.rb.x, rect.rb.y)
	step := 10 //调整步长
	height := rect.rb.y - rect.lt.y
	width := rect.rb.x - rect.lt.x
	for i := 0; i < width; i++ {
		robotgo.SaveBitmap(robotgo.GetPortion(bitmap, i, 0, step, height), fmt.Sprintf("./jx/target.png"))
		num := jx.Find("./jx/target.png", "./img/num", 5)
		if num != "" {
			fmt.Println(num)
		}

	}
	fmt.Println(time.Now())

}
