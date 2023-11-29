// 自动生成模板ServerDiscovery
package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ServerDiscovery 结构体  ServerDiscovery
type ServerDiscovery struct {
	global.GVA_MODEL
	Sn            string `json:"sn" form:"sn" gorm:"unique;column:sn;comment:序列号;size:64;"`                                      //SN
	Hostname      string `json:"hostname" form:"hostname" gorm:"column:hostname;comment:Hostname;size:255;"`                     //Hostname
	Ip            string `json:"ip" form:"ip" gorm:"column:ip;comment:IP;size:64;"`                                              //IP
	VersionOs     string `json:"versionOs" form:"versionOs" gorm:"column:version_os;comment:VersionOs;size:64;"`                 //操作系统版本
	VersionKernel string `json:"versionKernel" form:"versionKernel" gorm:"column:version_kernel;comment:VersionKernel;size:64;"` //内核版本
	Manufacturer  string `json:"manufacturer" form:"manufacturer" gorm:"column:manufacturer;comment:生产商;size:64;"`               //厂商
	ModelName     string `json:"modelName" form:"modelName" gorm:"column:model_name;comment:型号;size:64;"`                        //型号
	CpuSum        *int   `json:"cpuSum" form:"cpuSum" gorm:"column:cpu_sum;comment:CPU核数;size:16;"`                              //CPU
	Cpu           string `json:"cpu" form:"cpu" gorm:"column:cpu;comment:CPU信息;type:longtext;"`                                  //CPU信息
	MemorySum     *int   `json:"memorySum" form:"memorySum" gorm:"column:memory_sum;comment:内存汇总;size:16;"`                      //内存
	Memory        string `json:"memory" form:"memory" gorm:"column:memory;comment:内存信息;type:longtext;"`                          //内存信息
	NicInfo       string `json:"nicInfo" form:"nicInfo" gorm:"column:nic_info;comment:网卡信息;type:longtext;"`                      //网卡
	NicDevice     string `json:"nicDevice" form:"nicDevice" gorm:"column:nic_device;comment:NIC;type:longtext;"`                 //NIC
	Gpu           string `json:"gpu" form:"gpu" gorm:"column:gpu;comment:Gpu信息;type:longtext;"`                                  //Gpu
	Motherboard   string `json:"motherboard" form:"motherboard" gorm:"column:motherboard;comment:主板信息;size:64;"`                 //主板
	Raid          string `json:"raid" form:"raid" gorm:"column:raid;comment:Raid卡信息;size:64;"`                                   //Raid卡
	Disk          string `json:"disk" form:"disk" gorm:"column:disk;comment:硬盘信息;type:longtext;"`                                //硬盘信息
	ServerType    *int   `json:"serverType" form:"serverType" gorm:"column:server_type;comment:Server类型;size:8;"`                //Server类型
	VersionAgent  string `json:"versionAgent" form:"versionAgent" gorm:"column:version_agent;comment:Agent版本;size:64;"`          //Agent版本
}

// TableName ServerDiscovery ServerDiscovery自定义表名 server_discovery
func (ServerDiscovery) TableName() string {
	return "server_discovery"
}
