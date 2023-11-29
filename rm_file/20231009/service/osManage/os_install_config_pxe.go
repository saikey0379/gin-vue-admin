package osManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/osManage/request"
)

type PxeService struct {
}

// CreatePxe 创建PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (PxeConfigService *PxeService) CreatePxe(PxeConfig *osManage.Pxe) (err error) {
	err = global.GVA_DB.Create(PxeConfig).Error
	return err
}

// DeletePxe 删除PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (PxeConfigService *PxeService)DeletePxe(PxeConfig osManage.Pxe) (err error) {
	err = global.GVA_DB.Delete(&PxeConfig).Error
	return err
}

// DeletePxeByIds 批量删除PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (PxeConfigService *PxeService)DeletePxeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]osManage.Pxe{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePxe 更新PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (PxeConfigService *PxeService)UpdatePxe(PxeConfig osManage.Pxe) (err error) {
	err = global.GVA_DB.Save(&PxeConfig).Error
	return err
}

// GetPxe 根据id获取PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (PxeConfigService *PxeService)GetPxe(id uint) (PxeConfig osManage.Pxe, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&PxeConfig).Error
	return
}

// GetPxeInfoList 分页获取PXE配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (PxeConfigService *PxeService)GetPxeInfoList(info osManageReq.PxeSearch) (list []osManage.Pxe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&osManage.Pxe{})
    var PxeConfigs []osManage.Pxe
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
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
	
	err = db.Find(&PxeConfigs).Error
	return  PxeConfigs, total, err
}
