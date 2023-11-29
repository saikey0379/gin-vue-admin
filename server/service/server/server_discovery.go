package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/server"
	serverReq "github.com/flipped-aurora/gin-vue-admin/server/model/server/request"
)

type ServerDiscoveryService struct {
}

// CreateServerDiscovery 创建ServerDiscovery记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) CreateServerDiscovery(serverDiscovery *server.ServerDiscovery) (err error) {
	err = global.GVA_DB.Create(serverDiscovery).Error
	return err
}

// DeleteServerDiscovery 删除ServerDiscovery记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) DeleteServerDiscovery(serverDiscovery server.ServerDiscovery) (err error) {
	err = global.GVA_DB.Delete(&serverDiscovery).Error
	return err
}

// DeleteServerDiscoveryByIds 批量删除ServerDiscovery记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) DeleteServerDiscoveryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]server.ServerDiscovery{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateServerDiscovery 更新ServerDiscovery记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) UpdateServerDiscovery(serverDiscovery server.ServerDiscovery) (err error) {
	err = global.GVA_DB.Save(&serverDiscovery).Error
	return err
}

// GetServerDiscovery 根据id获取ServerDiscovery记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) GetServerDiscovery(id uint) (serverDiscovery server.ServerDiscovery, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&serverDiscovery).Error
	return
}

// GetServerDiscoveryBySn 根据sn获取ServerInfo记录
// Author [piexlmax](https://github.根据id获取ServerDiscovery记录/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) GetServerDiscoveryBySn(sn string) (serverDiscovery server.ServerDiscovery, err error) {
	err = global.GVA_DB.Where("sn = ?", sn).First(&serverDiscovery).Error
	return
}

// GetServerDiscoveryInfoList 分页获取ServerDiscovery记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverDiscoveryService *ServerDiscoveryService) GetServerDiscoveryInfoList(info serverReq.ServerDiscoverySearch) (list []server.ServerDiscovery, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&server.ServerDiscovery{})
	var serverDiscoverys []server.ServerDiscovery
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Sn != "" {
		db = db.Where("sn = ?", info.Sn)
	}
	if info.Hostname != "" {
		db = db.Where("hostname = ?", info.Hostname)
	}
	if info.Ip != "" {
		db = db.Where("ip = ?", info.Ip)
	}
	if info.VersionOs != "" {
		db = db.Where("version_os = ?", info.VersionOs)
	}
	if info.VersionKernel != "" {
		db = db.Where("version_kernel = ?", info.VersionKernel)
	}
	if info.Manufacturer != "" {
		db = db.Where("manufacturer = ?", info.Manufacturer)
	}
	if info.ModelName != "" {
		db = db.Where("model_name = ?", info.ModelName)
	}
	if info.CpuSum != nil {
		db = db.Where("cpu_sum = ?", info.CpuSum)
	}
	if info.MemorySum != nil {
		db = db.Where("memory_sum = ?", info.MemorySum)
	}
	if info.NicInfo != "" {
		db = db.Where("nic_info LIKE ?", "%"+info.NicInfo+"%")
	}
	if info.Gpu != "" {
		db = db.Where("gpu LIKE ?", "%"+info.Gpu+"%")
	}
	if info.Motherboard != "" {
		db = db.Where("motherboard LIKE ?", "%"+info.Motherboard+"%")
	}
	if info.ServerType != nil {
		db = db.Where("server_type = ?", info.ServerType)
	}
	if info.VersionAgent != "" {
		db = db.Where("version_agent = ?", info.VersionAgent)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["hostname"] = true
	orderMap["ip"] = true
	orderMap["version_os"] = true
	orderMap["version_kernel"] = true
	orderMap["manufacturer"] = true
	orderMap["cpu_sum"] = true
	orderMap["memory_sum"] = true
	orderMap["server_type"] = true
	orderMap["version_agent"] = true
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

	err = db.Find(&serverDiscoverys).Error
	return serverDiscoverys, total, err
}
