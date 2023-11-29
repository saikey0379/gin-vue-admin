// 自动生成模板DataCenter
package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心 结构体  DataCenter
type DataCenter struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //Name 
      Describe  string `json:"config" form:"config" gorm:"column:config;comment:PXE;type:text;"`  //描述 
}


// TableName 数据中心 DataCenter自定义表名 data_center
func (DataCenter) TableName() string {
  return "data_center"
}

