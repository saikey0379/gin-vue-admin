package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
)

type RuleLabelService struct {
}

// CreateRuleLabel 创建规则标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleLabelService *RuleLabelService) CreateRuleLabel(ruleLabel *monit.RuleLabel) (err error) {
	err = global.GVA_DB.Create(ruleLabel).Error
	return err
}

// DeleteRuleLabel 删除规则标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleLabelService *RuleLabelService)DeleteRuleLabel(ruleLabel monit.RuleLabel) (err error) {
	err = global.GVA_DB.Delete(&ruleLabel).Error
	return err
}

// DeleteRuleLabelByIds 批量删除规则标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleLabelService *RuleLabelService)DeleteRuleLabelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]monit.RuleLabel{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRuleLabel 更新规则标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleLabelService *RuleLabelService)UpdateRuleLabel(ruleLabel monit.RuleLabel) (err error) {
	err = global.GVA_DB.Save(&ruleLabel).Error
	return err
}

// GetRuleLabel 根据id获取规则标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleLabelService *RuleLabelService)GetRuleLabel(id uint) (ruleLabel monit.RuleLabel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ruleLabel).Error
	return
}

// GetRuleLabelInfoList 分页获取规则标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleLabelService *RuleLabelService)GetRuleLabelInfoList(info monitReq.RuleLabelSearch) (list []monit.RuleLabel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&monit.RuleLabel{})
    var ruleLabels []monit.RuleLabel
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
	
	err = db.Find(&ruleLabels).Error
	return  ruleLabels, total, err
}
