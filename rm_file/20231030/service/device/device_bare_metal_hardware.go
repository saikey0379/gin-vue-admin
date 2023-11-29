package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
)

type DeviceBareMetalHardwareService struct {
}

// CreateDeviceBareMetalHardware 创建裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) CreateDeviceBareMetalHardware(deviceBareMetalHardware *device.DeviceBareMetalHardware) (err error) {
	err = global.GVA_DB.Create(deviceBareMetalHardware).Error
	return err
}

// DeleteDeviceBareMetalHardware 删除裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) DeleteDeviceBareMetalHardware(deviceBareMetalHardware device.DeviceBareMetalHardware) (err error) {
	err = global.GVA_DB.Delete(&deviceBareMetalHardware).Error
	return err
}

// DeleteDeviceBareMetalHardwareByIds 批量删除裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) DeleteDeviceBareMetalHardwareByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]device.DeviceBareMetalHardware{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDeviceBareMetalHardware 更新裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) UpdateDeviceBareMetalHardware(deviceBareMetalHardware device.DeviceBareMetalHardware) (err error) {
	err = global.GVA_DB.Save(&deviceBareMetalHardware).Error
	return err
}

// GetDeviceBareMetalHardware 根据id获取裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) GetDeviceBareMetalHardware(id uint) (deviceBareMetalHardware device.DeviceBareMetalHardware, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&deviceBareMetalHardware).Error
	return
}

// GetDeviceBareMetalHardwareBySn 根据sn获取裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) GetDeviceBareMetalHardwareBySn(sn string) (deviceBareMetalHardware device.DeviceBareMetalHardware, err error) {
	err = global.GVA_DB.Where("sn = ?", sn).First(&deviceBareMetalHardware).Error
	return
}

// GetDeviceBareMetalHardwareInfoList 分页获取裸金属设备硬件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (deviceBareMetalHardwareService *DeviceBareMetalHardwareService) GetDeviceBareMetalHardwareInfoList(info deviceReq.DeviceBareMetalHardwareSearch) (list []device.DeviceBareMetalHardware, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&device.DeviceBareMetalHardware{})
	var deviceBareMetalHardwares []device.DeviceBareMetalHardware
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Sn != "" {
		db = db.Where("sn = ?", info.Sn)
	}
	if info.Manufacturer != "" {
		db = db.Where("manufacturer = ?", info.Manufacturer)
	}
	if info.Model != "" {
		db = db.Where("model = ?", info.Model)
	}
	if info.CpuSum != nil {
		db = db.Where("cpu_sum = ?", info.CpuSum)
	}
	if info.MemorySum != nil {
		db = db.Where("memory_sum = ?", info.MemorySum)
	}
	if info.Nic != "" {
		db = db.Where("nic LIKE ?", "%"+info.Nic+"%")
	}
	if info.Gpu != "" {
		db = db.Where("gpu LIKE ?", "%"+info.Gpu+"%")
	}
	if info.Motherboard != "" {
		db = db.Where("motherboard LIKE ?", "%"+info.Motherboard+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["manufacturer"] = true
	orderMap["cpu_sum"] = true
	orderMap["memory_sum"] = true
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

	err = db.Find(&deviceBareMetalHardwares).Error
	return deviceBareMetalHardwares, total, err
}
