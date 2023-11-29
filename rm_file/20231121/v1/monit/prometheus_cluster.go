package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/monit"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PrometheusClusterApi struct {
}

var prometheusClusterService = service.ServiceGroupApp.MonitServiceGroup.PrometheusClusterService


// CreatePrometheusCluster 创建集群管理
// @Tags PrometheusCluster
// @Summary 创建集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.PrometheusCluster true "创建集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /prometheusCluster/createPrometheusCluster [post]
func (prometheusClusterApi *PrometheusClusterApi) CreatePrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindJSON(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
        "SshUser":{utils.NotEmpty()},
        "SshRsa":{utils.NotEmpty()},
        "Nodes":{utils.NotEmpty()},
        "PathConf":{utils.NotEmpty()},
    }
	if err := utils.Verify(prometheusCluster, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := prometheusClusterService.CreatePrometheusCluster(&prometheusCluster); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePrometheusCluster 删除集群管理
// @Tags PrometheusCluster
// @Summary 删除集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.PrometheusCluster true "删除集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prometheusCluster/deletePrometheusCluster [delete]
func (prometheusClusterApi *PrometheusClusterApi) DeletePrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindJSON(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := prometheusClusterService.DeletePrometheusCluster(prometheusCluster); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePrometheusClusterByIds 批量删除集群管理
// @Tags PrometheusCluster
// @Summary 批量删除集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /prometheusCluster/deletePrometheusClusterByIds [delete]
func (prometheusClusterApi *PrometheusClusterApi) DeletePrometheusClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := prometheusClusterService.DeletePrometheusClusterByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePrometheusCluster 更新集群管理
// @Tags PrometheusCluster
// @Summary 更新集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.PrometheusCluster true "更新集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /prometheusCluster/updatePrometheusCluster [put]
func (prometheusClusterApi *PrometheusClusterApi) UpdatePrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindJSON(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
          "SshUser":{utils.NotEmpty()},
          "SshRsa":{utils.NotEmpty()},
          "Nodes":{utils.NotEmpty()},
          "PathConf":{utils.NotEmpty()},
      }
    if err := utils.Verify(prometheusCluster, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := prometheusClusterService.UpdatePrometheusCluster(prometheusCluster); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPrometheusCluster 用id查询集群管理
// @Tags PrometheusCluster
// @Summary 用id查询集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monit.PrometheusCluster true "用id查询集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prometheusCluster/findPrometheusCluster [get]
func (prometheusClusterApi *PrometheusClusterApi) FindPrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reprometheusCluster, err := prometheusClusterService.GetPrometheusCluster(prometheusCluster.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprometheusCluster": reprometheusCluster}, c)
	}
}

// GetPrometheusClusterList 分页获取集群管理列表
// @Tags PrometheusCluster
// @Summary 分页获取集群管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monitReq.PrometheusClusterSearch true "分页获取集群管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prometheusCluster/getPrometheusClusterList [get]
func (prometheusClusterApi *PrometheusClusterApi) GetPrometheusClusterList(c *gin.Context) {
	var pageInfo monitReq.PrometheusClusterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := prometheusClusterService.GetPrometheusClusterInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
