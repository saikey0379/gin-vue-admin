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

type IdcIpSegmentApi struct {
}

var idcIpSegmentService = service.ServiceGroupApp.IdcServiceGroup.IdcIpSegmentService

type ViewIpSegment struct {
	idc.IdcIpSegment
	IdcName string `json:"idcName" form:"idcName"`
}

// CreateIdcIpSegment 创建数据中心网段
// @Tags IdcIpSegment
// @Summary 创建数据中心网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSegment true "创建数据中心网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpSegment/createIdcIpSegment [post]
func (idcIpSegmentApi *IdcIpSegmentApi) CreateIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindJSON(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":       {utils.NotEmpty()},
		"Prefix":        {utils.NotEmpty()},
		"IdcId":         {utils.NotEmpty()},
		"IpSegmentType": {utils.NotEmpty()},
		"Name":          {utils.NotEmpty()},
		"Label":         {utils.NotEmpty()},
		"Describe":      {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpSegment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo idcReq.IdcIpSegmentSearch
	list, _, err := idcIpSegmentService.GetIdcIpSegmentInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if utils.IsNetworkConflict(idcIpSegment.Network, *idcIpSegment.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := idcIpSegmentService.CreateIdcIpSegment(&idcIpSegment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpSegment 删除数据中心网段
// @Tags IdcIpSegment
// @Summary 删除数据中心网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSegment true "删除数据中心网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSegment/deleteIdcIpSegment [delete]
func (idcIpSegmentApi *IdcIpSegmentApi) DeleteIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindJSON(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSegmentService.DeleteIdcIpSegment(idcIpSegment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpSegmentByIds 批量删除数据中心网段
// @Tags IdcIpSegment
// @Summary 批量删除数据中心网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpSegment/deleteIdcIpSegmentByIds [delete]
func (idcIpSegmentApi *IdcIpSegmentApi) DeleteIdcIpSegmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSegmentService.DeleteIdcIpSegmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpSegment 更新数据中心网段
// @Tags IdcIpSegment
// @Summary 更新数据中心网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSegment true "更新数据中心网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpSegment/updateIdcIpSegment [put]
func (idcIpSegmentApi *IdcIpSegmentApi) UpdateIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindJSON(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":       {utils.NotEmpty()},
		"Prefix":        {utils.NotEmpty()},
		"IdcId":         {utils.NotEmpty()},
		"IpSegmentType": {utils.NotEmpty()},
		"Name":          {utils.NotEmpty()},
		"Label":         {utils.NotEmpty()},
		"Describe":      {utils.NotEmpty()},
	}
	if err := utils.Verify(idcIpSegment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo idcReq.IdcIpSegmentSearch
	list, _, err := idcIpSegmentService.GetIdcIpSegmentInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if i.ID != idcIpSegment.ID && utils.IsNetworkConflict(idcIpSegment.Network, *idcIpSegment.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := idcIpSegmentService.UpdateIdcIpSegment(idcIpSegment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpSegment 用id查询数据中心网段
// @Tags IdcIpSegment
// @Summary 用id查询数据中心网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpSegment true "用id查询数据中心网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpSegment/findIdcIpSegment [get]
func (idcIpSegmentApi *IdcIpSegmentApi) FindIdcIpSegment(c *gin.Context) {
	var idcIpSegment idc.IdcIpSegment
	err := c.ShouldBindQuery(&idcIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcIpSegment, err := idcIpSegmentService.GetIdcIpSegment(idcIpSegment.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	var viewIpSegment ViewIpSegment
	viewIpSegment.IdcIpSegment = reidcIpSegment
	if *reidcIpSegment.IdcId > 0 {
		reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reidcIpSegment.IdcId))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}

		viewIpSegment.IdcName = reidcInfo.Name
	}
	response.OkWithData(gin.H{"reidcIpSegment": viewIpSegment}, c)
}

// GetIdcIpSegmentList 分页获取数据中心网段列表
// @Tags IdcIpSegment
// @Summary 分页获取数据中心网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpSegmentSearch true "分页获取数据中心网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpSegment/getIdcIpSegmentList [get]
func (idcIpSegmentApi *IdcIpSegmentApi) GetIdcIpSegmentList(c *gin.Context) {
	var pageInfo idcReq.IdcIpSegmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	idcId := c.Query("IdcId")
	if idcId != "" {
		id, err := strconv.Atoi(idcId)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("IdcId[%s]解析失败", idcId), c)
			return
		}
		pageInfo.IdcId = &id
	}

	ipSegmentType := c.Query("ipSegmentType")
	if ipSegmentType != "" {
		tp, err := strconv.Atoi(ipSegmentType)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("ipSegmentType[%s]解析失败", ipSegmentType), c)
			return
		}
		pageInfo.IpSegmentType = &tp
	}

	list, total, err := idcIpSegmentService.GetIdcIpSegmentInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var mapIdc = make(map[int]string)
	var viewIpSegments []ViewIpSegment
	for _, ipSegment := range list {
		var viewIpSegment ViewIpSegment
		viewIpSegment.IdcIpSegment = ipSegment
		if *ipSegment.IdcId > 0 {
			if mapIdc[*ipSegment.IdcId] == "" {
				reidcInfo, err := idcInfoService.GetIdcInfo(uint(*ipSegment.IdcId))
				if err != nil {
					global.GVA_LOG.Error("查询失败!", zap.Error(err))
					response.FailWithMessage("查询失败", c)
					return
				}
				mapIdc[*ipSegment.IdcId] = reidcInfo.Name
			}
			viewIpSegment.IdcName = mapIdc[*ipSegment.IdcId]
		}
		viewIpSegments = append(viewIpSegments, viewIpSegment)
	}

	response.OkWithDetailed(response.PageResult{
		List:     viewIpSegments,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
