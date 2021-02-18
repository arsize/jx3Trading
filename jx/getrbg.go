package jx

import "github.com/go-vgo/robotgo"

// GetRgb 获取屏幕像素点rgb色值
func GetRgb() [3]uint32 {
	var r, g, b uint32
	bitmap := robotgo.CaptureScreen(0, 0, 1000, 1000)
	defer robotgo.FreeBitmap(bitmap)
	pic := robotgo.ToImage(bitmap)
	for i := 0; i < 999; i++ {
		for j := 0; j < 999; j++ {
			r, g, b, _ = pic.At(i, j).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
		}
	}
	c := [3]uint32{r, g, b}
	return c

}
