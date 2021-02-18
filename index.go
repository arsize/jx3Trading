package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Point 屏幕坐标点
type Point struct {
	x, y int
}

// Rectangle 截图矩形
type Rectangle struct {
	lt, rb Point
}

// Config 配置文件
type Config struct {
	WindowWidth  int
	WindowHeight int
}

var config Config = Config{}

func init() {
	fileData, err := ioutil.ReadFile("./config.json")
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
	fmt.Println(config.WindowHeight)

	// fmt.Println(time.Now())
	// rect := Rectangle{Point{35, 366}, Point{154, 378}}
	// bitmap := robotgo.CaptureScreen(rect.lt.x, rect.lt.y, rect.rb.x, rect.rb.y)
	// step := 10 //调整步长
	// height := rect.rb.y - rect.lt.y
	// width := rect.rb.x - rect.lt.x
	// for i := 0; i < width; i++ {
	// 	robotgo.SaveBitmap(robotgo.GetPortion(bitmap, i, 0, step, height), fmt.Sprintf("./jx/target.png"))
	// 	num := jx.Find("./jx/target.png", "./img/num", 5)
	// 	if num != "" {
	// 		fmt.Println(num)
	// 	}

	// }
	// fmt.Println(time.Now())

}
