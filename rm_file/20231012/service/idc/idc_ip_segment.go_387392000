package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcIpSegmentService struct {
}

// CreateIdcIpSegment 创建数据中心IP网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSegmentService *IdcIpSegmentService) CreateIdcIpSegment(idcIpSegment *idc.IdcIpSegment) (err error) {
	err = global.GVA_DB.Create(idcIpSegment).Error
	return err
}

// DeleteIdcIpSegment 删除数据中心IP网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSegmentService *IdcIpSegmentService)DeleteIdcIpSegment(idcIpSegment idc.IdcIpSegment) (err error) {
	err = global.GVA_DB.Delete(&idcIpSegment).Error
	return err
}

// DeleteIdcIpSegmentByIds 批量删除数据中心IP网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSegmentService *IdcIpSegmentService)DeleteIdcIpSegmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcIpSegment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcIpSegment 更新数据中心IP网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSegmentService *IdcIpSegmentService)UpdateIdcIpSegment(idcIpSegment idc.IdcIpSegment) (err error) {
	err = global.GVA_DB.Save(&idcIpSegment).Error
	return err
}

// GetIdcIpSegment 根据id获取数据中心IP网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSegmentService *IdcIpSegmentService)GetIdcIpSegment(id uint) (idcIpSegment idc.IdcIpSegment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcIpSegment).Error
	return
}

// GetIdcIpSegmentInfoList 分页获取数据中心IP网段记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSegmentService *IdcIpSegmentService)GetIdcIpSegmentInfoList(info idcReq.IdcIpSegmentSearch) (list []idc.IdcIpSegment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcIpSegment{})
    var idcIpSegments []idc.IdcIpSegment
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Network != "" {
        db = db.Where("network LIKE ?","%"+ info.Network+"%")
    }
    if info.IdcId != nil {
        db = db.Where("idc_id = ?",info.IdcId)
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
         	orderMap["idc_id"] = true
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
	
	err = db.Find(&idcIpSegments).Error
	return  idcIpSegments, total, err
}
