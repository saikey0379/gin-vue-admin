// 自动生成模板RuleAnnotation
package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 注解标签 结构体  RuleAnnotation
type RuleAnnotation struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 注解标签 RuleAnnotation自定义表名 rule_annotation
func (RuleAnnotation) TableName() string {
  return "rule_annotation"
}

