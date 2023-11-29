package osInstall

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/known"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
	osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/server"
	request1 "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type OsInstallQueueApi struct {
}

var osInstallQueueService = service.ServiceGroupApp.OsInstallServiceGroup.OsInstallQueueService

// CreateOsInstallQueue 创建操作系统安装队列
// @Tags OsInstallQueue
// @Summary 创建操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallQueue true "创建操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /osInstallQueue/createOsInstallQueue [post]
func (osInstallQueueApi *OsInstallQueueApi) CreateOsInstallQueue(c *gin.Context) {
	var osInstallQueue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&osInstallQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":          {utils.NotEmpty()},
		"Hostname":    {utils.NotEmpty()},
		"Ip":          {utils.NotEmpty()},
		"NetworkId":   {utils.NotEmpty()},
		"HardwareId":  {utils.NotEmpty()},
		"PxeId":       {utils.NotEmpty()},
		"KickstartId": {utils.NotEmpty()},
		"Status":      {utils.NotEmpty()},
		"Manager":     {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallQueue, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallQueueService.CreateOsInstallQueue(&osInstallQueue); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOsInstallQueue 删除操作系统安装队列
// @Tags OsInstallQueue
// @Summary 删除操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallQueue true "删除操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallQueue/deleteOsInstallQueue [delete]
func (osInstallQueueApi *OsInstallQueueApi) DeleteOsInstallQueue(c *gin.Context) {
	var osInstallQueue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&osInstallQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallQueueService.DeleteOsInstallQueue(osInstallQueue); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOsInstallQueueByIds 批量删除操作系统安装队列
// @Tags OsInstallQueue
// @Summary 批量删除操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /osInstallQueue/deleteOsInstallQueueByIds [delete]
func (osInstallQueueApi *OsInstallQueueApi) DeleteOsInstallQueueByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallQueueService.DeleteOsInstallQueueByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOsInstallQueue 更新操作系统安装队列
// @Tags OsInstallQueue
// @Summary 更新操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallQueue true "更新操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /osInstallQueue/updateOsInstallQueue [put]
func (osInstallQueueApi *OsInstallQueueApi) UpdateOsInstallQueue(c *gin.Context) {
	var osInstallQueue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&osInstallQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":          {utils.NotEmpty()},
		"Hostname":    {utils.NotEmpty()},
		"Ip":          {utils.NotEmpty()},
		"NetworkId":   {utils.NotEmpty()},
		"HardwareId":  {utils.NotEmpty()},
		"PxeId":       {utils.NotEmpty()},
		"KickstartId": {utils.NotEmpty()},
		"Status":      {utils.NotEmpty()},
		"Manager":     {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallQueue, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osInstallQueueService.UpdateOsInstallQueue(osInstallQueue); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOsInstallQueue 用id查询操作系统安装队列
// @Tags OsInstallQueue
// @Summary 用id查询操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstall.OsInstallQueue true "用id查询操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /osInstallQueue/findOsInstallQueue [get]
func (osInstallQueueApi *OsInstallQueueApi) FindOsInstallQueue(c *gin.Context) {
	var osInstallQueue osInstall.OsInstallQueue
	err := c.ShouldBindQuery(&osInstallQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reosInstallQueue, err := osInstallQueueService.GetOsInstallQueue(osInstallQueue.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reosInstallQueue": reosInstallQueue}, c)
	}
}

// GetOsInstallQueueList 分页获取操作系统安装队列列表
// @Tags OsInstallQueue
// @Summary 分页获取操作系统安装队列列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstallReq.OsInstallQueueSearch true "分页获取操作系统安装队列列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /osInstallQueue/getOsInstallQueueList [get]
func (osInstallQueueApi *OsInstallQueueApi) GetOsInstallQueueList(c *gin.Context) {
	var pageInfo osInstallReq.OsInstallQueueSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := osInstallQueueService.GetOsInstallQueueInfoList(pageInfo); err != nil {
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

func (osInstallQueueApi *OsInstallQueueApi) OsInstallStart(c *gin.Context) {
	var osInstallQueue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&osInstallQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	osInstallQueue.ID = common.GenSnowFlakeID()
	*osInstallQueue.Status = 6

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	osInstallQueue.Manager = claimsReal.Username

	verify := utils.Rules{
		"Sn":          {utils.NotEmpty()},
		"Hostname":    {utils.NotEmpty()},
		"Ip":          {utils.NotEmpty()},
		"NetworkId":   {utils.NotEmpty()},
		"HardwareId":  {utils.NotEmpty()},
		"KickstartId": {utils.NotEmpty()},
	}
	if err := utils.Verify(osInstallQueue, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = osInstallQueueService.CreateOsInstallQueue(&osInstallQueue)
	if err != nil {
		global.GVA_LOG.Error("任务创建失败!", zap.Error(err))
		response.FailWithMessage("任务创建失败", c)
		return
	}

	kickstartConfig, err := osInstallConfigKickstartService.GetOsInstallConfigKickstart(uint(*osInstallQueue.KickstartId))
	if err != nil {
		global.GVA_LOG.Error("系统配置查询失败!", zap.Error(err))
		response.FailWithMessage("系统配置查询失败", c)
		return
	}

	pxeConfig, err := osInstallConfigPxeService.GetOsInstallConfigPxe(uint(*kickstartConfig.PxeId))
	if err != nil {
		global.GVA_LOG.Error("PXE配置查询失败!", zap.Error(err))
		response.FailWithMessage("PXE配置查询失败", c)
		return
	}

	var osInstallSystemService = service.ServiceGroupApp.OsInstallServiceGroup.OsInstallSystemService
	osInstallSystemService.DirPxe = global.GVA_CONFIG.OsInstall.DirPxe
	if osInstallSystemService.DirPxe == "" {
		osInstallSystemService.DirPxe = known.DefaultDirPxe
	}
	osInstallSystemService.Content = strings.Replace(pxeConfig.Content, "{sn}", osInstallQueue.Sn, -1)
	osInstallSystemService.Content = strings.Replace(osInstallSystemService.Content, "\r\n", "\n", -1)

	serverServiceGroup := service.ServiceGroupApp.ServerServiceGroup
	serverInfo, err := serverServiceGroup.GetServerInfoBySn(osInstallQueue.Sn)
	if err != nil {
		global.GVA_LOG.Error("Mac信息查询失败!", zap.Error(err))
		response.FailWithMessage("Mac信息查询失败", c)
		return
	}

	var nics []server.NicInfo
	err = json.Unmarshal([]byte(serverInfo.NicInfo), &nics)
	if err != nil {
		global.GVA_LOG.Error("Mac信息解析失败!", zap.Error(err))
		response.FailWithMessage("Mac信息解析失败", c)
		return
	}

	for _, i := range nics {
		mac := strings.ToLower(strings.TrimSpace(i.Mac))
		file := utils.GetPxeFileNameByMac(mac)
		osInstallSystemService.Files = append(osInstallSystemService.Files, file)
	}

	err = osInstallSystemService.CreateFilePxe()
	if err != nil {
		global.GVA_LOG.Error("PXE配置文件创建失败!", zap.Error(err))
		response.FailWithMessage("PXE配置文件创建失败", c)
		return
	}

	err = service.ServiceGroupApp.AgentServiceGroup.AgentOsInstall(osInstallQueue.Ip, osInstallQueue.Sn)
	if err != nil {
		global.GVA_LOG.Error("任务启动失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	*osInstallQueue.Status = 3
	if err := osInstallQueueService.UpdateOsInstallQueue(osInstallQueue); err != nil {
		global.GVA_LOG.Error("任务状态更新失败!", zap.Error(err))
		response.FailWithMessage("任务状态更新失败", c)
	} else {
		response.OkWithMessage(fmt.Sprintf("[%s]开始安装操作系统", osInstallQueue.Hostname), c)
	}
}

func (osInstallQueueApi *OsInstallQueueApi) GetDeviceInfoBySn(c *gin.Context) {
	sn := c.Query("sn")
	if sn == "" {
		response.FailWithMessage("SN不能为空", c)
		return
	}

	var osInstallQueue osInstall.OsInstallQueue
	err := c.ShouldBindQuery(&osInstallQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	device, err := osInstallQueueService.GetOsInstallQueueBySn(sn)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}

	var deviceInfo string
	if device.Hostname != "" {
		deviceInfo += "HOSTNAME=\"" + device.Hostname + "\""
	}
	if device.Ip != "" {
		deviceInfo += "\nIPADDR=\"" + device.Ip + "\""
	}

	ipSubnet, err := service.ServiceGroupApp.IdcServiceGroup.GetIdcIpSubnet(uint(*device.NetworkId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}

	if *ipSubnet.Prefix > 0 {
		deviceInfo += fmt.Sprintf("\nPREFIX=%d", *ipSubnet.Prefix)
	}
	if ipSubnet.Gateway != "" {
		deviceInfo += "\nGATEWAY=\"" + ipSubnet.Gateway + "\""
	}
	c.String(200, deviceInfo)
}
