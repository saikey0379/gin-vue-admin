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

type RuleAnnotationApi struct {
}

var ruleAnnotationService = service.ServiceGroupApp.MonitServiceGroup.RuleAnnotationService

// CreateRuleAnnotation 创建注解标签
// @Tags RuleAnnotation
// @Summary 创建注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleAnnotation true "创建注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleAnnotation/createRuleAnnotation [post]
func (ruleAnnotationApi *RuleAnnotationApi) CreateRuleAnnotation(c *gin.Context) {
	var ruleAnnotation monit.RuleAnnotation
	err := c.ShouldBindJSON(&ruleAnnotation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleAnnotation, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ruleAnnotation.ID = common.GenSnowFlakeID()

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	ruleAnnotation.Manager = claimsReal.Username

	if err := ruleAnnotationService.CreateRuleAnnotation(&ruleAnnotation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRuleAnnotation 删除注解标签
// @Tags RuleAnnotation
// @Summary 删除注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleAnnotation true "删除注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleAnnotation/deleteRuleAnnotation [delete]
func (ruleAnnotationApi *RuleAnnotationApi) DeleteRuleAnnotation(c *gin.Context) {
	var ruleAnnotation monit.RuleAnnotation
	err := c.ShouldBindJSON(&ruleAnnotation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleAnnotationService.DeleteRuleAnnotation(ruleAnnotation); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRuleAnnotationByIds 批量删除注解标签
// @Tags RuleAnnotation
// @Summary 批量删除注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ruleAnnotation/deleteRuleAnnotationByIds [delete]
func (ruleAnnotationApi *RuleAnnotationApi) DeleteRuleAnnotationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleAnnotationService.DeleteRuleAnnotationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRuleAnnotation 更新注解标签
// @Tags RuleAnnotation
// @Summary 更新注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleAnnotation true "更新注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleAnnotation/updateRuleAnnotation [put]
func (ruleAnnotationApi *RuleAnnotationApi) UpdateRuleAnnotation(c *gin.Context) {
	var ruleAnnotation monit.RuleAnnotation
	err := c.ShouldBindJSON(&ruleAnnotation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleAnnotation, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	ruleAnnotation.Manager = claimsReal.Username

	if err := ruleAnnotationService.UpdateRuleAnnotation(ruleAnnotation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRuleAnnotation 用id查询注解标签
// @Tags RuleAnnotation
// @Summary 用id查询注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monit.RuleAnnotation true "用id查询注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleAnnotation/findRuleAnnotation [get]
func (ruleAnnotationApi *RuleAnnotationApi) FindRuleAnnotation(c *gin.Context) {
	var ruleAnnotation monit.RuleAnnotation
	err := c.ShouldBindQuery(&ruleAnnotation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reruleAnnotation, err := ruleAnnotationService.GetRuleAnnotation(ruleAnnotation.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reruleAnnotation": reruleAnnotation}, c)
	}
}

// GetRuleAnnotationList 分页获取注解标签列表
// @Tags RuleAnnotation
// @Summary 分页获取注解标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monitReq.RuleAnnotationSearch true "分页获取注解标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleAnnotation/getRuleAnnotationList [get]
func (ruleAnnotationApi *RuleAnnotationApi) GetRuleAnnotationList(c *gin.Context) {
	var pageInfo monitReq.RuleAnnotationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ruleAnnotationService.GetRuleAnnotationInfoList(pageInfo); err != nil {
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
