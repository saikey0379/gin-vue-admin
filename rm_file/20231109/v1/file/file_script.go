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

type FileScriptApi struct {
}

var fileScriptService = service.ServiceGroupApp.FileServiceGroup.FileScriptService


// CreateFileScript 创建脚本文件
// @Tags FileScript
// @Summary 创建脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileScript true "创建脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileScript/createFileScript [post]
func (fileScriptApi *FileScriptApi) CreateFileScript(c *gin.Context) {
	var fileScript file.FileScript
	err := c.ShouldBindJSON(&fileScript)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Description":{utils.NotEmpty()},
        "Content":{utils.NotEmpty()},
        "Manager":{utils.NotEmpty()},
    }
	if err := utils.Verify(fileScript, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := fileScriptService.CreateFileScript(&fileScript); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileScript 删除脚本文件
// @Tags FileScript
// @Summary 删除脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileScript true "删除脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileScript/deleteFileScript [delete]
func (fileScriptApi *FileScriptApi) DeleteFileScript(c *gin.Context) {
	var fileScript file.FileScript
	err := c.ShouldBindJSON(&fileScript)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileScriptService.DeleteFileScript(fileScript); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileScriptByIds 批量删除脚本文件
// @Tags FileScript
// @Summary 批量删除脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileScript/deleteFileScriptByIds [delete]
func (fileScriptApi *FileScriptApi) DeleteFileScriptByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileScriptService.DeleteFileScriptByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileScript 更新脚本文件
// @Tags FileScript
// @Summary 更新脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FileScript true "更新脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileScript/updateFileScript [put]
func (fileScriptApi *FileScriptApi) UpdateFileScript(c *gin.Context) {
	var fileScript file.FileScript
	err := c.ShouldBindJSON(&fileScript)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Description":{utils.NotEmpty()},
          "Content":{utils.NotEmpty()},
          "Manager":{utils.NotEmpty()},
      }
    if err := utils.Verify(fileScript, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := fileScriptService.UpdateFileScript(fileScript); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileScript 用id查询脚本文件
// @Tags FileScript
// @Summary 用id查询脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FileScript true "用id查询脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileScript/findFileScript [get]
func (fileScriptApi *FileScriptApi) FindFileScript(c *gin.Context) {
	var fileScript file.FileScript
	err := c.ShouldBindQuery(&fileScript)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileScript, err := fileScriptService.GetFileScript(fileScript.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileScript": refileScript}, c)
	}
}

// GetFileScriptList 分页获取脚本文件列表
// @Tags FileScript
// @Summary 分页获取脚本文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FileScriptSearch true "分页获取脚本文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileScript/getFileScriptList [get]
func (fileScriptApi *FileScriptApi) GetFileScriptList(c *gin.Context) {
	var pageInfo fileReq.FileScriptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileScriptService.GetFileScriptInfoList(pageInfo); err != nil {
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
