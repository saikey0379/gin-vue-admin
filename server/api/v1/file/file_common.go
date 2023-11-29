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

type FileCommonApi struct {
}

var fileCommonService = service.ServiceGroupApp.FileServiceGroup.FileCommonService

// CreateFileCommon 创建普通文件
// @Tags FileCommon
// @Summary 创建普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommon true "创建普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileCommon/createFileCommon [post]
func (fileCommonApi *FileCommonApi) CreateFileCommon(c *gin.Context) {
	var fileCommon file.FileCommon
	err := c.ShouldBindJSON(&fileCommon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Md5":         {utils.NotEmpty()},
	}
	if err := utils.Verify(fileCommon, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileCommon.Manager = claimsReal.Username

	if err := fileCommonService.CreateFileCommon(&fileCommon); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileCommon 删除普通文件
// @Tags FileCommon
// @Summary 删除普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommon true "删除普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommon/deleteFileCommon [delete]
func (fileCommonApi *FileCommonApi) DeleteFileCommon(c *gin.Context) {
	var fileCommon file.FileCommon
	err := c.ShouldBindJSON(&fileCommon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileCommonService.DeleteFileCommon(fileCommon); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileCommonByIds 批量删除普通文件
// @Tags FileCommon
// @Summary 批量删除普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileCommon/deleteFileCommonByIds [delete]
func (fileCommonApi *FileCommonApi) DeleteFileCommonByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileCommonService.DeleteFileCommonByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileCommon 更新普通文件
// @Tags FileCommon
// @Summary 更新普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommon true "更新普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileCommon/updateFileCommon [put]
func (fileCommonApi *FileCommonApi) UpdateFileCommon(c *gin.Context) {
	var fileCommon file.FileCommon
	err := c.ShouldBindJSON(&fileCommon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Md5":         {utils.NotEmpty()},
	}
	if err := utils.Verify(fileCommon, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileCommon.Manager = claimsReal.Username

	if err := fileCommonService.UpdateFileCommon(fileCommon); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileCommon 用id查询普通文件
// @Tags FileCommon
// @Summary 用id查询普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileCommon true "用id查询普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileCommon/findFileCommon [get]
func (fileCommonApi *FileCommonApi) FindFileCommon(c *gin.Context) {
	var fileCommon file.FileCommon
	err := c.ShouldBindQuery(&fileCommon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileCommon, err := fileCommonService.GetFileCommon(fileCommon.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileCommon": refileCommon}, c)
	}
}

// GetFileCommonList 分页获取普通文件列表
// @Tags FileCommon
// @Summary 分页获取普通文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileCommonSearch true "分页获取普通文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileCommon/getFileCommonList [get]
func (fileCommonApi *FileCommonApi) GetFileCommonList(c *gin.Context) {
	var pageInfo fileReq.FileCommonSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileCommonService.GetFileCommonInfoList(pageInfo); err != nil {
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
