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

type OsInstallLogApi struct {
}

var osInstallLogService = service.ServiceGroupApp.OsInstallServiceGroup.OsInstallLogService

// CreateOsInstallLog 创建操作系统安装日志
// @Tags OsInstallLog
// @Summary 创建操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallLog true "创建操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /osInstallLog/createOsInstallLog [post]
func (osInstallLogApi *OsInstallLogApi) CreateOsInstallLog(c *gin.Context) {
	var osInstallLog osInstall.OsInstallLog
	err := c.ShouldBindJSON(&osInstallLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":      {utils.NotEmpty()},
		"QueueId": {utils.NotEmpty()},
		"Log":     {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallLog, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallLogService.CreateOsInstallLog(&osInstallLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOsInstallLog 删除操作系统安装日志
// @Tags OsInstallLog
// @Summary 删除操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallLog true "删除操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallLog/deleteOsInstallLog [delete]
func (osInstallLogApi *OsInstallLogApi) DeleteOsInstallLog(c *gin.Context) {
	var osInstallLog osInstall.OsInstallLog
	err := c.ShouldBindJSON(&osInstallLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallLogService.DeleteOsInstallLog(osInstallLog); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOsInstallLogByIds 批量删除操作系统安装日志
// @Tags OsInstallLog
// @Summary 批量删除操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /osInstallLog/deleteOsInstallLogByIds [delete]
func (osInstallLogApi *OsInstallLogApi) DeleteOsInstallLogByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallLogService.DeleteOsInstallLogByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOsInstallLog 更新操作系统安装日志
// @Tags OsInstallLog
// @Summary 更新操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallLog true "更新操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /osInstallLog/updateOsInstallLog [put]
func (osInstallLogApi *OsInstallLogApi) UpdateOsInstallLog(c *gin.Context) {
	var osInstallLog osInstall.OsInstallLog
	err := c.ShouldBindJSON(&osInstallLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":      {utils.NotEmpty()},
		"QueueId": {utils.NotEmpty()},
		"Log":     {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallLog, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallLogService.UpdateOsInstallLog(osInstallLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOsInstallLog 用id查询操作系统安装日志
// @Tags OsInstallLog
// @Summary 用id查询操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstall.OsInstallLog true "用id查询操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /osInstallLog/findOsInstallLog [get]
func (osInstallLogApi *OsInstallLogApi) FindOsInstallLog(c *gin.Context) {
	var osInstallLog osInstall.OsInstallLog
	err := c.ShouldBindQuery(&osInstallLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reosInstallLog, err := osInstallLogService.GetOsInstallLog(osInstallLog.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reosInstallLog": reosInstallLog}, c)
	}
}

// GetOsInstallLogList 分页获取操作系统安装日志列表
// @Tags OsInstallLog
// @Summary 分页获取操作系统安装日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstallReq.OsInstallLogSearch true "分页获取操作系统安装日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /osInstallLog/getOsInstallLogList [get]
func (osInstallLogApi *OsInstallLogApi) GetOsInstallLogList(c *gin.Context) {
	var pageInfo osInstallReq.OsInstallLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := osInstallLogService.GetOsInstallLogInfoList(pageInfo); err != nil {
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

func (osInstallLogApi *OsInstallLogApi) Report(c *gin.Context) {
	var osInstallLog osInstall.OsInstallLog
	err := c.ShouldBindJSON(&osInstallLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":  {utils.NotEmpty()},
		"Log": {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallLog, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reosInstallingQueue, err := osInstallQueueService.GetOsInstallQueueBySn(osInstallLog.Sn)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	*osInstallLog.QueueId = int(reosInstallingQueue.ID)
	if err := osInstallLogService.CreateOsInstallLog(&osInstallLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
