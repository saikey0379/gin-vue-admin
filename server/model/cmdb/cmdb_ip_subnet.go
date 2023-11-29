// 自动生成模板CmdbIpSubnet
package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 子网管理 结构体  CmdbIpSubnet
type CmdbIpSubnet struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //IP地址 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码 
      Gateway  string `json:"gateway" form:"gateway" gorm:"column:gateway;comment:网关;size:64;"`  //网关 
      RegionId  *int `json:"regionId" form:"regionId" gorm:"column:region_id;comment:区域;size:64;"`  //区域 
      SegmentId  *int `json:"segmentId" form:"segmentId" gorm:"column:segment_id;comment:网段ID;size:16;"`  //网段ID 
      VlanId  *int `json:"vlanId" form:"vlanId" gorm:"column:vlan_id;comment:VLAN;size:64;"`  //VLAN 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:网段标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 子网管理 CmdbIpSubnet自定义表名 cmdb_ip_subnet
func (CmdbIpSubnet) TableName() string {
  return "cmdb_ip_subnet"
}

