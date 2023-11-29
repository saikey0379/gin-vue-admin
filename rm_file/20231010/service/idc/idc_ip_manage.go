package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcIpManageService struct {
}

// CreateIdcIpManage 创建数据中心管理网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpManageService *IdcIpManageService) CreateIdcIpManage(idcIpManage *idc.IdcIpManage) (err error) {
	err = global.GVA_DB.Create(idcIpManage).Error
	return err
}

// DeleteIdcIpManage 删除数据中心管理网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpManageService *IdcIpManageService)DeleteIdcIpManage(idcIpManage idc.IdcIpManage) (err error) {
	err = global.GVA_DB.Delete(&idcIpManage).Error
	return err
}

// DeleteIdcIpManageByIds 批量删除数据中心管理网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpManageService *IdcIpManageService)DeleteIdcIpManageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcIpManage{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcIpManage 更新数据中心管理网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpManageService *IdcIpManageService)UpdateIdcIpManage(idcIpManage idc.IdcIpManage) (err error) {
	err = global.GVA_DB.Save(&idcIpManage).Error
	return err
}

// GetIdcIpManage 根据id获取数据中心管理网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpManageService *IdcIpManageService)GetIdcIpManage(id uint) (idcIpManage idc.IdcIpManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcIpManage).Error
	return
}

// GetIdcIpManageInfoList 分页获取数据中心管理网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpManageService *IdcIpManageService)GetIdcIpManageInfoList(info idcReq.IdcIpManageSearch) (list []idc.IdcIpManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcIpManage{})
    var idcIpManages []idc.IdcIpManage
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Network != "" {
        db = db.Where("network LIKE ?","%"+ info.Network+"%")
    }
    if info.Prefix != nil {
        db = db.Where("netmask = ?",info.Prefix)
    }
    if info.Gateway != "" {
        db = db.Where("gateway LIKE ?","%"+ info.Gateway+"%")
    }
    if info.IdcId != nil {
        db = db.Where("idc_id = ?",info.IdcId)
    }
    if info.VlanId != nil {
        db = db.Where("vlan_id = ?",info.VlanId)
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
         	orderMap["network"] = true
         	orderMap["netmask"] = true
         	orderMap["gateway"] = true
         	orderMap["idc_id"] = true
         	orderMap["vlan_id"] = true
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
	
	err = db.Find(&idcIpManages).Error
	return  idcIpManages, total, err
}
