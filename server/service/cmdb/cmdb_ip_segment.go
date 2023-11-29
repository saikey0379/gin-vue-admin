package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
)

type CmdbIpSegmentService struct {
}

// CreateCmdbIpSegment 创建网段管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSegmentService *CmdbIpSegmentService) CreateCmdbIpSegment(cmdbIpSegment *cmdb.CmdbIpSegment) (err error) {
	err = global.GVA_DB.Create(cmdbIpSegment).Error
	return err
}

// DeleteCmdbIpSegment 删除网段管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSegmentService *CmdbIpSegmentService)DeleteCmdbIpSegment(cmdbIpSegment cmdb.CmdbIpSegment) (err error) {
	err = global.GVA_DB.Delete(&cmdbIpSegment).Error
	return err
}

// DeleteCmdbIpSegmentByIds 批量删除网段管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSegmentService *CmdbIpSegmentService)DeleteCmdbIpSegmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cmdb.CmdbIpSegment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmdbIpSegment 更新网段管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSegmentService *CmdbIpSegmentService)UpdateCmdbIpSegment(cmdbIpSegment cmdb.CmdbIpSegment) (err error) {
	err = global.GVA_DB.Save(&cmdbIpSegment).Error
	return err
}

// GetCmdbIpSegment 根据id获取网段管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSegmentService *CmdbIpSegmentService)GetCmdbIpSegment(id uint) (cmdbIpSegment cmdb.CmdbIpSegment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmdbIpSegment).Error
	return
}

// GetCmdbIpSegmentInfoList 分页获取网段管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSegmentService *CmdbIpSegmentService)GetCmdbIpSegmentInfoList(info cmdbReq.CmdbIpSegmentSearch) (list []cmdb.CmdbIpSegment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cmdb.CmdbIpSegment{})
    var cmdbIpSegments []cmdb.CmdbIpSegment
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
         	orderMap["region_id"] = true
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
	
	err = db.Find(&cmdbIpSegments).Error
	return  cmdbIpSegments, total, err
}
