// 自动生成模板BareMetal
package bareMetal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 裸金属设备 结构体  BareMetal
type BareMetal struct {
      global.GVA_MODEL
      Sn  string `json:"sn" form:"sn" gorm:"column:sn;comment:序列号;size:64;"`  //SN 
      AssetId  string `json:"asset_id" form:"asset_id" gorm:"column:asset_id;comment:资产编号;size:64;"`  //资产编号 
      Hostname  string `json:"hostname" form:"hostname" gorm:"column:hostname;comment:主机名;size:256;"`  //实例名 
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:IP地址;size:256;"`  //业务IP 
      NetworkId  *int `json:"network_id" form:"network_id" gorm:"column:network_id;comment:网段ID;size:64;"`  //业务网络 
      ManageIp  string `json:"manage_ip" form:"manage_ip" gorm:"column:manage_ip;comment:管理IP;size:256;"`  //管理IP 
      ManageNetworkId  *int `json:"manage_network_id" form:"manage_network_id" gorm:"column:manage_network_id;comment:管理网段ID;size:64;"`  //管理网络 
      HardwareId  *int `json:"hardware_id" form:"hardware_id" gorm:"column:hardware_id;comment:硬件配置模版;size:64;"`  //硬件配置 
      KickstartId  *int `json:"kickstart_id" form:"kickstart_id" gorm:"column:kickstart_id;comment:Kickstart模版;size:64;"`  //系统 
      CabinetId  *int `json:"cabinet_id" form:"cabinet_id" gorm:"column:cabinet_id;comment:机柜ID;size:64;"`  //机柜 
      CabinetU  string `json:"cabinet_u" form:"cabinet_u" gorm:"column:cabinet_u;comment:机柜U位;size:64;"`  //U位 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:设备状态;size:64;"`  //状态 
      ManagerDev  string `json:"manager_dev" form:"manager_dev" gorm:"column:manager_dev;comment:开发负责人;size:64;"`  //开发 
      ManagerOps  string `json:"manager_ops" form:"manager_ops" gorm:"column:manager_ops;comment:运维负责人;size:64;"`  //运维 
      LabelId  *int `json:"label_id" form:"label_id" gorm:"column:label_id;comment:设备标签ID;size:64;"`  //标签 
      Describe  string `json:"device_describe" form:"device_describe" gorm:"column:device_describe;comment:设备描述;size:4096;"`  //描述 
}


// TableName 裸金属设备 BareMetal自定义表名 device_bare_metal
func (BareMetal) TableName() string {
  return "device_bare_metal"
}

