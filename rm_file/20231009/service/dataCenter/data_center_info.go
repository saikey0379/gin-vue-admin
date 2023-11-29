package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dataCenterReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter/request"
)

type DataCenterInfoService struct {
}

// CreateDataCenterInfo 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterInfoService) CreateDataCenterInfo(dc *dataCenter.DataCenterInfo) (err error) {
	err = global.GVA_DB.Create(dc).Error
	return err
}

// DeleteDataCenterInfo 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterInfoService)DeleteDataCenterInfo(dc dataCenter.DataCenterInfo) (err error) {
	err = global.GVA_DB.Delete(&dc).Error
	return err
}

// DeleteDataCenterInfoByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterInfoService)DeleteDataCenterInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dataCenter.DataCenterInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDataCenterInfo 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterInfoService)UpdateDataCenterInfo(dc dataCenter.DataCenterInfo) (err error) {
	err = global.GVA_DB.Save(&dc).Error
	return err
}

// GetDataCenterInfo 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterInfoService)GetDataCenterInfo(id uint) (dc dataCenter.DataCenterInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dc).Error
	return
}

// GetDataCenterInfoInfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (dcService *DataCenterInfoService)GetDataCenterInfoInfoList(info dataCenterReq.DataCenterInfoSearch) (list []dataCenter.DataCenterInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dataCenter.DataCenterInfo{})
    var dcs []dataCenter.DataCenterInfo
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
