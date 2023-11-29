// 自动生成模板IdcIpService
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心业务网段 结构体  IdcIpService
type IdcIpService struct {
      global.GVA_MODEL
      Network  string `json:"network" form:"network" gorm:"column:network;comment:网络号;size:64;"`  //IP地址 
      Prefix  *int `json:"netmask" form:"netmask" gorm:"column:netmask;comment:网段前缀;size:16;"`  //掩码 
      Gateway  string `json:"gateway" form:"gateway" gorm:"column:gateway;comment:网关;size:64;"`  //网关 
      DataCenterId  *int `json:"data_denter_id" form:"data_denter_id" gorm:"column:data_center_id;comment:机房;size:64;"`  //机房 
      Vlan  *int `json:"vlan" form:"vlan" gorm:"column:vlan;comment:VLAN;size:64;"`  //VLAN 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;type:text;"`  //描述 
}


// TableName 数据中心业务网段 IdcIpService自定义表名 idc_ip_service
func (IdcIpService) TableName() string {
  return "idc_ip_service"
}

