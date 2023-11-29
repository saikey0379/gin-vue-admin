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

type DataCenterApi struct {
}

var dcService = service.ServiceGroupApp.DataCenterServiceGroup.DataCenterService


// CreateDataCenter 创建数据中心
// @Tags DataCenter
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenter true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dc/createDataCenter [post]
func (dcApi *DataCenterApi) CreateDataCenter(c *gin.Context) {
	var dc dataCenter.DataCenter
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
	if err := dcService.CreateDataCenter(&dc); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataCenter 删除数据中心
// @Tags DataCenter
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenter true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dc/deleteDataCenter [delete]
func (dcApi *DataCenterApi) DeleteDataCenter(c *gin.Context) {
	var dc dataCenter.DataCenter
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dcService.DeleteDataCenter(dc); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataCenterByIds 批量删除数据中心
// @Tags DataCenter
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dc/deleteDataCenterByIds [delete]
func (dcApi *DataCenterApi) DeleteDataCenterByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dcService.DeleteDataCenterByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataCenter 更新数据中心
// @Tags DataCenter
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataCenter.DataCenter true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dc/updateDataCenter [put]
func (dcApi *DataCenterApi) UpdateDataCenter(c *gin.Context) {
	var dc dataCenter.DataCenter
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
	if err := dcService.UpdateDataCenter(dc); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataCenter 用id查询数据中心
// @Tags DataCenter
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenter.DataCenter true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dc/findDataCenter [get]
func (dcApi *DataCenterApi) FindDataCenter(c *gin.Context) {
	var dc dataCenter.DataCenter
	err := c.ShouldBindQuery(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redc, err := dcService.GetDataCenter(dc.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redc": redc}, c)
	}
}

// GetDataCenterList 分页获取数据中心列表
// @Tags DataCenter
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataCenterReq.DataCenterSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dc/getDataCenterList [get]
func (dcApi *DataCenterApi) GetDataCenterList(c *gin.Context) {
	var pageInfo dataCenterReq.DataCenterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dcService.GetDataCenterInfoList(pageInfo); err != nil {
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
