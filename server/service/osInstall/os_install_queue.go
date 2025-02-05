package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
	osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
)

type OsInstallQueueService struct {
}

// CreateOsInstallQueue 创建操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) CreateOsInstallQueue(osInstallQueue *osInstall.OsInstallQueue) (err error) {
	err = global.GVA_DB.Create(osInstallQueue).Error
	return err
}

// DeleteOsInstallQueue 删除操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) DeleteOsInstallQueue(osInstallQueue osInstall.OsInstallQueue) (err error) {
	err = global.GVA_DB.Delete(&osInstallQueue).Error
	return err
}

// DeleteOsInstallQueueByIds 批量删除操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) DeleteOsInstallQueueByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]osInstall.OsInstallQueue{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOsInstallQueue 更新操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) UpdateOsInstallQueue(osInstallQueue osInstall.OsInstallQueue) (err error) {
	err = global.GVA_DB.Save(&osInstallQueue).Error
	return err
}

// GetOsInstallQueue 根据id获取操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) GetOsInstallQueue(id uint) (osInstallQueue osInstall.OsInstallQueue, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&osInstallQueue).Error
	return
}

// GetOsInstallQueueBySn 根据sn获取操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) GetOsInstallQueueBySn(sn string) (osInstallQueue osInstall.OsInstallQueue, err error) {
	err = global.GVA_DB.Where("sn = ?", sn).Where("status = 3").First(&osInstallQueue).Error
	return
}

// GetOsInstallQueueInfoList 分页获取操作系统安装队列记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallQueueService *OsInstallQueueService) GetOsInstallQueueInfoList(info osInstallReq.OsInstallQueueSearch) (list []osInstall.OsInstallQueue, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&osInstall.OsInstallQueue{})
	var osInstallQueues []osInstall.OsInstallQueue
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Sn != "" {
		db = db.Where("sn LIKE ?", "%"+info.Sn+"%")
	}
	if info.Hostname != "" {
		db = db.Where("hostname LIKE ?", "%"+info.Hostname+"%")
	}
	if info.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+info.Ip+"%")
	}
	if info.NetworkId != nil {
		db = db.Where("network_id = ?", info.NetworkId)
	}
	if info.Manager != "" {
		db = db.Where("manager LIKE ?", "%"+info.Manager+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["hostname"] = true
	orderMap["ip"] = true
	orderMap["network_id"] = true
	orderMap["hardware_id"] = true
	orderMap["pxe_id"] = true
	orderMap["kickstart_id"] = true
	orderMap["status"] = true
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

	err = db.Find(&osInstallQueues).Error
	return osInstallQueues, total, err
}
