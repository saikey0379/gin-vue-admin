package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OsInstallConfigPxeApi struct {
}

var osInstallConfigPxeService = service.ServiceGroupApp.OsInstallServiceGroup.OsInstallConfigPxeService


// CreateOsInstallConfigPxe 创建PXE配置
// @Tags OsInstallConfigPxe
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallConfigPxe true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /osInstallConfigPxe/createOsInstallConfigPxe [post]
func (osInstallConfigPxeApi *OsInstallConfigPxeApi) CreateOsInstallConfigPxe(c *gin.Context) {
	var osInstallConfigPxe osInstall.OsInstallConfigPxe
	err := c.ShouldBindJSON(&osInstallConfigPxe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Content":{utils.NotEmpty()},
    }
	if err := utils.Verify(osInstallConfigPxe, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := osInstallConfigPxeService.CreateOsInstallConfigPxe(&osInstallConfigPxe); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOsInstallConfigPxe 删除PXE配置
// @Tags OsInstallConfigPxe
// @Summary 删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallConfigPxe true "删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallConfigPxe/deleteOsInstallConfigPxe [delete]
func (osInstallConfigPxeApi *OsInstallConfigPxeApi) DeleteOsInstallConfigPxe(c *gin.Context) {
	var osInstallConfigPxe osInstall.OsInstallConfigPxe
	err := c.ShouldBindJSON(&osInstallConfigPxe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallConfigPxeService.DeleteOsInstallConfigPxe(osInstallConfigPxe); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOsInstallConfigPxeByIds 批量删除PXE配置
// @Tags OsInstallConfigPxe
// @Summary 批量删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /osInstallConfigPxe/deleteOsInstallConfigPxeByIds [delete]
func (osInstallConfigPxeApi *OsInstallConfigPxeApi) DeleteOsInstallConfigPxeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallConfigPxeService.DeleteOsInstallConfigPxeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOsInstallConfigPxe 更新PXE配置
// @Tags OsInstallConfigPxe
// @Summary 更新PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallConfigPxe true "更新PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /osInstallConfigPxe/updateOsInstallConfigPxe [put]
func (osInstallConfigPxeApi *OsInstallConfigPxeApi) UpdateOsInstallConfigPxe(c *gin.Context) {
	var osInstallConfigPxe osInstall.OsInstallConfigPxe
	err := c.ShouldBindJSON(&osInstallConfigPxe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Content":{utils.NotEmpty()},
      }
    if err := utils.Verify(osInstallConfigPxe, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := osInstallConfigPxeService.UpdateOsInstallConfigPxe(osInstallConfigPxe); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOsInstallConfigPxe 用id查询PXE配置
// @Tags OsInstallConfigPxe
// @Summary 用id查询PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstall.OsInstallConfigPxe true "用id查询PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /osInstallConfigPxe/findOsInstallConfigPxe [get]
func (osInstallConfigPxeApi *OsInstallConfigPxeApi) FindOsInstallConfigPxe(c *gin.Context) {
	var osInstallConfigPxe osInstall.OsInstallConfigPxe
	err := c.ShouldBindQuery(&osInstallConfigPxe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reosInstallConfigPxe, err := osInstallConfigPxeService.GetOsInstallConfigPxe(osInstallConfigPxe.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reosInstallConfigPxe": reosInstallConfigPxe}, c)
	}
}

// GetOsInstallConfigPxeList 分页获取PXE配置列表
// @Tags OsInstallConfigPxe
// @Summary 分页获取PXE配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstallReq.OsInstallConfigPxeSearch true "分页获取PXE配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /osInstallConfigPxe/getOsInstallConfigPxeList [get]
func (osInstallConfigPxeApi *OsInstallConfigPxeApi) GetOsInstallConfigPxeList(c *gin.Context) {
	var pageInfo osInstallReq.OsInstallConfigPxeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := osInstallConfigPxeService.GetOsInstallConfigPxeInfoList(pageInfo); err != nil {
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
