package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/slb"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SlbCertApi struct {
}

var slbCertService = service.ServiceGroupApp.SlbServiceGroup.SlbCertService


// CreateSlbCert 创建证书管理
// @Tags SlbCert
// @Summary 创建证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbCert true "创建证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbCert/createSlbCert [post]
func (slbCertApi *SlbCertApi) CreateSlbCert(c *gin.Context) {
	var slbCert slb.SlbCert
	err := c.ShouldBindJSON(&slbCert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
        "ClusterIds":{utils.NotEmpty()},
        "CertFile":{utils.NotEmpty()},
        "CertContent":{utils.NotEmpty()},
        "KeyFile":{utils.NotEmpty()},
        "KeyContent":{utils.NotEmpty()},
    }
	if err := utils.Verify(slbCert, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := slbCertService.CreateSlbCert(&slbCert); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlbCert 删除证书管理
// @Tags SlbCert
// @Summary 删除证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbCert true "删除证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbCert/deleteSlbCert [delete]
func (slbCertApi *SlbCertApi) DeleteSlbCert(c *gin.Context) {
	var slbCert slb.SlbCert
	err := c.ShouldBindJSON(&slbCert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbCertService.DeleteSlbCert(slbCert); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlbCertByIds 批量删除证书管理
// @Tags SlbCert
// @Summary 批量删除证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slbCert/deleteSlbCertByIds [delete]
func (slbCertApi *SlbCertApi) DeleteSlbCertByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbCertService.DeleteSlbCertByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlbCert 更新证书管理
// @Tags SlbCert
// @Summary 更新证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbCert true "更新证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbCert/updateSlbCert [put]
func (slbCertApi *SlbCertApi) UpdateSlbCert(c *gin.Context) {
	var slbCert slb.SlbCert
	err := c.ShouldBindJSON(&slbCert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
          "ClusterIds":{utils.NotEmpty()},
          "CertFile":{utils.NotEmpty()},
          "CertContent":{utils.NotEmpty()},
          "KeyFile":{utils.NotEmpty()},
          "KeyContent":{utils.NotEmpty()},
      }
    if err := utils.Verify(slbCert, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := slbCertService.UpdateSlbCert(slbCert); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlbCert 用id查询证书管理
// @Tags SlbCert
// @Summary 用id查询证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slb.SlbCert true "用id查询证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbCert/findSlbCert [get]
func (slbCertApi *SlbCertApi) FindSlbCert(c *gin.Context) {
	var slbCert slb.SlbCert
	err := c.ShouldBindQuery(&slbCert)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslbCert, err := slbCertService.GetSlbCert(slbCert.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reslbCert": reslbCert}, c)
	}
}

// GetSlbCertList 分页获取证书管理列表
// @Tags SlbCert
// @Summary 分页获取证书管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slbReq.SlbCertSearch true "分页获取证书管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbCert/getSlbCertList [get]
func (slbCertApi *SlbCertApi) GetSlbCertList(c *gin.Context) {
	var pageInfo slbReq.SlbCertSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slbCertService.GetSlbCertInfoList(pageInfo); err != nil {
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
