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
)

type IdcIpPreemptApi struct {
}

var idcIpPreemptService = service.ServiceGroupApp.IdcServiceGroup.IdcIpPreemptService

type ViewIpPreempt struct {
	idc.IdcIpPreempt
	SubnetName  string `json:"subnetName"`
	SegmentId   int    `json:"segmentId"`
	SegmentName string `json:"segmentName"`
	IdcId       int    `json:"idcId"`
	IdcName     string `json:"idcName"`
}

// CreateIdcIpPreempt 创建数据中心地址预占
// @Tags IdcIpPreempt
// @Summary 创建数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpPreempt true "创建数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpPreempt/createIdcIpPreempt [post]
func (idcIpPreemptApi *IdcIpPreemptApi) CreateIdcIpPreempt(c *gin.Context) {
	var idcIpPreempt idc.IdcIpPreempt
	err := c.ShouldBindJSON(&idcIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":   {utils.NotEmpty()},
		"Prefix":    {utils.NotEmpty()},
		"IdcId":     {utils.NotEmpty()},
		"SegmentId": {utils.NotEmpty()},
		"SubnetId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Label":     {utils.NotEmpty()},
		"Describe":  {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpPreempt, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(uint(*idcIpPreempt.SubnetId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	if !utils.IsSubnetOf(reidcIpSubnet.Network, *reidcIpSubnet.Prefix, idcIpPreempt.Network, *idcIpPreempt.Prefix) {
		response.FailWithMessage(fmt.Sprintf("IP地址不在子网[%s/%d]内", reidcIpSubnet.Network, *reidcIpSubnet.Prefix), c)
		return
	}

	var pageInfo idcReq.IdcIpPreemptSearch
	pageInfo.SubnetId = idcIpPreempt.SubnetId
	list, _, err := idcIpPreemptService.GetIdcIpPreemptInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if utils.IsNetworkConflict(idcIpPreempt.Network, *idcIpPreempt.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := idcIpPreemptService.CreateIdcIpPreempt(&idcIpPreempt); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpPreempt 删除数据中心地址预占
// @Tags IdcIpPreempt
// @Summary 删除数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpPreempt true "删除数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpPreempt/deleteIdcIpPreempt [delete]
func (idcIpPreemptApi *IdcIpPreemptApi) DeleteIdcIpPreempt(c *gin.Context) {
	var idcIpPreempt idc.IdcIpPreempt
	err := c.ShouldBindJSON(&idcIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpPreemptService.DeleteIdcIpPreempt(idcIpPreempt); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpPreemptByIds 批量删除数据中心地址预占
// @Tags IdcIpPreempt
// @Summary 批量删除数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpPreempt/deleteIdcIpPreemptByIds [delete]
func (idcIpPreemptApi *IdcIpPreemptApi) DeleteIdcIpPreemptByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpPreemptService.DeleteIdcIpPreemptByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpPreempt 更新数据中心地址预占
// @Tags IdcIpPreempt
// @Summary 更新数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpPreempt true "更新数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpPreempt/updateIdcIpPreempt [put]
func (idcIpPreemptApi *IdcIpPreemptApi) UpdateIdcIpPreempt(c *gin.Context) {
	var idcIpPreempt idc.IdcIpPreempt
	err := c.ShouldBindJSON(&idcIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":   {utils.NotEmpty()},
		"Prefix":    {utils.NotEmpty()},
		"IdcId":     {utils.NotEmpty()},
		"SegmentId": {utils.NotEmpty()},
		"SubnetId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Label":     {utils.NotEmpty()},
		"Describe":  {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpPreempt, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(uint(*idcIpPreempt.SubnetId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	if !utils.IsSubnetOf(reidcIpSubnet.Network, *reidcIpSubnet.Prefix, idcIpPreempt.Network, *idcIpPreempt.Prefix) {
		response.FailWithMessage(fmt.Sprintf("IP地址不在子网[%s/%d]内", reidcIpSubnet.Network, *reidcIpSubnet.Prefix), c)
		return
	}

	var pageInfo idcReq.IdcIpPreemptSearch
	pageInfo.SubnetId = idcIpPreempt.SubnetId
	list, _, err := idcIpPreemptService.GetIdcIpPreemptInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if i.ID != idcIpPreempt.ID && utils.IsNetworkConflict(idcIpPreempt.Network, *idcIpPreempt.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := idcIpPreemptService.UpdateIdcIpPreempt(idcIpPreempt); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpPreempt 用id查询数据中心地址预占
// @Tags IdcIpPreempt
// @Summary 用id查询数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpPreempt true "用id查询数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpPreempt/findIdcIpPreempt [get]
func (idcIpPreemptApi *IdcIpPreemptApi) FindIdcIpPreempt(c *gin.Context) {
	var idcIpPreempt idc.IdcIpPreempt
	err := c.ShouldBindQuery(&idcIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpPreempt, err := idcIpPreemptService.GetIdcIpPreempt(idcIpPreempt.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(uint(*reidcIpPreempt.SubnetId))
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
	var viewIpPreempt = ViewIpPreempt{
		reidcIpPreempt,
		reidcIpSubnet.Name,
		int(reidcIpSegment.ID),
		reidcIpSegment.Name,
		int(reidcInfo.ID),
		reidcInfo.Name,
	}
	response.OkWithData(gin.H{"reidcIpPreempt": viewIpPreempt}, c)
}

// GetIdcIpPreemptList 分页获取数据中心地址预占列表
// @Tags IdcIpPreempt
// @Summary 分页获取数据中心地址预占列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpPreemptSearch true "分页获取数据中心地址预占列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpPreempt/getIdcIpPreemptList [get]
func (idcIpPreemptApi *IdcIpPreemptApi) GetIdcIpPreemptList(c *gin.Context) {
	var pageInfo idcReq.IdcIpPreemptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := idcIpPreemptService.GetIdcIpPreemptInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var mapIdcId = make(map[int]int)
	var mapIdcName = make(map[int]string)
	var mapSegmentId = make(map[int]int)
	var mapSegmentName = make(map[int]string)
	var mapSubnetName = make(map[int]string)

	var viewIpPreempts []ViewIpPreempt

	for _, ipPreempt := range list {
		if mapSubnetName[*ipPreempt.SubnetId] == "" {
			reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(uint(*ipPreempt.SubnetId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
				return
			}
			mapSubnetName[*ipPreempt.SubnetId] = reidcIpSubnet.Name

			if mapSegmentName[*ipPreempt.SubnetId] == "" {
				ipSegment, err := idcIpSegmentService.GetIdcIpSegment(uint(*reidcIpSubnet.SegmentId))
				if err != nil {
					global.GVA_LOG.Error("查询失败!", zap.Error(err))
					response.FailWithMessage("查询失败", c)
					return
				}
				mapSegmentId[*ipPreempt.SegmentId] = int(ipSegment.ID)
				mapSegmentName[*ipPreempt.SegmentId] = ipSegment.Name

				if mapIdcName[*ipSegment.IdcId] == "" {
					reidcInfo, err := idcInfoService.GetIdcInfo(uint(*ipSegment.IdcId))
					if err != nil {
						global.GVA_LOG.Error("查询失败!", zap.Error(err))
						response.FailWithMessage("查询失败", c)
						return
					}
					mapIdcId[*ipPreempt.IdcId] = int(reidcInfo.ID)
					mapIdcName[*ipPreempt.IdcId] = reidcInfo.Name
				}
			}
		}

		var viewIpPreempt = ViewIpPreempt{
			ipPreempt,
			mapSubnetName[*ipPreempt.SubnetId],
			mapSegmentId[*ipPreempt.SegmentId],
			mapSegmentName[*ipPreempt.SegmentId],
			mapIdcId[*ipPreempt.IdcId],
			mapIdcName[*ipPreempt.IdcId],
		}
		viewIpPreempts = append(viewIpPreempts, viewIpPreempt)
	}
	response.OkWithDetailed(response.PageResult{
		List:     viewIpPreempts,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
