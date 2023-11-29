package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
)

type IdcInfottService struct {
}

// CreateIdcInfott 创建数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfottService *IdcInfottService) CreateIdcInfott(idcInfott *idc.IdcInfott) (err error) {
	err = global.GVA_DB.Create(idcInfott).Error
	return err
}

// DeleteIdcInfott 删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfottService *IdcInfottService)DeleteIdcInfott(idcInfott idc.IdcInfott) (err error) {
	err = global.GVA_DB.Delete(&idcInfott).Error
	return err
}

// DeleteIdcInfottByIds 批量删除数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfottService *IdcInfottService)DeleteIdcInfottByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]idc.IdcInfott{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIdcInfott 更新数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfottService *IdcInfottService)UpdateIdcInfott(idcInfott idc.IdcInfott) (err error) {
	err = global.GVA_DB.Save(&idcInfott).Error
	return err
}

// GetIdcInfott 根据id获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfottService *IdcInfottService)GetIdcInfott(id uint) (idcInfott idc.IdcInfott, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&idcInfott).Error
	return
}

// GetIdcInfottInfoList 分页获取数据中心记录
// Author [piexlmax](https://github.com/piexlmax)
func (idcInfottService *IdcInfottService)GetIdcInfottInfoList(info idcReq.IdcInfottSearch) (list []idc.IdcInfott, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&idc.IdcInfott{})
    var idcInfotts []idc.IdcInfott
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
	
	err = db.Find(&idcInfotts).Error
	return  idcInfotts, total, err
}
