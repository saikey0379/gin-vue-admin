package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
)

type SlbAccesslistService struct {
}

// CreateSlbAccesslist 创建访问控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbAccesslistService *SlbAccesslistService) CreateSlbAccesslist(slbAccesslist *slb.SlbAccesslist) (err error) {
	err = global.GVA_DB.Create(slbAccesslist).Error
	return err
}

// DeleteSlbAccesslist 删除访问控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbAccesslistService *SlbAccesslistService)DeleteSlbAccesslist(slbAccesslist slb.SlbAccesslist) (err error) {
	err = global.GVA_DB.Delete(&slbAccesslist).Error
	return err
}

// DeleteSlbAccesslistByIds 批量删除访问控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbAccesslistService *SlbAccesslistService)DeleteSlbAccesslistByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]slb.SlbAccesslist{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlbAccesslist 更新访问控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbAccesslistService *SlbAccesslistService)UpdateSlbAccesslist(slbAccesslist slb.SlbAccesslist) (err error) {
	err = global.GVA_DB.Save(&slbAccesslist).Error
	return err
}

// GetSlbAccesslist 根据id获取访问控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbAccesslistService *SlbAccesslistService)GetSlbAccesslist(id uint) (slbAccesslist slb.SlbAccesslist, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slbAccesslist).Error
	return
}

// GetSlbAccesslistInfoList 分页获取访问控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbAccesslistService *SlbAccesslistService)GetSlbAccesslistInfoList(info slbReq.SlbAccesslistSearch) (list []slb.SlbAccesslist, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&slb.SlbAccesslist{})
    var slbAccesslists []slb.SlbAccesslist
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ClusterId != nil {
        db = db.Where("cluster_id = ?",info.ClusterId)
    }
    if info.DomainId != "" {
        db = db.Where("domain_id = ?",info.DomainId)
    }
    if info.RouteId != "" {
        db = db.Where("route_id = ?",info.RouteId)
    }
    if info.Nodes != "" {
        db = db.Where("nodes LIKE ?","%"+ info.Nodes+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&slbAccesslists).Error
	return  slbAccesslists, total, err
}
