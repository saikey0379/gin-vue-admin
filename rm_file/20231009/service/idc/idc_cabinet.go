package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcCabinetService struct {
}

// CreateIdcCabinet 创建机柜记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcCabinetService *IdcCabinetService) CreateIdcCabinet(idcCabinet *idc.IdcCabinet) (err error) {
	err = global.GVA_DB.Create(idcCabinet).Error
	return err
}

// DeleteIdcCabinet 删除机柜记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcCabinetService *IdcCabinetService)DeleteIdcCabinet(idcCabinet idc.IdcCabinet) (err error) {
	err = global.GVA_DB.Delete(&idcCabinet).Error
	return err
}

// DeleteIdcCabinetByIds 批量删除机柜记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcCabinetService *IdcCabinetService)DeleteIdcCabinetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcCabinet{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcCabinet 更新机柜记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcCabinetService *IdcCabinetService)UpdateIdcCabinet(idcCabinet idc.IdcCabinet) (err error) {
	err = global.GVA_DB.Save(&idcCabinet).Error
	return err
}

// GetIdcCabinet 根据id获取机柜记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcCabinetService *IdcCabinetService)GetIdcCabinet(id uint) (idcCabinet idc.IdcCabinet, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcCabinet).Error
	return
}

// GetIdcCabinetInfoList 分页获取机柜记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcCabinetService *IdcCabinetService)GetIdcCabinetInfoList(info idcReq.IdcCabinetSearch) (list []idc.IdcCabinet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcCabinet{})
    var idcCabinets []idc.IdcCabinet
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.RoomId != nil {
        db = db.Where("room_id = ?",info.RoomId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
         	orderMap["room_id"] = true
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
	
	err = db.Find(&idcCabinets).Error
	return  idcCabinets, total, err
}
