package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcIpSubnetService struct {
}

// CreateIdcIpSubnet 创建数据中心子网记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSubnetService *IdcIpSubnetService) CreateIdcIpSubnet(idcIpSubnet *idc.IdcIpSubnet) (err error) {
	err = global.GVA_DB.Create(idcIpSubnet).Error
	return err
}

// DeleteIdcIpSubnet 删除数据中心子网记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSubnetService *IdcIpSubnetService)DeleteIdcIpSubnet(idcIpSubnet idc.IdcIpSubnet) (err error) {
	err = global.GVA_DB.Delete(&idcIpSubnet).Error
	return err
}

// DeleteIdcIpSubnetByIds 批量删除数据中心子网记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSubnetService *IdcIpSubnetService)DeleteIdcIpSubnetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcIpSubnet{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcIpSubnet 更新数据中心子网记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSubnetService *IdcIpSubnetService)UpdateIdcIpSubnet(idcIpSubnet idc.IdcIpSubnet) (err error) {
	err = global.GVA_DB.Save(&idcIpSubnet).Error
	return err
}

// GetIdcIpSubnet 根据id获取数据中心子网记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSubnetService *IdcIpSubnetService)GetIdcIpSubnet(id uint) (idcIpSubnet idc.IdcIpSubnet, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcIpSubnet).Error
	return
}

// GetIdcIpSubnetInfoList 分页获取数据中心子网记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcIpSubnetService *IdcIpSubnetService)GetIdcIpSubnetInfoList(info idcReq.IdcIpSubnetSearch) (list []idc.IdcIpSubnet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcIpSubnet{})
    var idcIpSubnets []idc.IdcIpSubnet
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Network != "" {
        db = db.Where("network LIKE ?","%"+ info.Network+"%")
    }
    if info.IdcId != nil {
        db = db.Where("idc_id = ?",info.IdcId)
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
         	orderMap["idc_id"] = true
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
	
	err = db.Find(&idcIpSubnets).Error
	return  idcIpSubnets, total, err
}
