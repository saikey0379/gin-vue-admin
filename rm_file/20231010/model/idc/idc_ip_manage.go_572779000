// 自动生成模板IdcIpManage
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心管理网段 结构体  IdcIpManage
type IdcIpManage struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //IP地址 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码 
      Gateway  string `json:"gateway" form:"gateway" gorm:"column:gateway;comment:网关;size:64;"`  //网关 
      IdcId  *int `json:"idc_id" form:"idc_id" gorm:"column:idc_id;comment:机房;size:64;"`  //机房 
      VlanId  *int `json:"vlan_id" form:"vlan_id" gorm:"column:vlan_id;comment:VLAN;size:64;"`  //VLAN 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Label  string `json:"label" form:"label" gorm:"column:label;comment:网段标签;size:256;"`  //标签 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 数据中心管理网段 IdcIpManage自定义表名 idc_ip_manage
func (IdcIpManage) TableName() string {
  return "idc_ip_manage"
}

