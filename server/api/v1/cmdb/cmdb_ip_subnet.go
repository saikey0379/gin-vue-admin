package cmdb

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
	"strconv"
)

type CmdbIpSubnetApi struct {
}

var cmdbIpSubnetService = service.ServiceGroupApp.CmdbServiceGroup.CmdbIpSubnetService

type ViewIpSubnet struct {
	cmdb.CmdbIpSubnet
	SegmentName string `json:"segmentName"`
	RegionId       int    `json:"regionId"`
	RegionName     string `json:"regionName"`
}

// CreateCmdbIpSubnet 创建子网管理
// @Tags CmdbIpSubnet
// @Summary 创建子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpSubnet true "创建子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbIpSubnet/createCmdbIpSubnet [post]
func (cmdbIpSubnetApi *CmdbIpSubnetApi) CreateCmdbIpSubnet(c *gin.Context) {
	var cmdbIpSubnet cmdb.CmdbIpSubnet
	err := c.ShouldBindJSON(&cmdbIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "RegionId":{utils.NotEmpty()},
        "SegmentId":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(cmdbIpSubnet, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    }

	recmdbIpSegment, err := cmdbIpSegmentService.GetCmdbIpSegment(uint(*cmdbIpSubnet.SegmentId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	if !utils.IsSubnetOf(recmdbIpSegment.Network, *recmdbIpSegment.Prefix, cmdbIpSubnet.Network, *cmdbIpSubnet.Prefix) {
		response.FailWithMessage(fmt.Sprintf("地址不在网段[%s/%d]内", recmdbIpSegment.Network, *recmdbIpSegment.Prefix), c)
		return
	}

	if cmdbIpSubnet.Gateway != "" {
		if !utils.IsSubnetOf(cmdbIpSubnet.Network, *cmdbIpSubnet.Prefix, cmdbIpSubnet.Gateway, 32) {
			response.FailWithMessage(fmt.Sprintf("网关不在网段[%s/%d]内", recmdbIpSegment.Network, *recmdbIpSegment.Prefix), c)
			return
		}
	}

	var pageInfo cmdbReq.CmdbIpSubnetSearch
	pageInfo.SegmentId = cmdbIpSubnet.SegmentId
	list, _, err := cmdbIpSubnetService.GetCmdbIpSubnetInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if utils.IsNetworkConflict(cmdbIpSubnet.Network, *cmdbIpSubnet.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := cmdbIpSubnetService.CreateCmdbIpSubnet(&cmdbIpSubnet); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmdbIpSubnet 删除子网管理
// @Tags CmdbIpSubnet
// @Summary 删除子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpSubnet true "删除子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpSubnet/deleteCmdbIpSubnet [delete]
func (cmdbIpSubnetApi *CmdbIpSubnetApi) DeleteCmdbIpSubnet(c *gin.Context) {
	var cmdbIpSubnet cmdb.CmdbIpSubnet
	err := c.ShouldBindJSON(&cmdbIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbIpSubnetService.DeleteCmdbIpSubnet(cmdbIpSubnet); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmdbIpSubnetByIds 批量删除子网管理
// @Tags CmdbIpSubnet
// @Summary 批量删除子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmdbIpSubnet/deleteCmdbIpSubnetByIds [delete]
func (cmdbIpSubnetApi *CmdbIpSubnetApi) DeleteCmdbIpSubnetByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbIpSubnetService.DeleteCmdbIpSubnetByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmdbIpSubnet 更新子网管理
// @Tags CmdbIpSubnet
// @Summary 更新子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpSubnet true "更新子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbIpSubnet/updateCmdbIpSubnet [put]
func (cmdbIpSubnetApi *CmdbIpSubnetApi) UpdateCmdbIpSubnet(c *gin.Context) {
	var cmdbIpSubnet cmdb.CmdbIpSubnet
	err := c.ShouldBindJSON(&cmdbIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "RegionId":{utils.NotEmpty()},
          "SegmentId":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
    }
    if err := utils.Verify(cmdbIpSubnet, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }

	recmdbIpSegment, err := cmdbIpSegmentService.GetCmdbIpSegment(uint(*cmdbIpSubnet.SegmentId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	if !utils.IsSubnetOf(recmdbIpSegment.Network, *recmdbIpSegment.Prefix, cmdbIpSubnet.Network, *cmdbIpSubnet.Prefix) {
		response.FailWithMessage(fmt.Sprintf("地址不在网段[%s/%d]内", recmdbIpSegment.Network, *recmdbIpSegment.Prefix), c)
		return
	}

	if cmdbIpSubnet.Gateway != "" {
		if !utils.IsSubnetOf(cmdbIpSubnet.Network, *cmdbIpSubnet.Prefix, cmdbIpSubnet.Gateway, 32) {
			response.FailWithMessage(fmt.Sprintf("网关不在网段[%s/%d]内", recmdbIpSegment.Network, *recmdbIpSegment.Prefix), c)
			return
		}
	}

	var pageInfo cmdbReq.CmdbIpSubnetSearch
	pageInfo.SegmentId = cmdbIpSubnet.SegmentId
	list, _, err := cmdbIpSubnetService.GetCmdbIpSubnetInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for _, i := range list {
		if i.ID != cmdbIpSubnet.ID && utils.IsNetworkConflict(cmdbIpSubnet.Network, *cmdbIpSubnet.Prefix, i.Network, *i.Prefix) {
			global.GVA_LOG.Error(fmt.Sprintf("网段冲突![%s]", i.Name), zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("网段冲突[%s]", i.Name), c)
			return
		}
	}

	if err := cmdbIpSubnetService.UpdateCmdbIpSubnet(cmdbIpSubnet); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmdbIpSubnet 用id查询子网管理
// @Tags CmdbIpSubnet
// @Summary 用id查询子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdb.CmdbIpSubnet true "用id查询子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbIpSubnet/findCmdbIpSubnet [get]
func (cmdbIpSubnetApi *CmdbIpSubnetApi) FindCmdbIpSubnet(c *gin.Context) {
	var cmdbIpSubnet cmdb.CmdbIpSubnet
	err := c.ShouldBindQuery(&cmdbIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recmdbIpSubnet, err := cmdbIpSubnetService.GetCmdbIpSubnet(cmdbIpSubnet.ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	recmdbIpSegment, err := cmdbIpSegmentService.GetCmdbIpSegment(uint(*recmdbIpSubnet.SegmentId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	recmdbRegion, err := cmdbRegionService.GetCmdbRegion(uint(*recmdbIpSegment.RegionId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	var viewIpSubnet = ViewIpSubnet{
		recmdbIpSubnet,
		recmdbIpSegment.Name,
		int(recmdbRegion.ID),
		recmdbRegion.Name,
	}
	response.OkWithData(gin.H{"recmdbIpSubnet": viewIpSubnet}, c)
}

// GetCmdbIpSubnetList 分页获取子网管理列表
// @Tags CmdbIpSubnet
// @Summary 分页获取子网管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdbReq.CmdbIpSubnetSearch true "分页获取子网管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbIpSubnet/getCmdbIpSubnetList [get]
func (cmdbIpSubnetApi *CmdbIpSubnetApi) GetCmdbIpSubnetList(c *gin.Context) {
	var pageInfo cmdbReq.CmdbIpSubnetSearch
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

	list, total, err := cmdbIpSubnetService.GetCmdbIpSubnetInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
		return
    }

	var mapRegionId = make(map[int]int)
	var mapRegionName = make(map[int]string)
	var mapSegmentName = make(map[int]string)
	var viewIpSubnets []ViewIpSubnet

	for _, ipSubnet := range list {
		if mapSegmentName[*ipSubnet.SegmentId] == "" {
			ipSegment, err := cmdbIpSegmentService.GetCmdbIpSegment(uint(*ipSubnet.SegmentId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
				return
			}
			mapSegmentName[*ipSubnet.SegmentId] = ipSegment.Name

			if mapRegionName[*ipSegment.RegionId] == "" {
				recmdbRegion, err := cmdbRegionService.GetCmdbRegion(uint(*ipSegment.RegionId))
				if err != nil {
					global.GVA_LOG.Error("查询失败!", zap.Error(err))
					response.FailWithMessage("查询失败", c)
					return
				}
				mapRegionId[*ipSubnet.SegmentId] = int(recmdbRegion.ID)
				mapRegionName[*ipSubnet.SegmentId] = recmdbRegion.Name
			}
		}
		var viewIpSubnet = ViewIpSubnet{
			ipSubnet,
			mapSegmentName[*ipSubnet.SegmentId],
			mapRegionId[*ipSubnet.RegionId],
			mapRegionName[*ipSubnet.RegionId],
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

func (cmdbIpSubnetApi *CmdbIpSubnetApi) ValidateIpSubnet(c *gin.Context) {
	var cmdbIpSubnet cmdb.CmdbIpSubnet
	err := c.ShouldBindQuery(&cmdbIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recmdbIpSubnet, err := cmdbIpSubnetService.GetCmdbIpSubnet(cmdbIpSubnet.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	address := c.Query("address")
	if !utils.IsSubnetOf(recmdbIpSubnet.Network, *recmdbIpSubnet.Prefix, address, 32) {
		response.FailWithMessage(fmt.Sprintf("地址不在子网[%s/%d]内", recmdbIpSubnet.Network, *recmdbIpSubnet.Prefix), c)
		return
	}
	response.OkWithData(gin.H{"Message": "IP校验成功"}, c)
}