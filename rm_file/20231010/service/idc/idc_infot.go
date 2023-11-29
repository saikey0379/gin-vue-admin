package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcInfotService struct {
}

// CreateIdcInfot 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfotService *IdcInfotService) CreateIdcInfot(idcInfot *idc.IdcInfot) (err error) {
	err = global.GVA_DB.Create(idcInfot).Error
	return err
}

// DeleteIdcInfot 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfotService *IdcInfotService)DeleteIdcInfot(idcInfot idc.IdcInfot) (err error) {
	err = global.GVA_DB.Delete(&idcInfot).Error
	return err
}

// DeleteIdcInfotByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfotService *IdcInfotService)DeleteIdcInfotByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcInfot{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcInfot 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfotService *IdcInfotService)UpdateIdcInfot(idcInfot idc.IdcInfot) (err error) {
	err = global.GVA_DB.Save(&idcInfot).Error
	return err
}

// GetIdcInfot 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfotService *IdcInfotService)GetIdcInfot(id uint) (idcInfot idc.IdcInfot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcInfot).Error
	return
}

// GetIdcInfotInfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfotService *IdcInfotService)GetIdcInfotInfoList(info idcReq.IdcInfotSearch) (list []idc.IdcInfot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcInfot{})
    var idcInfots []idc.IdcInfot
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Label != "" {
        db = db.Where("label LIKE ?","%"+ info.Label+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
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
	
	err = db.Find(&idcInfots).Error
	return  idcInfots, total, err
}
