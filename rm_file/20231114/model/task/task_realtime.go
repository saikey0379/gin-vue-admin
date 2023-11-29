// 自动生成模板TaskRealtime
package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 实时任务 结构体  TaskRealtime
type TaskRealtime struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`                 //Name
	Manager   string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`         //负责人
	Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`    //描述
	Nodes     string `json:"nodes" form:"nodes" gorm:"column:nodes;comment:节点;type:longtext;"`          //节点
	FileType  *int   `json:"fileType" form:"fileType" gorm:"column:file_type;comment:文件类型;size:8;"`     //文件类型
	FileId    *int   `json:"fileId" form:"fileId" gorm:"column:file_id;comment:文件ID;size:16;"`          //文件ID
	FileMod   string `json:"fileMod" form:"fileMod" gorm:"column:file_mod;comment:文件权限;size:8;"`        //文件权限
	Parameter string `json:"parameter" form:"parameter" gorm:"column:parameter;comment:执行参数;size:256;"` //执行参数
	DestPath  string `json:"destPath" form:"destPath" gorm:"column:dest_path;comment:目标路径;size:256;"`   //目标路径
}

// TableName 实时任务 TaskRealtime自定义表名 task_realtime
func (TaskRealtime) TableName() string {
	return "task_realtime"
}
