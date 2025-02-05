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

type ServerDiscoveryApi struct {
}

var serverDiscoveryService = service.ServiceGroupApp.ServerServiceGroup.ServerDiscoveryService

// CreateServerDiscovery 创建ServerDiscovery
// @Tags ServerDiscovery
// @Summary 创建ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body server.ServerDiscovery true "创建ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serverDiscovery/createServerDiscovery [post]
func (serverDiscoveryApi *ServerDiscoveryApi) CreateServerDiscovery(c *gin.Context) {
	var serverDiscovery server.ServerDiscovery
	err := c.ShouldBindJSON(&serverDiscovery)
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
	if err := utils.Verify(serverDiscovery, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverDiscoveryService.CreateServerDiscovery(&serverDiscovery); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteServerDiscovery 删除ServerDiscovery
// @Tags ServerDiscovery
// @Summary 删除ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body server.ServerDiscovery true "删除ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverDiscovery/deleteServerDiscovery [delete]
func (serverDiscoveryApi *ServerDiscoveryApi) DeleteServerDiscovery(c *gin.Context) {
	var serverDiscovery server.ServerDiscovery
	err := c.ShouldBindJSON(&serverDiscovery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverDiscoveryService.DeleteServerDiscovery(serverDiscovery); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServerDiscoveryByIds 批量删除ServerDiscovery
// @Tags ServerDiscovery
// @Summary 批量删除ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /serverDiscovery/deleteServerDiscoveryByIds [delete]
func (serverDiscoveryApi *ServerDiscoveryApi) DeleteServerDiscoveryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverDiscoveryService.DeleteServerDiscoveryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateServerDiscovery 更新ServerDiscovery
// @Tags ServerDiscovery
// @Summary 更新ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body server.ServerDiscovery true "更新ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serverDiscovery/updateServerDiscovery [put]
func (serverDiscoveryApi *ServerDiscoveryApi) UpdateServerDiscovery(c *gin.Context) {
	var serverDiscovery server.ServerDiscovery
	err := c.ShouldBindJSON(&serverDiscovery)
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
	if err := utils.Verify(serverDiscovery, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverDiscoveryService.UpdateServerDiscovery(serverDiscovery); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindServerDiscovery 用id查询ServerDiscovery
// @Tags ServerDiscovery
// @Summary 用id查询ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query server.ServerDiscovery true "用id查询ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serverDiscovery/findServerDiscovery [get]
func (serverDiscoveryApi *ServerDiscoveryApi) FindServerDiscovery(c *gin.Context) {
	var serverDiscovery server.ServerDiscovery
	err := c.ShouldBindQuery(&serverDiscovery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reserverDiscovery, err := serverDiscoveryService.GetServerDiscovery(serverDiscovery.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserverDiscovery": reserverDiscovery}, c)
	}
}

// GetServerDiscoveryList 分页获取ServerDiscovery列表
// @Tags ServerDiscovery
// @Summary 分页获取ServerDiscovery列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query serverReq.ServerDiscoverySearch true "分页获取ServerDiscovery列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverDiscovery/getServerDiscoveryList [get]
func (serverDiscoveryApi *ServerDiscoveryApi) GetServerDiscoveryList(c *gin.Context) {
	var pageInfo serverReq.ServerDiscoverySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serverDiscoveryService.GetServerDiscoveryInfoList(pageInfo); err != nil {
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

func (serverDiscoveryApi *ServerDiscoveryApi) EntryServerDiscoveryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for _, i := range IDS.Ids {
		if reserverDiscovery, err := serverDiscoveryService.GetServerDiscovery(uint(i)); err != nil {
			global.GVA_LOG.Error("录入失败!", zap.Error(err))
			response.FailWithMessage("录入失败", c)
			return
		} else {
			var reserverInfo = &server.ServerInfo{
				Sn:            reserverDiscovery.Sn,
				Hostname:      reserverDiscovery.Hostname,
				Ip:            reserverDiscovery.Ip,
				VersionOs:     reserverDiscovery.VersionOs,
				VersionKernel: reserverDiscovery.VersionKernel,
				Manufacturer:  reserverDiscovery.Manufacturer,
				ModelName:     reserverDiscovery.ModelName,
				CpuSum:        reserverDiscovery.CpuSum,
				Cpu:           reserverDiscovery.Cpu,
				MemorySum:     reserverDiscovery.MemorySum,
				Memory:        reserverDiscovery.Memory,
				NicInfo:       reserverDiscovery.NicInfo,
				NicDevice:     reserverDiscovery.NicDevice,
				Gpu:           reserverDiscovery.Gpu,
				Motherboard:   reserverDiscovery.Motherboard,
				Raid:          reserverDiscovery.Raid,
				Disk:          reserverDiscovery.Disk,
				ServerType:    reserverDiscovery.ServerType,
				VersionAgent:  reserverDiscovery.VersionAgent,
			}
			if err := serverInfoService.CreateServerInfo(reserverInfo); err != nil {
				global.GVA_LOG.Error("录入失败!", zap.Error(err))
				response.FailWithMessage("录入失败", c)
				return
			}

		}
	}

	if err := serverDiscoveryService.DeleteServerDiscoveryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量录入失败", c)
		return
	}

	response.OkWithMessage("批量录入成功", c)

}
