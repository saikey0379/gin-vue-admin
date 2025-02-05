package monit

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
	request1 "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RuleGroupApi struct {
}

var ruleGroupService = service.ServiceGroupApp.MonitServiceGroup.RuleGroupService

// CreateRuleGroup 创建规则分组
// @Tags RuleGroup
// @Summary 创建规则分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleGroup true "创建规则分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleGroup/createRuleGroup [post]
func (ruleGroupApi *RuleGroupApi) CreateRuleGroup(c *gin.Context) {
	var ruleGroup monit.RuleGroup
	err := c.ShouldBindJSON(&ruleGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleGroup, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ruleGroup.ID = common.GenSnowFlakeID()

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	ruleGroup.Manager = claimsReal.Username

	if err := ruleGroupService.CreateRuleGroup(&ruleGroup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRuleGroup 删除规则分组
// @Tags RuleGroup
// @Summary 删除规则分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleGroup true "删除规则分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleGroup/deleteRuleGroup [delete]
func (ruleGroupApi *RuleGroupApi) DeleteRuleGroup(c *gin.Context) {
	var ruleGroup monit.RuleGroup
	err := c.ShouldBindJSON(&ruleGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleGroupService.DeleteRuleGroup(ruleGroup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRuleGroupByIds 批量删除规则分组
// @Tags RuleGroup
// @Summary 批量删除规则分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除规则分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ruleGroup/deleteRuleGroupByIds [delete]
func (ruleGroupApi *RuleGroupApi) DeleteRuleGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleGroupService.DeleteRuleGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRuleGroup 更新规则分组
// @Tags RuleGroup
// @Summary 更新规则分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleGroup true "更新规则分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleGroup/updateRuleGroup [put]
func (ruleGroupApi *RuleGroupApi) UpdateRuleGroup(c *gin.Context) {
	var ruleGroup monit.RuleGroup
	err := c.ShouldBindJSON(&ruleGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleGroup, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	ruleGroup.Manager = claimsReal.Username

	if err := ruleGroupService.UpdateRuleGroup(ruleGroup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRuleGroup 用id查询规则分组
// @Tags RuleGroup
// @Summary 用id查询规则分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monit.RuleGroup true "用id查询规则分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleGroup/findRuleGroup [get]
func (ruleGroupApi *RuleGroupApi) FindRuleGroup(c *gin.Context) {
	var ruleGroup monit.RuleGroup
	err := c.ShouldBindQuery(&ruleGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reruleGroup, err := ruleGroupService.GetRuleGroup(ruleGroup.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reruleGroup": reruleGroup}, c)
	}
}

// GetRuleGroupList 分页获取规则分组列表
// @Tags RuleGroup
// @Summary 分页获取规则分组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monitReq.RuleGroupSearch true "分页获取规则分组列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleGroup/getRuleGroupList [get]
func (ruleGroupApi *RuleGroupApi) GetRuleGroupList(c *gin.Context) {
	var pageInfo monitReq.RuleGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ruleGroupService.GetRuleGroupInfoList(pageInfo); err != nil {
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
