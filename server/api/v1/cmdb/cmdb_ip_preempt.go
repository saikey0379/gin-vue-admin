package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CmdbIpPreemptApi struct {
}

var cmdbIpPreemptService = service.ServiceGroupApp.CmdbServiceGroup.CmdbIpPreemptService


// CreateCmdbIpPreempt 创建预占地址
// @Tags CmdbIpPreempt
// @Summary 创建预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpPreempt true "创建预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbIpPreempt/createCmdbIpPreempt [post]
func (cmdbIpPreemptApi *CmdbIpPreemptApi) CreateCmdbIpPreempt(c *gin.Context) {
	var cmdbIpPreempt cmdb.CmdbIpPreempt
	err := c.ShouldBindJSON(&cmdbIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "RegionId":{utils.NotEmpty()},
        "SegmentId":{utils.NotEmpty()},
        "SubnetId":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(cmdbIpPreempt, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := cmdbIpPreemptService.CreateCmdbIpPreempt(&cmdbIpPreempt); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmdbIpPreempt 删除预占地址
// @Tags CmdbIpPreempt
// @Summary 删除预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpPreempt true "删除预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpPreempt/deleteCmdbIpPreempt [delete]
func (cmdbIpPreemptApi *CmdbIpPreemptApi) DeleteCmdbIpPreempt(c *gin.Context) {
	var cmdbIpPreempt cmdb.CmdbIpPreempt
	err := c.ShouldBindJSON(&cmdbIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbIpPreemptService.DeleteCmdbIpPreempt(cmdbIpPreempt); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmdbIpPreemptByIds 批量删除预占地址
// @Tags CmdbIpPreempt
// @Summary 批量删除预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmdbIpPreempt/deleteCmdbIpPreemptByIds [delete]
func (cmdbIpPreemptApi *CmdbIpPreemptApi) DeleteCmdbIpPreemptByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbIpPreemptService.DeleteCmdbIpPreemptByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmdbIpPreempt 更新预占地址
// @Tags CmdbIpPreempt
// @Summary 更新预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbIpPreempt true "更新预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbIpPreempt/updateCmdbIpPreempt [put]
func (cmdbIpPreemptApi *CmdbIpPreemptApi) UpdateCmdbIpPreempt(c *gin.Context) {
	var cmdbIpPreempt cmdb.CmdbIpPreempt
	err := c.ShouldBindJSON(&cmdbIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "RegionId":{utils.NotEmpty()},
          "SegmentId":{utils.NotEmpty()},
          "SubnetId":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(cmdbIpPreempt, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := cmdbIpPreemptService.UpdateCmdbIpPreempt(cmdbIpPreempt); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmdbIpPreempt 用id查询预占地址
// @Tags CmdbIpPreempt
// @Summary 用id查询预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdb.CmdbIpPreempt true "用id查询预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbIpPreempt/findCmdbIpPreempt [get]
func (cmdbIpPreemptApi *CmdbIpPreemptApi) FindCmdbIpPreempt(c *gin.Context) {
	var cmdbIpPreempt cmdb.CmdbIpPreempt
	err := c.ShouldBindQuery(&cmdbIpPreempt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recmdbIpPreempt, err := cmdbIpPreemptService.GetCmdbIpPreempt(cmdbIpPreempt.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmdbIpPreempt": recmdbIpPreempt}, c)
	}
}

// GetCmdbIpPreemptList 分页获取预占地址列表
// @Tags CmdbIpPreempt
// @Summary 分页获取预占地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdbReq.CmdbIpPreemptSearch true "分页获取预占地址列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbIpPreempt/getCmdbIpPreemptList [get]
func (cmdbIpPreemptApi *CmdbIpPreemptApi) GetCmdbIpPreemptList(c *gin.Context) {
	var pageInfo cmdbReq.CmdbIpPreemptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmdbIpPreemptService.GetCmdbIpPreemptInfoList(pageInfo); err != nil {
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
