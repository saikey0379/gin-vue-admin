// 自动生成模板OsInstallLog
package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 操作系统安装日志 结构体  OsInstallLog
type OsInstallLog struct {
      global.GVA_MODEL
      Sn  string `json:"sn" form:"sn" gorm:"column:sn;comment:SN;size:64;"`  //设备序列号 
      QueueId  *int `json:"queueId" form:"queueId" gorm:"column:queue_id;comment:装机队列ID;size:64;"`  //队列ID 
      Log  string `json:"log" form:"log" gorm:"column:log;comment:日志;size:4096;"`  //日志 
}


// TableName 操作系统安装日志 OsInstallLog自定义表名 os_install_log
func (OsInstallLog) TableName() string {
  return "os_install_log"
}

