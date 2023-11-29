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

type SlbAccesslistApi struct {
}

var slbAccesslistService = service.ServiceGroupApp.SlbServiceGroup.SlbAccesslistService


// CreateSlbAccesslist 创建访问控制
// @Tags SlbAccesslist
// @Summary 创建访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbAccesslist true "创建访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbAccesslist/createSlbAccesslist [post]
func (slbAccesslistApi *SlbAccesslistApi) CreateSlbAccesslist(c *gin.Context) {
	var slbAccesslist slb.SlbAccesslist
	err := c.ShouldBindJSON(&slbAccesslist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "ClusterId":{utils.NotEmpty()},
        "DomainId":{utils.NotEmpty()},
        "RouteId":{utils.NotEmpty()},
        "AccessType":{utils.NotEmpty()},
        "Nodes":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(slbAccesslist, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := slbAccesslistService.CreateSlbAccesslist(&slbAccesslist); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlbAccesslist 删除访问控制
// @Tags SlbAccesslist
// @Summary 删除访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbAccesslist true "删除访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbAccesslist/deleteSlbAccesslist [delete]
func (slbAccesslistApi *SlbAccesslistApi) DeleteSlbAccesslist(c *gin.Context) {
	var slbAccesslist slb.SlbAccesslist
	err := c.ShouldBindJSON(&slbAccesslist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbAccesslistService.DeleteSlbAccesslist(slbAccesslist); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlbAccesslistByIds 批量删除访问控制
// @Tags SlbAccesslist
// @Summary 批量删除访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slbAccesslist/deleteSlbAccesslistByIds [delete]
func (slbAccesslistApi *SlbAccesslistApi) DeleteSlbAccesslistByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbAccesslistService.DeleteSlbAccesslistByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlbAccesslist 更新访问控制
// @Tags SlbAccesslist
// @Summary 更新访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbAccesslist true "更新访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbAccesslist/updateSlbAccesslist [put]
func (slbAccesslistApi *SlbAccesslistApi) UpdateSlbAccesslist(c *gin.Context) {
	var slbAccesslist slb.SlbAccesslist
	err := c.ShouldBindJSON(&slbAccesslist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "ClusterId":{utils.NotEmpty()},
          "DomainId":{utils.NotEmpty()},
          "RouteId":{utils.NotEmpty()},
          "AccessType":{utils.NotEmpty()},
          "Nodes":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(slbAccesslist, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := slbAccesslistService.UpdateSlbAccesslist(slbAccesslist); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlbAccesslist 用id查询访问控制
// @Tags SlbAccesslist
// @Summary 用id查询访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slb.SlbAccesslist true "用id查询访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbAccesslist/findSlbAccesslist [get]
func (slbAccesslistApi *SlbAccesslistApi) FindSlbAccesslist(c *gin.Context) {
	var slbAccesslist slb.SlbAccesslist
	err := c.ShouldBindQuery(&slbAccesslist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslbAccesslist, err := slbAccesslistService.GetSlbAccesslist(slbAccesslist.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reslbAccesslist": reslbAccesslist}, c)
	}
}

// GetSlbAccesslistList 分页获取访问控制列表
// @Tags SlbAccesslist
// @Summary 分页获取访问控制列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slbReq.SlbAccesslistSearch true "分页获取访问控制列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbAccesslist/getSlbAccesslistList [get]
func (slbAccesslistApi *SlbAccesslistApi) GetSlbAccesslistList(c *gin.Context) {
	var pageInfo slbReq.SlbAccesslistSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slbAccesslistService.GetSlbAccesslistInfoList(pageInfo); err != nil {
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
