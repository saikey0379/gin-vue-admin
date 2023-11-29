package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
)

type CmdbRegionService struct {
}

// CreateCmdbRegion 创建区域信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbRegionService *CmdbRegionService) CreateCmdbRegion(cmdbRegion *cmdb.CmdbRegion) (err error) {
	err = global.GVA_DB.Create(cmdbRegion).Error
	return err
}

// DeleteCmdbRegion 删除区域信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbRegionService *CmdbRegionService)DeleteCmdbRegion(cmdbRegion cmdb.CmdbRegion) (err error) {
	err = global.GVA_DB.Delete(&cmdbRegion).Error
	return err
}

// DeleteCmdbRegionByIds 批量删除区域信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbRegionService *CmdbRegionService)DeleteCmdbRegionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cmdb.CmdbRegion{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmdbRegion 更新区域信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbRegionService *CmdbRegionService)UpdateCmdbRegion(cmdbRegion cmdb.CmdbRegion) (err error) {
	err = global.GVA_DB.Save(&cmdbRegion).Error
	return err
}

// GetCmdbRegion 根据id获取区域信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbRegionService *CmdbRegionService)GetCmdbRegion(id uint) (cmdbRegion cmdb.CmdbRegion, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmdbRegion).Error
	return
}

// GetCmdbRegionInfoList 分页获取区域信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbRegionService *CmdbRegionService)GetCmdbRegionInfoList(info cmdbReq.CmdbRegionSearch) (list []cmdb.CmdbRegion, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cmdb.CmdbRegion{})
    var cmdbRegions []cmdb.CmdbRegion
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.RegionType != nil {
        db = db.Where("region_type = ?",info.RegionType)
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
         	orderMap["region_type"] = true
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
	
	err = db.Find(&cmdbRegions).Error
	return  cmdbRegions, total, err
}
