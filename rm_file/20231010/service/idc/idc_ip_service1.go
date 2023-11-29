package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcIpService1Service struct {
}

// CreateIdcIpService1 创建数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpService1Service *IdcIpService1Service) CreateIdcIpService1(idcIpService1 *idc.IdcIpService1) (err error) {
	err = global.GVA_DB.Create(idcIpService1).Error
	return err
}

// DeleteIdcIpService1 删除数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpService1Service *IdcIpService1Service)DeleteIdcIpService1(idcIpService1 idc.IdcIpService1) (err error) {
	err = global.GVA_DB.Delete(&idcIpService1).Error
	return err
}

// DeleteIdcIpService1ByIds 批量删除数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpService1Service *IdcIpService1Service)DeleteIdcIpService1ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcIpService1{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcIpService1 更新数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpService1Service *IdcIpService1Service)UpdateIdcIpService1(idcIpService1 idc.IdcIpService1) (err error) {
	err = global.GVA_DB.Save(&idcIpService1).Error
	return err
}

// GetIdcIpService1 根据id获取数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpService1Service *IdcIpService1Service)GetIdcIpService1(id uint) (idcIpService1 idc.IdcIpService1, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcIpService1).Error
	return
}

// GetIdcIpService1InfoList 分页获取数据中心业务网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpService1Service *IdcIpService1Service)GetIdcIpService1InfoList(info idcReq.IdcIpService1Search) (list []idc.IdcIpService1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcIpService1{})
    var idcIpService1s []idc.IdcIpService1
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
    if info.Vlan != nil {
        db = db.Where("vlan = ?",info.Vlan)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Label != "" {
        db = db.Where("label LIKE ?","%"+ info.Label+"%")
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
         	orderMap["vlan"] = true
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
	
	err = db.Find(&idcIpService1s).Error
	return  idcIpService1s, total, err
}
