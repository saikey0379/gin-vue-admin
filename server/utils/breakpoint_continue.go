package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/known"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: BreakPointContinue
//@description: 断点续传
//@param: content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string
//@return: error, string

func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := known.TmpBreakPointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathC, err := makeFileContent(content, fileName, path, contentNumber)
	return pathC, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CheckMd5
//@description: 检查Md5
//@param: content []byte, chunkMd5 string
//@return: CanUpload bool

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // 可以继续上传
	} else {
		return false // 切片不完整，废弃
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: makeFileContent
//@description: 创建切片内容
//@param: content []byte, fileName string, FileDir string, contentNumber int
//@return: string, error

func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	if strings.Index(fileName, "..") > -1 || strings.Index(FileDir, "..") > -1 {
		return "", errors.New("文件名或路径不合法")
	}
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: makeFileContent
//@description: 创建切片文件
//@param: fileName string, FileMd5 string
//@return: error, string

func MakeFile(fileDir string, fileName string, fileMd5 string, fileMod os.FileMode) (string, error) {
	rd, err := os.ReadDir(known.TmpBreakPointDir + fileMd5)
	if err != nil {
		return fileDir + fileName, err
	}

	_ = os.MkdirAll(fileDir, os.ModePerm)
	fd, err := os.OpenFile(fileDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, fileMod)
	if err != nil {
		return fileDir + fileName, err
	}
	defer fd.Close()
	for k := range rd {
		content, _ := os.ReadFile(known.TmpBreakPointDir + fileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(fileDir + fileName)
			return fileDir + fileName, err
		}
	}
	return fileDir + fileName, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RemoveChunk
//@description: 移除切片
//@param: FileMd5 string
//@return: error

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(known.TmpBreakPointDir + FileMd5)
	return err
}
