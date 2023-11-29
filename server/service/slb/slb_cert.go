package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
)

type SlbCertService struct {
}

// CreateSlbCert 创建证书管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbCertService *SlbCertService) CreateSlbCert(slbCert *slb.SlbCert) (err error) {
	err = global.GVA_DB.Create(slbCert).Error
	return err
}

// DeleteSlbCert 删除证书管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbCertService *SlbCertService)DeleteSlbCert(slbCert slb.SlbCert) (err error) {
	err = global.GVA_DB.Delete(&slbCert).Error
	return err
}

// DeleteSlbCertByIds 批量删除证书管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbCertService *SlbCertService)DeleteSlbCertByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]slb.SlbCert{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlbCert 更新证书管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbCertService *SlbCertService)UpdateSlbCert(slbCert slb.SlbCert) (err error) {
	err = global.GVA_DB.Save(&slbCert).Error
	return err
}

// GetSlbCert 根据id获取证书管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbCertService *SlbCertService)GetSlbCert(id uint) (slbCert slb.SlbCert, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slbCert).Error
	return
}

// GetSlbCertInfoList 分页获取证书管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (slbCertService *SlbCertService)GetSlbCertInfoList(info slbReq.SlbCertSearch) (list []slb.SlbCert, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&slb.SlbCert{})
    var slbCerts []slb.SlbCert
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
    if info.CertFile != "" {
        db = db.Where("cert_file = ?",info.CertFile)
    }
    if info.KeyFile != "" {
        db = db.Where("key_file = ?",info.KeyFile)
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
	
	err = db.Find(&slbCerts).Error
	return  slbCerts, total, err
}
