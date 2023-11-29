package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
)

type RuleGroupService struct {
}

// CreateRuleGroup 创建规则分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleGroupService *RuleGroupService) CreateRuleGroup(ruleGroup *monit.RuleGroup) (err error) {
	err = global.GVA_DB.Create(ruleGroup).Error
	return err
}

// DeleteRuleGroup 删除规则分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleGroupService *RuleGroupService)DeleteRuleGroup(ruleGroup monit.RuleGroup) (err error) {
	err = global.GVA_DB.Delete(&ruleGroup).Error
	return err
}

// DeleteRuleGroupByIds 批量删除规则分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleGroupService *RuleGroupService)DeleteRuleGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]monit.RuleGroup{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRuleGroup 更新规则分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleGroupService *RuleGroupService)UpdateRuleGroup(ruleGroup monit.RuleGroup) (err error) {
	err = global.GVA_DB.Save(&ruleGroup).Error
	return err
}

// GetRuleGroup 根据id获取规则分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleGroupService *RuleGroupService)GetRuleGroup(id uint) (ruleGroup monit.RuleGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ruleGroup).Error
	return
}

// GetRuleGroupInfoList 分页获取规则分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleGroupService *RuleGroupService)GetRuleGroupInfoList(info monitReq.RuleGroupSearch) (list []monit.RuleGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&monit.RuleGroup{})
    var ruleGroups []monit.RuleGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Manager != "" {
        db = db.Where("manager = ?",info.Manager)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
         	orderMap["manager"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&ruleGroups).Error
	return  ruleGroups, total, err
}
