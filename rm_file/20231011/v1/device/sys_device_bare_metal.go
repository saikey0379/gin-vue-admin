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

type BareMetalApi struct {
}

var bmService = service.ServiceGroupApp.DeviceServiceGroup.BareMetalService


// CreateBareMetal 创建裸金属设备
// @Tags BareMetal
// @Summary 创建裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.BareMetal true "创建裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bm/createBareMetal [post]
func (bmApi *BareMetalApi) CreateBareMetal(c *gin.Context) {
	var bm device.BareMetal
	err := c.ShouldBindJSON(&bm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Sn":{utils.NotEmpty()},
        "AssetId":{utils.NotEmpty()},
        "Hostname":{utils.NotEmpty()},
        "Ip":{utils.NotEmpty()},
        "NetworkId":{utils.NotEmpty()},
        "ManageIp":{utils.NotEmpty()},
        "ManageNetworkId":{utils.NotEmpty()},
        "KickstartId":{utils.NotEmpty()},
        "CabinetId":{utils.NotEmpty()},
        "CabinetU":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
        "ManagerDev":{utils.NotEmpty()},
        "ManagerOps":{utils.NotEmpty()},
        "LabelId":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(bm, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := bmService.CreateBareMetal(&bm); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBareMetal 删除裸金属设备
// @Tags BareMetal
// @Summary 删除裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.BareMetal true "删除裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bm/deleteBareMetal [delete]
func (bmApi *BareMetalApi) DeleteBareMetal(c *gin.Context) {
	var bm device.BareMetal
	err := c.ShouldBindJSON(&bm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bmService.DeleteBareMetal(bm); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBareMetalByIds 批量删除裸金属设备
// @Tags BareMetal
// @Summary 批量删除裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bm/deleteBareMetalByIds [delete]
func (bmApi *BareMetalApi) DeleteBareMetalByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bmService.DeleteBareMetalByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBareMetal 更新裸金属设备
// @Tags BareMetal
// @Summary 更新裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body device.BareMetal true "更新裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bm/updateBareMetal [put]
func (bmApi *BareMetalApi) UpdateBareMetal(c *gin.Context) {
	var bm device.BareMetal
	err := c.ShouldBindJSON(&bm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Sn":{utils.NotEmpty()},
          "AssetId":{utils.NotEmpty()},
          "Hostname":{utils.NotEmpty()},
          "Ip":{utils.NotEmpty()},
          "NetworkId":{utils.NotEmpty()},
          "ManageIp":{utils.NotEmpty()},
          "ManageNetworkId":{utils.NotEmpty()},
          "KickstartId":{utils.NotEmpty()},
          "CabinetId":{utils.NotEmpty()},
          "CabinetU":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
          "ManagerDev":{utils.NotEmpty()},
          "ManagerOps":{utils.NotEmpty()},
          "LabelId":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(bm, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := bmService.UpdateBareMetal(bm); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBareMetal 用id查询裸金属设备
// @Tags BareMetal
// @Summary 用id查询裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query device.BareMetal true "用id查询裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bm/findBareMetal [get]
func (bmApi *BareMetalApi) FindBareMetal(c *gin.Context) {
	var bm device.BareMetal
	err := c.ShouldBindQuery(&bm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebm, err := bmService.GetBareMetal(bm.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebm": rebm}, c)
	}
}

// GetBareMetalList 分页获取裸金属设备列表
// @Tags BareMetal
// @Summary 分页获取裸金属设备列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query deviceReq.BareMetalSearch true "分页获取裸金属设备列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bm/getBareMetalList [get]
func (bmApi *BareMetalApi) GetBareMetalList(c *gin.Context) {
	var pageInfo deviceReq.BareMetalSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := bmService.GetBareMetalInfoList(pageInfo); err != nil {
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
