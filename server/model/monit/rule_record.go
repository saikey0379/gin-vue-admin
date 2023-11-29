// 自动生成模板RuleRecord
package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 规则配置 结构体  RuleRecord
type RuleRecord struct {
	global.GVA_MODEL
	Name          string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`                              //Name
	RuleGroupId   *int   `json:"ruleGroupId" form:"ruleGroupId" gorm:"column:rule_group_id;comment:分组ID;size:32;"`       //分组ID
	Manager       string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`                      //负责人
	Describe      string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`                 //描述
	Clusters      string `json:"clusters" form:"clusters" gorm:"column:clusters;comment:配置集群;size:256;"`                 //集群
	Record        string `json:"record" form:"record" gorm:"column:record;comment:记录;size:256;"`                         //记录
	Alert         string `json:"alert" form:"alert" gorm:"column:alert;comment:警报;size:64;"`                             //警报
	Expr          string `json:"expr" form:"expr" gorm:"column:expr;comment:表达式;type:longtext;"`                         //表达式
	For           string `json:"for" form:"for" gorm:"column:for;comment:持续时间;size:32;"`                                 //持续时间
	KeepFiringFor string `json:"keepFiringFor" form:"keepFiringFor" gorm:"column:keep_firing_for;comment:保持时间;size:32;"` //保持时间
	Labels        string `json:"labels" form:"labels" gorm:"column:labels;comment:标签;size:256;"`                         //标签
	Annotations   string `json:"annotations" form:"annotations" gorm:"column:annotations;comment:注解;size:256;"`          //注解
}

// TableName 规则配置 RuleRecord自定义表名 rule_record
func (RuleRecord) TableName() string {
	return "rule_record"
}
