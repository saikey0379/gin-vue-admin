package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cmdbReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmdb/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CmdbRegionApi struct {
}

var cmdbRegionService = service.ServiceGroupApp.CmdbServiceGroup.CmdbRegionService


// CreateCmdbRegion 创建区域信息
// @Tags CmdbRegion
// @Summary 创建区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbRegion true "创建区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbRegion/createCmdbRegion [post]
func (cmdbRegionApi *CmdbRegionApi) CreateCmdbRegion(c *gin.Context) {
	var cmdbRegion cmdb.CmdbRegion
	err := c.ShouldBindJSON(&cmdbRegion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "RegionType":{utils.NotEmpty()},
        "Label":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
    }
	if err := utils.Verify(cmdbRegion, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := cmdbRegionService.CreateCmdbRegion(&cmdbRegion); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmdbRegion 删除区域信息
// @Tags CmdbRegion
// @Summary 删除区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbRegion true "删除区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbRegion/deleteCmdbRegion [delete]
func (cmdbRegionApi *CmdbRegionApi) DeleteCmdbRegion(c *gin.Context) {
	var cmdbRegion cmdb.CmdbRegion
	err := c.ShouldBindJSON(&cmdbRegion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbRegionService.DeleteCmdbRegion(cmdbRegion); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmdbRegionByIds 批量删除区域信息
// @Tags CmdbRegion
// @Summary 批量删除区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmdbRegion/deleteCmdbRegionByIds [delete]
func (cmdbRegionApi *CmdbRegionApi) DeleteCmdbRegionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmdbRegionService.DeleteCmdbRegionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmdbRegion 更新区域信息
// @Tags CmdbRegion
// @Summary 更新区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmdb.CmdbRegion true "更新区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbRegion/updateCmdbRegion [put]
func (cmdbRegionApi *CmdbRegionApi) UpdateCmdbRegion(c *gin.Context) {
	var cmdbRegion cmdb.CmdbRegion
	err := c.ShouldBindJSON(&cmdbRegion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "RegionType":{utils.NotEmpty()},
          "Label":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
      }
    if err := utils.Verify(cmdbRegion, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := cmdbRegionService.UpdateCmdbRegion(cmdbRegion); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmdbRegion 用id查询区域信息
// @Tags CmdbRegion
// @Summary 用id查询区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdb.CmdbRegion true "用id查询区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbRegion/findCmdbRegion [get]
func (cmdbRegionApi *CmdbRegionApi) FindCmdbRegion(c *gin.Context) {
	var cmdbRegion cmdb.CmdbRegion
	err := c.ShouldBindQuery(&cmdbRegion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recmdbRegion, err := cmdbRegionService.GetCmdbRegion(cmdbRegion.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmdbRegion": recmdbRegion}, c)
	}
}

// GetCmdbRegionList 分页获取区域信息列表
// @Tags CmdbRegion
// @Summary 分页获取区域信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmdbReq.CmdbRegionSearch true "分页获取区域信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbRegion/getCmdbRegionList [get]
func (cmdbRegionApi *CmdbRegionApi) GetCmdbRegionList(c *gin.Context) {
	var pageInfo cmdbReq.CmdbRegionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmdbRegionService.GetCmdbRegionInfoList(pageInfo); err != nil {
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
