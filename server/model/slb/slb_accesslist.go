// 自动生成模板SlbAccesslist
package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 访问控制 结构体  SlbAccesslist
type SlbAccesslist struct {
      global.GVA_MODEL
      ClusterId  *int `json:"clusterId" form:"clusterId" gorm:"column:cluster_id;comment:集群ID;size:64;"`  //集群ID 
      DomainId  string `json:"domainId" form:"domainId" gorm:"column:domain_id;comment:负责人;size:64;"`  //域名 
      RouteId  string `json:"routeId" form:"routeId" gorm:"column:route_id;comment:Path;size:64;"`  //Path 
      AccessType  *int `json:"accessType" form:"accessType" gorm:"column:access_type;comment:类型;size:2;"`  //类型 
      Nodes  string `json:"nodes" form:"nodes" gorm:"column:nodes;comment:节点;size:64;"`  //节点 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 访问控制 SlbAccesslist自定义表名 slb_accesslist
func (SlbAccesslist) TableName() string {
  return "slb_accesslist"
}

