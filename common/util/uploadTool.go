// 图片上传工具类

package util

import (
	"fmt"
	"os"
)

// CreateDir 创建目录
func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		fmt.Println("err", err)
		return err
	}
	return nil
}

// IsExist 判断是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
