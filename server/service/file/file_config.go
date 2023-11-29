package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileConfigService struct {
}

// CreateFileConfig 创建配置文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) CreateFileConfig(fileConfig *file.FileConfig) (err error) {
	err = global.GVA_DB.Create(fileConfig).Error
	return err
}

// DeleteFileConfig 删除配置文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) DeleteFileConfig(fileConfig file.FileConfig) (err error) {
	err = global.GVA_DB.Delete(&fileConfig).Error
	return err
}

// DeleteFileConfigByIds 批量删除配置文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) DeleteFileConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FileConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFileConfig 更新配置文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) UpdateFileConfig(fileConfig file.FileConfig) (err error) {
	err = global.GVA_DB.Save(&fileConfig).Error
	return err
}

// GetFileConfig 根据id获取配置文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) GetFileConfig(id uint) (fileConfig file.FileConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileConfig).Error
	return
}

// GetFileConfigByName 根据id获取脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) GetFileConfigByName(name string) (fileConfig file.FileConfig, err error) {
	err = global.GVA_DB.Where("name = ?", name).First(&fileConfig).Error
	return
}

// GetFileConfigInfoList 分页获取配置文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileConfigService *FileConfigService) GetFileConfigInfoList(info fileReq.FileConfigSearch) (list []file.FileConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&file.FileConfig{})
	var fileConfigs []file.FileConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}
	if info.Manager != "" {
		db = db.Where("manager = ?", info.Manager)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["name"] = true
	orderMap["manager"] = true
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

	err = db.Find(&fileConfigs).Error
	return fileConfigs, total, err
}
