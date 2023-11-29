// 自动生成模板SlbDomain
package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 域名管理 结构体  SlbDomain
type SlbDomain struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      ClusterIds  string `json:"clusterIds" form:"clusterIds" gorm:"column:cluster_ids;comment:集群;size:256;"`  //集群 
      PortHttp  *int `json:"portHttp" form:"portHttp" gorm:"column:port_http;comment:HTTP;size:16;"`  //HTTP 
      PortHttps  *int `json:"portHttps" form:"portHttps" gorm:"column:port_https;comment:HTTPS;size:16;"`  //HTTPS 
      Http2  *bool `json:"http2" form:"http2" gorm:"column:http2;comment:HTTP2;"`  //HTTP2 
      CertId  *int `json:"certId" form:"certId" gorm:"column:cert_id;comment:证书;size:64;"`  //证书 
      Customize  string `json:"customize" form:"customize" gorm:"column:customize;comment:自定义配置;size:256;"`  //自定义配置 
      Locations  string `json:"locations" form:"locations" gorm:"column:locations;comment:Path;size:256;"`  //Path 
}


// TableName 域名管理 SlbDomain自定义表名 slb_domain
func (SlbDomain) TableName() string {
  return "slb_domain"
}

