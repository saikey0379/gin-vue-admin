package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceBareMetalHardwareApi struct {
}

var deviceBareMetalHardwareService = service.ServiceGroupApp.DeviceServiceGroup.DeviceBareMetalHardwareService

// CreateDeviceBareMetalHardware 创建裸金属设备硬件信息
// @Tags DeviceBareMetalHardware
// @Summary 创建裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceBareMetalHardware true "创建裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceBareMetalHardware/createDeviceBareMetalHardware [post]
func (deviceBareMetalHardwareApi *DeviceBareMetalHardwareApi) CreateDeviceBareMetalHardware(c *gin.Context) {
	var deviceBareMetalHardware device.DeviceBareMetalHardware
	err := c.ShouldBindJSON(&deviceBareMetalHardware)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":           {utils.NotEmpty()},
		"Manufacturer": {utils.NotEmpty()},
		"Model":        {utils.NotEmpty()},
		"CpuSum":       {utils.NotEmpty()},
		"Cpu":          {utils.NotEmpty()},
		"MemorySum":    {utils.NotEmpty()},
		"Memory":       {utils.NotEmpty()},
		"Nic":          {utils.NotEmpty()},
		"Gpu":          {utils.NotEmpty()},
		"Motherboard":  {utils.NotEmpty()},
		"Raid":         {utils.NotEmpty()},
		"Disk":         {utils.NotEmpty()},
	}
	if err := utils.Verify(deviceBareMetalHardware, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalHardwareService.CreateDeviceBareMetalHardware(&deviceBareMetalHardware); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDeviceBareMetalHardware 删除裸金属设备硬件信息
// @Tags DeviceBareMetalHardware
// @Summary 删除裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceBareMetalHardware true "删除裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceBareMetalHardware/deleteDeviceBareMetalHardware [delete]
func (deviceBareMetalHardwareApi *DeviceBareMetalHardwareApi) DeleteDeviceBareMetalHardware(c *gin.Context) {
	var deviceBareMetalHardware device.DeviceBareMetalHardware
	err := c.ShouldBindJSON(&deviceBareMetalHardware)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalHardwareService.DeleteDeviceBareMetalHardware(deviceBareMetalHardware); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceBareMetalHardwareByIds 批量删除裸金属设备硬件信息
// @Tags DeviceBareMetalHardware
// @Summary 批量删除裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /deviceBareMetalHardware/deleteDeviceBareMetalHardwareByIds [delete]
func (deviceBareMetalHardwareApi *DeviceBareMetalHardwareApi) DeleteDeviceBareMetalHardwareByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalHardwareService.DeleteDeviceBareMetalHardwareByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDeviceBareMetalHardware 更新裸金属设备硬件信息
// @Tags DeviceBareMetalHardware
// @Summary 更新裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceBareMetalHardware true "更新裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceBareMetalHardware/updateDeviceBareMetalHardware [put]
func (deviceBareMetalHardwareApi *DeviceBareMetalHardwareApi) UpdateDeviceBareMetalHardware(c *gin.Context) {
	var deviceBareMetalHardware device.DeviceBareMetalHardware
	err := c.ShouldBindJSON(&deviceBareMetalHardware)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":           {utils.NotEmpty()},
		"Manufacturer": {utils.NotEmpty()},
		"Model":        {utils.NotEmpty()},
		"CpuSum":       {utils.NotEmpty()},
		"Cpu":          {utils.NotEmpty()},
		"MemorySum":    {utils.NotEmpty()},
		"Memory":       {utils.NotEmpty()},
		"Nic":          {utils.NotEmpty()},
		"Gpu":          {utils.NotEmpty()},
		"Motherboard":  {utils.NotEmpty()},
		"Raid":         {utils.NotEmpty()},
		"Disk":         {utils.NotEmpty()},
	}
	if err := utils.Verify(deviceBareMetalHardware, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalHardwareService.UpdateDeviceBareMetalHardware(deviceBareMetalHardware); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDeviceBareMetalHardware 用id查询裸金属设备硬件信息
// @Tags DeviceBareMetalHardware
// @Summary 用id查询裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query device.DeviceBareMetalHardware true "用id查询裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceBareMetalHardware/findDeviceBareMetalHardware [get]
func (deviceBareMetalHardwareApi *DeviceBareMetalHardwareApi) FindDeviceBareMetalHardware(c *gin.Context) {
	var deviceBareMetalHardware device.DeviceBareMetalHardware
	err := c.ShouldBindQuery(&deviceBareMetalHardware)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redeviceBareMetalHardware, err := deviceBareMetalHardwareService.GetDeviceBareMetalHardware(deviceBareMetalHardware.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redeviceBareMetalHardware": redeviceBareMetalHardware}, c)
	}
}

// GetDeviceBareMetalHardwareList 分页获取裸金属设备硬件信息列表
// @Tags DeviceBareMetalHardware
// @Summary 分页获取裸金属设备硬件信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query deviceReq.DeviceBareMetalHardwareSearch true "分页获取裸金属设备硬件信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceBareMetalHardware/getDeviceBareMetalHardwareList [get]
func (deviceBareMetalHardwareApi *DeviceBareMetalHardwareApi) GetDeviceBareMetalHardwareList(c *gin.Context) {
	var pageInfo deviceReq.DeviceBareMetalHardwareSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := deviceBareMetalHardwareService.GetDeviceBareMetalHardwareInfoList(pageInfo); err != nil {
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
