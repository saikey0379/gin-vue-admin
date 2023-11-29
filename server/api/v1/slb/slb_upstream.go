package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/slb"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SlbUpstreamApi struct {
}

var slbUpstreamService = service.ServiceGroupApp.SlbServiceGroup.SlbUpstreamService


// CreateSlbUpstream 创建目标节点
// @Tags SlbUpstream
// @Summary 创建目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbUpstream true "创建目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbUpstream/createSlbUpstream [post]
func (slbUpstreamApi *SlbUpstreamApi) CreateSlbUpstream(c *gin.Context) {
	var slbUpstream slb.SlbUpstream
	err := c.ShouldBindJSON(&slbUpstream)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
        "ClusterIds":{utils.NotEmpty()},
        "Nodes":{utils.NotEmpty()},
    }
	if err := utils.Verify(slbUpstream, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := slbUpstreamService.CreateSlbUpstream(&slbUpstream); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlbUpstream 删除目标节点
// @Tags SlbUpstream
// @Summary 删除目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbUpstream true "删除目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbUpstream/deleteSlbUpstream [delete]
func (slbUpstreamApi *SlbUpstreamApi) DeleteSlbUpstream(c *gin.Context) {
	var slbUpstream slb.SlbUpstream
	err := c.ShouldBindJSON(&slbUpstream)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbUpstreamService.DeleteSlbUpstream(slbUpstream); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlbUpstreamByIds 批量删除目标节点
// @Tags SlbUpstream
// @Summary 批量删除目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slbUpstream/deleteSlbUpstreamByIds [delete]
func (slbUpstreamApi *SlbUpstreamApi) DeleteSlbUpstreamByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbUpstreamService.DeleteSlbUpstreamByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlbUpstream 更新目标节点
// @Tags SlbUpstream
// @Summary 更新目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbUpstream true "更新目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbUpstream/updateSlbUpstream [put]
func (slbUpstreamApi *SlbUpstreamApi) UpdateSlbUpstream(c *gin.Context) {
	var slbUpstream slb.SlbUpstream
	err := c.ShouldBindJSON(&slbUpstream)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
          "ClusterIds":{utils.NotEmpty()},
          "Nodes":{utils.NotEmpty()},
      }
    if err := utils.Verify(slbUpstream, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := slbUpstreamService.UpdateSlbUpstream(slbUpstream); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlbUpstream 用id查询目标节点
// @Tags SlbUpstream
// @Summary 用id查询目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slb.SlbUpstream true "用id查询目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbUpstream/findSlbUpstream [get]
func (slbUpstreamApi *SlbUpstreamApi) FindSlbUpstream(c *gin.Context) {
	var slbUpstream slb.SlbUpstream
	err := c.ShouldBindQuery(&slbUpstream)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslbUpstream, err := slbUpstreamService.GetSlbUpstream(slbUpstream.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reslbUpstream": reslbUpstream}, c)
	}
}

// GetSlbUpstreamList 分页获取目标节点列表
// @Tags SlbUpstream
// @Summary 分页获取目标节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slbReq.SlbUpstreamSearch true "分页获取目标节点列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbUpstream/getSlbUpstreamList [get]
func (slbUpstreamApi *SlbUpstreamApi) GetSlbUpstreamList(c *gin.Context) {
	var pageInfo slbReq.SlbUpstreamSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slbUpstreamService.GetSlbUpstreamInfoList(pageInfo); err != nil {
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
