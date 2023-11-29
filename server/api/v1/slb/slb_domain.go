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

type SlbDomainApi struct {
}

var slbDomainService = service.ServiceGroupApp.SlbServiceGroup.SlbDomainService


// CreateSlbDomain 创建域名管理
// @Tags SlbDomain
// @Summary 创建域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbDomain true "创建域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbDomain/createSlbDomain [post]
func (slbDomainApi *SlbDomainApi) CreateSlbDomain(c *gin.Context) {
	var slbDomain slb.SlbDomain
	err := c.ShouldBindJSON(&slbDomain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
        "ClusterIds":{utils.NotEmpty()},
        "Http2":{utils.NotEmpty()},
        "Locations":{utils.NotEmpty()},
    }
	if err := utils.Verify(slbDomain, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := slbDomainService.CreateSlbDomain(&slbDomain); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlbDomain 删除域名管理
// @Tags SlbDomain
// @Summary 删除域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbDomain true "删除域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbDomain/deleteSlbDomain [delete]
func (slbDomainApi *SlbDomainApi) DeleteSlbDomain(c *gin.Context) {
	var slbDomain slb.SlbDomain
	err := c.ShouldBindJSON(&slbDomain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbDomainService.DeleteSlbDomain(slbDomain); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlbDomainByIds 批量删除域名管理
// @Tags SlbDomain
// @Summary 批量删除域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slbDomain/deleteSlbDomainByIds [delete]
func (slbDomainApi *SlbDomainApi) DeleteSlbDomainByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbDomainService.DeleteSlbDomainByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlbDomain 更新域名管理
// @Tags SlbDomain
// @Summary 更新域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbDomain true "更新域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbDomain/updateSlbDomain [put]
func (slbDomainApi *SlbDomainApi) UpdateSlbDomain(c *gin.Context) {
	var slbDomain slb.SlbDomain
	err := c.ShouldBindJSON(&slbDomain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
          "ClusterIds":{utils.NotEmpty()},
          "Http2":{utils.NotEmpty()},
          "Locations":{utils.NotEmpty()},
      }
    if err := utils.Verify(slbDomain, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := slbDomainService.UpdateSlbDomain(slbDomain); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlbDomain 用id查询域名管理
// @Tags SlbDomain
// @Summary 用id查询域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slb.SlbDomain true "用id查询域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbDomain/findSlbDomain [get]
func (slbDomainApi *SlbDomainApi) FindSlbDomain(c *gin.Context) {
	var slbDomain slb.SlbDomain
	err := c.ShouldBindQuery(&slbDomain)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslbDomain, err := slbDomainService.GetSlbDomain(slbDomain.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reslbDomain": reslbDomain}, c)
	}
}

// GetSlbDomainList 分页获取域名管理列表
// @Tags SlbDomain
// @Summary 分页获取域名管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slbReq.SlbDomainSearch true "分页获取域名管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbDomain/getSlbDomainList [get]
func (slbDomainApi *SlbDomainApi) GetSlbDomainList(c *gin.Context) {
	var pageInfo slbReq.SlbDomainSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slbDomainService.GetSlbDomainInfoList(pageInfo); err != nil {
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
