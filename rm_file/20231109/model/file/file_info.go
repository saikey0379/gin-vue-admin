// 自动生成模板FileInfo
package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// PXE配置 结构体  FileInfo
type FileInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:256;"`  //名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:4096;"`  //描述 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:配置内容;size:4096;"`  //内容 
      Md5  string `json:"md5" form:"md5" gorm:"column:md5;comment:MD5;size:64;"`  //MD5 
      FileType  *int `json:"fileType" form:"fileType" gorm:"column:file_type;comment:文件类型;"`  //类型 
      Interpreter  *int `json:"interpreter" form:"interpreter" gorm:"column:interpreter;comment:解释器;"`  //解释器 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:256;"`  //负责人 
}


// TableName PXE配置 FileInfo自定义表名 file_info
func (FileInfo) TableName() string {
  return "file_info"
}

