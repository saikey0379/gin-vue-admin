package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
)

type SlbDomainService struct {
}

// CreateSlbDomain 创建域名管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbDomainService *SlbDomainService) CreateSlbDomain(slbDomain *slb.SlbDomain) (err error) {
	err = global.GVA_DB.Create(slbDomain).Error
	return err
}

// DeleteSlbDomain 删除域名管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbDomainService *SlbDomainService)DeleteSlbDomain(slbDomain slb.SlbDomain) (err error) {
	err = global.GVA_DB.Delete(&slbDomain).Error
	return err
}

// DeleteSlbDomainByIds 批量删除域名管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbDomainService *SlbDomainService)DeleteSlbDomainByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]slb.SlbDomain{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlbDomain 更新域名管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbDomainService *SlbDomainService)UpdateSlbDomain(slbDomain slb.SlbDomain) (err error) {
	err = global.GVA_DB.Save(&slbDomain).Error
	return err
}

// GetSlbDomain 根据id获取域名管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbDomainService *SlbDomainService)GetSlbDomain(id uint) (slbDomain slb.SlbDomain, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slbDomain).Error
	return
}

// GetSlbDomainInfoList 分页获取域名管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbDomainService *SlbDomainService)GetSlbDomainInfoList(info slbReq.SlbDomainSearch) (list []slb.SlbDomain, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&slb.SlbDomain{})
    var slbDomains []slb.SlbDomain
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Manager != "" {
        db = db.Where("manager = ?",info.Manager)
    }
    if info.ClusterIds != "" {
        db = db.Where("cluster_ids LIKE ?","%"+ info.ClusterIds+"%")
    }
    if info.CertId != nil {
        db = db.Where("cert_id = ?",info.CertId)
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
	
	err = db.Find(&slbDomains).Error
	return  slbDomains, total, err
}
