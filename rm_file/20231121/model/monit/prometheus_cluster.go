// 自动生成模板PrometheusCluster
package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 集群管理 结构体  PrometheusCluster
type PrometheusCluster struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      SshUser  string `json:"sshUser" form:"sshUser" gorm:"column:ssh_user;comment:用户名;size:32;"`  //用户名 
      SshRsa  string `json:"sshRsa" form:"sshRsa" gorm:"column:ssh_rsa;comment:RSA;size:64;"`  //RSA 
      Nodes  string `json:"nodes" form:"nodes" gorm:"column:nodes;comment:节点;size:64;"`  //节点 
      PathConf  string `json:"pathConf" form:"pathConf" gorm:"column:path_conf;comment:配置路径;size:256;"`  //配置路径 
}


// TableName 集群管理 PrometheusCluster自定义表名 prometheus_cluster
func (PrometheusCluster) TableName() string {
  return "prometheus_cluster"
}

