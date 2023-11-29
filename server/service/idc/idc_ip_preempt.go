package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcIpPreemptService struct {
}

// CreateIdcIpPreempt 创建数据中心地址预占记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpPreemptService *IdcIpPreemptService) CreateIdcIpPreempt(idcIpPreempt *idc.IdcIpPreempt) (err error) {
	err = global.GVA_DB.Create(idcIpPreempt).Error
	return err
}

// DeleteIdcIpPreempt 删除数据中心地址预占记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpPreemptService *IdcIpPreemptService)DeleteIdcIpPreempt(idcIpPreempt idc.IdcIpPreempt) (err error) {
	err = global.GVA_DB.Delete(&idcIpPreempt).Error
	return err
}

// DeleteIdcIpPreemptByIds 批量删除数据中心地址预占记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpPreemptService *IdcIpPreemptService)DeleteIdcIpPreemptByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcIpPreempt{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcIpPreempt 更新数据中心地址预占记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpPreemptService *IdcIpPreemptService)UpdateIdcIpPreempt(idcIpPreempt idc.IdcIpPreempt) (err error) {
	err = global.GVA_DB.Save(&idcIpPreempt).Error
	return err
}

// GetIdcIpPreempt 根据id获取数据中心地址预占记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpPreemptService *IdcIpPreemptService)GetIdcIpPreempt(id uint) (idcIpPreempt idc.IdcIpPreempt, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcIpPreempt).Error
	return
}

// GetIdcIpPreemptInfoList 分页获取数据中心地址预占记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpPreemptService *IdcIpPreemptService)GetIdcIpPreemptInfoList(info idcReq.IdcIpPreemptSearch) (list []idc.IdcIpPreempt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcIpPreempt{})
    var idcIpPreempts []idc.IdcIpPreempt
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
    if info.SegmentId != nil {
        db = db.Where("segment_id = ?",info.SegmentId)
    }
    if info.SubnetId != nil {
        db = db.Where("subnet_id = ?",info.SubnetId)
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
         	orderMap["idc_id"] = true
         	orderMap["segment_id"] = true
         	orderMap["subnet_id"] = true
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
	
	err = db.Find(&idcIpPreempts).Error
	return  idcIpPreempts, total, err
}
