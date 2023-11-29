// 自动生成模板Pxe
package osManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// PXE配置 结构体  Pxe
type Pxe struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //PXE 
      Config  string `json:"config" form:"config" gorm:"column:config;comment:PXE;type:text;"`  //PXE配置 
}


// TableName PXE配置 Pxe自定义表名 os_install_config_pxe
func (Pxe) TableName() string {
  return "os_install_config_pxe"
}

