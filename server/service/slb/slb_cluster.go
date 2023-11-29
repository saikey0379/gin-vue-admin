package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
)

type SlbClusterService struct {
}

// CreateSlbCluster 创建集群管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbClusterService *SlbClusterService) CreateSlbCluster(slbCluster *slb.SlbCluster) (err error) {
	err = global.GVA_DB.Create(slbCluster).Error
	return err
}

// DeleteSlbCluster 删除集群管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbClusterService *SlbClusterService)DeleteSlbCluster(slbCluster slb.SlbCluster) (err error) {
	err = global.GVA_DB.Delete(&slbCluster).Error
	return err
}

// DeleteSlbClusterByIds 批量删除集群管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbClusterService *SlbClusterService)DeleteSlbClusterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]slb.SlbCluster{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlbCluster 更新集群管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbClusterService *SlbClusterService)UpdateSlbCluster(slbCluster slb.SlbCluster) (err error) {
	err = global.GVA_DB.Save(&slbCluster).Error
	return err
}

// GetSlbCluster 根据id获取集群管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbClusterService *SlbClusterService)GetSlbCluster(id uint) (slbCluster slb.SlbCluster, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slbCluster).Error
	return
}

// GetSlbClusterInfoList 分页获取集群管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbClusterService *SlbClusterService)GetSlbClusterInfoList(info slbReq.SlbClusterSearch) (list []slb.SlbCluster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&slb.SlbCluster{})
    var slbClusters []slb.SlbCluster
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
	
	err = db.Find(&slbClusters).Error
	return  slbClusters, total, err
}
