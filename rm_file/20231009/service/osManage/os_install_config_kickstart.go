package osManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/osManage/request"
)

type ConfigKickstartService struct {
}

// CreateConfigKickstart 创建Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ksService *ConfigKickstartService) CreateConfigKickstart(ks *osManage.ConfigKickstart) (err error) {
	err = global.GVA_DB.Create(ks).Error
	return err
}

// DeleteConfigKickstart 删除Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ksService *ConfigKickstartService)DeleteConfigKickstart(ks osManage.ConfigKickstart) (err error) {
	err = global.GVA_DB.Delete(&ks).Error
	return err
}

// DeleteConfigKickstartByIds 批量删除Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ksService *ConfigKickstartService)DeleteConfigKickstartByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]osManage.ConfigKickstart{},"id in ?",ids.Ids).Error
	return err
}

// UpdateConfigKickstart 更新Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ksService *ConfigKickstartService)UpdateConfigKickstart(ks osManage.ConfigKickstart) (err error) {
	err = global.GVA_DB.Save(&ks).Error
	return err
}

// GetConfigKickstart 根据id获取Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ksService *ConfigKickstartService)GetConfigKickstart(id uint) (ks osManage.ConfigKickstart, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ks).Error
	return
}

// GetConfigKickstartInfoList 分页获取Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (ksService *ConfigKickstartService)GetConfigKickstartInfoList(info osManageReq.ConfigKickstartSearch) (list []osManage.ConfigKickstart, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&osManage.ConfigKickstart{})
    var kss []osManage.ConfigKickstart
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
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
	
	err = db.Find(&kss).Error
	return  kss, total, err
}
