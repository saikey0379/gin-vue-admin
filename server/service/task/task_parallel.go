package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
)

type TaskParallelService struct {
}

// CreateTaskParallel 创建并行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParallelService *TaskParallelService) CreateTaskParallel(taskParallel *task.TaskParallel) (err error) {
	err = global.GVA_DB.Create(taskParallel).Error
	return err
}

// DeleteTaskParallel 删除并行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParallelService *TaskParallelService)DeleteTaskParallel(taskParallel task.TaskParallel) (err error) {
	err = global.GVA_DB.Delete(&taskParallel).Error
	return err
}

// DeleteTaskParallelByIds 批量删除并行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParallelService *TaskParallelService)DeleteTaskParallelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]task.TaskParallel{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTaskParallel 更新并行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParallelService *TaskParallelService)UpdateTaskParallel(taskParallel task.TaskParallel) (err error) {
	err = global.GVA_DB.Save(&taskParallel).Error
	return err
}

// GetTaskParallel 根据id获取并行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParallelService *TaskParallelService)GetTaskParallel(id uint) (taskParallel task.TaskParallel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&taskParallel).Error
	return
}

// GetTaskParallelInfoList 分页获取并行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParallelService *TaskParallelService)GetTaskParallelInfoList(info taskReq.TaskParallelSearch) (list []task.TaskParallel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&task.TaskParallel{})
    var taskParallels []task.TaskParallel
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
	
	err = db.Find(&taskParallels).Error
	return  taskParallels, total, err
}
