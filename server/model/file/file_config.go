// 自动生成模板FileConfig
package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 配置文件 结构体  FileConfig
type FileConfig struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`                         //名称
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:4096;"` //描述
	Content     string `json:"content" form:"content" gorm:"column:content;comment:配置内容;type:longtext;"`       //内容
	Md5         string `json:"md5" form:"md5" gorm:"column:md5;comment:MD5;size:64;"`                          //MD5
	DataType    *int   `json:"dataType" form:"dataType" gorm:"column:data_type;comment:文件类型;"`                 //类型
	Manager     string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:256;"`             //负责人
}

// TableName 配置文件 FileConfig自定义表名 file_config
func (FileConfig) TableName() string {
	return "file_config"
}
