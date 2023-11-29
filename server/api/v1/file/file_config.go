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

type FileConfigApi struct {
}

var fileConfigService = service.ServiceGroupApp.FileServiceGroup.FileConfigService

// CreateFileConfig 创建配置文件
// @Tags FileConfig
// @Summary 创建配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileConfig true "创建配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileConfig/createFileConfig [post]
func (fileConfigApi *FileConfigApi) CreateFileConfig(c *gin.Context) {
	var fileConfig file.FileConfig
	err := c.ShouldBindJSON(&fileConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
		"DataType":    {utils.NotEmpty()},
	}
	if err := utils.Verify(fileConfig, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byte_content := []byte(fileConfig.Content)
	md5sum_content := md5.Sum(byte_content)
	fileConfig.Md5 = hex.EncodeToString(md5sum_content[:])

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileConfig.Manager = claimsReal.Username

	if err := fileConfigService.CreateFileConfig(&fileConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileConfig 删除配置文件
// @Tags FileConfig
// @Summary 删除配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileConfig true "删除配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileConfig/deleteFileConfig [delete]
func (fileConfigApi *FileConfigApi) DeleteFileConfig(c *gin.Context) {
	var fileConfig file.FileConfig
	err := c.ShouldBindJSON(&fileConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileConfigService.DeleteFileConfig(fileConfig); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileConfigByIds 批量删除配置文件
// @Tags FileConfig
// @Summary 批量删除配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileConfig/deleteFileConfigByIds [delete]
func (fileConfigApi *FileConfigApi) DeleteFileConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileConfigService.DeleteFileConfigByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileConfig 更新配置文件
// @Tags FileConfig
// @Summary 更新配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileConfig true "更新配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileConfig/updateFileConfig [put]
func (fileConfigApi *FileConfigApi) UpdateFileConfig(c *gin.Context) {
	var fileConfig file.FileConfig
	err := c.ShouldBindJSON(&fileConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
		"DataType":    {utils.NotEmpty()},
	}
	if err := utils.Verify(fileConfig, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byte_content := []byte(fileConfig.Content)
	md5sum_content := md5.Sum(byte_content)
	fileConfig.Md5 = hex.EncodeToString(md5sum_content[:])

	claims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	claimsReal := claims.(*request1.CustomClaims)
	fileConfig.Manager = claimsReal.Username

	if err := fileConfigService.UpdateFileConfig(fileConfig); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileConfig 用id查询配置文件
// @Tags FileConfig
// @Summary 用id查询配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileConfig true "用id查询配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileConfig/findFileConfig [get]
func (fileConfigApi *FileConfigApi) FindFileConfig(c *gin.Context) {
	var fileConfig file.FileConfig
	err := c.ShouldBindQuery(&fileConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileConfig, err := fileConfigService.GetFileConfig(fileConfig.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileConfig": refileConfig}, c)
	}
}

// GetFileConfigList 分页获取配置文件列表
// @Tags FileConfig
// @Summary 分页获取配置文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileConfigSearch true "分页获取配置文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileConfig/getFileConfigList [get]
func (fileConfigApi *FileConfigApi) GetFileConfigList(c *gin.Context) {
	var pageInfo fileReq.FileConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileConfigService.GetFileConfigInfoList(pageInfo); err != nil {
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

func (fileConfigApi *FileConfigApi) GetContent(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		response.FailWithMessage("name不能为空", c)
		return
	}

	var fileConfig file.FileConfig
	err := c.ShouldBindQuery(&fileConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if refileConfig, err := fileConfigService.GetFileConfigByName(name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.ResultContent([]byte(refileConfig.Content), c)
	}
}
