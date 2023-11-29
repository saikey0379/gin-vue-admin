// 自动生成模板IdcCabinet
package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 数据中心机柜 结构体  IdcCabinet
type IdcCabinet struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //机柜 
      RoomId  *int `json:"room_id" form:"room_id" gorm:"column:room_id;comment:房间号;size:64;"`  //房间 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;type:text;"`  //描述 
}


// TableName 数据中心机柜 IdcCabinet自定义表名 idc_cabinet
func (IdcCabinet) TableName() string {
  return "idc_cabinet"
}

