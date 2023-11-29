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

type DeviceBareMetalApi struct {
}

var deviceBareMetalService = service.ServiceGroupApp.DeviceServiceGroup.DeviceBareMetalService

// CreateDeviceBareMetal 创建裸金属设备
// @Tags DeviceBareMetal
// @Summary 创建裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceBareMetal true "创建裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceBareMetal/createDeviceBareMetal [post]
func (deviceBareMetalApi *DeviceBareMetalApi) CreateDeviceBareMetal(c *gin.Context) {
	var deviceBareMetal device.DeviceBareMetal
	err := c.ShouldBindJSON(&deviceBareMetal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":              {utils.NotEmpty()},
		"AssetId":         {utils.NotEmpty()},
		"Hostname":        {utils.NotEmpty()},
		"Ip":              {utils.NotEmpty()},
		"NetworkId":       {utils.NotEmpty()},
		"ManageIp":        {utils.NotEmpty()},
		"ManageNetworkId": {utils.NotEmpty()},
		"KickstartId":     {utils.NotEmpty()},
		"CabinetId":       {utils.NotEmpty()},
		"CabinetU":        {utils.NotEmpty()},
		"Status":          {utils.NotEmpty()},
		"ManagerDev":      {utils.NotEmpty()},
		"ManagerOps":      {utils.NotEmpty()},
		"Label":           {utils.NotEmpty()},
		"Describe":        {utils.NotEmpty()},
	}
	if err := utils.Verify(deviceBareMetal, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalService.CreateDeviceBareMetal(&deviceBareMetal); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDeviceBareMetal 删除裸金属设备
// @Tags DeviceBareMetal
// @Summary 删除裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceBareMetal true "删除裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceBareMetal/deleteDeviceBareMetal [delete]
func (deviceBareMetalApi *DeviceBareMetalApi) DeleteDeviceBareMetal(c *gin.Context) {
	var deviceBareMetal device.DeviceBareMetal
	err := c.ShouldBindJSON(&deviceBareMetal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalService.DeleteDeviceBareMetal(deviceBareMetal); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceBareMetalByIds 批量删除裸金属设备
// @Tags DeviceBareMetal
// @Summary 批量删除裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /deviceBareMetal/deleteDeviceBareMetalByIds [delete]
func (deviceBareMetalApi *DeviceBareMetalApi) DeleteDeviceBareMetalByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalService.DeleteDeviceBareMetalByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDeviceBareMetal 更新裸金属设备
// @Tags DeviceBareMetal
// @Summary 更新裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.DeviceBareMetal true "更新裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceBareMetal/updateDeviceBareMetal [put]
func (deviceBareMetalApi *DeviceBareMetalApi) UpdateDeviceBareMetal(c *gin.Context) {
	var deviceBareMetal device.DeviceBareMetal
	err := c.ShouldBindJSON(&deviceBareMetal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sn":              {utils.NotEmpty()},
		"AssetId":         {utils.NotEmpty()},
		"Hostname":        {utils.NotEmpty()},
		"Ip":              {utils.NotEmpty()},
		"NetworkId":       {utils.NotEmpty()},
		"ManageIp":        {utils.NotEmpty()},
		"ManageNetworkId": {utils.NotEmpty()},
		"KickstartId":     {utils.NotEmpty()},
		"CabinetId":       {utils.NotEmpty()},
		"CabinetU":        {utils.NotEmpty()},
		"Status":          {utils.NotEmpty()},
		"ManagerDev":      {utils.NotEmpty()},
		"ManagerOps":      {utils.NotEmpty()},
		"Label":           {utils.NotEmpty()},
		"Describe":        {utils.NotEmpty()},
	}
	if err := utils.Verify(deviceBareMetal, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := deviceBareMetalService.UpdateDeviceBareMetal(deviceBareMetal); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDeviceBareMetal 用id查询裸金属设备
// @Tags DeviceBareMetal
// @Summary 用id查询裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query device.DeviceBareMetal true "用id查询裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceBareMetal/findDeviceBareMetal [get]
func (deviceBareMetalApi *DeviceBareMetalApi) FindDeviceBareMetal(c *gin.Context) {
	var deviceBareMetal device.DeviceBareMetal
	err := c.ShouldBindQuery(&deviceBareMetal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redeviceBareMetal, err := deviceBareMetalService.GetDeviceBareMetal(deviceBareMetal.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redeviceBareMetal": redeviceBareMetal}, c)
	}
}

// GetDeviceBareMetalList 分页获取裸金属设备列表
// @Tags DeviceBareMetal
// @Summary 分页获取裸金属设备列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query deviceReq.DeviceBareMetalSearch true "分页获取裸金属设备列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceBareMetal/getDeviceBareMetalList [get]
func (deviceBareMetalApi *DeviceBareMetalApi) GetDeviceBareMetalList(c *gin.Context) {
	var pageInfo deviceReq.DeviceBareMetalSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := deviceBareMetalService.GetDeviceBareMetalInfoList(pageInfo); err != nil {
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
