// 自动生成模板OsInstallConfigKickstart
package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Kickstart配置 结构体  OsInstallConfigKickstart
type OsInstallConfigKickstart struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //Name 
      PxeId  *int `json:"pxe_id" form:"pxe_id" gorm:"column:pxe_id;comment:PXE模版;size:64;"`  //关联PXE 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:配置内容;size:4096;"`  //配置内容 
}


// TableName Kickstart配置 OsInstallConfigKickstart自定义表名 os_install_config_kickstart
func (OsInstallConfigKickstart) TableName() string {
  return "os_install_config_kickstart"
}

