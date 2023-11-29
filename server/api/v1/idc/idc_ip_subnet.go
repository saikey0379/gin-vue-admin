package idc

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type IdcIpSubnetApi struct {
}

var idcIpSubnetService = service.ServiceGroupApp.IdcServiceGroup.IdcIpSubnetService

type ViewIpSubnet struct {
	idc.IdcIpSubnet
	SegmentName string `json:"segmentName"`
	IdcId       int    `json:"idcId"`
	IdcName     string `json:"idcName"`
}

// CreateIdcIpSubnet 创建数据中心子网
// @Tags IdcIpSubnet
// @Summary 创建数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSubnet true "创建数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpSubnet/createIdcIpSubnet [post]
func (idcIpSubnetApi *IdcIpSubnetApi) CreateIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindJSON(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":   {utils.NotEmpty()},
		"Prefix":    {utils.NotEmpty()},
		"IdcId":     {utils.NotEmpty()},
		"SegmentId": {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Label":     {utils.NotEmpty()},
		"Describe":  {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpSubnet, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reidcIpSegment, err := idcIpSegmentService.GetIdcIpSegment(uint(*idcIpSubnet.SegmentId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	if !utils.IsSubnetOf(reidcIpSegment.Network, *reidcIpSegment.Prefix, idcIpSubnet.Network, *idcIpSubnet.Prefix) {
		response.FailWithMessage(fmt.Sprintf("地址不在网段[%s/%d]内", reidcIpSegment.Network, *reidcIpSegment.Prefix), c)
		return
	}

	if idcIpSubnet.Gateway != "" {
		if !utils.IsSubnetOf(idcIpSubnet.Network, *idcIpSubnet.Prefix, idcIpSubnet.Gateway, 32) {
			response.FailWithMessage(fmt.Sprintf("网关不在网段[%s/%d]内", reidcIpSegment.Network, *reidcIpSegment.Prefix), c)
			return
		}
	}

	var pageInfo idcReq.IdcIpSubnetSearch
	pageInfo.SegmentId = idcIpSubnet.SegmentId
	list, _, err := idcIpSubnetService.GetIdcIpSubnetInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if utils.IsNetworkConflict(idcIpSubnet.Network, *idcIpSubnet.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := idcIpSubnetService.CreateIdcIpSubnet(&idcIpSubnet); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpSubnet 删除数据中心子网
// @Tags IdcIpSubnet
// @Summary 删除数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSubnet true "删除数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSubnet/deleteIdcIpSubnet [delete]
func (idcIpSubnetApi *IdcIpSubnetApi) DeleteIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindJSON(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSubnetService.DeleteIdcIpSubnet(idcIpSubnet); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpSubnetByIds 批量删除数据中心子网
// @Tags IdcIpSubnet
// @Summary 批量删除数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpSubnet/deleteIdcIpSubnetByIds [delete]
func (idcIpSubnetApi *IdcIpSubnetApi) DeleteIdcIpSubnetByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSubnetService.DeleteIdcIpSubnetByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpSubnet 更新数据中心子网
// @Tags IdcIpSubnet
// @Summary 更新数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSubnet true "更新数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpSubnet/updateIdcIpSubnet [put]
func (idcIpSubnetApi *IdcIpSubnetApi) UpdateIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindJSON(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":   {utils.NotEmpty()},
		"Prefix":    {utils.NotEmpty()},
		"IdcId":     {utils.NotEmpty()},
		"SegmentId": {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Label":     {utils.NotEmpty()},
		"Describe":  {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpSubnet, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reidcIpSegment, err := idcIpSegmentService.GetIdcIpSegment(uint(*idcIpSubnet.SegmentId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	if !utils.IsSubnetOf(reidcIpSegment.Network, *reidcIpSegment.Prefix, idcIpSubnet.Network, *idcIpSubnet.Prefix) {
		response.FailWithMessage(fmt.Sprintf("地址不在网段[%s/%d]内", reidcIpSegment.Network, *reidcIpSegment.Prefix), c)
		return
	}

	if idcIpSubnet.Gateway != "" {
		if !utils.IsSubnetOf(idcIpSubnet.Network, *idcIpSubnet.Prefix, idcIpSubnet.Gateway, 32) {
			response.FailWithMessage(fmt.Sprintf("网关不在网段[%s/%d]内", reidcIpSegment.Network, *reidcIpSegment.Prefix), c)
			return
		}
	}

	var pageInfo idcReq.IdcIpSubnetSearch
	pageInfo.SegmentId = idcIpSubnet.SegmentId
	list, _, err := idcIpSubnetService.GetIdcIpSubnetInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if i.ID != idcIpSubnet.ID && utils.IsNetworkConflict(idcIpSubnet.Network, *idcIpSubnet.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := idcIpSubnetService.UpdateIdcIpSubnet(idcIpSubnet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpSubnet 用id查询数据中心子网
// @Tags IdcIpSubnet
// @Summary 用id查询数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpSubnet true "用id查询数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpSubnet/findIdcIpSubnet [get]
func (idcIpSubnetApi *IdcIpSubnetApi) FindIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindQuery(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(idcIpSubnet.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	reidcIpSegment, err := idcIpSegmentService.GetIdcIpSegment(uint(*reidcIpSubnet.SegmentId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reidcIpSegment.IdcId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	var viewIpSubnet = ViewIpSubnet{
		reidcIpSubnet,
		reidcIpSegment.Name,
		int(reidcInfo.ID),
		reidcInfo.Name,
	}
	response.OkWithData(gin.H{"reidcIpSubnet": viewIpSubnet}, c)
}

// GetIdcIpSubnetList 分页获取数据中心子网列表
// @Tags IdcIpSubnet
// @Summary 分页获取数据中心子网列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpSubnetSearch true "分页获取数据中心子网列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpSubnet/getIdcIpSubnetList [get]
func (idcIpSubnetApi *IdcIpSubnetApi) GetIdcIpSubnetList(c *gin.Context) {
	var pageInfo idcReq.IdcIpSubnetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	segmentId := c.Query("SegmentId")
	if segmentId != "" {
		id, err := strconv.Atoi(segmentId)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("SegmentId[%s]解析失败", segmentId), c)
			return
		}
		pageInfo.SegmentId = &id
	}

	list, total, err := idcIpSubnetService.GetIdcIpSubnetInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var mapIdcId = make(map[int]int)
	var mapIdcName = make(map[int]string)
	var mapSegmentName = make(map[int]string)
	var viewIpSubnets []ViewIpSubnet

	for _, ipSubnet := range list {
		if mapSegmentName[*ipSubnet.SegmentId] == "" {
			ipSegment, err := idcIpSegmentService.GetIdcIpSegment(uint(*ipSubnet.SegmentId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
				return
			}
			mapSegmentName[*ipSubnet.SegmentId] = ipSegment.Name

			if mapIdcName[*ipSegment.IdcId] == "" {
				reidcInfo, err := idcInfoService.GetIdcInfo(uint(*ipSegment.IdcId))
				if err != nil {
					global.GVA_LOG.Error("查询失败!", zap.Error(err))
					response.FailWithMessage("查询失败", c)
					return
				}
				mapIdcId[*ipSubnet.SegmentId] = int(reidcInfo.ID)
				mapIdcName[*ipSubnet.SegmentId] = reidcInfo.Name
			}
		}
		var viewIpSubnet = ViewIpSubnet{
			ipSubnet,
			mapSegmentName[*ipSubnet.SegmentId],
			mapIdcId[*ipSubnet.IdcId],
			mapIdcName[*ipSubnet.IdcId],
		}
		viewIpSubnets = append(viewIpSubnets, viewIpSubnet)
	}

	response.OkWithDetailed(response.PageResult{
		List:     viewIpSubnets,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (idcIpSubnetApi *IdcIpSubnetApi) ValidateIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindQuery(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(idcIpSubnet.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	address := c.Query("address")
	if !utils.IsSubnetOf(reidcIpSubnet.Network, *reidcIpSubnet.Prefix, address, 32) {
		response.FailWithMessage(fmt.Sprintf("地址不在子网[%s/%d]内", reidcIpSubnet.Network, *reidcIpSubnet.Prefix), c)
		return
	}
	response.OkWithData(gin.H{"Message": "IP校验成功"}, c)
}
