package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcIpServiceService struct {
}

// CreateIdcIpService 创建数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpServiceService *IdcIpServiceService) CreateIdcIpService(idcIpService *idc.IdcIpService) (err error) {
	err = global.GVA_DB.Create(idcIpService).Error
	return err
}

// DeleteIdcIpService 删除数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpServiceService *IdcIpServiceService) DeleteIdcIpService(idcIpService idc.IdcIpService) (err error) {
	err = global.GVA_DB.Delete(&idcIpService).Error
	return err
}

// DeleteIdcIpServiceByIds 批量删除数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpServiceService *IdcIpServiceService) DeleteIdcIpServiceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcIpService{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIdcIpService 更新数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpServiceService *IdcIpServiceService) UpdateIdcIpService(idcIpService idc.IdcIpService) (err error) {
	err = global.GVA_DB.Save(&idcIpService).Error
	return err
}

// GetIdcIpService 根据id获取数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpServiceService *IdcIpServiceService) GetIdcIpService(id uint) (idcIpService idc.IdcIpService, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcIpService).Error
	return
}

// GetIdcIpServiceInfoList 分页获取数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpServiceService *IdcIpServiceService) GetIdcIpServiceInfoList(info idcReq.IdcIpServiceSearch) (list []idc.IdcIpService, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&idc.IdcIpService{})
	var idcIpServices []idc.IdcIpService
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Network != "" {
		db = db.Where("network LIKE ?", "%"+info.Network+"%")
	}
	if info.Prefix != nil {
		db = db.Where("netmask = ?", info.Prefix)
	}
	if info.Gateway != "" {
		db = db.Where("gateway LIKE ?", "%"+info.Gateway+"%")
	}
	if info.IdcId != nil {
		db = db.Where("idc_id = ?", info.IdcId)
	}
	if info.VlanId != nil {
		db = db.Where("vlan_id = ?", info.VlanId)
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
	orderMap["network"] = true
	orderMap["netmask"] = true
	orderMap["gateway"] = true
	orderMap["idc_id"] = true
	orderMap["vlan_id"] = true
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

	err = db.Find(&idcIpServices).Error
	return idcIpServices, total, err
}
