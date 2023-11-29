package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/idc"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type IdcIpServiceApi struct {
}

var idcIpServiceService = service.ServiceGroupApp.IdcServiceGroup.IdcIpServiceService


// CreateIdcIpService 创建数据中心业务网段
// @Tags IdcIpService
// @Summary 创建数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService true "创建数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpService/createIdcIpService [post]
func (idcIpServiceApi *IdcIpServiceApi) CreateIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindJSON(&idcIpService)
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
	if err := utils.Verify(idcIpService, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcIpServiceService.CreateIdcIpService(&idcIpService); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpService 删除数据中心业务网段
// @Tags IdcIpService
// @Summary 删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService true "删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService/deleteIdcIpService [delete]
func (idcIpServiceApi *IdcIpServiceApi) DeleteIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindJSON(&idcIpService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpServiceService.DeleteIdcIpService(idcIpService); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpServiceByIds 批量删除数据中心业务网段
// @Tags IdcIpService
// @Summary 批量删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpService/deleteIdcIpServiceByIds [delete]
func (idcIpServiceApi *IdcIpServiceApi) DeleteIdcIpServiceByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpServiceService.DeleteIdcIpServiceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpService 更新数据中心业务网段
// @Tags IdcIpService
// @Summary 更新数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService true "更新数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpService/updateIdcIpService [put]
func (idcIpServiceApi *IdcIpServiceApi) UpdateIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindJSON(&idcIpService)
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
    if err := utils.Verify(idcIpService, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcIpServiceService.UpdateIdcIpService(idcIpService); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpService 用id查询数据中心业务网段
// @Tags IdcIpService
// @Summary 用id查询数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpService true "用id查询数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpService/findIdcIpService [get]
func (idcIpServiceApi *IdcIpServiceApi) FindIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindQuery(&idcIpService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcIpService, err := idcIpServiceService.GetIdcIpService(idcIpService.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcIpService": reidcIpService}, c)
	}
}

// GetIdcIpServiceList 分页获取数据中心业务网段列表
// @Tags IdcIpService
// @Summary 分页获取数据中心业务网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpServiceSearch true "分页获取数据中心业务网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpService/getIdcIpServiceList [get]
func (idcIpServiceApi *IdcIpServiceApi) GetIdcIpServiceList(c *gin.Context) {
	var pageInfo idcReq.IdcIpServiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcIpServiceService.GetIdcIpServiceInfoList(pageInfo); err != nil {
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
