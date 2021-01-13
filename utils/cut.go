package utils

import (
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
)

// CutImg 根据矩形坐标切图
func CutImg(pointArr [2]Point) {
	point1 := pointArr[0]
	point2 := pointArr[1]
	dir, _ := os.Getwd()
	state := FileExist(dir + "/screen.png")
	if state {
		fmt.Println("img has found")
	} else {
		fmt.Println("create img bitmap")
		bitmap := robotgo.CaptureScreen(point1.x, point1.y, (point2.x - point1.x), (point2.y - point1.y))
		robotgo.SaveBitmap(bitmap, `screen.png`)
	}

}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
