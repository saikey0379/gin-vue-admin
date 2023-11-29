package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
	osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
)

type OsInstallConfigPxeService struct {
}

// CreateOsInstallConfigPxe 创建PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigPxeService *OsInstallConfigPxeService) CreateOsInstallConfigPxe(osInstallConfigPxe *osInstall.OsInstallConfigPxe) (err error) {
	err = global.GVA_DB.Create(osInstallConfigPxe).Error
	return err
}

// DeleteOsInstallConfigPxe 删除PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigPxeService *OsInstallConfigPxeService) DeleteOsInstallConfigPxe(osInstallConfigPxe osInstall.OsInstallConfigPxe) (err error) {
	err = global.GVA_DB.Delete(&osInstallConfigPxe).Error
	return err
}

// DeleteOsInstallConfigPxeByIds 批量删除PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigPxeService *OsInstallConfigPxeService) DeleteOsInstallConfigPxeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]osInstall.OsInstallConfigPxe{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOsInstallConfigPxe 更新PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigPxeService *OsInstallConfigPxeService) UpdateOsInstallConfigPxe(osInstallConfigPxe osInstall.OsInstallConfigPxe) (err error) {
	err = global.GVA_DB.Save(&osInstallConfigPxe).Error
	return err
}

// GetOsInstallConfigPxe 根据id获取PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigPxeService *OsInstallConfigPxeService) GetOsInstallConfigPxe(id uint) (osInstallConfigPxe osInstall.OsInstallConfigPxe, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&osInstallConfigPxe).Error
	return
}

// GetOsInstallConfigPxeInfoList 分页获取PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (osInstallConfigPxeService *OsInstallConfigPxeService) GetOsInstallConfigPxeInfoList(info osInstallReq.OsInstallConfigPxeSearch) (list []osInstall.OsInstallConfigPxe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&osInstall.OsInstallConfigPxe{})
	var osInstallConfigPxes []osInstall.OsInstallConfigPxe
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
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

	err = db.Find(&osInstallConfigPxes).Error
	return osInstallConfigPxes, total, err
}
