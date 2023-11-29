package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
)

type BareMetalService struct {
}

// CreateBareMetal 创建裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmService *BareMetalService) CreateBareMetal(bm *device.BareMetal) (err error) {
	err = global.GVA_DB.Create(bm).Error
	return err
}

// DeleteBareMetal 删除裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmService *BareMetalService)DeleteBareMetal(bm device.BareMetal) (err error) {
	err = global.GVA_DB.Delete(&bm).Error
	return err
}

// DeleteBareMetalByIds 批量删除裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmService *BareMetalService)DeleteBareMetalByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]device.BareMetal{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBareMetal 更新裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmService *BareMetalService)UpdateBareMetal(bm device.BareMetal) (err error) {
	err = global.GVA_DB.Save(&bm).Error
	return err
}

// GetBareMetal 根据id获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmService *BareMetalService)GetBareMetal(id uint) (bm device.BareMetal, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&bm).Error
	return
}

// GetBareMetalInfoList 分页获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmService *BareMetalService)GetBareMetalInfoList(info deviceReq.BareMetalSearch) (list []device.BareMetal, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&device.BareMetal{})
    var bms []device.BareMetal
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Sn != "" {
        db = db.Where("sn LIKE ?","%"+ info.Sn+"%")
    }
    if info.AssetId != "" {
        db = db.Where("asset_id LIKE ?","%"+ info.AssetId+"%")
    }
    if info.Hostname != "" {
        db = db.Where("hostname LIKE ?","%"+ info.Hostname+"%")
    }
    if info.Ip != "" {
        db = db.Where("ip LIKE ?","%"+ info.Ip+"%")
    }
    if info.NetworkId != nil {
        db = db.Where("network_id = ?",info.NetworkId)
    }
    if info.ManageIp != "" {
        db = db.Where("manage_ip LIKE ?","%"+ info.ManageIp+"%")
    }
    if info.ManageNetworkId != nil {
        db = db.Where("manage_network_id = ?",info.ManageNetworkId)
    }
    if info.HardwareId != nil {
        db = db.Where("hardware_id = ?",info.HardwareId)
    }
    if info.KickstartId != nil {
        db = db.Where("kickstart_id = ?",info.KickstartId)
    }
    if info.CabinetId != nil {
        db = db.Where("cabinet_id = ?",info.CabinetId)
    }
    if info.CabinetU != "" {
        db = db.Where("cabinet_u LIKE ?","%"+ info.CabinetU+"%")
    }
    if info.ManagerDev != "" {
        db = db.Where("manager_dev LIKE ?","%"+ info.ManagerDev+"%")
    }
    if info.ManagerOps != "" {
        db = db.Where("manager_ops LIKE ?","%"+ info.ManagerOps+"%")
    }
    if info.LabelId != nil {
        db = db.Where("label_id = ?",info.LabelId)
    }
    if info.Describe != "" {
        db = db.Where("device_describe LIKE ?","%"+ info.Describe+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["asset_id"] = true
         	orderMap["hostname"] = true
         	orderMap["ip"] = true
         	orderMap["network_id"] = true
         	orderMap["manage_ip"] = true
         	orderMap["manage_network_id"] = true
         	orderMap["hardware_id"] = true
         	orderMap["kickstart_id"] = true
         	orderMap["cabinet_id"] = true
         	orderMap["cabinet_u"] = true
         	orderMap["status"] = true
         	orderMap["manager_dev"] = true
         	orderMap["manager_ops"] = true
         	orderMap["label_id"] = true
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
	
	err = db.Find(&bms).Error
	return  bms, total, err
}
