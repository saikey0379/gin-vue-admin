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

type FileCommadBlacklistApi struct {
}

var fileCommandBlacklistService = service.ServiceGroupApp.FileServiceGroup.FileCommadBlacklistService

// CreateFileCommadBlacklist 创建命令黑名单
// @Tags FileCommadBlacklist
// @Summary 创建命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommadBlacklist true "创建命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileCommandBlacklist/createFileCommadBlacklist [post]
func (fileCommandBlacklistApi *FileCommadBlacklistApi) CreateFileCommadBlacklist(c *gin.Context) {
	var fileCommandBlacklist file.FileCommadBlacklist
	err := c.ShouldBindJSON(&fileCommandBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
	}
	if err := utils.Verify(fileCommandBlacklist, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byte_content := []byte(fileCommandBlacklist.Content)
	md5sum_content := md5.Sum(byte_content)
	fileCommandBlacklist.Md5 = hex.EncodeToString(md5sum_content[:])

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileCommandBlacklist.Manager = claimsReal.Username

	if err := fileCommandBlacklistService.CreateFileCommadBlacklist(&fileCommandBlacklist); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileCommadBlacklist 删除命令黑名单
// @Tags FileCommadBlacklist
// @Summary 删除命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommadBlacklist true "删除命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommandBlacklist/deleteFileCommadBlacklist [delete]
func (fileCommandBlacklistApi *FileCommadBlacklistApi) DeleteFileCommadBlacklist(c *gin.Context) {
	var fileCommandBlacklist file.FileCommadBlacklist
	err := c.ShouldBindJSON(&fileCommandBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileCommandBlacklistService.DeleteFileCommadBlacklist(fileCommandBlacklist); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileCommadBlacklistByIds 批量删除命令黑名单
// @Tags FileCommadBlacklist
// @Summary 批量删除命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileCommandBlacklist/deleteFileCommadBlacklistByIds [delete]
func (fileCommandBlacklistApi *FileCommadBlacklistApi) DeleteFileCommadBlacklistByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileCommandBlacklistService.DeleteFileCommadBlacklistByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileCommadBlacklist 更新命令黑名单
// @Tags FileCommadBlacklist
// @Summary 更新命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileCommadBlacklist true "更新命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileCommandBlacklist/updateFileCommadBlacklist [put]
func (fileCommandBlacklistApi *FileCommadBlacklistApi) UpdateFileCommadBlacklist(c *gin.Context) {
	var fileCommandBlacklist file.FileCommadBlacklist
	err := c.ShouldBindJSON(&fileCommandBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
	}
	if err := utils.Verify(fileCommandBlacklist, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byte_content := []byte(fileCommandBlacklist.Content)
	md5sum_content := md5.Sum(byte_content)
	fileCommandBlacklist.Md5 = hex.EncodeToString(md5sum_content[:])

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileCommandBlacklist.Manager = claimsReal.Username

	if err := fileCommandBlacklistService.UpdateFileCommadBlacklist(fileCommandBlacklist); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileCommadBlacklist 用id查询命令黑名单
// @Tags FileCommadBlacklist
// @Summary 用id查询命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileCommadBlacklist true "用id查询命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileCommandBlacklist/findFileCommadBlacklist [get]
func (fileCommandBlacklistApi *FileCommadBlacklistApi) FindFileCommadBlacklist(c *gin.Context) {
	var fileCommandBlacklist file.FileCommadBlacklist
	err := c.ShouldBindQuery(&fileCommandBlacklist)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileCommandBlacklist, err := fileCommandBlacklistService.GetFileCommadBlacklist(fileCommandBlacklist.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileCommandBlacklist": refileCommandBlacklist}, c)
	}
}

// GetFileCommadBlacklistList 分页获取命令黑名单列表
// @Tags FileCommadBlacklist
// @Summary 分页获取命令黑名单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileCommadBlacklistSearch true "分页获取命令黑名单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileCommandBlacklist/getFileCommadBlacklistList [get]
func (fileCommandBlacklistApi *FileCommadBlacklistApi) GetFileCommadBlacklistList(c *gin.Context) {
	var pageInfo fileReq.FileCommadBlacklistSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileCommandBlacklistService.GetFileCommadBlacklistInfoList(pageInfo); err != nil {
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
