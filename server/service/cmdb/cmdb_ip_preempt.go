package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
)

type CmdbIpPreemptService struct {
}

// CreateCmdbIpPreempt 创建预占地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpPreemptService *CmdbIpPreemptService) CreateCmdbIpPreempt(cmdbIpPreempt *cmdb.CmdbIpPreempt) (err error) {
	err = global.GVA_DB.Create(cmdbIpPreempt).Error
	return err
}

// DeleteCmdbIpPreempt 删除预占地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpPreemptService *CmdbIpPreemptService)DeleteCmdbIpPreempt(cmdbIpPreempt cmdb.CmdbIpPreempt) (err error) {
	err = global.GVA_DB.Delete(&cmdbIpPreempt).Error
	return err
}

// DeleteCmdbIpPreemptByIds 批量删除预占地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpPreemptService *CmdbIpPreemptService)DeleteCmdbIpPreemptByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cmdb.CmdbIpPreempt{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmdbIpPreempt 更新预占地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpPreemptService *CmdbIpPreemptService)UpdateCmdbIpPreempt(cmdbIpPreempt cmdb.CmdbIpPreempt) (err error) {
	err = global.GVA_DB.Save(&cmdbIpPreempt).Error
	return err
}

// GetCmdbIpPreempt 根据id获取预占地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpPreemptService *CmdbIpPreemptService)GetCmdbIpPreempt(id uint) (cmdbIpPreempt cmdb.CmdbIpPreempt, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmdbIpPreempt).Error
	return
}

// GetCmdbIpPreemptInfoList 分页获取预占地址记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpPreemptService *CmdbIpPreemptService)GetCmdbIpPreemptInfoList(info cmdbReq.CmdbIpPreemptSearch) (list []cmdb.CmdbIpPreempt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cmdb.CmdbIpPreempt{})
    var cmdbIpPreempts []cmdb.CmdbIpPreempt
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Network != "" {
        db = db.Where("network LIKE ?","%"+ info.Network+"%")
    }
    if info.RegionId != nil {
        db = db.Where("region_id = ?",info.RegionId)
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
         	orderMap["region_id"] = true
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
	
	err = db.Find(&cmdbIpPreempts).Error
	return  cmdbIpPreempts, total, err
}
