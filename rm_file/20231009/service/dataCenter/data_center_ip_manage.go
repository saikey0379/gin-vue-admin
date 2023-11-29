package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dataCenterReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter/request"
)

type DataCenterIpManageService struct {
}

// CreateDataCenterIpManage 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (ipManageService *DataCenterIpManageService) CreateDataCenterIpManage(ipManage *dataCenter.DataCenterIpManage) (err error) {
	err = global.GVA_DB.Create(ipManage).Error
	return err
}

// DeleteDataCenterIpManage 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (ipManageService *DataCenterIpManageService)DeleteDataCenterIpManage(ipManage dataCenter.DataCenterIpManage) (err error) {
	err = global.GVA_DB.Delete(&ipManage).Error
	return err
}

// DeleteDataCenterIpManageByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (ipManageService *DataCenterIpManageService)DeleteDataCenterIpManageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dataCenter.DataCenterIpManage{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDataCenterIpManage 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (ipManageService *DataCenterIpManageService)UpdateDataCenterIpManage(ipManage dataCenter.DataCenterIpManage) (err error) {
	err = global.GVA_DB.Save(&ipManage).Error
	return err
}

// GetDataCenterIpManage 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (ipManageService *DataCenterIpManageService)GetDataCenterIpManage(id uint) (ipManage dataCenter.DataCenterIpManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ipManage).Error
	return
}

// GetDataCenterIpManageInfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (ipManageService *DataCenterIpManageService)GetDataCenterIpManageInfoList(info dataCenterReq.DataCenterIpManageSearch) (list []dataCenter.DataCenterIpManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dataCenter.DataCenterIpManage{})
    var ipManages []dataCenter.DataCenterIpManage
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Network != "" {
        db = db.Where("network LIKE ?","%"+ info.Network+"%")
    }
    if info.Prefix != nil {
        db = db.Where("netmask = ?",info.Prefix)
    }
    if info.Gateway != "" {
        db = db.Where("gateway LIKE ?","%"+ info.Gateway+"%")
    }
    if info.DataCenterId != nil {
        db = db.Where("data_center_id = ?",info.DataCenterId)
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
         	orderMap["network"] = true
         	orderMap["netmask"] = true
         	orderMap["gateway"] = true
         	orderMap["data_center_id"] = true
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
	
	err = db.Find(&ipManages).Error
	return  ipManages, total, err
}
