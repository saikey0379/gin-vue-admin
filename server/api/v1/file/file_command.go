package file

import (
	"crypto/md5"
	"encoding/hex"
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

type FileCommandApi struct {
}

var fileCommandService = service.ServiceGroupApp.FileServiceGroup.FileCommandService

// CreateFileCommand 创建命令信息
// @Tags FileCommand
// @Summary 创建命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommand true "创建命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileCommand/createFileCommand [post]
func (fileCommandApi *FileCommandApi) CreateFileCommand(c *gin.Context) {
	var fileCommand file.FileCommand
	err := c.ShouldBindJSON(&fileCommand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
	}
	if err := utils.Verify(fileCommand, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byte_content := []byte(fileCommand.Content)
	md5sum_content := md5.Sum(byte_content)
	fileCommand.Md5 = hex.EncodeToString(md5sum_content[:])

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileCommand.Manager = claimsReal.Username

	if err := fileCommandService.CreateFileCommand(&fileCommand); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileCommand 删除命令信息
// @Tags FileCommand
// @Summary 删除命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommand true "删除命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommand/deleteFileCommand [delete]
func (fileCommandApi *FileCommandApi) DeleteFileCommand(c *gin.Context) {
	var fileCommand file.FileCommand
	err := c.ShouldBindJSON(&fileCommand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileCommandService.DeleteFileCommand(fileCommand); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileCommandByIds 批量删除命令信息
// @Tags FileCommand
// @Summary 批量删除命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileCommand/deleteFileCommandByIds [delete]
func (fileCommandApi *FileCommandApi) DeleteFileCommandByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileCommandService.DeleteFileCommandByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileCommand 更新命令信息
// @Tags FileCommand
// @Summary 更新命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommand true "更新命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileCommand/updateFileCommand [put]
func (fileCommandApi *FileCommandApi) UpdateFileCommand(c *gin.Context) {
	var fileCommand file.FileCommand
	err := c.ShouldBindJSON(&fileCommand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
	}
	if err := utils.Verify(fileCommand, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byte_content := []byte(fileCommand.Content)
	md5sum_content := md5.Sum(byte_content)
	fileCommand.Md5 = hex.EncodeToString(md5sum_content[:])

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileCommand.Manager = claimsReal.Username

	if err := fileCommandService.UpdateFileCommand(fileCommand); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileCommand 用id查询命令信息
// @Tags FileCommand
// @Summary 用id查询命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileCommand true "用id查询命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileCommand/findFileCommand [get]
func (fileCommandApi *FileCommandApi) FindFileCommand(c *gin.Context) {
	var fileCommand file.FileCommand
	err := c.ShouldBindQuery(&fileCommand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileCommand, err := fileCommandService.GetFileCommand(fileCommand.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileCommand": refileCommand}, c)
	}
}

// GetFileCommandList 分页获取命令信息列表
// @Tags FileCommand
// @Summary 分页获取命令信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileCommandSearch true "分页获取命令信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileCommand/getFileCommandList [get]
func (fileCommandApi *FileCommandApi) GetFileCommandList(c *gin.Context) {
	var pageInfo fileReq.FileCommandSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileCommandService.GetFileCommandInfoList(pageInfo); err != nil {
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
