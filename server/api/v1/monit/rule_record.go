package monit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/prometheus/prometheus/promql/parser"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type RuleRecordApi struct {
}

var ruleRecordService = service.ServiceGroupApp.MonitServiceGroup.RuleRecordService

type RuleRecordReqResp struct {
	monit.RuleRecord
	ClusterArray    []int           `json:"clusterArray"`
	LabelArray      []utils.RuleMap `json:"labelArray"`
	AnnotationArray []utils.RuleMap `json:"annotationArray"`
}

// CreateRuleRecord 创建规则配置
// @Tags RuleRecord
// @Summary 创建规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleRecord true "创建规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleRecord/createRuleRecord [post]
func (ruleRecordApi *RuleRecordApi) CreateRuleRecord(c *gin.Context) {
	var ruleRecordRequest RuleRecordReqResp
	err := c.ShouldBindJSON(&ruleRecordRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"RuleGroupId":  {utils.NotEmpty()},
		"ClusterArray": {utils.NotEmptyArray()},
		"Describe":     {utils.NotEmpty()},
		"Record":       {utils.NotEmpty()},
		"Expr":         {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleRecordRequest, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// Verify Rule
	expr, err := parser.ParseExpr(ruleRecordRequest.Expr)
	if err != nil {
		global.GVA_LOG.Error("Rule.Expr异常!", zap.Error(err))
		response.FailWithMessage("Rule.Expr异常", c)
		return
	}

	var rule = rulefmt.Rule{
		Record: ruleRecordRequest.Record,
		Expr:   expr.String(),
	}

	if ruleRecordRequest.Alert != "" {
		rule.Alert = ruleRecordRequest.Alert
	}

	if ruleRecordRequest.For != "" {
		ruleFor, err := time.ParseDuration(ruleRecordRequest.For)
		if err != nil {
			global.GVA_LOG.Error("Rule.For异常!", zap.Error(err))
			response.FailWithMessage("Rule.For异常", c)
			return
		}
		rule.For = model.Duration(ruleFor)
	}

	if ruleRecordRequest.KeepFiringFor != "" {
		ruleKeepFiringFor, err := time.ParseDuration(ruleRecordRequest.KeepFiringFor)
		if err != nil {
			global.GVA_LOG.Error("Rule.KeepFiringFor异常!", zap.Error(err))
			response.FailWithMessage("Rule.KeepFiringFor异常", c)
			return
		}
		rule.KeepFiringFor = model.Duration(ruleKeepFiringFor)
	}

	if len(ruleRecordRequest.LabelArray) > 0 {
		labelMap := make(map[string]string)
		for _, i := range ruleRecordRequest.LabelArray {
			labelMap[i.Key] = i.Value
		}
		rule.Labels = labelMap
	}

	if len(ruleRecordRequest.AnnotationArray) > 0 {
		annotationMap := make(map[string]string)
		for _, i := range ruleRecordRequest.AnnotationArray {
			annotationMap[i.Key] = i.Value
		}
		rule.Annotations = annotationMap
	}

	_, err = yaml.Marshal(rule)
	if err != nil {
		global.GVA_LOG.Error("Rule序列化失败!", zap.Error(err))
		response.FailWithMessage("Rule序列化失败", c)
		return
	}

	// modRuleRecord
	byteLabels, err := json.Marshal(ruleRecordRequest.LabelArray)
	if err != nil {
		global.GVA_LOG.Error("Rule.Labels序列化失败!", zap.Error(err))
		response.FailWithMessage("Rule.Labels序列化失败", c)
		return
	}

	byteAnnotations, err := json.Marshal(ruleRecordRequest.AnnotationArray)
	if err != nil {
		global.GVA_LOG.Error("Rule.Annotations序列化失败!", zap.Error(err))
		response.FailWithMessage("Rule.Annotations序列化失败", c)
		return
	}

	var clusters []string
	for _, i := range ruleRecordRequest.ClusterArray {
		cluster := strconv.Itoa(i)
		clusters = append(clusters, cluster)
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)

	var ruleRecord = monit.RuleRecord{
		GVA_MODEL:     ruleRecordRequest.GVA_MODEL,
		Name:          ruleRecordRequest.Name,
		RuleGroupId:   ruleRecordRequest.RuleGroupId,
		Clusters:      strings.Join(clusters, ","),
		Manager:       claimsReal.Username,
		Describe:      ruleRecordRequest.Describe,
		Record:        ruleRecordRequest.Record,
		Alert:         ruleRecordRequest.Alert,
		Expr:          expr.String(),
		For:           ruleRecordRequest.For,
		Labels:        string(byteLabels),
		Annotations:   string(byteAnnotations),
		KeepFiringFor: ruleRecordRequest.KeepFiringFor,
	}

	ruleRecord.ID = common.GenSnowFlakeID()

	if err := ruleRecordService.CreateRuleRecord(&ruleRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRuleRecord 删除规则配置
// @Tags RuleRecord
// @Summary 删除规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleRecord true "删除规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleRecord/deleteRuleRecord [delete]
func (ruleRecordApi *RuleRecordApi) DeleteRuleRecord(c *gin.Context) {
	var ruleRecord monit.RuleRecord
	err := c.ShouldBindJSON(&ruleRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleRecordService.DeleteRuleRecord(ruleRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRuleRecordByIds 批量删除规则配置
// @Tags RuleRecord
// @Summary 批量删除规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ruleRecord/deleteRuleRecordByIds [delete]
func (ruleRecordApi *RuleRecordApi) DeleteRuleRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ruleRecordService.DeleteRuleRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRuleRecord 更新规则配置
// @Tags RuleRecord
// @Summary 更新规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.RuleRecord true "更新规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleRecord/updateRuleRecord [put]
func (ruleRecordApi *RuleRecordApi) UpdateRuleRecord(c *gin.Context) {
	var ruleRecordRequest RuleRecordReqResp
	err := c.ShouldBindJSON(&ruleRecordRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"RuleGroupId":  {utils.NotEmpty()},
		"ClusterArray": {utils.NotEmptyArray()},
		"Describe":     {utils.NotEmpty()},
		"Record":       {utils.NotEmpty()},
		"Expr":         {utils.NotEmpty()},
	}
	if err := utils.Verify(ruleRecordRequest, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// Verify Rule
	expr, err := parser.ParseExpr(ruleRecordRequest.Expr)
	if err != nil {
		global.GVA_LOG.Error("Rule.Expr异常!", zap.Error(err))
		response.FailWithMessage("Rule.Expr异常", c)
		return
	}

	var rule = rulefmt.Rule{
		Record: ruleRecordRequest.Record,
		Expr:   expr.String(),
	}

	if ruleRecordRequest.Alert != "" {
		rule.Alert = ruleRecordRequest.Alert
	}

	if ruleRecordRequest.For != "" {
		ruleFor, err := time.ParseDuration(ruleRecordRequest.For)
		if err != nil {
			global.GVA_LOG.Error("Rule.For异常!", zap.Error(err))
			response.FailWithMessage("Rule.For异常", c)
			return
		}
		rule.For = model.Duration(ruleFor)
	}

	if ruleRecordRequest.KeepFiringFor != "" {
		ruleKeepFiringFor, err := time.ParseDuration(ruleRecordRequest.KeepFiringFor)
		if err != nil {
			global.GVA_LOG.Error("Rule.KeepFiringFor异常!", zap.Error(err))
			response.FailWithMessage("Rule.KeepFiringFor异常", c)
			return
		}
		rule.KeepFiringFor = model.Duration(ruleKeepFiringFor)
	}

	if len(ruleRecordRequest.LabelArray) > 0 {
		labelMap := make(map[string]string)
		for _, i := range ruleRecordRequest.LabelArray {
			labelMap[i.Key] = i.Value
		}
		rule.Labels = labelMap
	}

	if len(ruleRecordRequest.AnnotationArray) > 0 {
		annotationMap := make(map[string]string)
		for _, i := range ruleRecordRequest.AnnotationArray {
			annotationMap[i.Key] = i.Value
		}
		rule.Annotations = annotationMap
	}

	_, err = yaml.Marshal(rule)
	if err != nil {
		global.GVA_LOG.Error("Rule序列化失败!", zap.Error(err))
		response.FailWithMessage("Rule序列化失败", c)
		return
	}

	// modRuleRecord
	byteLabels, err := json.Marshal(ruleRecordRequest.LabelArray)
	if err != nil {
		global.GVA_LOG.Error("Rule.Labels序列化失败!", zap.Error(err))
		response.FailWithMessage("Rule.Labels序列化失败", c)
		return
	}

	byteAnnotations, err := json.Marshal(ruleRecordRequest.AnnotationArray)
	if err != nil {
		global.GVA_LOG.Error("Rule.Annotations序列化失败!", zap.Error(err))
		response.FailWithMessage("Rule.Annotations序列化失败", c)
		return
	}

	var clusters []string
	for _, i := range ruleRecordRequest.ClusterArray {
		cluster := strconv.Itoa(i)
		clusters = append(clusters, cluster)
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)

	var ruleRecord = monit.RuleRecord{
		GVA_MODEL:     ruleRecordRequest.GVA_MODEL,
		Name:          ruleRecordRequest.Name,
		RuleGroupId:   ruleRecordRequest.RuleGroupId,
		Clusters:      strings.Join(clusters, ","),
		Manager:       claimsReal.Username,
		Describe:      ruleRecordRequest.Describe,
		Record:        ruleRecordRequest.Record,
		Alert:         ruleRecordRequest.Alert,
		Expr:          expr.String(),
		For:           ruleRecordRequest.For,
		Labels:        string(byteLabels),
		Annotations:   string(byteAnnotations),
		KeepFiringFor: ruleRecordRequest.KeepFiringFor,
	}
	fmt.Println(ruleRecord.RuleGroupId)

	if err := ruleRecordService.UpdateRuleRecord(ruleRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRuleRecord 用id查询规则配置
// @Tags RuleRecord
// @Summary 用id查询规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monit.RuleRecord true "用id查询规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleRecord/findRuleRecord [get]
func (ruleRecordApi *RuleRecordApi) FindRuleRecord(c *gin.Context) {
	var ruleRecord monit.RuleRecord
	err := c.ShouldBindQuery(&ruleRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reruleRecord, err := ruleRecordService.GetRuleRecord(ruleRecord.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var clustersInt []int
	for _, i := range strings.Split(reruleRecord.Clusters, ",") {
		cluster, err := strconv.Atoi(i)
		if err != nil {
			global.GVA_LOG.Error("集群ID异常!", zap.Error(err))
			response.FailWithMessage("集群ID异常", c)
			return
		}

		clustersInt = append(clustersInt, cluster)
	}
	var ruleRecordResp = RuleRecordReqResp{
		RuleRecord:   reruleRecord,
		ClusterArray: clustersInt,
	}

	err = json.Unmarshal([]byte(reruleRecord.Labels), &ruleRecordResp.LabelArray)
	if err != nil {
		global.GVA_LOG.Error("LabelArray反序列化失败!", zap.Error(err))
		response.FailWithMessage("LabelArray反序列化失败", c)
		return
	}

	err = json.Unmarshal([]byte(reruleRecord.Annotations), &ruleRecordResp.AnnotationArray)
	if err != nil {
		global.GVA_LOG.Error("AnnotationArray反序列化失败!", zap.Error(err))
		response.FailWithMessage("AnnotationArray反序列化失败", c)
		return
	}

	response.OkWithData(gin.H{"reruleRecord": ruleRecordResp}, c)
}

// GetRuleRecordList 分页获取规则配置列表
// @Tags RuleRecord
// @Summary 分页获取规则配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monitReq.RuleRecordSearch true "分页获取规则配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleRecord/getRuleRecordList [get]
func (ruleRecordApi *RuleRecordApi) GetRuleRecordList(c *gin.Context) {
	var pageInfo monitReq.RuleRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ruleRecordService.GetRuleRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var ruleRecords []RuleRecordReqResp

	for _, i := range list {

		var clustersInt []int
		for _, i := range strings.Split(i.Clusters, ",") {
			cluster, err := strconv.Atoi(i)
			if err != nil {
				global.GVA_LOG.Error("集群ID异常!", zap.Error(err))
				response.FailWithMessage("集群ID异常", c)
				return
			}

			clustersInt = append(clustersInt, cluster)
		}
		var ruleRecordResp = RuleRecordReqResp{
			RuleRecord:   i,
			ClusterArray: clustersInt,
		}

		err = json.Unmarshal([]byte(i.Labels), &ruleRecordResp.LabelArray)
		if err != nil {
			global.GVA_LOG.Error("Rule.Labels反序列化失败!", zap.Error(err))
			response.FailWithMessage("Rule.Labels反序列化失败", c)
			return
		}

		err = json.Unmarshal([]byte(i.Annotations), &ruleRecordResp.AnnotationArray)
		if err != nil {
			global.GVA_LOG.Error("Rule.Annotations反序列化失败!", zap.Error(err))
			response.FailWithMessage("Rule.Annotations反序列化失败", c)
			return
		}
		ruleRecords = append(ruleRecords, ruleRecordResp)
	}

	response.OkWithDetailed(response.PageResult{
		List:     ruleRecords,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
