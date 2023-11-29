// 自动生成模板IdcRoom
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 数据中心房间 结构体  IdcRoom
type IdcRoom struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`                //Name
	Label    string `json:"label" form:"label" gorm:"column:label;comment:房间标签;size:256;"`         //标签
	IdcId    *int   `json:"idc_id" form:"idc_id" gorm:"column:idc_id;comment:机房ID;size:64;"`       //IDC
	Describe string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:256;"` //描述
}

// TableName 数据中心房间 IdcRoom自定义表名 idc_room
func (IdcRoom) TableName() string {
	return "idc_room"
}
