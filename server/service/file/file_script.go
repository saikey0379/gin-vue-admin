package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileScriptService struct {
}

// CreateFileScript 创建脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) CreateFileScript(fileScript *file.FileScript) (err error) {
	err = global.GVA_DB.Create(fileScript).Error
	return err
}

// DeleteFileScript 删除脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) DeleteFileScript(fileScript file.FileScript) (err error) {
	err = global.GVA_DB.Delete(&fileScript).Error
	return err
}

// DeleteFileScriptByIds 批量删除脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) DeleteFileScriptByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FileScript{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFileScript 更新脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) UpdateFileScript(fileScript file.FileScript) (err error) {
	err = global.GVA_DB.Save(&fileScript).Error
	return err
}

// GetFileScript 根据id获取脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) GetFileScript(id uint) (fileScript file.FileScript, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileScript).Error
	return
}

// GetFileScriptByName 根据id获取脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) GetFileScriptByName(name string) (fileScript file.FileScript, err error) {
	err = global.GVA_DB.Where("name = ?", name).First(&fileScript).Error
	return
}

// GetFileScriptInfoList 分页获取脚本文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileScriptService *FileScriptService) GetFileScriptInfoList(info fileReq.FileScriptSearch) (list []file.FileScript, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&file.FileScript{})
	var fileScripts []file.FileScript
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

	err = db.Find(&fileScripts).Error
	return fileScripts, total, err
}
