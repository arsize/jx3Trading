package utils

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Point 为一个像素坐标点
type Point struct{ x, y int }

var breakFlag = false
var tp, lp, rp Point

// GetMatrix 鼠标左键获取矩阵四个坐标
func GetMatrix(chan1 chan interface{}) {
START:
	fmt.Println("开始采集：")
	a := robotgo.AddEvent("a")
	ok := robotgo.AddEvent("mleft") //robot库会触发两次，是个BUG，用x限定
	if ok && a {
		tp.x, tp.y = robotgo.GetMousePos()
		if lp.x == 0 {
			lp = tp
			fmt.Println("采集到左上角坐标：", lp)
			goto START
		} else if rp.x == 0 {
			rp = tp
			fmt.Println("采集到右下角坐标：", rp)
			temp := [2]Point{lp, rp}
			chan1 <- temp
		}
	}
}
