// 自动生成模板TaskCrontab
package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 计划任务 结构体  TaskCrontab
type TaskCrontab struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:任务名;size:256;"`  //Name 
      Manager  string `json:"manager" form:"manager" gorm:"column:manager;comment:负责人;size:64;"`  //负责人 
      Describe  string `json:"describe" form:"describe" gorm:"column:describe;comment:PXE;size:4096;"`  //描述 
      Nodes  string `json:"nodes" form:"nodes" gorm:"column:nodes;comment:节点;size:64;"`  //节点 
      FileType  *int `json:"fileType" form:"fileType" gorm:"column:file_type;comment:文件类型;size:8;"`  //文件类型 
      FileId  *int `json:"fileId" form:"fileId" gorm:"column:file_id;comment:文件ID;size:16;"`  //文件ID 
      FileMod  string `json:"fileMod" form:"fileMod" gorm:"column:file_mod;comment:文件权限;size:8;"`  //文件权限 
      Parameter  string `json:"parameter" form:"parameter" gorm:"column:parameter;comment:执行参数;size:256;"`  //执行参数 
      DestPath  string `json:"destPath" form:"destPath" gorm:"column:dest_path;comment:目标路径;size:256;"`  //目标路径 
      CronMinute  string `json:"cronMinute" form:"cronMinute" gorm:"column:cron_minute;comment:Minute;size:4;"`  //Minute 
      CronHour  string `json:"cronHour" form:"cronHour" gorm:"column:cron_hour;comment:Hour;size:4;"`  //Hour 
      CronDayOfMonth  string `json:"cronDayOfMonth" form:"cronDayOfMonth" gorm:"column:cron_day_of_month;comment:DayOfMonth;size:4;"`  //DayOfMonth 
      CronMonth  string `json:"cronMonth" form:"cronMonth" gorm:"column:cron_month;comment:Month;size:4;"`  //Month 
      CronDayOfWeek  string `json:"cronDayOfWeek" form:"cronDayOfWeek" gorm:"column:cron_day_of_week;comment:DayOfWeek;size:4;"`  //DayOfWeek 
}


// TableName 计划任务 TaskCrontab自定义表名 task_crontab
func (TaskCrontab) TableName() string {
  return "task_crontab"
}

