package osManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/osManage"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    osManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/osManage/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PxeApi struct {
}

var PxeConfigService = service.ServiceGroupApp.OsManageServiceGroup.PxeService


// CreatePxe 创建PXE配置
// @Tags Pxe
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osManage.Pxe true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PxeConfig/createPxe [post]
func (PxeConfigApi *PxeApi) CreatePxe(c *gin.Context) {
	var PxeConfig osManage.Pxe
	err := c.ShouldBindJSON(&PxeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Config":{utils.NotEmpty()},
    }
	if err := utils.Verify(PxeConfig, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := PxeConfigService.CreatePxe(&PxeConfig); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePxe 删除PXE配置
// @Tags Pxe
// @Summary 删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osManage.Pxe true "删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PxeConfig/deletePxe [delete]
func (PxeConfigApi *PxeApi) DeletePxe(c *gin.Context) {
	var PxeConfig osManage.Pxe
	err := c.ShouldBindJSON(&PxeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PxeConfigService.DeletePxe(PxeConfig); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePxeByIds 批量删除PXE配置
// @Tags Pxe
// @Summary 批量删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PxeConfig/deletePxeByIds [delete]
func (PxeConfigApi *PxeApi) DeletePxeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PxeConfigService.DeletePxeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePxe 更新PXE配置
// @Tags Pxe
// @Summary 更新PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body osManage.Pxe true "更新PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PxeConfig/updatePxe [put]
func (PxeConfigApi *PxeApi) UpdatePxe(c *gin.Context) {
	var PxeConfig osManage.Pxe
	err := c.ShouldBindJSON(&PxeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Config":{utils.NotEmpty()},
      }
    if err := utils.Verify(PxeConfig, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := PxeConfigService.UpdatePxe(PxeConfig); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPxe 用id查询PXE配置
// @Tags Pxe
// @Summary 用id查询PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osManage.Pxe true "用id查询PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PxeConfig/findPxe [get]
func (PxeConfigApi *PxeApi) FindPxe(c *gin.Context) {
	var PxeConfig osManage.Pxe
	err := c.ShouldBindQuery(&PxeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rePxeConfig, err := PxeConfigService.GetPxe(PxeConfig.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePxeConfig": rePxeConfig}, c)
	}
}

// GetPxeList 分页获取PXE配置列表
// @Tags Pxe
// @Summary 分页获取PXE配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query osManageReq.PxeSearch true "分页获取PXE配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PxeConfig/getPxeList [get]
func (PxeConfigApi *PxeApi) GetPxeList(c *gin.Context) {
	var pageInfo osManageReq.PxeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := PxeConfigService.GetPxeInfoList(pageInfo); err != nil {
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
