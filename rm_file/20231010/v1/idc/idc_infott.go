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

type IdcInfottApi struct {
}

var idcInfottService = service.ServiceGroupApp.IdcServiceGroup.IdcInfottService


// CreateIdcInfott 创建数据中心
// @Tags IdcInfott
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfott true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfott/createIdcInfott [post]
func (idcInfottApi *IdcInfottApi) CreateIdcInfott(c *gin.Context) {
	var idcInfott idc.IdcInfott
	err := c.ShouldBindJSON(&idcInfott)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcInfott, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcInfottService.CreateIdcInfott(&idcInfott); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcInfott 删除数据中心
// @Tags IdcInfott
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfott true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfott/deleteIdcInfott [delete]
func (idcInfottApi *IdcInfottApi) DeleteIdcInfott(c *gin.Context) {
	var idcInfott idc.IdcInfott
	err := c.ShouldBindJSON(&idcInfott)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfottService.DeleteIdcInfott(idcInfott); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcInfottByIds 批量删除数据中心
// @Tags IdcInfott
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcInfott/deleteIdcInfottByIds [delete]
func (idcInfottApi *IdcInfottApi) DeleteIdcInfottByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfottService.DeleteIdcInfottByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcInfott 更新数据中心
// @Tags IdcInfott
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfott true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfott/updateIdcInfott [put]
func (idcInfottApi *IdcInfottApi) UpdateIdcInfott(c *gin.Context) {
	var idcInfott idc.IdcInfott
	err := c.ShouldBindJSON(&idcInfott)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcInfott, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcInfottService.UpdateIdcInfott(idcInfott); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcInfott 用id查询数据中心
// @Tags IdcInfott
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcInfott true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfott/findIdcInfott [get]
func (idcInfottApi *IdcInfottApi) FindIdcInfott(c *gin.Context) {
	var idcInfott idc.IdcInfott
	err := c.ShouldBindQuery(&idcInfott)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcInfott, err := idcInfottService.GetIdcInfott(idcInfott.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcInfott": reidcInfott}, c)
	}
}

// GetIdcInfottList 分页获取数据中心列表
// @Tags IdcInfott
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcInfottSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfott/getIdcInfottList [get]
func (idcInfottApi *IdcInfottApi) GetIdcInfottList(c *gin.Context) {
	var pageInfo idcReq.IdcInfottSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcInfottService.GetIdcInfottInfoList(pageInfo); err != nil {
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
