package utils

import (
	"fmt"
	"os"

	"github.com/otiai10/gosseract"
)

// OrcHandle 对图片进行orc识别
func OrcHandle() {
	client := gosseract.NewClient()
	defer client.Close()
	dir, _ := os.Getwd()
	client.SetImage(dir + "/screen.png")
	text, _ := client.Text()
	fmt.Println(text)
}
