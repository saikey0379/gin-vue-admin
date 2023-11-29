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

type IdcInfoApi struct {
}

var idcInfoService = service.ServiceGroupApp.IdcServiceGroup.IdcInfoService


// CreateIdcInfo 创建数据中心
// @Tags IdcInfo
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfo true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfo/createIdcInfo [post]
func (idcInfoApi *IdcInfoApi) CreateIdcInfo(c *gin.Context) {
	var idcInfo idc.IdcInfo
	err := c.ShouldBindJSON(&idcInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcInfo, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcInfoService.CreateIdcInfo(&idcInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcInfo 删除数据中心
// @Tags IdcInfo
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfo true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfo/deleteIdcInfo [delete]
func (idcInfoApi *IdcInfoApi) DeleteIdcInfo(c *gin.Context) {
	var idcInfo idc.IdcInfo
	err := c.ShouldBindJSON(&idcInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfoService.DeleteIdcInfo(idcInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcInfoByIds 批量删除数据中心
// @Tags IdcInfo
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcInfo/deleteIdcInfoByIds [delete]
func (idcInfoApi *IdcInfoApi) DeleteIdcInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcInfoService.DeleteIdcInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcInfo 更新数据中心
// @Tags IdcInfo
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcInfo true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfo/updateIdcInfo [put]
func (idcInfoApi *IdcInfoApi) UpdateIdcInfo(c *gin.Context) {
	var idcInfo idc.IdcInfo
	err := c.ShouldBindJSON(&idcInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcInfo, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcInfoService.UpdateIdcInfo(idcInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcInfo 用id查询数据中心
// @Tags IdcInfo
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcInfo true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfo/findIdcInfo [get]
func (idcInfoApi *IdcInfoApi) FindIdcInfo(c *gin.Context) {
	var idcInfo idc.IdcInfo
	err := c.ShouldBindQuery(&idcInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcInfo, err := idcInfoService.GetIdcInfo(idcInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcInfo": reidcInfo}, c)
	}
}

// GetIdcInfoList 分页获取数据中心列表
// @Tags IdcInfo
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcInfoSearch true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfo/getIdcInfoList [get]
func (idcInfoApi *IdcInfoApi) GetIdcInfoList(c *gin.Context) {
	var pageInfo idcReq.IdcInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcInfoService.GetIdcInfoInfoList(pageInfo); err != nil {
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
