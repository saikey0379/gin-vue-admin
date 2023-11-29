package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcRoomService struct {
}

// CreateIdcRoom 创建数据中心房间记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *IdcRoomService) CreateIdcRoom(room *idc.IdcRoom) (err error) {
	err = global.GVA_DB.Create(room).Error
	return err
}

// DeleteIdcRoom 删除数据中心房间记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *IdcRoomService) DeleteIdcRoom(room idc.IdcRoom) (err error) {
	err = global.GVA_DB.Delete(&room).Error
	return err
}

// DeleteIdcRoomByIds 批量删除数据中心房间记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *IdcRoomService) DeleteIdcRoomByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcRoom{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIdcRoom 更新数据中心房间记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *IdcRoomService) UpdateIdcRoom(room idc.IdcRoom) (err error) {
	err = global.GVA_DB.Save(&room).Error
	return err
}

// GetIdcRoom 根据id获取数据中心房间记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *IdcRoomService) GetIdcRoom(id uint) (room idc.IdcRoom, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&room).Error
	return
}

// GetIdcRoomInfoList 分页获取数据中心房间记录
// Author [piexlmax](https://github.com/piexlmax)
func (roomService *IdcRoomService) GetIdcRoomInfoList(info idcReq.IdcRoomSearch) (list []idc.IdcRoom, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&idc.IdcRoom{})
	var rooms []idc.IdcRoom
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
	if info.IdcId != nil {
		db = db.Where("idc_id = ?", info.IdcId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["name"] = true
	orderMap["label"] = true
	orderMap["idc_id"] = true
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

	err = db.Find(&rooms).Error
	return rooms, total, err
}
