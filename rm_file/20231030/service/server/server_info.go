package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/server"
	serverReq "github.com/flipped-aurora/gin-vue-admin/server/model/server/request"
)

type ServerInfoService struct {
}

// CreateServerInfo 创建ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) CreateServerInfo(serverInfo *server.ServerInfo) (err error) {
	err = global.GVA_DB.Create(serverInfo).Error
	return err
}

// DeleteServerInfo 删除ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) DeleteServerInfo(serverInfo server.ServerInfo) (err error) {
	err = global.GVA_DB.Delete(&serverInfo).Error
	return err
}

// DeleteServerInfoByIds 批量删除ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) DeleteServerInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]server.ServerInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateServerInfo 更新ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) UpdateServerInfo(serverInfo server.ServerInfo) (err error) {
	err = global.GVA_DB.Save(&serverInfo).Error
	return err
}

// GetServerInfo 根据id获取ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) GetServerInfo(id uint) (serverInfo server.ServerInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&serverInfo).Error
	return
}

// GetServerInfoBySn 根据sn获取ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) GetServerInfoBySn(sn string) (serverInfo server.ServerInfo, err error) {
	err = global.GVA_DB.Where("sn = ?", sn).First(&serverInfo).Error
	return
}

// GetServerInfoInfoList 分页获取ServerInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverInfoService *ServerInfoService) GetServerInfoInfoList(info serverReq.ServerInfoSearch) (list []server.ServerInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&server.ServerInfo{})
	var serverInfos []server.ServerInfo
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
	if info.ServerType != nil {
		db = db.Where("server_type = ?", info.ServerType)
	}
	if info.AgentVersion != "" {
		db = db.Where("agent_version = ?", info.AgentVersion)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["hostname"] = true
	orderMap["ip"] = true
	orderMap["manufacturer"] = true
	orderMap["cpu_sum"] = true
	orderMap["memory_sum"] = true
	orderMap["server_type"] = true
	orderMap["agent_version"] = true
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

	err = db.Find(&serverInfos).Error
	return serverInfos, total, err
}
