// 自动生成模板FileCommand
package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 命令信息 结构体  FileCommand
type FileCommand struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`                         //名称
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:4096;"` //描述
	Content     string `json:"content" form:"content" gorm:"column:content;comment:配置内容;type:longtext;"`       //内容
	Md5         string `json:"md5" form:"md5" gorm:"column:md5;comment:MD5;size:64;"`                          //MD5
	Manager     string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:256;"`             //负责人
}

// TableName 命令信息 FileCommand自定义表名 file_command
func (FileCommand) TableName() string {
	return "file_command"
}
