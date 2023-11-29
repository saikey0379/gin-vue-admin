package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/idc"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type IdcCabinetApi struct {
}

var idcCabinetService = service.ServiceGroupApp.IdcServiceGroup.IdcCabinetService


// CreateIdcCabinet 创建机柜
// @Tags IdcCabinet
// @Summary 创建机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcCabinet true "创建机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcCabinet/createIdcCabinet [post]
func (idcCabinetApi *IdcCabinetApi) CreateIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindJSON(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "RoomId":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcCabinet, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcCabinetService.CreateIdcCabinet(&idcCabinet); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcCabinet 删除机柜
// @Tags IdcCabinet
// @Summary 删除机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcCabinet true "删除机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcCabinet/deleteIdcCabinet [delete]
func (idcCabinetApi *IdcCabinetApi) DeleteIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindJSON(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcCabinetService.DeleteIdcCabinet(idcCabinet); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcCabinetByIds 批量删除机柜
// @Tags IdcCabinet
// @Summary 批量删除机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcCabinet/deleteIdcCabinetByIds [delete]
func (idcCabinetApi *IdcCabinetApi) DeleteIdcCabinetByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcCabinetService.DeleteIdcCabinetByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcCabinet 更新机柜
// @Tags IdcCabinet
// @Summary 更新机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcCabinet true "更新机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcCabinet/updateIdcCabinet [put]
func (idcCabinetApi *IdcCabinetApi) UpdateIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindJSON(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "RoomId":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcCabinet, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcCabinetService.UpdateIdcCabinet(idcCabinet); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcCabinet 用id查询机柜
// @Tags IdcCabinet
// @Summary 用id查询机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcCabinet true "用id查询机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcCabinet/findIdcCabinet [get]
func (idcCabinetApi *IdcCabinetApi) FindIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindQuery(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcCabinet, err := idcCabinetService.GetIdcCabinet(idcCabinet.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcCabinet": reidcCabinet}, c)
	}
}

// GetIdcCabinetList 分页获取机柜列表
// @Tags IdcCabinet
// @Summary 分页获取机柜列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcCabinetSearch true "分页获取机柜列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcCabinet/getIdcCabinetList [get]
func (idcCabinetApi *IdcCabinetApi) GetIdcCabinetList(c *gin.Context) {
	var pageInfo idcReq.IdcCabinetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcCabinetService.GetIdcCabinetInfoList(pageInfo); err != nil {
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
