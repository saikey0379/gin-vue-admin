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

type RuleLabelApi struct {
}

var ruleLabelService = service.ServiceGroupApp.MonitServiceGroup.RuleLabelService

// CreateRuleLabel 创建规则标签
// @Tags RuleLabel
// @Summary 创建规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleLabel true "创建规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleLabel/createRuleLabel [post]
func (ruleLabelApi *RuleLabelApi) CreateRuleLabel(c *gin.Context) {
	var ruleLabel monit.RuleLabel
	err := c.ShouldBindJSON(&ruleLabel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleLabel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ruleLabel.ID = common.GenSnowFlakeID()

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	ruleLabel.Manager = claimsReal.Username

	if err := ruleLabelService.CreateRuleLabel(&ruleLabel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRuleLabel 删除规则标签
// @Tags RuleLabel
// @Summary 删除规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleLabel true "删除规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleLabel/deleteRuleLabel [delete]
func (ruleLabelApi *RuleLabelApi) DeleteRuleLabel(c *gin.Context) {
	var ruleLabel monit.RuleLabel
	err := c.ShouldBindJSON(&ruleLabel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleLabelService.DeleteRuleLabel(ruleLabel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRuleLabelByIds 批量删除规则标签
// @Tags RuleLabel
// @Summary 批量删除规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ruleLabel/deleteRuleLabelByIds [delete]
func (ruleLabelApi *RuleLabelApi) DeleteRuleLabelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleLabelService.DeleteRuleLabelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRuleLabel 更新规则标签
// @Tags RuleLabel
// @Summary 更新规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleLabel true "更新规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleLabel/updateRuleLabel [put]
func (ruleLabelApi *RuleLabelApi) UpdateRuleLabel(c *gin.Context) {
	var ruleLabel monit.RuleLabel
	err := c.ShouldBindJSON(&ruleLabel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleLabel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	ruleLabel.Manager = claimsReal.Username

	if err := ruleLabelService.UpdateRuleLabel(ruleLabel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRuleLabel 用id查询规则标签
// @Tags RuleLabel
// @Summary 用id查询规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monit.RuleLabel true "用id查询规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleLabel/findRuleLabel [get]
func (ruleLabelApi *RuleLabelApi) FindRuleLabel(c *gin.Context) {
	var ruleLabel monit.RuleLabel
	err := c.ShouldBindQuery(&ruleLabel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reruleLabel, err := ruleLabelService.GetRuleLabel(ruleLabel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reruleLabel": reruleLabel}, c)
	}
}

// GetRuleLabelList 分页获取规则标签列表
// @Tags RuleLabel
// @Summary 分页获取规则标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monitReq.RuleLabelSearch true "分页获取规则标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleLabel/getRuleLabelList [get]
func (ruleLabelApi *RuleLabelApi) GetRuleLabelList(c *gin.Context) {
	var pageInfo monitReq.RuleLabelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ruleLabelService.GetRuleLabelInfoList(pageInfo); err != nil {
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
