package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/known"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	slbReq "github.com/flipped-aurora/gin-vue-admin/server/model/slb/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"os"
	"path"
)

type SlbClusterApi struct {
}

var slbClusterService = service.ServiceGroupApp.SlbServiceGroup.SlbClusterService

// CreateSlbCluster 创建集群管理
// @Tags SlbCluster
// @Summary 创建集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbCluster true "创建集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbCluster/createSlbCluster [post]
func (slbClusterApi *SlbClusterApi) CreateSlbCluster(c *gin.Context) {
	var slbCluster slb.SlbCluster
	err := c.ShouldBindJSON(&slbCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
		"SshUser":  {utils.NotEmpty()},
		"SshRsa":   {utils.NotEmpty()},
		"Nodes":    {utils.NotEmpty()},
		"PathConf": {utils.NotEmpty()},
		"PathSsl":  {utils.NotEmpty()},
		"ExecTest": {utils.NotEmpty()},
		"ExecLoad": {utils.NotEmpty()},
	}
	if err := utils.Verify(slbCluster, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbClusterService.CreateSlbCluster(&slbCluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlbCluster 删除集群管理
// @Tags SlbCluster
// @Summary 删除集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbCluster true "删除集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbCluster/deleteSlbCluster [delete]
func (slbClusterApi *SlbClusterApi) DeleteSlbCluster(c *gin.Context) {
	var slbCluster slb.SlbCluster
	err := c.ShouldBindJSON(&slbCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbClusterService.DeleteSlbCluster(slbCluster); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlbClusterByIds 批量删除集群管理
// @Tags SlbCluster
// @Summary 批量删除集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slbCluster/deleteSlbClusterByIds [delete]
func (slbClusterApi *SlbClusterApi) DeleteSlbClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbClusterService.DeleteSlbClusterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlbCluster 更新集群管理
// @Tags SlbCluster
// @Summary 更新集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body slb.SlbCluster true "更新集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbCluster/updateSlbCluster [put]
func (slbClusterApi *SlbClusterApi) UpdateSlbCluster(c *gin.Context) {
	var slbCluster slb.SlbCluster
	err := c.ShouldBindJSON(&slbCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
		"SshUser":  {utils.NotEmpty()},
		"SshRsa":   {utils.NotEmpty()},
		"Nodes":    {utils.NotEmpty()},
		"PathConf": {utils.NotEmpty()},
		"PathSsl":  {utils.NotEmpty()},
		"ExecTest": {utils.NotEmpty()},
		"ExecLoad": {utils.NotEmpty()},
	}
	if err := utils.Verify(slbCluster, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slbClusterService.UpdateSlbCluster(slbCluster); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlbCluster 用id查询集群管理
// @Tags SlbCluster
// @Summary 用id查询集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slb.SlbCluster true "用id查询集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbCluster/findSlbCluster [get]
func (slbClusterApi *SlbClusterApi) FindSlbCluster(c *gin.Context) {
	var slbCluster slb.SlbCluster
	err := c.ShouldBindQuery(&slbCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslbCluster, err := slbClusterService.GetSlbCluster(slbCluster.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reslbCluster": reslbCluster}, c)
	}
}

// GetSlbClusterList 分页获取集群管理列表
// @Tags SlbCluster
// @Summary 分页获取集群管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query slbReq.SlbClusterSearch true "分页获取集群管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbCluster/getSlbClusterList [get]
func (slbClusterApi *SlbClusterApi) GetSlbClusterList(c *gin.Context) {
	var pageInfo slbReq.SlbClusterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slbClusterService.GetSlbClusterInfoList(pageInfo); err != nil {
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

func (slbClusterApi *SlbClusterApi) UploadSlbRsa(c *gin.Context) {
	r := c.Request
	r.ParseForm()
	file, handle, err := r.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
		return
	}

	filename := r.FormValue("filename")

	if !utils.FileExist(known.RootSlbRsa) {
		err := os.MkdirAll(known.RootSlbRsa, 0700)
		if err != nil {
			global.GVA_LOG.Error("上传失败!", zap.Error(err))
			response.FailWithMessage("上传失败", c)
			return
		}
	}

	SSHSrc := path.Join(known.RootSlbRsa, handle.Filename)
	if utils.FileExist(SSHSrc) {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
		return
	}

	SSHDst := path.Join(known.RootSlbRsa, filename)
	if utils.FileExist(SSHDst) {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
		return
	}

	f, err := os.OpenFile(SSHSrc, os.O_WRONLY|os.O_CREATE, 0600)
	io.Copy(f, file)
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
		return
	}
	defer f.Close()
	defer file.Close()
	err = os.Rename(SSHSrc, SSHDst)
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
		return
	} else {
		response.OkWithMessage("上传成功", c)
	}
}
