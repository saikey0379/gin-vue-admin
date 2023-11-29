// 自动生成模板RuleLabel
package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 规则标签 结构体  RuleLabel
type RuleLabel struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
}


// TableName 规则标签 RuleLabel自定义表名 rule_label
func (RuleLabel) TableName() string {
  return "rule_label"
}

