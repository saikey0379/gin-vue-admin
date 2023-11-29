package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
)

type DeviceBareMetalService struct {
}

// CreateDeviceBareMetal 创建裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) CreateDeviceBareMetal(deviceBareMetal *device.DeviceBareMetal) (err error) {
	err = global.GVA_DB.Create(deviceBareMetal).Error
	return err
}

// DeleteDeviceBareMetal 删除裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) DeleteDeviceBareMetal(deviceBareMetal device.DeviceBareMetal) (err error) {
	err = global.GVA_DB.Delete(&deviceBareMetal).Error
	return err
}

// DeleteDeviceBareMetalByIds 批量删除裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) DeleteDeviceBareMetalByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]device.DeviceBareMetal{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDeviceBareMetal 更新裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) UpdateDeviceBareMetal(deviceBareMetal device.DeviceBareMetal) (err error) {
	err = global.GVA_DB.Save(&deviceBareMetal).Error
	return err
}

// GetDeviceBareMetal 根据id获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) GetDeviceBareMetal(id uint) (deviceBareMetal device.DeviceBareMetal, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&deviceBareMetal).Error
	return
}

// GetDeviceBareMetalBySn 根据sn获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) GetDeviceBareMetalBySn(sn string) (deviceBareMetal device.DeviceBareMetal, err error) {
	err = global.GVA_DB.Where("sn = ?", sn).First(&deviceBareMetal).Error
	return
}

// GetDeviceBareMetalInfoList 分页获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) GetDeviceBareMetalInfoList(info deviceReq.DeviceBareMetalSearch) (list []device.DeviceBareMetal, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&device.DeviceBareMetal{})
	var deviceBareMetals []device.DeviceBareMetal
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Sn != "" {
		db = db.Where("sn LIKE ?", "%"+info.Sn+"%")
	}
	if info.AssetId != "" {
		db = db.Where("asset_id LIKE ?", "%"+info.AssetId+"%")
	}
	if info.Hostname != "" {
		db = db.Where("hostname LIKE ?", "%"+info.Hostname+"%")
	}
	if info.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+info.Ip+"%")
	}
	if info.NetworkId != nil {
		db = db.Where("network_id = ?", info.NetworkId)
	}
	if info.ManageIp != "" {
		db = db.Where("manage_ip LIKE ?", "%"+info.ManageIp+"%")
	}
	if info.ManageNetworkId != nil {
		db = db.Where("manage_network_id = ?", info.ManageNetworkId)
	}
	if info.HardwareId != nil {
		db = db.Where("hardware_id = ?", info.HardwareId)
	}
	if info.KickstartId != nil {
		db = db.Where("kickstart_id = ?", info.KickstartId)
	}
	if info.CabinetId != nil {
		db = db.Where("cabinet_id = ?", info.CabinetId)
	}
	if info.ManagerDev != "" {
		db = db.Where("manager_dev LIKE ?", "%"+info.ManagerDev+"%")
	}
	if info.ManagerOps != "" {
		db = db.Where("manager_ops LIKE ?", "%"+info.ManagerOps+"%")
	}
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Describe != "" {
		db = db.Where("describe LIKE ?", "%"+info.Describe+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
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
	orderMap["status"] = true
	orderMap["manager_dev"] = true
	orderMap["manager_ops"] = true
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

	err = db.Find(&deviceBareMetals).Error
	return deviceBareMetals, total, err
}

// CountDeviceBareMetalByIpService 根据业务ip获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) CountDeviceBareMetalByIpService(ip string) (total int64, err error) {
	err = global.GVA_DB.Model(&device.DeviceBareMetal{}).Where("ip = ?", ip).Count(&total).Error
	return
}

// CountDeviceBareMetalByIpManage 根据管理ip获取裸金属设备记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalService *DeviceBareMetalService) CountDeviceBareMetalByIpManage(ip string) (total int64, err error) {
	err = global.GVA_DB.Model(&device.DeviceBareMetal{}).Where("manage_ip = ?", ip).Count(&total).Error
	return
}
