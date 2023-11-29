// 自动生成模板ServerDiscovery
package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Server发现 结构体  ServerDiscovery
type ServerDiscovery struct {
      global.GVA_MODEL
      Sn  string `json:"sn" form:"sn" gorm:"column:sn;comment:序列号;size:64;"`  //SN 
      Hostname  string `json:"hostname" form:"hostname" gorm:"column:hostname;comment:Hostname;size:255;"`  //Hostname 
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:IP;size:64;"`  //IP 
      Manufacturer  string `json:"manufacturer" form:"manufacturer" gorm:"column:manufacturer;comment:生产商;size:256;"`  //厂商 
      Model  string `json:"model" form:"model" gorm:"column:model;comment:型号;size:256;"`  //型号 
      CpuSum  *int `json:"cpuSum" form:"cpuSum" gorm:"column:cpu_sum;comment:CPU核数;size:16;"`  //CPU 
      Cpu  string `json:"cpu" form:"cpu" gorm:"column:cpu;comment:CPU信息;size:4096;"`  //CPU信息 
      MemorySum  *int `json:"memorySum" form:"memorySum" gorm:"column:memory_sum;comment:内存汇总;size:16;"`  //内存 
      Memory  string `json:"memory" form:"memory" gorm:"column:memory;comment:内存信息;size:4096;"`  //内存信息 
      Nic  string `json:"nic" form:"nic" gorm:"column:nic;comment:网卡信息;size:4096;"`  //网卡 
      Gpu  string `json:"gpu" form:"gpu" gorm:"column:gpu;comment:Gpu信息;size:4096;"`  //Gpu 
      Motherboard  string `json:"motherboard" form:"motherboard" gorm:"column:motherboard;comment:主板信息;size:4096;"`  //主板 
      Raid  string `json:"raid" form:"raid" gorm:"column:raid;comment:Raid卡信息;size:4096;"`  //Raid卡 
      Disk  string `json:"disk" form:"disk" gorm:"column:status;comment:硬盘信息;size:4096;"`  //硬盘信息 
      AgentVersion  string `json:"agentVersion" form:"agentVersion" gorm:"column:agent_version;comment:Agent版本;size:64;"`  //Agent版本 
}


// TableName Server发现 ServerDiscovery自定义表名 server_discovery
func (ServerDiscovery) TableName() string {
  return "server_discovery"
}

