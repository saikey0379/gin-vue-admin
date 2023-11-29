package idc

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	idcReq "github.com/flipped-aurora/gin-vue-admin/server/model/idc/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type IdcCabinetApi struct {
}

type CabinetInfo struct {
	idc.IdcCabinet
	RoomName string `json:"room_name" form:"room_name"`
	IdcId    int    `json:"idc_id" form:"idc_id"`
	IdcName  string `json:"idc_name" form:"idc_name"`
}

var idcCabinetService = service.ServiceGroupApp.IdcServiceGroup.IdcCabinetService

// CreateIdcCabinet 创建数据中心机柜
// @Tags IdcCabinet
// @Summary 创建数据中心机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcCabinet true "创建数据中心机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcCabinet/createIdcCabinet [post]
func (idcCabinetApi *IdcCabinetApi) CreateIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindJSON(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"RoomId":   {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(idcCabinet, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcCabinetService.CreateIdcCabinet(&idcCabinet); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdcCabinet 删除数据中心机柜
// @Tags IdcCabinet
// @Summary 删除数据中心机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcCabinet true "删除数据中心机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcCabinet/deleteIdcCabinet [delete]
func (idcCabinetApi *IdcCabinetApi) DeleteIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindJSON(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcCabinetService.DeleteIdcCabinet(idcCabinet); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdcCabinetByIds 批量删除数据中心机柜
// @Tags IdcCabinet
// @Summary 批量删除数据中心机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idcCabinet/deleteIdcCabinetByIds [delete]
func (idcCabinetApi *IdcCabinetApi) DeleteIdcCabinetByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcCabinetService.DeleteIdcCabinetByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdcCabinet 更新数据中心机柜
// @Tags IdcCabinet
// @Summary 更新数据中心机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body idc.IdcCabinet true "更新数据中心机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcCabinet/updateIdcCabinet [put]
func (idcCabinetApi *IdcCabinetApi) UpdateIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindJSON(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Label":    {utils.NotEmpty()},
		"RoomId":   {utils.NotEmpty()},
		"Describe": {utils.NotEmpty()},
	}
	if err := utils.Verify(idcCabinet, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := idcCabinetService.UpdateIdcCabinet(idcCabinet); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdcCabinet 用id查询数据中心机柜
// @Tags IdcCabinet
// @Summary 用id查询数据中心机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idc.IdcCabinet true "用id查询数据中心机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcCabinet/findIdcCabinet [get]
func (idcCabinetApi *IdcCabinetApi) FindIdcCabinet(c *gin.Context) {
	var idcCabinet idc.IdcCabinet
	err := c.ShouldBindQuery(&idcCabinet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reidcCabinet, err := idcCabinetService.GetIdcCabinet(idcCabinet.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	reroom, err := roomService.GetIdcRoom(uint(*reidcCabinet.RoomId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reroom.IdcId))
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var cabinetInfo = CabinetInfo{
		reidcCabinet,
		reroom.Name,
		*reroom.IdcId,
		reidcInfo.Name,
	}
	response.OkWithData(gin.H{"reidcCabinet": cabinetInfo}, c)
}

// GetIdcCabinetList 分页获取数据中心机柜列表
// @Tags IdcCabinet
// @Summary 分页获取数据中心机柜列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query idcReq.IdcCabinetSearch true "分页获取数据中心机柜列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcCabinet/getIdcCabinetList [get]
func (idcCabinetApi *IdcCabinetApi) GetIdcCabinetList(c *gin.Context) {
	var pageInfo idcReq.IdcCabinetSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	roomId := c.Query("RoomId")
	if roomId != "" {
		id, err := strconv.Atoi(roomId)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("RoomId[%s]解析失败", roomId), c)
			return
		}
		pageInfo.RoomId = &id
	}

	list, total, err := idcCabinetService.GetIdcCabinetInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	var mapIdcId = make(map[int]int)
	var mapIdcName = make(map[int]string)
	var mapRoomName = make(map[int]string)
	var cabinetsInfo []CabinetInfo
	for _, cabinet := range list {
		if mapRoomName[*cabinet.RoomId] == "" {
			reroom, err := roomService.GetIdcRoom(uint(*cabinet.RoomId))
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
				return
			}
			mapRoomName[*cabinet.RoomId] = reroom.Name
			mapIdcId[*cabinet.RoomId] = *reroom.IdcId

			if mapIdcName[*reroom.IdcId] == "" {
				reidcInfo, err := idcInfoService.GetIdcInfo(uint(*reroom.IdcId))
				if err != nil {
					global.GVA_LOG.Error("查询失败!", zap.Error(err))
					response.FailWithMessage("查询失败", c)
					return
				}
				mapIdcName[*cabinet.RoomId] = reidcInfo.Name
			}
		}
		var cabinetInfo = CabinetInfo{
			cabinet,
			mapRoomName[*cabinet.RoomId],
			mapIdcId[*cabinet.RoomId],
			mapIdcName[*cabinet.RoomId],
		}
		cabinetsInfo = append(cabinetsInfo, cabinetInfo)
	}
	response.OkWithDetailed(response.PageResult{
		List:     cabinetsInfo,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
