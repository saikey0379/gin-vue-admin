package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/agent"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/service/config"
	"github.com/flipped-aurora/gin-vue-admin/server/service/device"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/file"
	filehandle "github.com/flipped-aurora/gin-vue-admin/server/service/fileHandle"
	"github.com/flipped-aurora/gin-vue-admin/server/service/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/service/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/service/osInstall"
	"github.com/flipped-aurora/gin-vue-admin/server/service/server"
	"github.com/flipped-aurora/gin-vue-admin/server/service/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/task"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	AgentServiceGroup      agent.ServiceGroup
	OsInstallServiceGroup  osInstall.ServiceGroup
	IdcServiceGroup        idc.ServiceGroup
	DeviceServiceGroup     device.ServiceGroup
	ServerServiceGroup     server.ServiceGroup
	FileServiceGroup       file.ServiceGroup
	FileHandleServiceGroup filehandle.ServiceGroup
	TaskServiceGroup       task.ServiceGroup
	SlbServiceGroup        slb.ServiceGroup
	ConfigServiceGroup     config.ServiceGroup
	CmdbServiceGroup       cmdb.ServiceGroup
	MonitServiceGroup      monit.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
