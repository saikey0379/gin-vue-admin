package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileCommonService struct {
}

// CreateFileCommon 创建普通文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommonService *FileCommonService) CreateFileCommon(fileCommon *file.FileCommon) (err error) {
	err = global.GVA_DB.Create(fileCommon).Error
	return err
}

// DeleteFileCommon 删除普通文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommonService *FileCommonService)DeleteFileCommon(fileCommon file.FileCommon) (err error) {
	err = global.GVA_DB.Delete(&fileCommon).Error
	return err
}

// DeleteFileCommonByIds 批量删除普通文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommonService *FileCommonService)DeleteFileCommonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FileCommon{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFileCommon 更新普通文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommonService *FileCommonService)UpdateFileCommon(fileCommon file.FileCommon) (err error) {
	err = global.GVA_DB.Save(&fileCommon).Error
	return err
}

// GetFileCommon 根据id获取普通文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommonService *FileCommonService)GetFileCommon(id uint) (fileCommon file.FileCommon, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileCommon).Error
	return
}

// GetFileCommonInfoList 分页获取普通文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileCommonService *FileCommonService)GetFileCommonInfoList(info fileReq.FileCommonSearch) (list []file.FileCommon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&file.FileCommon{})
    var fileCommons []file.FileCommon
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
	
	err = db.Find(&fileCommons).Error
	return  fileCommons, total, err
}
