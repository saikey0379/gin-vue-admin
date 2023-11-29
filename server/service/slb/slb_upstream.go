package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
)

type SlbUpstreamService struct {
}

// CreateSlbUpstream 创建目标节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbUpstreamService *SlbUpstreamService) CreateSlbUpstream(slbUpstream *slb.SlbUpstream) (err error) {
	err = global.GVA_DB.Create(slbUpstream).Error
	return err
}

// DeleteSlbUpstream 删除目标节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbUpstreamService *SlbUpstreamService)DeleteSlbUpstream(slbUpstream slb.SlbUpstream) (err error) {
	err = global.GVA_DB.Delete(&slbUpstream).Error
	return err
}

// DeleteSlbUpstreamByIds 批量删除目标节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbUpstreamService *SlbUpstreamService)DeleteSlbUpstreamByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]slb.SlbUpstream{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlbUpstream 更新目标节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbUpstreamService *SlbUpstreamService)UpdateSlbUpstream(slbUpstream slb.SlbUpstream) (err error) {
	err = global.GVA_DB.Save(&slbUpstream).Error
	return err
}

// GetSlbUpstream 根据id获取目标节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbUpstreamService *SlbUpstreamService)GetSlbUpstream(id uint) (slbUpstream slb.SlbUpstream, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slbUpstream).Error
	return
}

// GetSlbUpstreamInfoList 分页获取目标节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbUpstreamService *SlbUpstreamService)GetSlbUpstreamInfoList(info slbReq.SlbUpstreamSearch) (list []slb.SlbUpstream, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&slb.SlbUpstream{})
    var slbUpstreams []slb.SlbUpstream
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Manager != "" {
        db = db.Where("manager = ?",info.Manager)
    }
    if info.ClusterIds != "" {
        db = db.Where("cluster_ids LIKE ?","%"+ info.ClusterIds+"%")
    }
    if info.Nodes != "" {
        db = db.Where("nodes LIKE ?","%"+ info.Nodes+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
         	orderMap["manager"] = true
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
	
	err = db.Find(&slbUpstreams).Error
	return  slbUpstreams, total, err
}
