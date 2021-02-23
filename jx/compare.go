package jx

import (
	"image/png"
	"os"

	"github.com/corona10/goimagehash"
)

// Compare 图片相似度比较2
func Compare(path1, path2 string) int {
	file1, _ := os.Open(path1)
	file2, _ := os.Open(path2)
	defer file1.Close()
	defer file2.Close()
	img1, _ := png.Decode(file1)
	img2, _ := png.Decode(file2)
	hash1, _ := goimagehash.DifferenceHash(img1)
	hash2, _ := goimagehash.DifferenceHash(img2)
	distance, _ := hash1.Distance(hash2)
	return distance

}
