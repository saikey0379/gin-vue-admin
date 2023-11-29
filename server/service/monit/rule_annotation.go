package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
)

type RuleAnnotationService struct {
}

// CreateRuleAnnotation 创建注解标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleAnnotationService *RuleAnnotationService) CreateRuleAnnotation(ruleAnnotation *monit.RuleAnnotation) (err error) {
	err = global.GVA_DB.Create(ruleAnnotation).Error
	return err
}

// DeleteRuleAnnotation 删除注解标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleAnnotationService *RuleAnnotationService)DeleteRuleAnnotation(ruleAnnotation monit.RuleAnnotation) (err error) {
	err = global.GVA_DB.Delete(&ruleAnnotation).Error
	return err
}

// DeleteRuleAnnotationByIds 批量删除注解标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleAnnotationService *RuleAnnotationService)DeleteRuleAnnotationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]monit.RuleAnnotation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRuleAnnotation 更新注解标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleAnnotationService *RuleAnnotationService)UpdateRuleAnnotation(ruleAnnotation monit.RuleAnnotation) (err error) {
	err = global.GVA_DB.Save(&ruleAnnotation).Error
	return err
}

// GetRuleAnnotation 根据id获取注解标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleAnnotationService *RuleAnnotationService)GetRuleAnnotation(id uint) (ruleAnnotation monit.RuleAnnotation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ruleAnnotation).Error
	return
}

// GetRuleAnnotationInfoList 分页获取注解标签记录
// Author [piexlmax](https://github.com/piexlmax)
func (ruleAnnotationService *RuleAnnotationService)GetRuleAnnotationInfoList(info monitReq.RuleAnnotationSearch) (list []monit.RuleAnnotation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&monit.RuleAnnotation{})
    var ruleAnnotations []monit.RuleAnnotation
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
	
	err = db.Find(&ruleAnnotations).Error
	return  ruleAnnotations, total, err
}
