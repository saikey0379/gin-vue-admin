package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
)

type OsInstallConfigKickstartService struct {
}

// CreateOsInstallConfigKickstart 创建Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigKickstartService *OsInstallConfigKickstartService) CreateOsInstallConfigKickstart(osInstallConfigKickstart *osInstall.OsInstallConfigKickstart) (err error) {
	err = global.GVA_DB.Create(osInstallConfigKickstart).Error
	return err
}

// DeleteOsInstallConfigKickstart 删除Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigKickstartService *OsInstallConfigKickstartService)DeleteOsInstallConfigKickstart(osInstallConfigKickstart osInstall.OsInstallConfigKickstart) (err error) {
	err = global.GVA_DB.Delete(&osInstallConfigKickstart).Error
	return err
}

// DeleteOsInstallConfigKickstartByIds 批量删除Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigKickstartService *OsInstallConfigKickstartService)DeleteOsInstallConfigKickstartByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]osInstall.OsInstallConfigKickstart{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOsInstallConfigKickstart 更新Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigKickstartService *OsInstallConfigKickstartService)UpdateOsInstallConfigKickstart(osInstallConfigKickstart osInstall.OsInstallConfigKickstart) (err error) {
	err = global.GVA_DB.Save(&osInstallConfigKickstart).Error
	return err
}

// GetOsInstallConfigKickstart 根据id获取Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigKickstartService *OsInstallConfigKickstartService)GetOsInstallConfigKickstart(id uint) (osInstallConfigKickstart osInstall.OsInstallConfigKickstart, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&osInstallConfigKickstart).Error
	return
}

// GetOsInstallConfigKickstartInfoList 分页获取Kickstart配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigKickstartService *OsInstallConfigKickstartService)GetOsInstallConfigKickstartInfoList(info osInstallReq.OsInstallConfigKickstartSearch) (list []osInstall.OsInstallConfigKickstart, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&osInstall.OsInstallConfigKickstart{})
    var osInstallConfigKickstarts []osInstall.OsInstallConfigKickstart
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
	
	err = db.Find(&osInstallConfigKickstarts).Error
	return  osInstallConfigKickstarts, total, err
}
