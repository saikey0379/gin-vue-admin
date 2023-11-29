package utils

import (
	"io/ioutil"
	"os"
	"strconv"
)

func StringToFileMod(fileMod string) (mod os.FileMode, err error) {
	fm, err := strconv.ParseUint(fileMod, 8, 32)
	if err != nil {
		return mod, err
	}
	return os.FileMode(fm), err
}

func CreateFile(dir string, file string, content string) error {
	if !FileExist(dir) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}

	//文件已存在,先删除
	if FileExist(dir + "/" + file) {
		err := os.Remove(dir + "/" + file)
		if err != nil {
			return err
		}
	}

	bytes := []byte(content)
	err := ioutil.WriteFile(dir+"/"+file, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(dir string, file string) error {
	//文件已存在,先删除
	if FileExist(dir + "/" + file) {
		err := os.Remove(dir + "/" + file)
		if err != nil {
			return err
		}
	}

	return nil
}
