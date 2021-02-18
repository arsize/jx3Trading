package main

import (
	"fmt"

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

var pos Point

func main() {
	location()
	// pre()

}

// 预处理：像素切片
func pre() {
	fmt.Println("正在预处理，请稍后...")
	rect := Rectangle{Point{25, 576}, Point{113, 588}}
	bitmap := robotgo.CaptureScreen(rect.lt.x, rect.lt.y, rect.rb.x, rect.rb.y)
	step := 10 //调整步长
	width := rect.rb.x - rect.lt.x
	height := rect.rb.y - rect.lt.y
	for i := 0; i < width; i++ {
		robotgo.SaveBitmap(robotgo.GetPortion(bitmap, i, 0, step, height), fmt.Sprintf("../img/%d.png", i))
	}
	fmt.Println("预处理结束！")
	defer robotgo.FreeBitmap(bitmap)
}

// 坐标定位点
func location() {
	for {
		pos.x, pos.y = robotgo.GetMousePos()
		fmt.Println(pos.x, pos.y)
	}
}
