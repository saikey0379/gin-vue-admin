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

type DataCenterInfoApi struct {
}

var dcService = service.ServiceGroupApp.DataCenterServiceGroup.DataCenterInfoService


// CreateDataCenterInfo 创建数据中心
// @Tags DataCenterInfo
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterInfo true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dc/createDataCenterInfo [post]
func (dcApi *DataCenterInfoApi) CreateDataCenterInfo(c *gin.Context) {
	var dc dataCenter.DataCenterInfo
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(dc, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := dcService.CreateDataCenterInfo(&dc); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataCenterInfo 删除数据中心
// @Tags DataCenterInfo
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterInfo true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dc/deleteDataCenterInfo [delete]
func (dcApi *DataCenterInfoApi) DeleteDataCenterInfo(c *gin.Context) {
	var dc dataCenter.DataCenterInfo
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dcService.DeleteDataCenterInfo(dc); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataCenterInfoByIds 批量删除数据中心
// @Tags DataCenterInfo
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dc/deleteDataCenterInfoByIds [delete]
func (dcApi *DataCenterInfoApi) DeleteDataCenterInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dcService.DeleteDataCenterInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataCenterInfo 更新数据中心
// @Tags DataCenterInfo
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenterInfo true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dc/updateDataCenterInfo [put]
func (dcApi *DataCenterInfoApi) UpdateDataCenterInfo(c *gin.Context) {
	var dc dataCenter.DataCenterInfo
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(dc, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := dcService.UpdateDataCenterInfo(dc); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataCenterInfo 用id查询数据中心
// @Tags DataCenterInfo
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenter.DataCenterInfo true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dc/findDataCenterInfo [get]
func (dcApi *DataCenterInfoApi) FindDataCenterInfo(c *gin.Context) {
	var dc dataCenter.DataCenterInfo
	err := c.ShouldBindQuery(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redc, err := dcService.GetDataCenterInfo(dc.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redc": redc}, c)
	}
}

// GetDataCenterInfoList 分页获取数据中心列表
// @Tags DataCenterInfo
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenterReq.DataCenterInfoSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dc/getDataCenterInfoList [get]
func (dcApi *DataCenterInfoApi) GetDataCenterInfoList(c *gin.Context) {
	var pageInfo dataCenterReq.DataCenterInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dcService.GetDataCenterInfoInfoList(pageInfo); err != nil {
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
