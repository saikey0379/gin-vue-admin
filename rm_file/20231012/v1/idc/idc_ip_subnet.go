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

type IdcIpSubnetApi struct {
}

var idcIpSubnetService = service.ServiceGroupApp.IdcServiceGroup.IdcIpSubnetService


// CreateIdcIpSubnet 创建数据中心子网
// @Tags IdcIpSubnet
// @Summary 创建数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSubnet true "创建数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpSubnet/createIdcIpSubnet [post]
func (idcIpSubnetApi *IdcIpSubnetApi) CreateIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindJSON(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Network":{utils.NotEmpty()},
        "Prefix":{utils.NotEmpty()},
        "IdcId":{utils.NotEmpty()},
        "Name":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(idcIpSubnet, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := idcIpSubnetService.CreateIdcIpSubnet(&idcIpSubnet); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcIpSubnet 删除数据中心子网
// @Tags IdcIpSubnet
// @Summary 删除数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSubnet true "删除数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSubnet/deleteIdcIpSubnet [delete]
func (idcIpSubnetApi *IdcIpSubnetApi) DeleteIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindJSON(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSubnetService.DeleteIdcIpSubnet(idcIpSubnet); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcIpSubnetByIds 批量删除数据中心子网
// @Tags IdcIpSubnet
// @Summary 批量删除数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcIpSubnet/deleteIdcIpSubnetByIds [delete]
func (idcIpSubnetApi *IdcIpSubnetApi) DeleteIdcIpSubnetByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcIpSubnetService.DeleteIdcIpSubnetByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcIpSubnet 更新数据中心子网
// @Tags IdcIpSubnet
// @Summary 更新数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcIpSubnet true "更新数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpSubnet/updateIdcIpSubnet [put]
func (idcIpSubnetApi *IdcIpSubnetApi) UpdateIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindJSON(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Network":{utils.NotEmpty()},
          "Prefix":{utils.NotEmpty()},
          "IdcId":{utils.NotEmpty()},
          "Name":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(idcIpSubnet, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := idcIpSubnetService.UpdateIdcIpSubnet(idcIpSubnet); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcIpSubnet 用id查询数据中心子网
// @Tags IdcIpSubnet
// @Summary 用id查询数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcIpSubnet true "用id查询数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpSubnet/findIdcIpSubnet [get]
func (idcIpSubnetApi *IdcIpSubnetApi) FindIdcIpSubnet(c *gin.Context) {
	var idcIpSubnet idc.IdcIpSubnet
	err := c.ShouldBindQuery(&idcIpSubnet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reidcIpSubnet, err := idcIpSubnetService.GetIdcIpSubnet(idcIpSubnet.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidcIpSubnet": reidcIpSubnet}, c)
	}
}

// GetIdcIpSubnetList 分页获取数据中心子网列表
// @Tags IdcIpSubnet
// @Summary 分页获取数据中心子网列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcIpSubnetSearch true "分页获取数据中心子网列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpSubnet/getIdcIpSubnetList [get]
func (idcIpSubnetApi *IdcIpSubnetApi) GetIdcIpSubnetList(c *gin.Context) {
	var pageInfo idcReq.IdcIpSubnetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := idcIpSubnetService.GetIdcIpSubnetInfoList(pageInfo); err != nil {
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
