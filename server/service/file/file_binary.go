package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
)

type FileBinaryService struct {
}

// CreateFileBinary 创建可执行程序记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileBinaryService *FileBinaryService) CreateFileBinary(fileBinary *file.FileBinary) (err error) {
	err = global.GVA_DB.Create(fileBinary).Error
	return err
}

// DeleteFileBinary 删除可执行程序记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileBinaryService *FileBinaryService)DeleteFileBinary(fileBinary file.FileBinary) (err error) {
	err = global.GVA_DB.Delete(&fileBinary).Error
	return err
}

// DeleteFileBinaryByIds 批量删除可执行程序记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileBinaryService *FileBinaryService)DeleteFileBinaryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FileBinary{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFileBinary 更新可执行程序记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileBinaryService *FileBinaryService)UpdateFileBinary(fileBinary file.FileBinary) (err error) {
	err = global.GVA_DB.Save(&fileBinary).Error
	return err
}

// GetFileBinary 根据id获取可执行程序记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileBinaryService *FileBinaryService)GetFileBinary(id uint) (fileBinary file.FileBinary, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileBinary).Error
	return
}

// GetFileBinaryInfoList 分页获取可执行程序记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileBinaryService *FileBinaryService)GetFileBinaryInfoList(info fileReq.FileBinarySearch) (list []file.FileBinary, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&file.FileBinary{})
    var fileBinarys []file.FileBinary
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
	
	err = db.Find(&fileBinarys).Error
	return  fileBinarys, total, err
}
