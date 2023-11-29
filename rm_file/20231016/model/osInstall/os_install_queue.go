// 自动生成模板OsInstallQueue
package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 操作系统安装队列 结构体  OsInstallQueue
type OsInstallQueue struct {
      global.GVA_MODEL
      Sn  string `json:"sn" form:"sn" gorm:"column:sn;comment:序列号;size:64;"`  //序列号 
      Hostname  string `json:"hostname" form:"hostname" gorm:"column:hostname;comment:主机名;size:256;"`  //主机名 
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:IP地址;size:256;"`  //IP地址 
      ManageIp  string `json:"manage_ip" form:"manage_ip" gorm:"column:manage_ip;comment:管理IP;size:256;"`  //管理IP 
      NetworkId  *int `json:"network_id" form:"network_id" gorm:"column:network_id;comment:网段ID;size:64;"`  //网段ID 
      ManageNetworkId  *int `json:"manage_network_id" form:"manage_network_id" gorm:"column:manage_network_id;comment:管理网段ID;size:64;"`  //管理网段ID 
      HardwareId  *int `json:"hardware_id" form:"hardware_id" gorm:"column:hardware_id;comment:硬件配置模版;size:64;"`  //硬件配置 
      PxeId  *int `json:"pxe_id" form:"pxe_id" gorm:"column:pxe_id;comment:PXE模版;size:64;"`  //PXE 
      KickstartId  *int `json:"kickstart_id" form:"kickstart_id" gorm:"column:kickstart_id;comment:Kickstart模版;size:64;"`  //Kickstart 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:设备状态;size:64;"`  //设备状态 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:运维负责人;size:64;"`  //任务负责人 
}


// TableName 操作系统安装队列 OsInstallQueue自定义表名 os_install_queue
func (OsInstallQueue) TableName() string {
  return "os_install_queue"
}

