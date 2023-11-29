// 自动生成模板CmdbIpPreempt
package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 预占地址 结构体  CmdbIpPreempt
type CmdbIpPreempt struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //IP地址 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码 
      RegionId  *int `json:"regionId" form:"regionId" gorm:"column:region_id;comment:机房;size:64;"`  //区域 
      SegmentId  *int `json:"segmentId" form:"segmentId" gorm:"column:segment_id;comment:网段ID;size:16;"`  //网段 
      SubnetId  *int `json:"subnetId" form:"subnetId" gorm:"column:subnet_id;comment:VLAN;size:16;"`  //子网 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:网段标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 预占地址 CmdbIpPreempt自定义表名 cmdb_ip_preempt
func (CmdbIpPreempt) TableName() string {
  return "cmdb_ip_preempt"
}

