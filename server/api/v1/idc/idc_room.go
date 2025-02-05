package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IdcRoomApi struct {
}

type RoomInfo struct {
	idc.IdcRoom
	IdcName string `json:"idc_name" form:"idc_name"`
}

var roomService = service.ServiceGroupApp.IdcServiceGroup.IdcRoomService

// CreateIdcRoom 创建数据中心房间
// @Tags IdcRoom
// @Summary 创建数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcRoom true "创建数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /room/createIdcRoom [post]
func (roomApi *IdcRoomApi) CreateIdcRoom(c *gin.Context) {
	var room idc.IdcRoom
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"IdcId":    {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(room, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.CreateIdcRoom(&room); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcRoom 删除数据中心房间
// @Tags IdcRoom
// @Summary 删除数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcRoom true "删除数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteIdcRoom [delete]
func (roomApi *IdcRoomApi) DeleteIdcRoom(c *gin.Context) {
	var room idc.IdcRoom
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.DeleteIdcRoom(room); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcRoomByIds 批量删除数据中心房间
// @Tags IdcRoom
// @Summary 批量删除数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /room/deleteIdcRoomByIds [delete]
func (roomApi *IdcRoomApi) DeleteIdcRoomByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.DeleteIdcRoomByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcRoom 更新数据中心房间
// @Tags IdcRoom
// @Summary 更新数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcRoom true "更新数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /room/updateIdcRoom [put]
func (roomApi *IdcRoomApi) UpdateIdcRoom(c *gin.Context) {
	var room idc.IdcRoom
	err := c.ShouldBindJSON(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"IdcId":    {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(room, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roomService.UpdateIdcRoom(room); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcRoom 用id查询数据中心房间
// @Tags IdcRoom
// @Summary 用id查询数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcRoom true "用id查询数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /room/findIdcRoom [get]
func (roomApi *IdcRoomApi) FindIdcRoom(c *gin.Context) {
	var room idc.IdcRoom
	err := c.ShouldBindQuery(&room)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reroom, err := roomService.GetIdcRoom(room.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reroom.IdcId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}

	var roomInfo = RoomInfo{
		reroom,
		reidcInfo.Name,
	}
	response.OkWithData(gin.H{"reroom": roomInfo}, c)
}

// GetIdcRoomList 分页获取数据中心房间列表
// @Tags IdcRoom
// @Summary 分页获取数据中心房间列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcRoomSearch true "分页获取数据中心房间列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/getIdcRoomList [get]
func (roomApi *IdcRoomApi) GetIdcRoomList(c *gin.Context) {
	var pageInfo idcReq.IdcRoomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := roomService.GetIdcRoomInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var mapIdc = make(map[int]string)
	var roomInfos []RoomInfo
	for _, room := range list {
		if mapIdc[*room.IdcId] == "" {
			reidcInfo, err := idcInfoService.GetIdcInfo(uint(*room.IdcId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
			}
			mapIdc[*room.IdcId] = reidcInfo.Name
		}
		var roomInfo = RoomInfo{
			room,
			mapIdc[*room.IdcId],
		}
		roomInfos = append(roomInfos, roomInfo)
	}

	response.OkWithDetailed(response.PageResult{
		List:     roomInfos,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
