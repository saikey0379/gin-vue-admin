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

type IdcIpManageApi struct {
}

var idcIpManageService = service.ServiceGroupApp.IdcServiceGroup.IdcIpManageService


// CreateIdcIpManage 创建数据中心管理网段
// @Tags IdcIpManage
// @Summary 创建数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpManage true "创建数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpManage/createIdcIpManage [post]
func (idcIpManageApi *IdcIpManageApi) CreateIdcIpManage(c *gin.Context) {
	var idcIpManage idc.IdcIpManage
	err := c.ShouldBindJSON(&idcIpManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "Gateway":{utils.NotEmpty()},
        "IdcId":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcIpManage, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcIpManageService.CreateIdcIpManage(&idcIpManage); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpManage 删除数据中心管理网段
// @Tags IdcIpManage
// @Summary 删除数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpManage true "删除数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpManage/deleteIdcIpManage [delete]
func (idcIpManageApi *IdcIpManageApi) DeleteIdcIpManage(c *gin.Context) {
	var idcIpManage idc.IdcIpManage
	err := c.ShouldBindJSON(&idcIpManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpManageService.DeleteIdcIpManage(idcIpManage); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpManageByIds 批量删除数据中心管理网段
// @Tags IdcIpManage
// @Summary 批量删除数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpManage/deleteIdcIpManageByIds [delete]
func (idcIpManageApi *IdcIpManageApi) DeleteIdcIpManageByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpManageService.DeleteIdcIpManageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpManage 更新数据中心管理网段
// @Tags IdcIpManage
// @Summary 更新数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpManage true "更新数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpManage/updateIdcIpManage [put]
func (idcIpManageApi *IdcIpManageApi) UpdateIdcIpManage(c *gin.Context) {
	var idcIpManage idc.IdcIpManage
	err := c.ShouldBindJSON(&idcIpManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "Gateway":{utils.NotEmpty()},
          "IdcId":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcIpManage, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcIpManageService.UpdateIdcIpManage(idcIpManage); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpManage 用id查询数据中心管理网段
// @Tags IdcIpManage
// @Summary 用id查询数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpManage true "用id查询数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpManage/findIdcIpManage [get]
func (idcIpManageApi *IdcIpManageApi) FindIdcIpManage(c *gin.Context) {
	var idcIpManage idc.IdcIpManage
	err := c.ShouldBindQuery(&idcIpManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcIpManage, err := idcIpManageService.GetIdcIpManage(idcIpManage.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcIpManage": reidcIpManage}, c)
	}
}

// GetIdcIpManageList 分页获取数据中心管理网段列表
// @Tags IdcIpManage
// @Summary 分页获取数据中心管理网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpManageSearch true "分页获取数据中心管理网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpManage/getIdcIpManageList [get]
func (idcIpManageApi *IdcIpManageApi) GetIdcIpManageList(c *gin.Context) {
	var pageInfo idcReq.IdcIpManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcIpManageService.GetIdcIpManageInfoList(pageInfo); err != nil {
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
