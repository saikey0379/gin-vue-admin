package bareMetal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bareMetal"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    bareMetalReq "github.com/flipped-aurora/gin-vue-admin/server/model/bareMetal/request"
)

type BareMetalService struct {
}

// CreateBareMetal 创建裸金属管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (DeviceBareMetalService *BareMetalService) CreateBareMetal(DeviceBareMetal *bareMetal.BareMetal) (err error) {
	err = global.GVA_DB.Create(DeviceBareMetal).Error
	return err
}

// DeleteBareMetal 删除裸金属管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (DeviceBareMetalService *BareMetalService)DeleteBareMetal(DeviceBareMetal bareMetal.BareMetal) (err error) {
	err = global.GVA_DB.Delete(&DeviceBareMetal).Error
	return err
}

// DeleteBareMetalByIds 批量删除裸金属管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (DeviceBareMetalService *BareMetalService)DeleteBareMetalByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]bareMetal.BareMetal{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBareMetal 更新裸金属管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (DeviceBareMetalService *BareMetalService)UpdateBareMetal(DeviceBareMetal bareMetal.BareMetal) (err error) {
	err = global.GVA_DB.Save(&DeviceBareMetal).Error
	return err
}

// GetBareMetal 根据id获取裸金属管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (DeviceBareMetalService *BareMetalService)GetBareMetal(id uint) (DeviceBareMetal bareMetal.BareMetal, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&DeviceBareMetal).Error
	return
}

// GetBareMetalInfoList 分页获取裸金属管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (DeviceBareMetalService *BareMetalService)GetBareMetalInfoList(info bareMetalReq.BareMetalSearch) (list []bareMetal.BareMetal, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&bareMetal.BareMetal{})
    var DeviceBareMetals []bareMetal.BareMetal
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
    if info.ManageIp != "" {
        db = db.Where("manage_ip LIKE ?","%"+ info.ManageIp+"%")
    }
    if info.NetworkId != nil {
        db = db.Where("network_id = ?",info.NetworkId)
    }
    if info.ManageNetworkId != nil {
        db = db.Where("manage_network_id = ?",info.ManageNetworkId)
    }
    if info.HardwareId != nil {
        db = db.Where("hardware_id = ?",info.HardwareId)
    }
    if info.PxeId != nil {
        db = db.Where("pxe_id = ?",info.PxeId)
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
    if info.DeviceDescribe != "" {
        db = db.Where("device_describe LIKE ?","%"+ info.DeviceDescribe+"%")
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
         	orderMap["manage_ip"] = true
         	orderMap["network_id"] = true
         	orderMap["manage_network_id"] = true
         	orderMap["hardware_id"] = true
         	orderMap["pxe_id"] = true
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
	
	err = db.Find(&DeviceBareMetals).Error
	return  DeviceBareMetals, total, err
}
