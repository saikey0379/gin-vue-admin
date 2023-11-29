// 自动生成模板FileScript
package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 脚本文件 结构体  FileScript
type FileScript struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`                         //名称
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:4096;"` //描述
	Interpreter *int   `json:"interpreter" form:"interpreter" gorm:"column:interpreter;comment:解释器;"`          //解释器
	Content     string `json:"content" form:"content" gorm:"column:content;comment:配置内容;type:longtext;"`       //内容
	Md5         string `json:"md5" form:"md5" gorm:"column:md5;comment:MD5;size:64;"`                          //MD5
	Manager     string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:256;"`             //负责人
}

// TableName 脚本文件 FileScript自定义表名 file_script
func (FileScript) TableName() string {
	return "file_script"
}
