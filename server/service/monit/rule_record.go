package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
)

type RuleRecordService struct {
}

// CreateRuleRecord 创建规则配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleRecordService *RuleRecordService) CreateRuleRecord(ruleRecord *monit.RuleRecord) (err error) {
	err = global.GVA_DB.Create(ruleRecord).Error
	return err
}

// DeleteRuleRecord 删除规则配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleRecordService *RuleRecordService) DeleteRuleRecord(ruleRecord monit.RuleRecord) (err error) {
	err = global.GVA_DB.Delete(&ruleRecord).Error
	return err
}

// DeleteRuleRecordByIds 批量删除规则配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleRecordService *RuleRecordService) DeleteRuleRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]monit.RuleRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRuleRecord 更新规则配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleRecordService *RuleRecordService) UpdateRuleRecord(ruleRecord monit.RuleRecord) (err error) {
	err = global.GVA_DB.Save(&ruleRecord).Debug().Error
	return err
}

// GetRuleRecord 根据id获取规则配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleRecordService *RuleRecordService) GetRuleRecord(id uint) (ruleRecord monit.RuleRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ruleRecord).Debug().Error
	return
}

// GetRuleRecordInfoList 分页获取规则配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleRecordService *RuleRecordService) GetRuleRecordInfoList(info monitReq.RuleRecordSearch) (list []monit.RuleRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&monit.RuleRecord{})
	var ruleRecords []monit.RuleRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RuleGroupId != nil {
		db = db.Where("rule_group_id = ?", info.RuleGroupId)
	}
	if info.Clusters != "" {
		db = db.Where("clusters LIKE ?", "%"+info.Clusters+"%")
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Manager != "" {
		db = db.Where("manager = ?", info.Manager)
	}
	if info.Record != "" {
		db = db.Where("record LIKE ?", "%"+info.Record+"%")
	}
	if info.Alert != "" {
		db = db.Where("alert LIKE ?", "%"+info.Alert+"%")
	}
	if info.Expr != "" {
		db = db.Where("expr LIKE ?", "%"+info.Expr+"%")
	}
	if info.Labels != "" {
		db = db.Where("labels LIKE ?", "%"+info.Labels+"%")
	}
	if info.Annotations != "" {
		db = db.Where("annotations LIKE ?", "%"+info.Annotations+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["name"] = true
	orderMap["manager"] = true
	orderMap["record"] = true
	orderMap["alert"] = true
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

	err = db.Find(&ruleRecords).Error
	return ruleRecords, total, err
}
