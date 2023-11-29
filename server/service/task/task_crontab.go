package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
)

type TaskCrontabService struct {
}

// CreateTaskCrontab 创建计划任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskCrontabService *TaskCrontabService) CreateTaskCrontab(taskCrontab *task.TaskCrontab) (err error) {
	err = global.GVA_DB.Create(taskCrontab).Error
	return err
}

// DeleteTaskCrontab 删除计划任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskCrontabService *TaskCrontabService)DeleteTaskCrontab(taskCrontab task.TaskCrontab) (err error) {
	err = global.GVA_DB.Delete(&taskCrontab).Error
	return err
}

// DeleteTaskCrontabByIds 批量删除计划任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskCrontabService *TaskCrontabService)DeleteTaskCrontabByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]task.TaskCrontab{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTaskCrontab 更新计划任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskCrontabService *TaskCrontabService)UpdateTaskCrontab(taskCrontab task.TaskCrontab) (err error) {
	err = global.GVA_DB.Save(&taskCrontab).Error
	return err
}

// GetTaskCrontab 根据id获取计划任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskCrontabService *TaskCrontabService)GetTaskCrontab(id uint) (taskCrontab task.TaskCrontab, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&taskCrontab).Error
	return
}

// GetTaskCrontabInfoList 分页获取计划任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskCrontabService *TaskCrontabService)GetTaskCrontabInfoList(info taskReq.TaskCrontabSearch) (list []task.TaskCrontab, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&task.TaskCrontab{})
    var taskCrontabs []task.TaskCrontab
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Manager != "" {
        db = db.Where("manager = ?",info.Manager)
    }
    if info.Nodes != "" {
        db = db.Where("nodes LIKE ?","%"+ info.Nodes+"%")
    }
    if info.FileType != nil {
        db = db.Where("file_type = ?",info.FileType)
    }
    if info.FileId != nil {
        db = db.Where("file_id = ?",info.FileId)
    }
    if info.DestPath != "" {
        db = db.Where("dest_path = ?",info.DestPath)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
         	orderMap["manager"] = true
         	orderMap["file_type"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&taskCrontabs).Error
	return  taskCrontabs, total, err
}
