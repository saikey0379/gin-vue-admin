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

type DataCenterIpManageApi struct {
}

var ipManageService = service.ServiceGroupApp.DataCenterServiceGroup.DataCenterIpManageService


// CreateDataCenterIpManage 创建数据中心
// @Tags DataCenterIpManage
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterIpManage true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ipManage/createDataCenterIpManage [post]
func (ipManageApi *DataCenterIpManageApi) CreateDataCenterIpManage(c *gin.Context) {
	var ipManage dataCenter.DataCenterIpManage
	err := c.ShouldBindJSON(&ipManage)
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
	if err := utils.Verify(ipManage, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ipManageService.CreateDataCenterIpManage(&ipManage); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataCenterIpManage 删除数据中心
// @Tags DataCenterIpManage
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterIpManage true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ipManage/deleteDataCenterIpManage [delete]
func (ipManageApi *DataCenterIpManageApi) DeleteDataCenterIpManage(c *gin.Context) {
	var ipManage dataCenter.DataCenterIpManage
	err := c.ShouldBindJSON(&ipManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ipManageService.DeleteDataCenterIpManage(ipManage); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataCenterIpManageByIds 批量删除数据中心
// @Tags DataCenterIpManage
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ipManage/deleteDataCenterIpManageByIds [delete]
func (ipManageApi *DataCenterIpManageApi) DeleteDataCenterIpManageByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ipManageService.DeleteDataCenterIpManageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataCenterIpManage 更新数据中心
// @Tags DataCenterIpManage
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterIpManage true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ipManage/updateDataCenterIpManage [put]
func (ipManageApi *DataCenterIpManageApi) UpdateDataCenterIpManage(c *gin.Context) {
	var ipManage dataCenter.DataCenterIpManage
	err := c.ShouldBindJSON(&ipManage)
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
    if err := utils.Verify(ipManage, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ipManageService.UpdateDataCenterIpManage(ipManage); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataCenterIpManage 用id查询数据中心
// @Tags DataCenterIpManage
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenter.DataCenterIpManage true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ipManage/findDataCenterIpManage [get]
func (ipManageApi *DataCenterIpManageApi) FindDataCenterIpManage(c *gin.Context) {
	var ipManage dataCenter.DataCenterIpManage
	err := c.ShouldBindQuery(&ipManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reipManage, err := ipManageService.GetDataCenterIpManage(ipManage.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reipManage": reipManage}, c)
	}
}

// GetDataCenterIpManageList 分页获取数据中心列表
// @Tags DataCenterIpManage
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenterReq.DataCenterIpManageSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ipManage/getDataCenterIpManageList [get]
func (ipManageApi *DataCenterIpManageApi) GetDataCenterIpManageList(c *gin.Context) {
	var pageInfo dataCenterReq.DataCenterIpManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ipManageService.GetDataCenterIpManageInfoList(pageInfo); err != nil {
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
