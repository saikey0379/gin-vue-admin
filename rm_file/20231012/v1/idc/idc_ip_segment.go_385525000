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

type IdcIpSegmentApi struct {
}

var idcIpSegmentService = service.ServiceGroupApp.IdcServiceGroup.IdcIpSegmentService


// CreateIdcIpSegment 创建数据中心IP网段
// @Tags IdcIpSegment
// @Summary 创建数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSegment true "创建数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpSegment/createIdcIpSegment [post]
func (idcIpSegmentApi *IdcIpSegmentApi) CreateIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindJSON(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "IdcId":{utils.NotEmpty()},
        "IpSegmentType":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcIpSegment, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcIpSegmentService.CreateIdcIpSegment(&idcIpSegment); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpSegment 删除数据中心IP网段
// @Tags IdcIpSegment
// @Summary 删除数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSegment true "删除数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSegment/deleteIdcIpSegment [delete]
func (idcIpSegmentApi *IdcIpSegmentApi) DeleteIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindJSON(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSegmentService.DeleteIdcIpSegment(idcIpSegment); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpSegmentByIds 批量删除数据中心IP网段
// @Tags IdcIpSegment
// @Summary 批量删除数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpSegment/deleteIdcIpSegmentByIds [delete]
func (idcIpSegmentApi *IdcIpSegmentApi) DeleteIdcIpSegmentByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSegmentService.DeleteIdcIpSegmentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpSegment 更新数据中心IP网段
// @Tags IdcIpSegment
// @Summary 更新数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSegment true "更新数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpSegment/updateIdcIpSegment [put]
func (idcIpSegmentApi *IdcIpSegmentApi) UpdateIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindJSON(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "IdcId":{utils.NotEmpty()},
          "IpSegmentType":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcIpSegment, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcIpSegmentService.UpdateIdcIpSegment(idcIpSegment); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpSegment 用id查询数据中心IP网段
// @Tags IdcIpSegment
// @Summary 用id查询数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpSegment true "用id查询数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpSegment/findIdcIpSegment [get]
func (idcIpSegmentApi *IdcIpSegmentApi) FindIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindQuery(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcIpSegment, err := idcIpSegmentService.GetIdcIpSegment(idcIpSegment.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcIpSegment": reidcIpSegment}, c)
	}
}

// GetIdcIpSegmentList 分页获取数据中心IP网段列表
// @Tags IdcIpSegment
// @Summary 分页获取数据中心IP网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpSegmentSearch true "分页获取数据中心IP网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpSegment/getIdcIpSegmentList [get]
func (idcIpSegmentApi *IdcIpSegmentApi) GetIdcIpSegmentList(c *gin.Context) {
	var pageInfo idcReq.IdcIpSegmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcIpSegmentService.GetIdcIpSegmentInfoList(pageInfo); err != nil {
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
