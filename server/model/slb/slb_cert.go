// 自动生成模板SlbCert
package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 证书管理 结构体  SlbCert
type SlbCert struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      ClusterIds  string `json:"clusterIds" form:"clusterIds" gorm:"column:cluster_ids;comment:相关集群;size:256;"`  //相关集群 
      CertFile  string `json:"certFile" form:"certFile" gorm:"column:cert_file;comment:证书名称;size:64;"`  //证书名称 
      CertContent  string `json:"certContent" form:"certContent" gorm:"column:cert_content;comment:证书内容;size:256;"`  //证书内容 
      KeyFile  string `json:"keyFile" form:"keyFile" gorm:"column:key_file;comment:秘钥名称;size:256;"`  //秘钥名称 
      KeyContent  string `json:"keyContent" form:"keyContent" gorm:"column:key_content;comment:秘钥内容;size:256;"`  //秘钥内容 
}


// TableName 证书管理 SlbCert自定义表名 slb_cert
func (SlbCert) TableName() string {
  return "slb_cert"
}

