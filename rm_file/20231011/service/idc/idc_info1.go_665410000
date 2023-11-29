package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcInfo1Service struct {
}

// CreateIdcInfo1 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfo1Service *IdcInfo1Service) CreateIdcInfo1(idcInfo1 *idc.IdcInfo1) (err error) {
	err = global.GVA_DB.Create(idcInfo1).Error
	return err
}

// DeleteIdcInfo1 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfo1Service *IdcInfo1Service)DeleteIdcInfo1(idcInfo1 idc.IdcInfo1) (err error) {
	err = global.GVA_DB.Delete(&idcInfo1).Error
	return err
}

// DeleteIdcInfo1ByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfo1Service *IdcInfo1Service)DeleteIdcInfo1ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcInfo1{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcInfo1 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfo1Service *IdcInfo1Service)UpdateIdcInfo1(idcInfo1 idc.IdcInfo1) (err error) {
	err = global.GVA_DB.Save(&idcInfo1).Error
	return err
}

// GetIdcInfo1 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfo1Service *IdcInfo1Service)GetIdcInfo1(id uint) (idcInfo1 idc.IdcInfo1, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcInfo1).Error
	return
}

// GetIdcInfo1InfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfo1Service *IdcInfo1Service)GetIdcInfo1InfoList(info idcReq.IdcInfo1Search) (list []idc.IdcInfo1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcInfo1{})
    var idcInfo1s []idc.IdcInfo1
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
	
	err = db.Find(&idcInfo1s).Error
	return  idcInfo1s, total, err
}
