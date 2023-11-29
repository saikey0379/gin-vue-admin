package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IdcIpServiceApi struct {
}

var idcIpServiceService = service.ServiceGroupApp.IdcServiceGroup.IdcIpServiceService

type ViewIpService struct {
	idc.IdcIpService
	IdcName    string `json:"idc_name" form:"idc_name"`
	IpAllocate int    `json:"ip_allocate" form:"ip_allocate"`
	IpTotal    int    `json:"ip_total" form:"ip_total"`
}

// CreateIdcIpService 创建数据中心业务网段
// @Tags IdcIpService
// @Summary 创建数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService true "创建数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpService/createIdcIpService [post]
func (idcIpServiceApi *IdcIpServiceApi) CreateIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindJSON(&idcIpService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":  {utils.NotEmpty()},
		"Prefix":   {utils.NotEmpty()},
		"IdcId":    {utils.NotEmpty()},
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpService, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	network, err := utils.GetNetworkAddress(idcIpService.Network, *idcIpService.Prefix)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	idcIpService.Network = network.String()
	if err := idcIpServiceService.CreateIdcIpService(&idcIpService); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpService 删除数据中心业务网段
// @Tags IdcIpService
// @Summary 删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService true "删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService/deleteIdcIpService [delete]
func (idcIpServiceApi *IdcIpServiceApi) DeleteIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindJSON(&idcIpService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpServiceService.DeleteIdcIpService(idcIpService); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpServiceByIds 批量删除数据中心业务网段
// @Tags IdcIpService
// @Summary 批量删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpService/deleteIdcIpServiceByIds [delete]
func (idcIpServiceApi *IdcIpServiceApi) DeleteIdcIpServiceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpServiceService.DeleteIdcIpServiceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpService 更新数据中心业务网段
// @Tags IdcIpService
// @Summary 更新数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpService true "更新数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpService/updateIdcIpService [put]
func (idcIpServiceApi *IdcIpServiceApi) UpdateIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindJSON(&idcIpService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":  {utils.NotEmpty()},
		"Prefix":   {utils.NotEmpty()},
		"IdcId":    {utils.NotEmpty()},
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpService, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	network, err := utils.GetNetworkAddress(idcIpService.Network, *idcIpService.Prefix)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	idcIpService.Network = network.String()
	if err := idcIpServiceService.UpdateIdcIpService(idcIpService); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpService 用id查询数据中心业务网段
// @Tags IdcIpService
// @Summary 用id查询数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpService true "用id查询数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpService/findIdcIpService [get]
func (idcIpServiceApi *IdcIpServiceApi) FindIdcIpService(c *gin.Context) {
	var idcIpService idc.IdcIpService
	err := c.ShouldBindQuery(&idcIpService)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpService, err := idcIpServiceService.GetIdcIpService(idcIpService.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reidcIpService.IdcId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	count := 0
	var viewIpService = ViewIpService{
		reidcIpService,
		reidcInfo.Name,
		count,
		utils.GetTotalIPsInSubnet(reidcIpService.Network, *reidcIpService.Prefix),
	}
	response.OkWithData(gin.H{"reidcIpService": viewIpService}, c)
}

// GetIdcIpServiceList 分页获取数据中心业务网段列表
// @Tags IdcIpService
// @Summary 分页获取数据中心业务网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpServiceSearch true "分页获取数据中心业务网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpService/getIdcIpServiceList [get]
func (idcIpServiceApi *IdcIpServiceApi) GetIdcIpServiceList(c *gin.Context) {
	var pageInfo idcReq.IdcIpServiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := idcIpServiceService.GetIdcIpServiceInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var mapIdc = make(map[int]string)
	var viewIpServices []ViewIpService
	for _, ipService := range list {
		if mapIdc[*ipService.IdcId] == "" {
			reidcInfo, err := idcInfoService.GetIdcInfo(uint(*ipService.IdcId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
				return
			}
			mapIdc[*ipService.IdcId] = reidcInfo.Name
		}
		count := 0
		var viewIpService = ViewIpService{
			ipService,
			mapIdc[*ipService.IdcId],
			count,
			utils.GetTotalIPsInSubnet(ipService.Network, *ipService.Prefix),
		}
		viewIpServices = append(viewIpServices, viewIpService)
	}

	response.OkWithDetailed(response.PageResult{
		List:     viewIpServices,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (idcIpServiceApi *IdcIpServiceApi) ValidateIpService(c *gin.Context) {
	var pageInfo idcReq.IdcIpServiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpService, err := idcIpServiceService.GetIdcIpService(pageInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	//	if !utils.IsIPInSubnet(pageInfo.Address, reidcIpService.Network, *reidcIpService.Prefix) {
	if !utils.IsSubnetOf(reidcIpService.Network, *reidcIpService.Prefix, pageInfo.Address, 32) {
		response.FailWithMessage("IP地址不在网段内", c)
		return
	}

	var deviceBareMetalService = service.ServiceGroupApp.DeviceServiceGroup.DeviceBareMetalService
	count, err := deviceBareMetalService.CountDeviceBareMetalByIpService(pageInfo.Address)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	if count > 0 {
		response.FailWithMessage("IP地址已使用", c)
		return
	}

	response.OkWithData(gin.H{"status": true}, c)
}
