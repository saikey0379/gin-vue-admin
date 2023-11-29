// 自动生成模板SlbUpstream
package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 目标节点 结构体  SlbUpstream
type SlbUpstream struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      ClusterIds  string `json:"clusterIds" form:"clusterIds" gorm:"column:cluster_ids;comment:集群;size:256;"`  //集群 
      Customize  string `json:"customize" form:"customize" gorm:"column:customize;comment:自定义配置;size:64;"`  //自定义配置 
      Nodes  string `json:"nodes" form:"nodes" gorm:"column:nodes;comment:节点;size:64;"`  //节点 
}


// TableName 目标节点 SlbUpstream自定义表名 slb_upstream
func (SlbUpstream) TableName() string {
  return "slb_upstream"
}

