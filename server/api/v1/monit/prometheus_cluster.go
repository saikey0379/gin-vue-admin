package monit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/known"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
	monitReq "github.com/flipped-aurora/gin-vue-admin/server/model/monit/request"
	request1 "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type PrometheusClusterApi struct {
}

var prometheusClusterService = service.ServiceGroupApp.MonitServiceGroup.PrometheusClusterService

type PromRulesResponse struct {
	ID                 int    `json:"ID"`
	Content            string `json:"content"`
	TotalGroups        int    `json:"totalGroups"`
	TotalRules         int    `json:"totalRules"`
	TotalCurrentGroups int    `json:"totalCurrentGroups"`
	TotalCurrentRules  int    `json:"totalCurrentRules"`
}

// CreatePrometheusCluster 创建监控集群
// @Tags PrometheusCluster
// @Summary 创建监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.PrometheusCluster true "创建监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /prometheusCluster/createPrometheusCluster [post]
func (prometheusClusterApi *PrometheusClusterApi) CreatePrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindJSON(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
		"SshUser":  {utils.NotEmpty()},
		"SshRsa":   {utils.NotEmpty()},
		"Nodes":    {utils.NotEmpty()},
		"PathConf": {utils.NotEmpty()},
		"ExecLoad": {utils.NotEmpty()},
	}
	if err := utils.Verify(prometheusCluster, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	prometheusCluster.Manager = claimsReal.Username

	if err := prometheusClusterService.CreatePrometheusCluster(&prometheusCluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePrometheusCluster 删除监控集群
// @Tags PrometheusCluster
// @Summary 删除监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.PrometheusCluster true "删除监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prometheusCluster/deletePrometheusCluster [delete]
func (prometheusClusterApi *PrometheusClusterApi) DeletePrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindJSON(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := prometheusClusterService.DeletePrometheusCluster(prometheusCluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePrometheusClusterByIds 批量删除监控集群
// @Tags PrometheusCluster
// @Summary 批量删除监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /prometheusCluster/deletePrometheusClusterByIds [delete]
func (prometheusClusterApi *PrometheusClusterApi) DeletePrometheusClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := prometheusClusterService.DeletePrometheusClusterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePrometheusCluster 更新监控集群
// @Tags PrometheusCluster
// @Summary 更新监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body monit.PrometheusCluster true "更新监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /prometheusCluster/updatePrometheusCluster [put]
func (prometheusClusterApi *PrometheusClusterApi) UpdatePrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindJSON(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
		"SshUser":  {utils.NotEmpty()},
		"SshRsa":   {utils.NotEmpty()},
		"Nodes":    {utils.NotEmpty()},
		"PathConf": {utils.NotEmpty()},
		"ExecLoad": {utils.NotEmpty()},
	}
	if err := utils.Verify(prometheusCluster, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	prometheusCluster.Manager = claimsReal.Username

	if err := prometheusClusterService.UpdatePrometheusCluster(prometheusCluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPrometheusCluster 用id查询监控集群
// @Tags PrometheusCluster
// @Summary 用id查询监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monit.PrometheusCluster true "用id查询监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prometheusCluster/findPrometheusCluster [get]
func (prometheusClusterApi *PrometheusClusterApi) FindPrometheusCluster(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reprometheusCluster, err := prometheusClusterService.GetPrometheusCluster(prometheusCluster.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprometheusCluster": reprometheusCluster}, c)
	}
}

// GetPrometheusClusterList 分页获取监控集群列表
// @Tags PrometheusCluster
// @Summary 分页获取监控集群列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query monitReq.PrometheusClusterSearch true "分页获取监控集群列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prometheusCluster/getPrometheusClusterList [get]
func (prometheusClusterApi *PrometheusClusterApi) GetPrometheusClusterList(c *gin.Context) {
	var pageInfo monitReq.PrometheusClusterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := prometheusClusterService.GetPrometheusClusterInfoList(pageInfo); err != nil {
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

func (prometheusClusterApi *PrometheusClusterApi) FindPromYamlListByClusterId(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ruleGroups, _, err := ruleGroupService.GetRuleGroupInfoList(monitReq.RuleGroupSearch{})
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	type PromYamlRuleGroup struct {
		ID      uint   `json:"ID"`
		Name    string `json:"name"`
		Content string `json:"content"`
		Total   int    `json:"totalRules"`
	}

	type PromYamlCluster struct {
		ID     int                 `json:"ID"`
		Groups []PromYamlRuleGroup `json:"groups"`
		Total  int64               `json:"totalGroups"`
	}
	var promYamlCluster = PromYamlCluster{
		ID: int(prometheusCluster.ID),
	}

	for _, ruleGroup := range ruleGroups {
		var pageInfo monitReq.RuleRecordSearch
		pageInfo.Clusters = strconv.Itoa(int(prometheusCluster.ID))
		ruleGroupId := int(ruleGroup.ID)
		pageInfo.RuleGroupId = &ruleGroupId

		rules, _, err := ruleRecordService.GetRuleRecordInfoList(pageInfo)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
			return
		}

		var promRules []rulefmt.Rule
		for _, rule := range rules {
			var promRule = rulefmt.Rule{
				Record: rule.Record,
				Alert:  rule.Alert,
				Expr:   rule.Expr,
			}

			if rule.For != "" {
				ruleFor, err := time.ParseDuration(rule.For)
				if err != nil {
					global.GVA_LOG.Error("Rule.For异常!", zap.Error(err))
					response.FailWithMessage("Rule.For异常", c)
					return
				}
				promRule.For = model.Duration(ruleFor)
			}
			if rule.KeepFiringFor != "" {
				ruleKeepFiringFor, err := time.ParseDuration(rule.KeepFiringFor)
				if err != nil {
					global.GVA_LOG.Error("Rule.KeepFiringFor异常!", zap.Error(err))
					response.FailWithMessage("Rule.KeepFiringFor异常", c)
					return
				}
				promRule.KeepFiringFor = model.Duration(ruleKeepFiringFor)
			}

			var labelArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Labels), &labelArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Labels反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Labels反序列化失败", c)
				return
			}

			var annotationArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Annotations), &annotationArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Annotations反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Annotations反序列化失败", c)
				return
			}

			if len(labelArray) > 0 {
				labelMap := make(map[string]string)
				for _, i := range labelArray {
					labelMap[i.Key] = i.Value
				}
				promRule.Labels = labelMap
			}

			if len(annotationArray) > 0 {
				annotationMap := make(map[string]string)
				for _, i := range annotationArray {
					annotationMap[i.Key] = i.Value
				}
				promRule.Annotations = annotationMap
			}
			promRules = append(promRules, promRule)
		}

		if len(promRules) > 0 {
			promYamlCluster.Total += 1

			var promYamlRuleGroup = PromYamlRuleGroup{
				ID:    ruleGroup.ID,
				Name:  ruleGroup.Name,
				Total: len(promRules),
			}

			promRuleGroup := utils.RuleGroup{
				Name:  ruleGroup.Name,
				Rules: promRules,
			}

			if ruleGroup.Limit != nil {
				promRuleGroup.Limit = *ruleGroup.Limit
			}

			if ruleGroup.Interval != "" {
				ruleGroupInterval, err := time.ParseDuration(ruleGroup.Interval)
				if err != nil {
					global.GVA_LOG.Error("RuleGroup.Interval异常!", zap.Error(err))
					response.FailWithMessage("RuleGroup.Interval异常", c)
					return
				}
				promRuleGroup.Interval = model.Duration(ruleGroupInterval)
			}

			promRuleGroupByte, err := yaml.Marshal(promRuleGroup)
			if err != nil {
				global.GVA_LOG.Error("RuleGroups序列化失败!", zap.Error(err))
				response.FailWithMessage("RuleGroups序列化失败", c)
				return
			}

			promYamlRuleGroup.Content = string(promRuleGroupByte)

			promYamlCluster.Groups = append(promYamlCluster.Groups, promYamlRuleGroup)
		}
	}

	response.OkWithData(gin.H{"promYamlCluster": promYamlCluster}, c)
}

func (prometheusClusterApi *PrometheusClusterApi) FindPromYamlDetailByClusterId(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ruleGroups, _, err := ruleGroupService.GetRuleGroupInfoList(monitReq.RuleGroupSearch{})
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var promYamlCluster = PromRulesResponse{
		ID: int(prometheusCluster.ID),
	}

	var promYamlRuleGroups utils.RuleGroups
	for _, ruleGroup := range ruleGroups {
		var pageInfo monitReq.RuleRecordSearch
		pageInfo.Clusters = strconv.Itoa(int(prometheusCluster.ID))
		ruleGroupId := int(ruleGroup.ID)
		pageInfo.RuleGroupId = &ruleGroupId

		rules, _, err := ruleRecordService.GetRuleRecordInfoList(pageInfo)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
			return
		}

		var promRules []rulefmt.Rule
		for _, rule := range rules {
			var promRule = rulefmt.Rule{
				Record: rule.Record,
				Alert:  rule.Alert,
				Expr:   rule.Expr,
			}

			if rule.For != "" {
				ruleFor, err := time.ParseDuration(rule.For)
				if err != nil {
					global.GVA_LOG.Error("Rule.For异常!", zap.Error(err))
					response.FailWithMessage("Rule.For异常", c)
					return
				}
				promRule.For = model.Duration(ruleFor)
			}
			if rule.KeepFiringFor != "" {
				ruleKeepFiringFor, err := time.ParseDuration(rule.KeepFiringFor)
				if err != nil {
					global.GVA_LOG.Error("Rule.KeepFiringFor异常!", zap.Error(err))
					response.FailWithMessage("Rule.KeepFiringFor异常", c)
					return
				}
				promRule.KeepFiringFor = model.Duration(ruleKeepFiringFor)
			}

			var labelArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Labels), &labelArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Labels反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Labels反序列化失败", c)
				return
			}

			var annotationArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Annotations), &annotationArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Annotations反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Annotations反序列化失败", c)
				return
			}

			if len(labelArray) > 0 {
				labelMap := make(map[string]string)
				for _, i := range labelArray {
					labelMap[i.Key] = i.Value
				}
				promRule.Labels = labelMap
			}

			if len(annotationArray) > 0 {
				annotationMap := make(map[string]string)
				for _, i := range annotationArray {
					annotationMap[i.Key] = i.Value
				}
				promRule.Annotations = annotationMap
			}
			promRules = append(promRules, promRule)
		}

		if len(promRules) > 0 {
			promYamlCluster.TotalGroups += 1
			promYamlCluster.TotalRules += len(promRules)

			promRuleGroup := utils.RuleGroup{
				Name:  ruleGroup.Name,
				Rules: promRules,
			}

			if ruleGroup.Limit != nil {
				promRuleGroup.Limit = *ruleGroup.Limit
			}

			if ruleGroup.Interval != "" {
				ruleGroupInterval, err := time.ParseDuration(ruleGroup.Interval)
				if err != nil {
					global.GVA_LOG.Error("RuleGroup.Interval异常!", zap.Error(err))
					response.FailWithMessage("RuleGroup.Interval异常", c)
					return
				}
				promRuleGroup.Interval = model.Duration(ruleGroupInterval)
			}

			promYamlRuleGroups.Groups = append(promYamlRuleGroups.Groups, promRuleGroup)
		}
	}

	promYamlRuleGroupsByte, err := yaml.Marshal(promYamlRuleGroups)
	if err != nil {
		global.GVA_LOG.Error("RuleGroups序列化失败!", zap.Error(err))
		response.FailWithMessage("RuleGroups序列化失败", c)
		return
	}

	promYamlCluster.Content = string(promYamlRuleGroupsByte)
	response.OkWithData(gin.H{"promYamlCluster": promYamlCluster}, c)
}

func (prometheusClusterApi *PrometheusClusterApi) FindPromYamlDetailCurrentByClusterId(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reprometheusCluster, err := prometheusClusterService.GetPrometheusCluster(prometheusCluster.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var promYamlCluster = PromRulesResponse{
		ID: int(prometheusCluster.ID),
	}

	// GetCurrentPromRuleYaml
	node := strings.Split(reprometheusCluster.Nodes, ":")
	port, err := strconv.Atoi(node[1])
	if err != nil {
		global.GVA_LOG.Error("Node信息异常!", zap.Error(err))
		response.FailWithMessage("Node信息异常", c)
		return
	}
	var sshClient = &utils.SSHClient{
		PrivateKey: known.ReturnFileDir("rsa", "") + reprometheusCluster.SshRsa,
		Address:    node[0],
		Port:       port,
		User:       reprometheusCluster.SshUser,
	}
	err = sshClient.CreateClient()
	if err != nil {
		global.GVA_LOG.Error("服务器连接失败!", zap.Error(err))
		response.FailWithMessage("服务器连接失败", c)
		return
	}

	promRuleYamlTargetFileName := fmt.Sprintf("prometheus-rules_%d.yaml", reprometheusCluster.ID)

	cmdGetCurrentPromRulesYaml := fmt.Sprintf("/usr/bin/cat %s/%s", reprometheusCluster.PathConf, promRuleYamlTargetFileName)
	promRuleGroupsByteCurrent, err := sshClient.RunShell(cmdGetCurrentPromRulesYaml)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("获取当前配置异常![%s]", promRuleGroupsByteCurrent), zap.Error(err))
		response.FailWithMessage("获取当前配置异常", c)
		return
	}

	var promRuleGroupsCurrent utils.RuleGroups
	err = yaml.Unmarshal(promRuleGroupsByteCurrent, &promRuleGroupsCurrent)
	if err != nil {
		global.GVA_LOG.Error("当前配置解析异常!", zap.Error(err))
		response.FailWithMessage("当前配置解析异常", c)
		return
	}

	promYamlCluster.TotalGroups = len(promRuleGroupsCurrent.Groups)
	for _, i := range promRuleGroupsCurrent.Groups {
		promYamlCluster.TotalRules += len(i.Rules)
	}

	promYamlCluster.Content = string(promRuleGroupsByteCurrent)
	response.OkWithData(gin.H{"promYamlCluster": promYamlCluster}, c)
}

func (prometheusClusterApi *PrometheusClusterApi) DiffPromYamlByClusterId(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reprometheusCluster, err := prometheusClusterService.GetPrometheusCluster(prometheusCluster.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var promYamlCluster = PromRulesResponse{
		ID: int(prometheusCluster.ID),
	}

	// GetCurrentPromRuleYaml
	node := strings.Split(reprometheusCluster.Nodes, ":")
	port, err := strconv.Atoi(node[1])
	if err != nil {
		global.GVA_LOG.Error("Node信息异常!", zap.Error(err))
		response.FailWithMessage("Node信息异常", c)
		return
	}
	var sshClient = &utils.SSHClient{
		PrivateKey: known.ReturnFileDir("rsa", "") + reprometheusCluster.SshRsa,
		Address:    node[0],
		Port:       port,
		User:       reprometheusCluster.SshUser,
	}
	err = sshClient.CreateClient()
	if err != nil {
		global.GVA_LOG.Error("服务器连接失败!", zap.Error(err))
		response.FailWithMessage("服务器连接失败", c)
		return
	}
	promRuleYamlTargetFileName := fmt.Sprintf("prometheus-rules_%d.yaml", reprometheusCluster.ID)

	cmdGetCurrentPromRulesYaml := fmt.Sprintf("/usr/bin/cat %s/%s", reprometheusCluster.PathConf, promRuleYamlTargetFileName)
	promRuleGroupsByteCurrent, err := sshClient.RunShell(cmdGetCurrentPromRulesYaml)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("获取当前配置异常![%s]", promRuleGroupsByteCurrent), zap.Error(err))
		response.FailWithMessage("获取当前配置异常", c)
		return
	}

	var promRuleGroupsCurrent utils.RuleGroups
	err = yaml.Unmarshal(promRuleGroupsByteCurrent, &promRuleGroupsCurrent)
	if err != nil {
		global.GVA_LOG.Error("当前配置解析异常!", zap.Error(err))
		response.FailWithMessage("当前配置解析异常", c)
		return
	}

	promYamlCluster.TotalCurrentGroups = len(promRuleGroupsCurrent.Groups)
	for _, i := range promRuleGroupsCurrent.Groups {
		promYamlCluster.TotalCurrentRules += len(i.Rules)
	}

	promRuleGroupsYamlCurrent := string(promRuleGroupsByteCurrent)
	// GetTargetPromYaml
	ruleGroups, _, err := ruleGroupService.GetRuleGroupInfoList(monitReq.RuleGroupSearch{})
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var promRuleGroupsTarget utils.RuleGroups
	for _, ruleGroup := range ruleGroups {
		var pageInfo monitReq.RuleRecordSearch
		pageInfo.Clusters = strconv.Itoa(int(prometheusCluster.ID))
		ruleGroupId := int(ruleGroup.ID)
		pageInfo.RuleGroupId = &ruleGroupId

		rules, _, err := ruleRecordService.GetRuleRecordInfoList(pageInfo)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
			return
		}

		var promRules []rulefmt.Rule
		for _, rule := range rules {
			var promRule = rulefmt.Rule{
				Record: rule.Record,
				Alert:  rule.Alert,
				Expr:   rule.Expr,
			}

			if rule.For != "" {
				ruleFor, err := time.ParseDuration(rule.For)
				if err != nil {
					global.GVA_LOG.Error("Rule.For异常!", zap.Error(err))
					response.FailWithMessage("Rule.For异常", c)
					return
				}
				promRule.For = model.Duration(ruleFor)
			}
			if rule.KeepFiringFor != "" {
				ruleKeepFiringFor, err := time.ParseDuration(rule.KeepFiringFor)
				if err != nil {
					global.GVA_LOG.Error("Rule.KeepFiringFor异常!", zap.Error(err))
					response.FailWithMessage("Rule.KeepFiringFor异常", c)
					return
				}
				promRule.KeepFiringFor = model.Duration(ruleKeepFiringFor)
			}

			var labelArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Labels), &labelArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Labels反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Labels反序列化失败", c)
				return
			}

			var annotationArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Annotations), &annotationArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Annotations反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Annotations反序列化失败", c)
				return
			}

			if len(labelArray) > 0 {
				labelMap := make(map[string]string)
				for _, i := range labelArray {
					labelMap[i.Key] = i.Value
				}
				promRule.Labels = labelMap
			}

			if len(annotationArray) > 0 {
				annotationMap := make(map[string]string)
				for _, i := range annotationArray {
					annotationMap[i.Key] = i.Value
				}
				promRule.Annotations = annotationMap
			}
			promRules = append(promRules, promRule)
		}

		if len(promRules) > 0 {
			promYamlCluster.TotalGroups += 1
			promYamlCluster.TotalRules += len(promRules)
			promRules = utils.SortRulesByName(promRules)

			promRuleGroup := utils.RuleGroup{
				Name:  ruleGroup.Name,
				Rules: promRules,
			}

			if ruleGroup.Limit != nil {
				promRuleGroup.Limit = *ruleGroup.Limit
			}

			if ruleGroup.Interval != "" {
				ruleGroupInterval, err := time.ParseDuration(ruleGroup.Interval)
				if err != nil {
					global.GVA_LOG.Error("RuleGroup.Interval异常!", zap.Error(err))
					response.FailWithMessage("RuleGroup.Interval异常", c)
					return
				}
				promRuleGroup.Interval = model.Duration(ruleGroupInterval)
			}

			promRuleGroupsTarget.Groups = append(promRuleGroupsTarget.Groups, promRuleGroup)
		}
	}

	promRuleGroupsByteTarget, err := yaml.Marshal(promRuleGroupsTarget)
	if err != nil {
		global.GVA_LOG.Error("RuleGroups序列化失败!", zap.Error(err))
		response.FailWithMessage("RuleGroups序列化失败", c)
		return
	}

	promRuleGroupsYamlTarget := string(promRuleGroupsByteTarget)

	diff := cmp.Diff(promRuleGroupsYamlCurrent, promRuleGroupsYamlTarget)

	promYamlCluster.Content = string(diff)
	global.GVA_LOG.Debug(fmt.Sprintf("DiffPrometheusRules![%s]", promYamlCluster.Content))
	response.OkWithData(gin.H{"promYamlCluster": promYamlCluster}, c)
}

func (prometheusClusterApi *PrometheusClusterApi) UpdatePrometheusRulesByClusterId(c *gin.Context) {
	var prometheusCluster monit.PrometheusCluster
	err := c.ShouldBindQuery(&prometheusCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	reprometheusCluster, err := prometheusClusterService.GetPrometheusCluster(prometheusCluster.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var promYamlCluster = PromRulesResponse{
		ID: int(prometheusCluster.ID),
	}

	// GetCurrentPromRuleYaml
	node := strings.Split(reprometheusCluster.Nodes, ":")
	port, err := strconv.Atoi(node[1])
	if err != nil {
		global.GVA_LOG.Error("Node信息异常!", zap.Error(err))
		response.FailWithMessage("Node信息异常", c)
		return
	}
	var sshClient = &utils.SSHClient{
		PrivateKey: known.ReturnFileDir("rsa", "") + reprometheusCluster.SshRsa,
		Address:    node[0],
		Port:       port,
		User:       reprometheusCluster.SshUser,
	}
	err = sshClient.CreateClient()
	if err != nil {
		global.GVA_LOG.Error("服务器连接失败!", zap.Error(err))
		response.FailWithMessage("服务器连接失败", c)
		return
	}

	promRuleYamlTargetFileName := fmt.Sprintf("prometheus-rules_%d.yaml", reprometheusCluster.ID)

	cmdGetCurrentPromRulesYaml := fmt.Sprintf("/usr/bin/cat %s/%s", reprometheusCluster.PathConf, promRuleYamlTargetFileName)
	global.GVA_LOG.Debug(cmdGetCurrentPromRulesYaml)
	fmt.Println(cmdGetCurrentPromRulesYaml)

	promRuleGroupsByteCurrent, err := sshClient.RunShell(cmdGetCurrentPromRulesYaml)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("获取当前配置异常![%s]", promRuleGroupsByteCurrent), zap.Error(err))
		response.FailWithMessage("获取当前配置异常", c)
		return
	}

	promRuleYamlBakDir := known.ReturnFileBakDir("config", "prometheus-rules")
	promRuleYamlBakContent := string(promRuleGroupsByteCurrent)
	promRuleYamlBakFileName := fmt.Sprintf("prometheus-rules_%d_%s_%d.yaml", reprometheusCluster.ID, reprometheusCluster.Name, time.Now().Unix())

	err = utils.CreateFile(promRuleYamlBakDir, promRuleYamlBakFileName, promRuleYamlBakContent)
	if err != nil {
		global.GVA_LOG.Error("配置备份失败!", zap.Error(err))
		response.FailWithMessage("配置备份失败", c)
		return
	}

	// GetTargetPromYaml
	ruleGroups, _, err := ruleGroupService.GetRuleGroupInfoList(monitReq.RuleGroupSearch{})
	if err != nil {
		global.GVA_LOG.Error("获取分组信息失败!", zap.Error(err))
		response.FailWithMessage("获取分组信息失败", c)
		return
	}

	var promRuleGroupsTarget utils.RuleGroups
	for _, ruleGroup := range ruleGroups {
		var pageInfo monitReq.RuleRecordSearch
		pageInfo.Clusters = strconv.Itoa(int(prometheusCluster.ID))
		ruleGroupId := int(ruleGroup.ID)
		pageInfo.RuleGroupId = &ruleGroupId

		rules, _, err := ruleRecordService.GetRuleRecordInfoList(pageInfo)
		if err != nil {
			global.GVA_LOG.Error("获取规则信息失败!", zap.Error(err))
			response.FailWithMessage("获取规则信息失败", c)
			return
		}

		var promRules []rulefmt.Rule
		for _, rule := range rules {
			var promRule = rulefmt.Rule{
				Record: rule.Record,
				Alert:  rule.Alert,
				Expr:   rule.Expr,
			}

			if rule.For != "" {
				ruleFor, err := time.ParseDuration(rule.For)
				if err != nil {
					global.GVA_LOG.Error("Rule.For异常!", zap.Error(err))
					response.FailWithMessage("Rule.For异常", c)
					return
				}
				promRule.For = model.Duration(ruleFor)
			}
			if rule.KeepFiringFor != "" {
				ruleKeepFiringFor, err := time.ParseDuration(rule.KeepFiringFor)
				if err != nil {
					global.GVA_LOG.Error("Rule.KeepFiringFor异常!", zap.Error(err))
					response.FailWithMessage("Rule.KeepFiringFor异常", c)
					return
				}
				promRule.KeepFiringFor = model.Duration(ruleKeepFiringFor)
			}

			var labelArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Labels), &labelArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Labels反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Labels反序列化失败", c)
				return
			}

			var annotationArray []utils.RuleMap
			err = json.Unmarshal([]byte(rule.Annotations), &annotationArray)
			if err != nil {
				global.GVA_LOG.Error("Rule.Annotations反序列化失败!", zap.Error(err))
				response.FailWithMessage("Rule.Annotations反序列化失败", c)
				return
			}

			if len(labelArray) > 0 {
				labelMap := make(map[string]string)
				for _, i := range labelArray {
					labelMap[i.Key] = i.Value
				}
				promRule.Labels = labelMap
			}

			if len(annotationArray) > 0 {
				annotationMap := make(map[string]string)
				for _, i := range annotationArray {
					annotationMap[i.Key] = i.Value
				}
				promRule.Annotations = annotationMap
			}
			promRules = append(promRules, promRule)
		}

		if len(promRules) > 0 {
			promYamlCluster.TotalGroups += 1
			promYamlCluster.TotalRules += len(promRules)
			promRules = utils.SortRulesByName(promRules)

			promRuleGroup := utils.RuleGroup{
				Name:  ruleGroup.Name,
				Rules: promRules,
			}

			if ruleGroup.Limit != nil {
				promRuleGroup.Limit = *ruleGroup.Limit
			}

			if ruleGroup.Interval != "" {
				ruleGroupInterval, err := time.ParseDuration(ruleGroup.Interval)
				if err != nil {
					global.GVA_LOG.Error("RuleGroup.Interval异常!", zap.Error(err))
					response.FailWithMessage("RuleGroup.Interval异常", c)
					return
				}
				promRuleGroup.Interval = model.Duration(ruleGroupInterval)
			}

			promRuleGroupsTarget.Groups = append(promRuleGroupsTarget.Groups, promRuleGroup)
		}
	}

	promRuleGroupsByteTarget, err := yaml.Marshal(promRuleGroupsTarget)
	if err != nil {
		global.GVA_LOG.Error("RuleGroups序列化失败!", zap.Error(err))
		response.FailWithMessage("RuleGroups序列化失败", c)
		return
	}

	promRuleYamlTargetLocalDir := known.ReturnFileDir("config", "prometheus-rules")
	promRuleYamlTargetContent := string(promRuleGroupsByteTarget)

	err = utils.CreateFile(promRuleYamlTargetLocalDir, promRuleYamlTargetFileName, promRuleYamlTargetContent)
	if err != nil {
		global.GVA_LOG.Error("配置生成失败!", zap.Error(err))
		response.FailWithMessage("配置生成失败", c)
		return
	}

	output, err := sshClient.FileSync(fmt.Sprintf("%s/%s", promRuleYamlTargetLocalDir, promRuleYamlTargetFileName), reprometheusCluster.PathConf, promRuleYamlTargetFileName)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("配置同步失败![%s]", string(output)), zap.Error(err))
		response.FailWithMessage("配置同步失败", c)
		return
	}

	global.GVA_LOG.Debug(string(output))

	response.OkWithMessage("更新成功", c)
}
