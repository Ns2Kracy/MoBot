package util

import (
	"io/ioutil"
	"os"
)

/**
 * 读取文件
 */
func ReadFile(path string) []byte {
	dataByte, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return dataByte
}

/**
 * 判断文件是否存在
 */
func FileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func CheckDir(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			return err
		}
	}
	// check again
	_, err := os.Stat(path)
	return err
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
