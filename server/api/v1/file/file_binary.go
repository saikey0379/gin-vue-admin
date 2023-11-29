package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
	request1 "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type FileBinaryApi struct {
}

var fileBinaryService = service.ServiceGroupApp.FileServiceGroup.FileBinaryService

// CreateFileBinary 创建可执行程序
// @Tags FileBinary
// @Summary 创建可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileBinary true "创建可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileBinary/createFileBinary [post]
func (fileBinaryApi *FileBinaryApi) CreateFileBinary(c *gin.Context) {
	var fileBinary file.FileBinary
	err := c.ShouldBindJSON(&fileBinary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Md5":         {utils.NotEmpty()},
	}
	if err := utils.Verify(fileBinary, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileBinary.Manager = claimsReal.Username

	if err := fileBinaryService.CreateFileBinary(&fileBinary); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileBinary 删除可执行程序
// @Tags FileBinary
// @Summary 删除可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileBinary true "删除可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileBinary/deleteFileBinary [delete]
func (fileBinaryApi *FileBinaryApi) DeleteFileBinary(c *gin.Context) {
	var fileBinary file.FileBinary
	err := c.ShouldBindJSON(&fileBinary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileBinaryService.DeleteFileBinary(fileBinary); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileBinaryByIds 批量删除可执行程序
// @Tags FileBinary
// @Summary 批量删除可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileBinary/deleteFileBinaryByIds [delete]
func (fileBinaryApi *FileBinaryApi) DeleteFileBinaryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileBinaryService.DeleteFileBinaryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileBinary 更新可执行程序
// @Tags FileBinary
// @Summary 更新可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileBinary true "更新可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileBinary/updateFileBinary [put]
func (fileBinaryApi *FileBinaryApi) UpdateFileBinary(c *gin.Context) {
	var fileBinary file.FileBinary
	err := c.ShouldBindJSON(&fileBinary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Md5":         {utils.NotEmpty()},
	}
	if err := utils.Verify(fileBinary, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileBinary.Manager = claimsReal.Username

	if err := fileBinaryService.UpdateFileBinary(fileBinary); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileBinary 用id查询可执行程序
// @Tags FileBinary
// @Summary 用id查询可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileBinary true "用id查询可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileBinary/findFileBinary [get]
func (fileBinaryApi *FileBinaryApi) FindFileBinary(c *gin.Context) {
	var fileBinary file.FileBinary
	err := c.ShouldBindQuery(&fileBinary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileBinary, err := fileBinaryService.GetFileBinary(fileBinary.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileBinary": refileBinary}, c)
	}
}

// GetFileBinaryList 分页获取可执行程序列表
// @Tags FileBinary
// @Summary 分页获取可执行程序列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileBinarySearch true "分页获取可执行程序列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileBinary/getFileBinaryList [get]
func (fileBinaryApi *FileBinaryApi) GetFileBinaryList(c *gin.Context) {
	var pageInfo fileReq.FileBinarySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileBinaryService.GetFileBinaryInfoList(pageInfo); err != nil {
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
