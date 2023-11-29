package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/server"
	serverReq "github.com/flipped-aurora/gin-vue-admin/server/model/server/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServerInfoApi struct {
}

var serverInfoService = service.ServiceGroupApp.ServerServiceGroup.ServerInfoService

// CreateServerInfo 创建ServerInfo
// @Tags ServerInfo
// @Summary 创建ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body server.ServerInfo true "创建ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serverInfo/createServerInfo [post]
func (serverInfoApi *ServerInfoApi) CreateServerInfo(c *gin.Context) {
	var serverInfo server.ServerInfo
	err := c.ShouldBindJSON(&serverInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":            {utils.NotEmpty()},
		"Hostname":      {utils.NotEmpty()},
		"Ip":            {utils.NotEmpty()},
		"VersionOs":     {utils.NotEmpty()},
		"VersionKernel": {utils.NotEmpty()},
		"Manufacturer":  {utils.NotEmpty()},
		"ModelName":     {utils.NotEmpty()},
		"CpuSum":        {utils.NotEmpty()},
		"Cpu":           {utils.NotEmpty()},
		"MemorySum":     {utils.NotEmpty()},
		"Memory":        {utils.NotEmpty()},
		"NicInfo":       {utils.NotEmpty()},
		"NicDevice":     {utils.NotEmpty()},
		"Gpu":           {utils.NotEmpty()},
		"Motherboard":   {utils.NotEmpty()},
		"Raid":          {utils.NotEmpty()},
		"Disk":          {utils.NotEmpty()},
		"ServerType":    {utils.NotEmpty()},
		"VersionAgent":  {utils.NotEmpty()},
	}
	if err := utils.Verify(serverInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverInfoService.CreateServerInfo(&serverInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteServerInfo 删除ServerInfo
// @Tags ServerInfo
// @Summary 删除ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body server.ServerInfo true "删除ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverInfo/deleteServerInfo [delete]
func (serverInfoApi *ServerInfoApi) DeleteServerInfo(c *gin.Context) {
	var serverInfo server.ServerInfo
	err := c.ShouldBindJSON(&serverInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverInfoService.DeleteServerInfo(serverInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServerInfoByIds 批量删除ServerInfo
// @Tags ServerInfo
// @Summary 批量删除ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /serverInfo/deleteServerInfoByIds [delete]
func (serverInfoApi *ServerInfoApi) DeleteServerInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverInfoService.DeleteServerInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateServerInfo 更新ServerInfo
// @Tags ServerInfo
// @Summary 更新ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body server.ServerInfo true "更新ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serverInfo/updateServerInfo [put]
func (serverInfoApi *ServerInfoApi) UpdateServerInfo(c *gin.Context) {
	var serverInfo server.ServerInfo
	err := c.ShouldBindJSON(&serverInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":            {utils.NotEmpty()},
		"Hostname":      {utils.NotEmpty()},
		"Ip":            {utils.NotEmpty()},
		"VersionOs":     {utils.NotEmpty()},
		"VersionKernel": {utils.NotEmpty()},
		"Manufacturer":  {utils.NotEmpty()},
		"ModelName":     {utils.NotEmpty()},
		"CpuSum":        {utils.NotEmpty()},
		"Cpu":           {utils.NotEmpty()},
		"MemorySum":     {utils.NotEmpty()},
		"Memory":        {utils.NotEmpty()},
		"NicInfo":       {utils.NotEmpty()},
		"NicDevice":     {utils.NotEmpty()},
		"Gpu":           {utils.NotEmpty()},
		"Motherboard":   {utils.NotEmpty()},
		"Raid":          {utils.NotEmpty()},
		"Disk":          {utils.NotEmpty()},
		"ServerType":    {utils.NotEmpty()},
		"VersionAgent":  {utils.NotEmpty()},
	}
	if err := utils.Verify(serverInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverInfoService.UpdateServerInfo(serverInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindServerInfo 用id查询ServerInfo
// @Tags ServerInfo
// @Summary 用id查询ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query server.ServerInfo true "用id查询ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serverInfo/findServerInfo [get]
func (serverInfoApi *ServerInfoApi) FindServerInfo(c *gin.Context) {
	var serverInfo server.ServerInfo
	err := c.ShouldBindQuery(&serverInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reserverInfo, err := serverInfoService.GetServerInfo(serverInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserverInfo": reserverInfo}, c)
	}
}

// GetServerInfoList 分页获取ServerInfo列表
// @Tags ServerInfo
// @Summary 分页获取ServerInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query serverReq.ServerInfoSearch true "分页获取ServerInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverInfo/getServerInfoList [get]
func (serverInfoApi *ServerInfoApi) GetServerInfoList(c *gin.Context) {
	var pageInfo serverReq.ServerInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serverInfoService.GetServerInfoInfoList(pageInfo); err != nil {
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

func (serverInfoApi *ServerInfoApi) Report(c *gin.Context) {
	var serverInfo server.ServerInfo
	err := c.ShouldBindJSON(&serverInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":            {utils.NotEmpty()},
		"Hostname":      {utils.NotEmpty()},
		"Ip":            {utils.NotEmpty()},
		"VersionOs":     {utils.NotEmpty()},
		"VersionKernel": {utils.NotEmpty()},
		"Manufacturer":  {utils.NotEmpty()},
		"ModelName":     {utils.NotEmpty()},
		"CpuSum":        {utils.NotEmpty()},
		"Cpu":           {utils.NotEmpty()},
		"MemorySum":     {utils.NotEmpty()},
		"Memory":        {utils.NotEmpty()},
		"NicInfo":       {utils.NotEmpty()},
		"NicDevice":     {utils.NotEmpty()},
		"Disk":          {utils.NotEmpty()},
		"ServerType":    {utils.NotEmpty()},
		"VersionAgent":  {utils.NotEmpty()},
	}
	if err := utils.Verify(serverInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = serverInfoService.GetServerInfoBySn(serverInfo.Sn)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))

		_, err = serverDiscoveryService.GetServerDiscoveryBySn(serverInfo.Sn)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))

			var serverDiscovery = &server.ServerDiscovery{
				Sn:            serverInfo.Sn,
				Hostname:      serverInfo.Hostname,
				Ip:            serverInfo.Ip,
				VersionOs:     serverInfo.VersionOs,
				VersionKernel: serverInfo.VersionKernel,
				Manufacturer:  serverInfo.Manufacturer,
				ModelName:     serverInfo.ModelName,
				CpuSum:        serverInfo.CpuSum,
				Cpu:           serverInfo.Cpu,
				MemorySum:     serverInfo.MemorySum,
				Memory:        serverInfo.Memory,
				NicInfo:       serverInfo.NicInfo,
				NicDevice:     serverInfo.NicDevice,
				Disk:          serverInfo.Disk,
				ServerType:    serverInfo.ServerType,
				VersionAgent:  serverInfo.VersionAgent,
			}
			if err := serverDiscoveryService.CreateServerDiscovery(serverDiscovery); err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
				response.FailWithMessage("创建失败", c)
			} else {
				response.OkWithMessage("创建成功", c)
			}
		} else {
			response.FailWithMessage("已注册!", c)
		}
	} else {
		if err := serverInfoService.UpdateServerInfo(serverInfo); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
		} else {
			response.OkWithMessage("更新成功", c)
		}
	}
}

func (serverInfoApi *ServerInfoApi) GetServerInfoSnListNotExistsInBareMetal(c *gin.Context) {
	var pageInfo serverReq.ServerInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, err := serverInfoService.GetServerInfoSnListNotExistsInBareMetal()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	type Sn struct {
		Sn string
	}
	var snList []Sn
	for _, i := range list {
		snList = append(snList, Sn{
			i,
		})
	}

	response.OkWithData(gin.H{"list": snList}, c)
}
