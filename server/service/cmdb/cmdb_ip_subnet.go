package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
)

type CmdbIpSubnetService struct {
}

// CreateCmdbIpSubnet 创建子网管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSubnetService *CmdbIpSubnetService) CreateCmdbIpSubnet(cmdbIpSubnet *cmdb.CmdbIpSubnet) (err error) {
	err = global.GVA_DB.Create(cmdbIpSubnet).Error
	return err
}

// DeleteCmdbIpSubnet 删除子网管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSubnetService *CmdbIpSubnetService)DeleteCmdbIpSubnet(cmdbIpSubnet cmdb.CmdbIpSubnet) (err error) {
	err = global.GVA_DB.Delete(&cmdbIpSubnet).Error
	return err
}

// DeleteCmdbIpSubnetByIds 批量删除子网管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSubnetService *CmdbIpSubnetService)DeleteCmdbIpSubnetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cmdb.CmdbIpSubnet{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmdbIpSubnet 更新子网管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSubnetService *CmdbIpSubnetService)UpdateCmdbIpSubnet(cmdbIpSubnet cmdb.CmdbIpSubnet) (err error) {
	err = global.GVA_DB.Save(&cmdbIpSubnet).Error
	return err
}

// GetCmdbIpSubnet 根据id获取子网管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSubnetService *CmdbIpSubnetService)GetCmdbIpSubnet(id uint) (cmdbIpSubnet cmdb.CmdbIpSubnet, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmdbIpSubnet).Error
	return
}

// GetCmdbIpSubnetInfoList 分页获取子网管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmdbIpSubnetService *CmdbIpSubnetService)GetCmdbIpSubnetInfoList(info cmdbReq.CmdbIpSubnetSearch) (list []cmdb.CmdbIpSubnet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cmdb.CmdbIpSubnet{})
    var cmdbIpSubnets []cmdb.CmdbIpSubnet
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
	
	err = db.Find(&cmdbIpSubnets).Error
	return  cmdbIpSubnets, total, err
}
