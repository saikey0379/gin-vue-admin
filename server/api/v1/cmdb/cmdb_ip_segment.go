package cmdb

import "C"
import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type CmdbIpSegmentApi struct {
}

var cmdbIpSegmentService = service.ServiceGroupApp.CmdbServiceGroup.CmdbIpSegmentService

type ViewIpSegment struct {
	cmdb.CmdbIpSegment
	RegionName string `json:"regionName" form:"regionName"`
}

// CreateCmdbIpSegment 创建网段管理
// @Tags CmdbIpSegment
// @Summary 创建网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpSegment true "创建网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbIpSegment/createCmdbIpSegment [post]
func (cmdbIpSegmentApi *CmdbIpSegmentApi) CreateCmdbIpSegment(c *gin.Context) {
	var cmdbIpSegment cmdb.CmdbIpSegment
	err := c.ShouldBindJSON(&cmdbIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":       {utils.NotEmpty()},
		"Prefix":        {utils.NotEmpty()},
		"RegionId":      {utils.NotEmpty()},
		"IpSegmentType": {utils.NotEmpty()},
		"Name":          {utils.NotEmpty()},
		"Label":         {utils.NotEmpty()},
		"Describe":      {utils.NotEmpty()},
	}
	if err := utils.Verify(cmdbIpSegment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo cmdbReq.CmdbIpSegmentSearch
	list, _, err := cmdbIpSegmentService.GetCmdbIpSegmentInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if utils.IsNetworkConflict(cmdbIpSegment.Network, *cmdbIpSegment.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := cmdbIpSegmentService.CreateCmdbIpSegment(&cmdbIpSegment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmdbIpSegment 删除网段管理
// @Tags CmdbIpSegment
// @Summary 删除网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpSegment true "删除网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpSegment/deleteCmdbIpSegment [delete]
func (cmdbIpSegmentApi *CmdbIpSegmentApi) DeleteCmdbIpSegment(c *gin.Context) {
	var cmdbIpSegment cmdb.CmdbIpSegment
	err := c.ShouldBindJSON(&cmdbIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbIpSegmentService.DeleteCmdbIpSegment(cmdbIpSegment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmdbIpSegmentByIds 批量删除网段管理
// @Tags CmdbIpSegment
// @Summary 批量删除网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmdbIpSegment/deleteCmdbIpSegmentByIds [delete]
func (cmdbIpSegmentApi *CmdbIpSegmentApi) DeleteCmdbIpSegmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbIpSegmentService.DeleteCmdbIpSegmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmdbIpSegment 更新网段管理
// @Tags CmdbIpSegment
// @Summary 更新网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpSegment true "更新网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbIpSegment/updateCmdbIpSegment [put]
func (cmdbIpSegmentApi *CmdbIpSegmentApi) UpdateCmdbIpSegment(c *gin.Context) {
	var cmdbIpSegment cmdb.CmdbIpSegment
	err := c.ShouldBindJSON(&cmdbIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Network":       {utils.NotEmpty()},
		"Prefix":        {utils.NotEmpty()},
		"RegionId":      {utils.NotEmpty()},
		"IpSegmentType": {utils.NotEmpty()},
		"Name":          {utils.NotEmpty()},
		"Label":         {utils.NotEmpty()},
		"Describe":      {utils.NotEmpty()},
	}
	if err := utils.Verify(cmdbIpSegment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo cmdbReq.CmdbIpSegmentSearch
	list, _, err := cmdbIpSegmentService.GetCmdbIpSegmentInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if i.ID != cmdbIpSegment.ID && utils.IsNetworkConflict(cmdbIpSegment.Network, *cmdbIpSegment.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := cmdbIpSegmentService.UpdateCmdbIpSegment(cmdbIpSegment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmdbIpSegment 用id查询网段管理
// @Tags CmdbIpSegment
// @Summary 用id查询网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdb.CmdbIpSegment true "用id查询网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbIpSegment/findCmdbIpSegment [get]
func (cmdbIpSegmentApi *CmdbIpSegmentApi) FindCmdbIpSegment(c *gin.Context) {
	var cmdbIpSegment cmdb.CmdbIpSegment
	err := c.ShouldBindQuery(&cmdbIpSegment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recmdbIpSegment, err := cmdbIpSegmentService.GetCmdbIpSegment(cmdbIpSegment.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var viewIpSegment ViewIpSegment
	viewIpSegment.CmdbIpSegment = recmdbIpSegment
	if *recmdbIpSegment.RegionId > 0 {
		recmdbRegion, err := cmdbRegionService.GetCmdbRegion(uint(*recmdbIpSegment.RegionId))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}

		viewIpSegment.RegionName = recmdbRegion.Name
	}
	response.OkWithData(gin.H{"recmdbIpSegment": viewIpSegment}, c)
}

// GetCmdbIpSegmentList 分页获取网段管理列表
// @Tags CmdbIpSegment
// @Summary 分页获取网段管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdbReq.CmdbIpSegmentSearch true "分页获取网段管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbIpSegment/getCmdbIpSegmentList [get]
func (cmdbIpSegmentApi *CmdbIpSegmentApi) GetCmdbIpSegmentList(c *gin.Context) {
	var pageInfo cmdbReq.CmdbIpSegmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	regionId := c.Query("RegionId")
	if regionId != "" {
		id, err := strconv.Atoi(regionId)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("RegionId[%s]解析失败", regionId), c)
			return
		}
		pageInfo.RegionId = &id
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

	list, total, err := cmdbIpSegmentService.GetCmdbIpSegmentInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var mapRegion = make(map[int]string)
	var viewIpSegments []ViewIpSegment
	for _, ipSegment := range list {
		var viewIpSegment ViewIpSegment
		viewIpSegment.CmdbIpSegment = ipSegment
		if *ipSegment.RegionId > 0 {
			if mapRegion[*ipSegment.RegionId] == "" {
				reregion, err := cmdbRegionService.GetCmdbRegion(uint(*ipSegment.RegionId))
				if err != nil {
					global.GVA_LOG.Error("查询失败!", zap.Error(err))
					response.FailWithMessage("查询失败", c)
					return
				}
				mapRegion[*ipSegment.RegionId] = reregion.Name
			}
			viewIpSegment.RegionName = mapRegion[*ipSegment.RegionId]
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
