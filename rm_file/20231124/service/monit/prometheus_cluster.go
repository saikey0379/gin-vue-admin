package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
)

type PrometheusClusterService struct {
}

// CreatePrometheusCluster 创建监控集群记录
// Author [piexlmax](https://github.com/piexlmax)
func (prometheusClusterService *PrometheusClusterService) CreatePrometheusCluster(prometheusCluster *monit.PrometheusCluster) (err error) {
	err = global.GVA_DB.Create(prometheusCluster).Error
	return err
}

// DeletePrometheusCluster 删除监控集群记录
// Author [piexlmax](https://github.com/piexlmax)
func (prometheusClusterService *PrometheusClusterService)DeletePrometheusCluster(prometheusCluster monit.PrometheusCluster) (err error) {
	err = global.GVA_DB.Delete(&prometheusCluster).Error
	return err
}

// DeletePrometheusClusterByIds 批量删除监控集群记录
// Author [piexlmax](https://github.com/piexlmax)
func (prometheusClusterService *PrometheusClusterService)DeletePrometheusClusterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]monit.PrometheusCluster{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePrometheusCluster 更新监控集群记录
// Author [piexlmax](https://github.com/piexlmax)
func (prometheusClusterService *PrometheusClusterService)UpdatePrometheusCluster(prometheusCluster monit.PrometheusCluster) (err error) {
	err = global.GVA_DB.Save(&prometheusCluster).Error
	return err
}

// GetPrometheusCluster 根据id获取监控集群记录
// Author [piexlmax](https://github.com/piexlmax)
func (prometheusClusterService *PrometheusClusterService)GetPrometheusCluster(id uint) (prometheusCluster monit.PrometheusCluster, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&prometheusCluster).Error
	return
}

// GetPrometheusClusterInfoList 分页获取监控集群记录
// Author [piexlmax](https://github.com/piexlmax)
func (prometheusClusterService *PrometheusClusterService)GetPrometheusClusterInfoList(info monitReq.PrometheusClusterSearch) (list []monit.PrometheusCluster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&monit.PrometheusCluster{})
    var prometheusClusters []monit.PrometheusCluster
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
    if info.Nodes != "" {
        db = db.Where("nodes LIKE ?","%"+ info.Nodes+"%")
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
	
	err = db.Find(&prometheusClusters).Error
	return  prometheusClusters, total, err
}
