// 自动生成模板IdcIpPreempt
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心地址预占 结构体  IdcIpPreempt
type IdcIpPreempt struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //IP地址 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码 
      IdcId  *int `json:"idcId" form:"idcId" gorm:"column:idc_id;comment:机房;size:64;"`  //机房 
      SegmentId  *int `json:"segmentId" form:"segmentId" gorm:"column:segment_id;comment:网段ID;size:16;"`  //网段 
      SubnetId  *int `json:"subnetId" form:"subnetId" gorm:"column:subnet_id;comment:VLAN;size:16;"`  //子网 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:网段标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 数据中心地址预占 IdcIpPreempt自定义表名 idc_ip_preempt
func (IdcIpPreempt) TableName() string {
  return "idc_ip_preempt"
}

