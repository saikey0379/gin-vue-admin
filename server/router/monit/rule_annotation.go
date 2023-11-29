package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RuleAnnotationRouter struct {
}

// InitRuleAnnotationRouter 初始化 注解标签 路由信息
func (s *RuleAnnotationRouter) InitRuleAnnotationRouter(Router *gin.RouterGroup) {
	ruleAnnotationRouter := Router.Group("ruleAnnotation").Use(middleware.OperationRecord())
	ruleAnnotationRouterWithoutRecord := Router.Group("ruleAnnotation")
	var ruleAnnotationApi = v1.ApiGroupApp.MonitApiGroup.RuleAnnotationApi
	{
		ruleAnnotationRouter.POST("createRuleAnnotation", ruleAnnotationApi.CreateRuleAnnotation)   // 新建注解标签
		ruleAnnotationRouter.DELETE("deleteRuleAnnotation", ruleAnnotationApi.DeleteRuleAnnotation) // 删除注解标签
		ruleAnnotationRouter.DELETE("deleteRuleAnnotationByIds", ruleAnnotationApi.DeleteRuleAnnotationByIds) // 批量删除注解标签
		ruleAnnotationRouter.PUT("updateRuleAnnotation", ruleAnnotationApi.UpdateRuleAnnotation)    // 更新注解标签
	}
	{
		ruleAnnotationRouterWithoutRecord.GET("findRuleAnnotation", ruleAnnotationApi.FindRuleAnnotation)        // 根据ID获取注解标签
		ruleAnnotationRouterWithoutRecord.GET("getRuleAnnotationList", ruleAnnotationApi.GetRuleAnnotationList)  // 获取注解标签列表
	}
}
