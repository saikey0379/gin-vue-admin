package initialize

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath))

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitChatGptRouter(PrivateGroup)

		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)

	}
	{
		osInstallRouter := router.RouterGroupApp.OsInstall

		osInstallRouter.InitOsInstallConfigPxeRouter(PrivateGroup)

		osInstallRouter.InitOsInstallConfigKickstartRouter(PrivateGroup)
		osInstallRouter.InitOsInstallConfigKickstartRouterPublic(PublicGroup)

		osInstallRouter.InitOsInstallQueueRouter(PrivateGroup)
		osInstallRouter.InitOsInstallQueueRouterPublic(PublicGroup)

		osInstallRouter.InitOsInstallLogRouterPublic(PublicGroup)
		osInstallRouter.InitOsInstallLogRouter(PrivateGroup)

	}
	{
		idcRouter := router.RouterGroupApp.Idc

		idcRouter.InitIdcInfoRouter(PrivateGroup)

		idcRouter.InitIdcCabinetRouter(PrivateGroup)
		idcRouter.InitIdcRoomRouter(PrivateGroup)
		idcRouter.InitIdcIpSegmentRouter(PrivateGroup)

		idcRouter.InitIdcIpSubnetRouter(PrivateGroup)
		idcRouter.InitIdcIpSubnetRouterPublic(PublicGroup)

		idcRouter.InitIdcIpPreemptRouter(PrivateGroup)
	}
	{
		deviceRouter := router.RouterGroupApp.Device

		deviceRouter.InitDeviceBareMetalRouter(PrivateGroup)
	}
	{
		serverRouter := router.RouterGroupApp.Server

		serverRouter.InitServerDiscoveryRouter(PrivateGroup)
		serverRouter.InitServerInfoRouter(PrivateGroup)
		serverRouter.InitServerInfoRouterPublic(PublicGroup)

	}
	{
		fileRouter := router.RouterGroupApp.File

		fileRouter.InitFileConfigRouter(PrivateGroup)
		fileRouter.InitFileConfigRouterPublic(PublicGroup)

		fileRouter.InitFileBinaryRouter(PrivateGroup)
		fileRouter.InitFileCommonRouter(PrivateGroup)
		fileRouter.InitFileCommandRouter(PrivateGroup)

		fileRouter.InitFileScriptRouter(PrivateGroup)
		fileRouter.InitFileScriptRouterPublic(PublicGroup)
		fileRouter.InitFileCommadBlacklistRouter(PrivateGroup)

	}
	{
		taskRouter := router.RouterGroupApp.Task

		taskRouter.InitTaskParallelRouter(PrivateGroup)

		taskRouter.InitTaskTimedRouter(PrivateGroup)
		taskRouter.InitTaskCrontabRouter(PrivateGroup)
		taskRouter.InitTaskRealtimeRouter(PrivateGroup)
	}
	{
		slbRouter := router.RouterGroupApp.Slb
		slbRouter.InitSlbClusterRouter(PrivateGroup)
		slbRouter.InitSlbCertRouter(PrivateGroup)
		slbRouter.InitSlbAccesslistRouter(PrivateGroup)
		slbRouter.InitSlbUpstreamRouter(PrivateGroup)
		slbRouter.InitSlbDomainRouter(PrivateGroup)
	}
	{
		cmdbRouter := router.RouterGroupApp.Cmdb

		cmdbRouter.InitCmdbIpSegmentRouter(PrivateGroup)
		cmdbRouter.InitCmdbIpSubnetRouter(PrivateGroup)
		cmdbRouter.InitCmdbIpSubnetRouterPublic(PublicGroup)

		cmdbRouter.InitCmdbIpPreemptRouter(PrivateGroup)
		cmdbRouter.InitCmdbRegionRouter(PrivateGroup)
	}
	{
		monitRouter := router.RouterGroupApp.Monit

		monitRouter.InitRuleLabelRouter(PrivateGroup)
		monitRouter.InitRuleAnnotationRouter(PrivateGroup)

		monitRouter.InitRuleGroupRouter(PrivateGroup)
		monitRouter.InitPrometheusClusterRouter(PrivateGroup)
		monitRouter.InitRuleRecordRouter(PrivateGroup)

	}
	{
		fileHandleRouter := router.RouterGroupApp.FileHandle

		fileHandleRouter.InitFileResumeTramsRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
