package fileHandle

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	FileResumeTransApi
}

var (
	fileResumeTransService = service.ServiceGroupApp.FileHandleServiceGroup.FileResumeTransService
)
