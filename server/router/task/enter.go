package task

type RouterGroup struct {
	TaskParallelRouter
	TaskTimedRouter
	TaskCrontabRouter
	TaskRealtimeRouter
}
