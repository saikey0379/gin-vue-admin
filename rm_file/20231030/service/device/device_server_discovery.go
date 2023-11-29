package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
)

type DeviceServerDiscoveryService struct {
}

// CreateDeviceServerDiscovery 创建Server发现记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceServerDiscoveryService *DeviceServerDiscoveryService) CreateDeviceServerDiscovery(deviceServerDiscovery *device.DeviceServerDiscovery) (err error) {
	err = global.GVA_DB.Create(deviceServerDiscovery).Error
	return err
}

// DeleteDeviceServerDiscovery 删除Server发现记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceServerDiscoveryService *DeviceServerDiscoveryService)DeleteDeviceServerDiscovery(deviceServerDiscovery device.DeviceServerDiscovery) (err error) {
	err = global.GVA_DB.Delete(&deviceServerDiscovery).Error
	return err
}

// DeleteDeviceServerDiscoveryByIds 批量删除Server发现记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceServerDiscoveryService *DeviceServerDiscoveryService)DeleteDeviceServerDiscoveryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]device.DeviceServerDiscovery{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDeviceServerDiscovery 更新Server发现记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceServerDiscoveryService *DeviceServerDiscoveryService)UpdateDeviceServerDiscovery(deviceServerDiscovery device.DeviceServerDiscovery) (err error) {
	err = global.GVA_DB.Save(&deviceServerDiscovery).Error
	return err
}

// GetDeviceServerDiscovery 根据id获取Server发现记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceServerDiscoveryService *DeviceServerDiscoveryService)GetDeviceServerDiscovery(id uint) (deviceServerDiscovery device.DeviceServerDiscovery, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&deviceServerDiscovery).Error
	return
}

// GetDeviceServerDiscoveryInfoList 分页获取Server发现记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceServerDiscoveryService *DeviceServerDiscoveryService)GetDeviceServerDiscoveryInfoList(info deviceReq.DeviceServerDiscoverySearch) (list []device.DeviceServerDiscovery, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&device.DeviceServerDiscovery{})
    var deviceServerDiscoverys []device.DeviceServerDiscovery
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Sn != "" {
        db = db.Where("sn LIKE ?","%"+ info.Sn+"%")
    }
    if info.Ip != "" {
        db = db.Where("ip LIKE ?","%"+ info.Ip+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["ip"] = true
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
	
	err = db.Find(&deviceServerDiscoverys).Error
	return  deviceServerDiscoverys, total, err
}
