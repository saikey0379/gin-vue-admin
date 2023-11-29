// 自动生成模板CmdbRegion
package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 区域信息 结构体  CmdbRegion
type CmdbRegion struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //区域 
      RegionType  *int `json:"regionType" form:"regionType" gorm:"column:region_type;comment:区域类型;"`  //类型 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:IDC标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 区域信息 CmdbRegion自定义表名 cmdb_region
func (CmdbRegion) TableName() string {
  return "cmdb_region"
}

