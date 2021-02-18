package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	hstr := ""
	fmt.Println(time.Now())

	bitmap := robotgo.CaptureScreen(0, 0, 1000, 1000)
	defer robotgo.FreeBitmap(bitmap)
	for i := 0; i < 999; i++ {
		for j := 0; j < 999; j++ {
			hstr = robotgo.GetColors(bitmap, i, j)
		}
	}

	fmt.Println(hstr)
	fmt.Println(time.Now())

}
