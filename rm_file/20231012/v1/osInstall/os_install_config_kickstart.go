package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
	osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OsInstallConfigKickstartApi struct {
}

var osInstallConfigKickstartService = service.ServiceGroupApp.OsInstallServiceGroup.OsInstallConfigKickstartService

type ViewOsInstallConfigKickstart struct {
	osInstall.OsInstallConfigKickstart
	PxeName string `json:"pxe_name" form:"pxe_name"`
}

// CreateOsInstallConfigKickstart 创建Kickstart配置
// @Tags OsInstallConfigKickstart
// @Summary 创建Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallConfigKickstart true "创建Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /osInstallConfigKickstart/createOsInstallConfigKickstart [post]
func (osInstallConfigKickstartApi *OsInstallConfigKickstartApi) CreateOsInstallConfigKickstart(c *gin.Context) {
	var osInstallConfigKickstart osInstall.OsInstallConfigKickstart
	err := c.ShouldBindJSON(&osInstallConfigKickstart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":    {utils.NotEmpty()},
		"PxeId":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallConfigKickstart, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallConfigKickstartService.CreateOsInstallConfigKickstart(&osInstallConfigKickstart); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOsInstallConfigKickstart 删除Kickstart配置
// @Tags OsInstallConfigKickstart
// @Summary 删除Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallConfigKickstart true "删除Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallConfigKickstart/deleteOsInstallConfigKickstart [delete]
func (osInstallConfigKickstartApi *OsInstallConfigKickstartApi) DeleteOsInstallConfigKickstart(c *gin.Context) {
	var osInstallConfigKickstart osInstall.OsInstallConfigKickstart
	err := c.ShouldBindJSON(&osInstallConfigKickstart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallConfigKickstartService.DeleteOsInstallConfigKickstart(osInstallConfigKickstart); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOsInstallConfigKickstartByIds 批量删除Kickstart配置
// @Tags OsInstallConfigKickstart
// @Summary 批量删除Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /osInstallConfigKickstart/deleteOsInstallConfigKickstartByIds [delete]
func (osInstallConfigKickstartApi *OsInstallConfigKickstartApi) DeleteOsInstallConfigKickstartByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallConfigKickstartService.DeleteOsInstallConfigKickstartByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOsInstallConfigKickstart 更新Kickstart配置
// @Tags OsInstallConfigKickstart
// @Summary 更新Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallConfigKickstart true "更新Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /osInstallConfigKickstart/updateOsInstallConfigKickstart [put]
func (osInstallConfigKickstartApi *OsInstallConfigKickstartApi) UpdateOsInstallConfigKickstart(c *gin.Context) {
	var osInstallConfigKickstart osInstall.OsInstallConfigKickstart
	err := c.ShouldBindJSON(&osInstallConfigKickstart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":    {utils.NotEmpty()},
		"PxeId":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallConfigKickstart, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallConfigKickstartService.UpdateOsInstallConfigKickstart(osInstallConfigKickstart); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOsInstallConfigKickstart 用id查询Kickstart配置
// @Tags OsInstallConfigKickstart
// @Summary 用id查询Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstall.OsInstallConfigKickstart true "用id查询Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /osInstallConfigKickstart/findOsInstallConfigKickstart [get]
func (osInstallConfigKickstartApi *OsInstallConfigKickstartApi) FindOsInstallConfigKickstart(c *gin.Context) {
	var osInstallConfigKickstart osInstall.OsInstallConfigKickstart
	err := c.ShouldBindQuery(&osInstallConfigKickstart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reosInstallConfigKickstart, err := osInstallConfigKickstartService.GetOsInstallConfigKickstart(osInstallConfigKickstart.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	reosInstallConfigPxe, err := osInstallConfigPxeService.GetOsInstallConfigPxe(uint(*reosInstallConfigKickstart.PxeId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	var viewOsInstallConfigKickstart = ViewOsInstallConfigKickstart{
		reosInstallConfigKickstart,
		reosInstallConfigPxe.Name,
	}
	response.OkWithData(gin.H{"reosInstallConfigKickstart": viewOsInstallConfigKickstart}, c)

}

// GetOsInstallConfigKickstartList 分页获取Kickstart配置列表
// @Tags OsInstallConfigKickstart
// @Summary 分页获取Kickstart配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstallReq.OsInstallConfigKickstartSearch true "分页获取Kickstart配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /osInstallConfigKickstart/getOsInstallConfigKickstartList [get]
func (osInstallConfigKickstartApi *OsInstallConfigKickstartApi) GetOsInstallConfigKickstartList(c *gin.Context) {
	var pageInfo osInstallReq.OsInstallConfigKickstartSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := osInstallConfigKickstartService.GetOsInstallConfigKickstartInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var mapPxe = make(map[int]string)
	var viewOsInstallConfigKickstarts []ViewOsInstallConfigKickstart
	for _, osInstallConfigKickstart := range list {
		if mapPxe[*osInstallConfigKickstart.PxeId] == "" {
			reosInstallConfigPxe, err := osInstallConfigPxeService.GetOsInstallConfigPxe(uint(*osInstallConfigKickstart.PxeId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
			}
			mapPxe[*osInstallConfigKickstart.PxeId] = reosInstallConfigPxe.Name
		}
		var viewOsInstallConfigKickstart = ViewOsInstallConfigKickstart{
			osInstallConfigKickstart,
			mapPxe[*osInstallConfigKickstart.PxeId],
		}
		viewOsInstallConfigKickstarts = append(viewOsInstallConfigKickstarts, viewOsInstallConfigKickstart)
	}
	response.OkWithDetailed(response.PageResult{
		List:     viewOsInstallConfigKickstarts,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
