package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
)

type TaskTimedService struct {
}

// CreateTaskTimed 创建定时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskTimedService *TaskTimedService) CreateTaskTimed(taskTimed *task.TaskTimed) (err error) {
	err = global.GVA_DB.Create(taskTimed).Error
	return err
}

// DeleteTaskTimed 删除定时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskTimedService *TaskTimedService)DeleteTaskTimed(taskTimed task.TaskTimed) (err error) {
	err = global.GVA_DB.Delete(&taskTimed).Error
	return err
}

// DeleteTaskTimedByIds 批量删除定时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskTimedService *TaskTimedService)DeleteTaskTimedByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]task.TaskTimed{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTaskTimed 更新定时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskTimedService *TaskTimedService)UpdateTaskTimed(taskTimed task.TaskTimed) (err error) {
	err = global.GVA_DB.Save(&taskTimed).Error
	return err
}

// GetTaskTimed 根据id获取定时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskTimedService *TaskTimedService)GetTaskTimed(id uint) (taskTimed task.TaskTimed, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&taskTimed).Error
	return
}

// GetTaskTimedInfoList 分页获取定时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskTimedService *TaskTimedService)GetTaskTimedInfoList(info taskReq.TaskTimedSearch) (list []task.TaskTimed, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&task.TaskTimed{})
    var taskTimeds []task.TaskTimed
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
	
	err = db.Find(&taskTimeds).Error
	return  taskTimeds, total, err
}
