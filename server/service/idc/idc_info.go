package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcInfoService struct {
}

// CreateIdcInfo 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfoService *IdcInfoService) CreateIdcInfo(idcInfo *idc.IdcInfo) (err error) {
	err = global.GVA_DB.Create(idcInfo).Error
	return err
}

// DeleteIdcInfo 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfoService *IdcInfoService) DeleteIdcInfo(idcInfo idc.IdcInfo) (err error) {
	err = global.GVA_DB.Delete(&idcInfo).Error
	return err
}

// DeleteIdcInfoByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfoService *IdcInfoService) DeleteIdcInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIdcInfo 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfoService *IdcInfoService) UpdateIdcInfo(idcInfo idc.IdcInfo) (err error) {
	err = global.GVA_DB.Save(&idcInfo).Error
	return err
}

// GetIdcInfo 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfoService *IdcInfoService) GetIdcInfo(id uint) (idcInfo idc.IdcInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcInfo).Error
	return
}

// GetIdcInfoInfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfoService *IdcInfoService) GetIdcInfoInfoList(info idcReq.IdcInfoSearch) (list []idc.IdcInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&idc.IdcInfo{})
	var idcInfos []idc.IdcInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["name"] = true
	orderMap["label"] = true
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

	err = db.Find(&idcInfos).Error
	return idcInfos, total, err
}
