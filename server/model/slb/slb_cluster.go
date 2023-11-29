// 自动生成模板SlbCluster
package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 集群管理 结构体  SlbCluster
type SlbCluster struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      SshUser  string `json:"sshUser" form:"sshUser" gorm:"column:ssh_user;comment:用户名;size:32;"`  //用户名 
      SshRsa  string `json:"sshRsa" form:"sshRsa" gorm:"column:ssh_rsa;comment:RSA;size:64;"`  //RSA 
      Nodes  string `json:"nodes" form:"nodes" gorm:"column:nodes;comment:节点;size:64;"`  //节点 
      PathConf  string `json:"pathConf" form:"pathConf" gorm:"column:path_conf;comment:配置路径;size:256;"`  //配置路径 
      PathSsl  string `json:"pathSsl" form:"pathSsl" gorm:"column:path_ssl;comment:证书路径;size:256;"`  //证书路径 
      ExecTest  string `json:"execTest" form:"execTest" gorm:"column:exec_test;comment:测试命令;size:256;"`  //测试命令 
      ExecLoad  string `json:"execLoad" form:"execLoad" gorm:"column:exec_load;comment:加载命令;size:256;"`  //加载命令 
}


// TableName 集群管理 SlbCluster自定义表名 slb_cluster
func (SlbCluster) TableName() string {
  return "slb_cluster"
}

