package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/file"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    fileReq "github.com/flipped-aurora/gin-vue-admin/server/model/file/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type FileInfoApi struct {
}

var fileInfoService = service.ServiceGroupApp.FileServiceGroup.FileInfoService


// CreateFileInfo 创建PXE配置
// @Tags FileInfo
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileInfo true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileInfo/createFileInfo [post]
func (fileInfoApi *FileInfoApi) CreateFileInfo(c *gin.Context) {
	var fileInfo file.FileInfo
	err := c.ShouldBindJSON(&fileInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Description":{utils.NotEmpty()},
        "Content":{utils.NotEmpty()},
        "FileType":{utils.NotEmpty()},
        "Manager":{utils.NotEmpty()},
    }
	if err := utils.Verify(fileInfo, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := fileInfoService.CreateFileInfo(&fileInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileInfo 删除PXE配置
// @Tags FileInfo
// @Summary 删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileInfo true "删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileInfo/deleteFileInfo [delete]
func (fileInfoApi *FileInfoApi) DeleteFileInfo(c *gin.Context) {
	var fileInfo file.FileInfo
	err := c.ShouldBindJSON(&fileInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileInfoService.DeleteFileInfo(fileInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileInfoByIds 批量删除PXE配置
// @Tags FileInfo
// @Summary 批量删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileInfo/deleteFileInfoByIds [delete]
func (fileInfoApi *FileInfoApi) DeleteFileInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileInfoService.DeleteFileInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileInfo 更新PXE配置
// @Tags FileInfo
// @Summary 更新PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileInfo true "更新PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileInfo/updateFileInfo [put]
func (fileInfoApi *FileInfoApi) UpdateFileInfo(c *gin.Context) {
	var fileInfo file.FileInfo
	err := c.ShouldBindJSON(&fileInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Description":{utils.NotEmpty()},
          "Content":{utils.NotEmpty()},
          "FileType":{utils.NotEmpty()},
          "Manager":{utils.NotEmpty()},
      }
    if err := utils.Verify(fileInfo, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := fileInfoService.UpdateFileInfo(fileInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileInfo 用id查询PXE配置
// @Tags FileInfo
// @Summary 用id查询PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileInfo true "用id查询PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileInfo/findFileInfo [get]
func (fileInfoApi *FileInfoApi) FindFileInfo(c *gin.Context) {
	var fileInfo file.FileInfo
	err := c.ShouldBindQuery(&fileInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileInfo, err := fileInfoService.GetFileInfo(fileInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileInfo": refileInfo}, c)
	}
}

// GetFileInfoList 分页获取PXE配置列表
// @Tags FileInfo
// @Summary 分页获取PXE配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileInfoSearch true "分页获取PXE配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileInfo/getFileInfoList [get]
func (fileInfoApi *FileInfoApi) GetFileInfoList(c *gin.Context) {
	var pageInfo fileReq.FileInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileInfoService.GetFileInfoInfoList(pageInfo); err != nil {
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
