package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dataCenterReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter/request"
)

type DataCenterService struct {
}

// CreateDataCenter 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterService) CreateDataCenter(dc *dataCenter.DataCenter) (err error) {
	err = global.GVA_DB.Create(dc).Error
	return err
}

// DeleteDataCenter 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterService)DeleteDataCenter(dc dataCenter.DataCenter) (err error) {
	err = global.GVA_DB.Delete(&dc).Error
	return err
}

// DeleteDataCenterByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterService)DeleteDataCenterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dataCenter.DataCenter{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDataCenter 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterService)UpdateDataCenter(dc dataCenter.DataCenter) (err error) {
	err = global.GVA_DB.Save(&dc).Error
	return err
}

// GetDataCenter 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterService)GetDataCenter(id uint) (dc dataCenter.DataCenter, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dc).Error
	return
}

// GetDataCenterInfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterService)GetDataCenterInfoList(info dataCenterReq.DataCenterSearch) (list []dataCenter.DataCenter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dataCenter.DataCenter{})
    var dcs []dataCenter.DataCenter
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
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
	
	err = db.Find(&dcs).Error
	return  dcs, total, err
}
