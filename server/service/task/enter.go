package task

type ServiceGroup struct {
	TaskParallelService
	TaskTimedService
	TaskCrontabService
	TaskRealtimeService
}
