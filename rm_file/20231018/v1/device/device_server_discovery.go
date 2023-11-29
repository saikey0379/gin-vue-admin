package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/device"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DeviceServerDiscoveryApi struct {
}

var deviceServerDiscoveryService = service.ServiceGroupApp.DeviceServiceGroup.DeviceServerDiscoveryService


// CreateDeviceServerDiscovery 创建Server发现
// @Tags DeviceServerDiscovery
// @Summary 创建Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceServerDiscovery true "创建Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceServerDiscovery/createDeviceServerDiscovery [post]
func (deviceServerDiscoveryApi *DeviceServerDiscoveryApi) CreateDeviceServerDiscovery(c *gin.Context) {
	var deviceServerDiscovery device.DeviceServerDiscovery
	err := c.ShouldBindJSON(&deviceServerDiscovery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Sn":{utils.NotEmpty()},
        "Hostname":{utils.NotEmpty()},
        "Ip":{utils.NotEmpty()},
    }
	if err := utils.Verify(deviceServerDiscovery, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := deviceServerDiscoveryService.CreateDeviceServerDiscovery(&deviceServerDiscovery); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDeviceServerDiscovery 删除Server发现
// @Tags DeviceServerDiscovery
// @Summary 删除Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceServerDiscovery true "删除Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceServerDiscovery/deleteDeviceServerDiscovery [delete]
func (deviceServerDiscoveryApi *DeviceServerDiscoveryApi) DeleteDeviceServerDiscovery(c *gin.Context) {
	var deviceServerDiscovery device.DeviceServerDiscovery
	err := c.ShouldBindJSON(&deviceServerDiscovery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceServerDiscoveryService.DeleteDeviceServerDiscovery(deviceServerDiscovery); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceServerDiscoveryByIds 批量删除Server发现
// @Tags DeviceServerDiscovery
// @Summary 批量删除Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /deviceServerDiscovery/deleteDeviceServerDiscoveryByIds [delete]
func (deviceServerDiscoveryApi *DeviceServerDiscoveryApi) DeleteDeviceServerDiscoveryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceServerDiscoveryService.DeleteDeviceServerDiscoveryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDeviceServerDiscovery 更新Server发现
// @Tags DeviceServerDiscovery
// @Summary 更新Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceServerDiscovery true "更新Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceServerDiscovery/updateDeviceServerDiscovery [put]
func (deviceServerDiscoveryApi *DeviceServerDiscoveryApi) UpdateDeviceServerDiscovery(c *gin.Context) {
	var deviceServerDiscovery device.DeviceServerDiscovery
	err := c.ShouldBindJSON(&deviceServerDiscovery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Sn":{utils.NotEmpty()},
          "Hostname":{utils.NotEmpty()},
          "Ip":{utils.NotEmpty()},
      }
    if err := utils.Verify(deviceServerDiscovery, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := deviceServerDiscoveryService.UpdateDeviceServerDiscovery(deviceServerDiscovery); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDeviceServerDiscovery 用id查询Server发现
// @Tags DeviceServerDiscovery
// @Summary 用id查询Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query device.DeviceServerDiscovery true "用id查询Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceServerDiscovery/findDeviceServerDiscovery [get]
func (deviceServerDiscoveryApi *DeviceServerDiscoveryApi) FindDeviceServerDiscovery(c *gin.Context) {
	var deviceServerDiscovery device.DeviceServerDiscovery
	err := c.ShouldBindQuery(&deviceServerDiscovery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redeviceServerDiscovery, err := deviceServerDiscoveryService.GetDeviceServerDiscovery(deviceServerDiscovery.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redeviceServerDiscovery": redeviceServerDiscovery}, c)
	}
}

// GetDeviceServerDiscoveryList 分页获取Server发现列表
// @Tags DeviceServerDiscovery
// @Summary 分页获取Server发现列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query deviceReq.DeviceServerDiscoverySearch true "分页获取Server发现列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceServerDiscovery/getDeviceServerDiscoveryList [get]
func (deviceServerDiscoveryApi *DeviceServerDiscoveryApi) GetDeviceServerDiscoveryList(c *gin.Context) {
	var pageInfo deviceReq.DeviceServerDiscoverySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := deviceServerDiscoveryService.GetDeviceServerDiscoveryInfoList(pageInfo); err != nil {
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
