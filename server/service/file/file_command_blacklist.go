package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileCommadBlacklistService struct {
}

// CreateFileCommadBlacklist 创建命令黑名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandBlacklistService *FileCommadBlacklistService) CreateFileCommadBlacklist(fileCommandBlacklist *file.FileCommadBlacklist) (err error) {
	err = global.GVA_DB.Create(fileCommandBlacklist).Error
	return err
}

// DeleteFileCommadBlacklist 删除命令黑名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandBlacklistService *FileCommadBlacklistService)DeleteFileCommadBlacklist(fileCommandBlacklist file.FileCommadBlacklist) (err error) {
	err = global.GVA_DB.Delete(&fileCommandBlacklist).Error
	return err
}

// DeleteFileCommadBlacklistByIds 批量删除命令黑名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandBlacklistService *FileCommadBlacklistService)DeleteFileCommadBlacklistByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FileCommadBlacklist{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFileCommadBlacklist 更新命令黑名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandBlacklistService *FileCommadBlacklistService)UpdateFileCommadBlacklist(fileCommandBlacklist file.FileCommadBlacklist) (err error) {
	err = global.GVA_DB.Save(&fileCommandBlacklist).Error
	return err
}

// GetFileCommadBlacklist 根据id获取命令黑名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandBlacklistService *FileCommadBlacklistService)GetFileCommadBlacklist(id uint) (fileCommandBlacklist file.FileCommadBlacklist, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileCommandBlacklist).Error
	return
}

// GetFileCommadBlacklistInfoList 分页获取命令黑名单记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandBlacklistService *FileCommadBlacklistService)GetFileCommadBlacklistInfoList(info fileReq.FileCommadBlacklistSearch) (list []file.FileCommadBlacklist, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&file.FileCommadBlacklist{})
    var fileCommandBlacklists []file.FileCommadBlacklist
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Description != "" {
        db = db.Where("description LIKE ?","%"+ info.Description+"%")
    }
    if info.Manager != "" {
        db = db.Where("manager = ?",info.Manager)
    }
	err = db.Count(&total).Error
	if err!=nil {
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
	
	err = db.Find(&fileCommandBlacklists).Error
	return  fileCommandBlacklists, total, err
}
