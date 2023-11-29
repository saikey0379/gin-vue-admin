// 自动生成模板RuleGroup
package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 规则分组 结构体  RuleGroup
type RuleGroup struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      Interval  string `json:"interval" form:"interval" gorm:"column:interval;comment:周期;size:16;"`  //周期 
      Limit  *int `json:"limit" form:"limit" gorm:"column:limit;comment:限制;size:32;"`  //限制 
}


// TableName 规则分组 RuleGroup自定义表名 rule_group
func (RuleGroup) TableName() string {
  return "rule_group"
}

