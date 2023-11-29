package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    dataCenterReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataCenter/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DataCenterIpServiceApi struct {
}

var ipServiceService = service.ServiceGroupApp.DataCenterServiceGroup.DataCenterIpServiceService


// CreateDataCenterIpService 创建数据中心
// @Tags DataCenterIpService
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterIpService true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ipService/createDataCenterIpService [post]
func (ipServiceApi *DataCenterIpServiceApi) CreateDataCenterIpService(c *gin.Context) {
	var ipService dataCenter.DataCenterIpService
	err := c.ShouldBindJSON(&ipService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "Gateway":{utils.NotEmpty()},
        "DataCenterId":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(ipService, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ipServiceService.CreateDataCenterIpService(&ipService); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataCenterIpService 删除数据中心
// @Tags DataCenterIpService
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterIpService true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ipService/deleteDataCenterIpService [delete]
func (ipServiceApi *DataCenterIpServiceApi) DeleteDataCenterIpService(c *gin.Context) {
	var ipService dataCenter.DataCenterIpService
	err := c.ShouldBindJSON(&ipService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ipServiceService.DeleteDataCenterIpService(ipService); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataCenterIpServiceByIds 批量删除数据中心
// @Tags DataCenterIpService
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ipService/deleteDataCenterIpServiceByIds [delete]
func (ipServiceApi *DataCenterIpServiceApi) DeleteDataCenterIpServiceByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ipServiceService.DeleteDataCenterIpServiceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataCenterIpService 更新数据中心
// @Tags DataCenterIpService
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterIpService true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ipService/updateDataCenterIpService [put]
func (ipServiceApi *DataCenterIpServiceApi) UpdateDataCenterIpService(c *gin.Context) {
	var ipService dataCenter.DataCenterIpService
	err := c.ShouldBindJSON(&ipService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "Gateway":{utils.NotEmpty()},
          "DataCenterId":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(ipService, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ipServiceService.UpdateDataCenterIpService(ipService); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataCenterIpService 用id查询数据中心
// @Tags DataCenterIpService
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenter.DataCenterIpService true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ipService/findDataCenterIpService [get]
func (ipServiceApi *DataCenterIpServiceApi) FindDataCenterIpService(c *gin.Context) {
	var ipService dataCenter.DataCenterIpService
	err := c.ShouldBindQuery(&ipService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reipService, err := ipServiceService.GetDataCenterIpService(ipService.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reipService": reipService}, c)
	}
}

// GetDataCenterIpServiceList 分页获取数据中心列表
// @Tags DataCenterIpService
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenterReq.DataCenterIpServiceSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ipService/getDataCenterIpServiceList [get]
func (ipServiceApi *DataCenterIpServiceApi) GetDataCenterIpServiceList(c *gin.Context) {
	var pageInfo dataCenterReq.DataCenterIpServiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ipServiceService.GetDataCenterIpServiceInfoList(pageInfo); err != nil {
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
