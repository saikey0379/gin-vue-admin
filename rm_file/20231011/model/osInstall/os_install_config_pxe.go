// 自动生成模板OsInstallConfigPxe
package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// PXE配置 结构体  OsInstallConfigPxe
type OsInstallConfigPxe struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //Name 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:PXE;size:1;"`  //Content 
}


// TableName PXE配置 OsInstallConfigPxe自定义表名 os_install_config_pxe
func (OsInstallConfigPxe) TableName() string {
  return "os_install_config_pxe"
}

