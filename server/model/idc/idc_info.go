// 自动生成模板IdcInfo
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心 结构体  IdcInfo
type IdcInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //IDC 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:IDC标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 数据中心 IdcInfo自定义表名 idc_info
func (IdcInfo) TableName() string {
  return "idc_info"
}

