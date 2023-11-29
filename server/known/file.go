package known

import "fmt"

const (
	TmpBreakPointDir  = "/tmp/breakPoint/"
	TmpResumeTransDir = "/tmp/resumeTrans/"
)

func ReturnFileDir(fileType string, modelType string) (fileDir string) {
	switch fileType {
	case "config":
		switch modelType {
		case "prometheus-rules", "slb":
			return RootDir + fmt.Sprintf("/config/%s/", modelType)
		}
	case "rsa":
		return RootDir + "/.ssh/"
	case "static":
		return RootDir + "/static/"
	}
	return fileDir
}

func ReturnFileBakDir(fileType string, modelType string) (fileDir string) {
	switch fileType {
	case "config":
		switch modelType {
		case "prometheus-rules", "slb":
			return RootDir + fmt.Sprintf("/config/%s/bak", modelType)
		}
	}
	return fileDir
}
