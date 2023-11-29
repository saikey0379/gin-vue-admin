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

type IdcIpService1Api struct {
}

var idcIpService1Service = service.ServiceGroupApp.IdcServiceGroup.IdcIpService1Service


// CreateIdcIpService1 创建数据中心业务网段
// @Tags IdcIpService1
// @Summary 创建数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService1 true "创建数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpService1/createIdcIpService1 [post]
func (idcIpService1Api *IdcIpService1Api) CreateIdcIpService1(c *gin.Context) {
	var idcIpService1 idc.IdcIpService1
	err := c.ShouldBindJSON(&idcIpService1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "IdcId":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcIpService1, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcIpService1Service.CreateIdcIpService1(&idcIpService1); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpService1 删除数据中心业务网段
// @Tags IdcIpService1
// @Summary 删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService1 true "删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService1/deleteIdcIpService1 [delete]
func (idcIpService1Api *IdcIpService1Api) DeleteIdcIpService1(c *gin.Context) {
	var idcIpService1 idc.IdcIpService1
	err := c.ShouldBindJSON(&idcIpService1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpService1Service.DeleteIdcIpService1(idcIpService1); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpService1ByIds 批量删除数据中心业务网段
// @Tags IdcIpService1
// @Summary 批量删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpService1/deleteIdcIpService1ByIds [delete]
func (idcIpService1Api *IdcIpService1Api) DeleteIdcIpService1ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpService1Service.DeleteIdcIpService1ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpService1 更新数据中心业务网段
// @Tags IdcIpService1
// @Summary 更新数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService1 true "更新数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpService1/updateIdcIpService1 [put]
func (idcIpService1Api *IdcIpService1Api) UpdateIdcIpService1(c *gin.Context) {
	var idcIpService1 idc.IdcIpService1
	err := c.ShouldBindJSON(&idcIpService1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "IdcId":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcIpService1, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcIpService1Service.UpdateIdcIpService1(idcIpService1); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpService1 用id查询数据中心业务网段
// @Tags IdcIpService1
// @Summary 用id查询数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpService1 true "用id查询数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpService1/findIdcIpService1 [get]
func (idcIpService1Api *IdcIpService1Api) FindIdcIpService1(c *gin.Context) {
	var idcIpService1 idc.IdcIpService1
	err := c.ShouldBindQuery(&idcIpService1)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcIpService1, err := idcIpService1Service.GetIdcIpService1(idcIpService1.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcIpService1": reidcIpService1}, c)
	}
}

// GetIdcIpService1List 分页获取数据中心业务网段列表
// @Tags IdcIpService1
// @Summary 分页获取数据中心业务网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpService1Search true "分页获取数据中心业务网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpService1/getIdcIpService1List [get]
func (idcIpService1Api *IdcIpService1Api) GetIdcIpService1List(c *gin.Context) {
	var pageInfo idcReq.IdcIpService1Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcIpService1Service.GetIdcIpService1InfoList(pageInfo); err != nil {
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
