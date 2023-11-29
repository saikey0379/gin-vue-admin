// 自动生成模板CmdbIpSegment
package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 网段管理 结构体  CmdbIpSegment
type CmdbIpSegment struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //网络号 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码位 
      RegionId  *int `json:"regionId" form:"regionId" gorm:"column:region_id;comment:机房;size:64;"`  //区域 
      IpSegmentType  *int `json:"ipSegmentType" form:"ipSegmentType" gorm:"column:ip_segment_type;comment:网段类型;"`  //类型 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:网段标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 网段管理 CmdbIpSegment自定义表名 cmdb_ip_segment
func (CmdbIpSegment) TableName() string {
  return "cmdb_ip_segment"
}

