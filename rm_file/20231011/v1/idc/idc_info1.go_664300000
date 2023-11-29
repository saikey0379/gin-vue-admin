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

type IdcInfo1Api struct {
}

var idcInfo1Service = service.ServiceGroupApp.IdcServiceGroup.IdcInfo1Service


// CreateIdcInfo1 创建数据中心
// @Tags IdcInfo1
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfo1 true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfo1/createIdcInfo1 [post]
func (idcInfo1Api *IdcInfo1Api) CreateIdcInfo1(c *gin.Context) {
	var idcInfo1 idc.IdcInfo1
	err := c.ShouldBindJSON(&idcInfo1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcInfo1, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcInfo1Service.CreateIdcInfo1(&idcInfo1); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcInfo1 删除数据中心
// @Tags IdcInfo1
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfo1 true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfo1/deleteIdcInfo1 [delete]
func (idcInfo1Api *IdcInfo1Api) DeleteIdcInfo1(c *gin.Context) {
	var idcInfo1 idc.IdcInfo1
	err := c.ShouldBindJSON(&idcInfo1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfo1Service.DeleteIdcInfo1(idcInfo1); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcInfo1ByIds 批量删除数据中心
// @Tags IdcInfo1
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcInfo1/deleteIdcInfo1ByIds [delete]
func (idcInfo1Api *IdcInfo1Api) DeleteIdcInfo1ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfo1Service.DeleteIdcInfo1ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcInfo1 更新数据中心
// @Tags IdcInfo1
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfo1 true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfo1/updateIdcInfo1 [put]
func (idcInfo1Api *IdcInfo1Api) UpdateIdcInfo1(c *gin.Context) {
	var idcInfo1 idc.IdcInfo1
	err := c.ShouldBindJSON(&idcInfo1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcInfo1, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcInfo1Service.UpdateIdcInfo1(idcInfo1); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcInfo1 用id查询数据中心
// @Tags IdcInfo1
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcInfo1 true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfo1/findIdcInfo1 [get]
func (idcInfo1Api *IdcInfo1Api) FindIdcInfo1(c *gin.Context) {
	var idcInfo1 idc.IdcInfo1
	err := c.ShouldBindQuery(&idcInfo1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcInfo1, err := idcInfo1Service.GetIdcInfo1(idcInfo1.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcInfo1": reidcInfo1}, c)
	}
}

// GetIdcInfo1List 分页获取数据中心列表
// @Tags IdcInfo1
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcInfo1Search true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfo1/getIdcInfo1List [get]
func (idcInfo1Api *IdcInfo1Api) GetIdcInfo1List(c *gin.Context) {
	var pageInfo idcReq.IdcInfo1Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcInfo1Service.GetIdcInfo1InfoList(pageInfo); err != nil {
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
