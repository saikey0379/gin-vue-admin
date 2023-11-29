// 自动生成模板DeviceServerDiscovery
package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Server发现 结构体  DeviceServerDiscovery
type DeviceServerDiscovery struct {
      global.GVA_MODEL
      Sn  string `json:"sn" form:"sn" gorm:"column:sn;comment:序列号;size:64;"`  //SN 
      Manufacturer  string `json:"manufacturer" form:"manufacturer" gorm:"column:manufacturer;comment:生产商;size:64;"`  //生产商 
      Model  string `json:"model" form:"model" gorm:"column:model;comment:型号;size:64;"`  //型号 
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:IP地址;size:256;"`  //业务IP 
}


// TableName Server发现 DeviceServerDiscovery自定义表名 device_server_discovery
func (DeviceServerDiscovery) TableName() string {
  return "device_server_discovery"
}

