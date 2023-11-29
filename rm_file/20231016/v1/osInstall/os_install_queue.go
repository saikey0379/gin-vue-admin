package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osInstallReq "github.com/flipped-aurora/gin-vue-admin/server/model/osInstall/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OsInstallQueueApi struct {
}

var queueService = service.ServiceGroupApp.OsInstallServiceGroup.OsInstallQueueService


// CreateOsInstallQueue 创建操作系统安装队列
// @Tags OsInstallQueue
// @Summary 创建操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallQueue true "创建操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /queue/createOsInstallQueue [post]
func (queueApi *OsInstallQueueApi) CreateOsInstallQueue(c *gin.Context) {
	var queue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&queue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Sn":{utils.NotEmpty()},
        "Hostname":{utils.NotEmpty()},
        "Ip":{utils.NotEmpty()},
        "ManageIp":{utils.NotEmpty()},
        "NetworkId":{utils.NotEmpty()},
        "ManageNetworkId":{utils.NotEmpty()},
        "HardwareId":{utils.NotEmpty()},
        "PxeId":{utils.NotEmpty()},
        "KickstartId":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
        "Manager":{utils.NotEmpty()},
    }
	if err := utils.Verify(queue, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := queueService.CreateOsInstallQueue(&queue); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOsInstallQueue 删除操作系统安装队列
// @Tags OsInstallQueue
// @Summary 删除操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallQueue true "删除操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /queue/deleteOsInstallQueue [delete]
func (queueApi *OsInstallQueueApi) DeleteOsInstallQueue(c *gin.Context) {
	var queue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&queue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := queueService.DeleteOsInstallQueue(queue); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOsInstallQueueByIds 批量删除操作系统安装队列
// @Tags OsInstallQueue
// @Summary 批量删除操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /queue/deleteOsInstallQueueByIds [delete]
func (queueApi *OsInstallQueueApi) DeleteOsInstallQueueByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := queueService.DeleteOsInstallQueueByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOsInstallQueue 更新操作系统安装队列
// @Tags OsInstallQueue
// @Summary 更新操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osInstall.OsInstallQueue true "更新操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /queue/updateOsInstallQueue [put]
func (queueApi *OsInstallQueueApi) UpdateOsInstallQueue(c *gin.Context) {
	var queue osInstall.OsInstallQueue
	err := c.ShouldBindJSON(&queue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Sn":{utils.NotEmpty()},
          "Hostname":{utils.NotEmpty()},
          "Ip":{utils.NotEmpty()},
          "ManageIp":{utils.NotEmpty()},
          "NetworkId":{utils.NotEmpty()},
          "ManageNetworkId":{utils.NotEmpty()},
          "HardwareId":{utils.NotEmpty()},
          "PxeId":{utils.NotEmpty()},
          "KickstartId":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
          "Manager":{utils.NotEmpty()},
      }
    if err := utils.Verify(queue, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := queueService.UpdateOsInstallQueue(queue); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOsInstallQueue 用id查询操作系统安装队列
// @Tags OsInstallQueue
// @Summary 用id查询操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstall.OsInstallQueue true "用id查询操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /queue/findOsInstallQueue [get]
func (queueApi *OsInstallQueueApi) FindOsInstallQueue(c *gin.Context) {
	var queue osInstall.OsInstallQueue
	err := c.ShouldBindQuery(&queue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if requeue, err := queueService.GetOsInstallQueue(queue.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requeue": requeue}, c)
	}
}

// GetOsInstallQueueList 分页获取操作系统安装队列列表
// @Tags OsInstallQueue
// @Summary 分页获取操作系统安装队列列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osInstallReq.OsInstallQueueSearch true "分页获取操作系统安装队列列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /queue/getOsInstallQueueList [get]
func (queueApi *OsInstallQueueApi) GetOsInstallQueueList(c *gin.Context) {
	var pageInfo osInstallReq.OsInstallQueueSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := queueService.GetOsInstallQueueInfoList(pageInfo); err != nil {
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
