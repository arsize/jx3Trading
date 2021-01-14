package main

import (
	"utils"
)

func main() {
	chan1 := make(chan [2]utils.Point)
	defer close(chan1)
	// 获取矩阵坐标
	go utils.GetMatrix(chan1)
	// 截图扫描区域
	utils.CutImg(<-chan1)
	// 图片二值化处理
	utils.OrcHandle()

}
