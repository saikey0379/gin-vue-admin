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

type IdcIpManageApi struct {
}

type ViewIpManage struct {
	idc.IdcIpManage
	IdcName    string `json:"idc_name" form:"idc_name"`
	IpAllocate int    `json:"ip_allocate" form:"ip_allocate"`
	IpTotal    int    `json:"ip_total" form:"ip_total"`
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
		"Network":  {utils.NotEmpty()},
		"Prefix":   {utils.NotEmpty()},
		"IdcId":    {utils.NotEmpty()},
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpManage, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	network, err := utils.GetNetworkAddress(idcIpManage.Network, *idcIpManage.Prefix)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	idcIpManage.Network = network.String()
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
		"Network":  {utils.NotEmpty()},
		"Prefix":   {utils.NotEmpty()},
		"IdcId":    {utils.NotEmpty()},
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpManage, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	network, err := utils.GetNetworkAddress(idcIpManage.Network, *idcIpManage.Prefix)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	idcIpManage.Network = network.String()
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
	reidcIpManage, err := idcIpManageService.GetIdcIpManage(idcIpManage.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reidcIpManage.IdcId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	count := 0
	var viewIpManage = ViewIpManage{
		reidcIpManage,
		reidcInfo.Name,
		count,
		utils.GetTotalIPsInSubnet(reidcIpManage.Network, *reidcIpManage.Prefix),
	}
	response.OkWithData(gin.H{"reidcIpManage": viewIpManage}, c)
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
	list, total, err := idcIpManageService.GetIdcIpManageInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var mapIdc = make(map[int]string)
	var viewIpManages []ViewIpManage
	for _, ipManage := range list {
		if mapIdc[*ipManage.IdcId] == "" {
			reidcInfo, err := idcInfoService.GetIdcInfo(uint(*ipManage.IdcId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
			}
			mapIdc[*ipManage.IdcId] = reidcInfo.Name
		}
		count := 0
		var viewIpManage = ViewIpManage{
			ipManage,
			mapIdc[*ipManage.IdcId],
			count,
			utils.GetTotalIPsInSubnet(ipManage.Network, *ipManage.Prefix),
		}
		viewIpManages = append(viewIpManages, viewIpManage)
	}

	response.OkWithDetailed(response.PageResult{
		List:     viewIpManages,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (idcIpManageApi *IdcIpManageApi) ValidateIpManage(c *gin.Context) {
	var pageInfo idcReq.IdcIpManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpManage, err := idcIpManageService.GetIdcIpManage(pageInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	//	if !utils.IsIPInSubnet(pageInfo.Address, reidcIpManage.Network, *reidcIpManage.Prefix) {
	if !utils.IsSubnetOf(reidcIpManage.Network, *reidcIpManage.Prefix, pageInfo.Address, 32) {
		response.FailWithMessage("IP地址不在网段内", c)
		return
	}

	var deviceBareMetalService = service.ServiceGroupApp.DeviceServiceGroup.DeviceBareMetalService
	count, err := deviceBareMetalService.CountDeviceBareMetalByIpManage(pageInfo.Address)
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
