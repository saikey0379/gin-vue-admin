package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
)

type OsInstallLogService struct {
}

// CreateOsInstallLog 创建操作系统安装日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallLogService *OsInstallLogService) CreateOsInstallLog(osInstallLog *osInstall.OsInstallLog) (err error) {
	err = global.GVA_DB.Create(osInstallLog).Error
	return err
}

// DeleteOsInstallLog 删除操作系统安装日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallLogService *OsInstallLogService)DeleteOsInstallLog(osInstallLog osInstall.OsInstallLog) (err error) {
	err = global.GVA_DB.Delete(&osInstallLog).Error
	return err
}

// DeleteOsInstallLogByIds 批量删除操作系统安装日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallLogService *OsInstallLogService)DeleteOsInstallLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]osInstall.OsInstallLog{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOsInstallLog 更新操作系统安装日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallLogService *OsInstallLogService)UpdateOsInstallLog(osInstallLog osInstall.OsInstallLog) (err error) {
	err = global.GVA_DB.Save(&osInstallLog).Error
	return err
}

// GetOsInstallLog 根据id获取操作系统安装日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallLogService *OsInstallLogService)GetOsInstallLog(id uint) (osInstallLog osInstall.OsInstallLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&osInstallLog).Error
	return
}

// GetOsInstallLogInfoList 分页获取操作系统安装日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallLogService *OsInstallLogService)GetOsInstallLogInfoList(info osInstallReq.OsInstallLogSearch) (list []osInstall.OsInstallLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&osInstall.OsInstallLog{})
    var osInstallLogs []osInstall.OsInstallLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Sn != "" {
        db = db.Where("sn = ?",info.Sn)
    }
    if info.QueueId != nil {
        db = db.Where("queue_id = ?",info.QueueId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&osInstallLogs).Error
	return  osInstallLogs, total, err
}
