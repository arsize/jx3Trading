package jx

import (
	"os"
	"path/filepath"
	"strings"
)

// Find 扫描
func Find(target, fp string, level int) string {
	temp := walk(target, fp, level)
	if temp != "" {
		return temp
	}
	return ""

}
func walk(target, fp string, level int) string {
	var linuxPath = toLinux(fp)
	var num string
	filepath.Walk(linuxPath, func(path string, info os.FileInfo, err error) error {
		pth := toLinux(path)
		if IsFile(pth) {
			pth = "./" + pth
			distance := Compare(target, "./"+path)

			if distance < level {
				filenameWithSuffix := filepath.Base(pth)
				fileName := strings.TrimSuffix(filenameWithSuffix, filepath.Ext(filenameWithSuffix))
				num = fileName
				return nil
			}
		}
		return nil
	})
	return num
}

func toLinux(basePath string) string {
	return strings.ReplaceAll(basePath, "\\", "/")
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)

}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
