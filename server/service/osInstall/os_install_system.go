package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OsInstallSystemService struct {
	Content string
	Files   []string
	DirPxe  string
}

func (osInstallSystemService *OsInstallSystemService) CreateFilePxe() (err error) {
	for _, i := range osInstallSystemService.Files {
		err = utils.CreateFile(osInstallSystemService.DirPxe, i, osInstallSystemService.Content)
		if err != nil {
			return err
		}
	}

	return err
}

func (osInstallSystemService *OsInstallSystemService) RemoveFilePxe() (err error) {
	for _, i := range osInstallSystemService.Files {
		err = utils.RemoveFile(osInstallSystemService.DirPxe, i)
		if err != nil {
			return err
		}
	}

	return err
}
