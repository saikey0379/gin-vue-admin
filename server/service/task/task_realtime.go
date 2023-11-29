package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
)

type TaskRealtimeService struct {
}

// CreateTaskRealtime 创建实时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskRealtimeService *TaskRealtimeService) CreateTaskRealtime(taskRealtime *task.TaskRealtime) (err error) {
	err = global.GVA_DB.Create(taskRealtime).Error
	return err
}

// DeleteTaskRealtime 删除实时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskRealtimeService *TaskRealtimeService)DeleteTaskRealtime(taskRealtime task.TaskRealtime) (err error) {
	err = global.GVA_DB.Delete(&taskRealtime).Error
	return err
}

// DeleteTaskRealtimeByIds 批量删除实时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskRealtimeService *TaskRealtimeService)DeleteTaskRealtimeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]task.TaskRealtime{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTaskRealtime 更新实时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskRealtimeService *TaskRealtimeService)UpdateTaskRealtime(taskRealtime task.TaskRealtime) (err error) {
	err = global.GVA_DB.Save(&taskRealtime).Error
	return err
}

// GetTaskRealtime 根据id获取实时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskRealtimeService *TaskRealtimeService)GetTaskRealtime(id uint) (taskRealtime task.TaskRealtime, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&taskRealtime).Error
	return
}

// GetTaskRealtimeInfoList 分页获取实时任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskRealtimeService *TaskRealtimeService)GetTaskRealtimeInfoList(info taskReq.TaskRealtimeSearch) (list []task.TaskRealtime, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&task.TaskRealtime{})
    var taskRealtimes []task.TaskRealtime
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
	
	err = db.Find(&taskRealtimes).Error
	return  taskRealtimes, total, err
}
