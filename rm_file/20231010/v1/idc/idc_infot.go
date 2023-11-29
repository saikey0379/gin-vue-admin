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

type IdcInfotApi struct {
}

var idcInfotService = service.ServiceGroupApp.IdcServiceGroup.IdcInfotService


// CreateIdcInfot 创建数据中心
// @Tags IdcInfot
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfot true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfot/createIdcInfot [post]
func (idcInfotApi *IdcInfotApi) CreateIdcInfot(c *gin.Context) {
	var idcInfot idc.IdcInfot
	err := c.ShouldBindJSON(&idcInfot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcInfot, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcInfotService.CreateIdcInfot(&idcInfot); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcInfot 删除数据中心
// @Tags IdcInfot
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfot true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfot/deleteIdcInfot [delete]
func (idcInfotApi *IdcInfotApi) DeleteIdcInfot(c *gin.Context) {
	var idcInfot idc.IdcInfot
	err := c.ShouldBindJSON(&idcInfot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfotService.DeleteIdcInfot(idcInfot); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcInfotByIds 批量删除数据中心
// @Tags IdcInfot
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcInfot/deleteIdcInfotByIds [delete]
func (idcInfotApi *IdcInfotApi) DeleteIdcInfotByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfotService.DeleteIdcInfotByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcInfot 更新数据中心
// @Tags IdcInfot
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfot true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfot/updateIdcInfot [put]
func (idcInfotApi *IdcInfotApi) UpdateIdcInfot(c *gin.Context) {
	var idcInfot idc.IdcInfot
	err := c.ShouldBindJSON(&idcInfot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcInfot, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcInfotService.UpdateIdcInfot(idcInfot); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcInfot 用id查询数据中心
// @Tags IdcInfot
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcInfot true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfot/findIdcInfot [get]
func (idcInfotApi *IdcInfotApi) FindIdcInfot(c *gin.Context) {
	var idcInfot idc.IdcInfot
	err := c.ShouldBindQuery(&idcInfot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcInfot, err := idcInfotService.GetIdcInfot(idcInfot.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcInfot": reidcInfot}, c)
	}
}

// GetIdcInfotList 分页获取数据中心列表
// @Tags IdcInfot
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcInfotSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfot/getIdcInfotList [get]
func (idcInfotApi *IdcInfotApi) GetIdcInfotList(c *gin.Context) {
	var pageInfo idcReq.IdcInfotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcInfotService.GetIdcInfotInfoList(pageInfo); err != nil {
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
