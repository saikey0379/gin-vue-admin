// 自动生成模板IdcIpSegment
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心网段 结构体  IdcIpSegment
type IdcIpSegment struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //网络号 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码位 
      IdcId  *int `json:"idcId" form:"idcId" gorm:"column:idc_id;comment:机房;size:64;"`  //机房 
      IpSegmentType  *int `json:"ipSegmentType" form:"ipSegmentType" gorm:"column:ip_segment_type;comment:网段类型;"`  //类型 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:网段标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 数据中心网段 IdcIpSegment自定义表名 idc_ip_segment
func (IdcIpSegment) TableName() string {
  return "idc_ip_segment"
}

