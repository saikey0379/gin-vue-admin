package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileCommandService struct {
}

// CreateFileCommand 创建命令信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandService *FileCommandService) CreateFileCommand(fileCommand *file.FileCommand) (err error) {
	err = global.GVA_DB.Create(fileCommand).Error
	return err
}

// DeleteFileCommand 删除命令信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandService *FileCommandService)DeleteFileCommand(fileCommand file.FileCommand) (err error) {
	err = global.GVA_DB.Delete(&fileCommand).Error
	return err
}

// DeleteFileCommandByIds 批量删除命令信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandService *FileCommandService)DeleteFileCommandByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FileCommand{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFileCommand 更新命令信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandService *FileCommandService)UpdateFileCommand(fileCommand file.FileCommand) (err error) {
	err = global.GVA_DB.Save(&fileCommand).Error
	return err
}

// GetFileCommand 根据id获取命令信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandService *FileCommandService)GetFileCommand(id uint) (fileCommand file.FileCommand, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileCommand).Error
	return
}

// GetFileCommandInfoList 分页获取命令信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommandService *FileCommandService)GetFileCommandInfoList(info fileReq.FileCommandSearch) (list []file.FileCommand, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&file.FileCommand{})
    var fileCommands []file.FileCommand
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
	
	err = db.Find(&fileCommands).Error
	return  fileCommands, total, err
}
