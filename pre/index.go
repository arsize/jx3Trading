package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

// LeftTopPos 解析json
type LeftTopPos struct {
	X, Y int
}

// Config 配置文件
type Config struct {
	LeftTopPos   LeftTopPos
	WindowWidth  int
	WindowHeight int
	ScanWidth    int
}

var config Config = Config{}

func init() {
	fileData, err := ioutil.ReadFile("../config.json")
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}
	config = Config{}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}
}

func main() {
	// location()
	pre()

}

// 预处理：像素切片
func pre() {
	fmt.Println("正在预处理，请稍后...")
	leftTop := Point{config.LeftTopPos.X, config.LeftTopPos.Y}
	rightBottom := Point{config.LeftTopPos.X + config.ScanWidth, config.LeftTopPos.Y + config.WindowHeight}
	rect := Rectangle{leftTop, rightBottom}
	bitmap := robotgo.CaptureScreen(rect.lt.x, rect.lt.y, rect.rb.x, rect.rb.y)
	for i := 0; i < config.ScanWidth; i++ {
		robotgo.SaveBitmap(robotgo.GetPortion(bitmap, i, 0, config.WindowWidth, config.WindowHeight), fmt.Sprintf("../img/%d.png", i))
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
