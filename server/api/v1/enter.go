package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/config"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/device"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/file"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/fileHandle"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/osInstall"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/server"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/task"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	OsInstallApiGroup  osInstall.ApiGroup
	IdcApiGroup        idc.ApiGroup
	DeviceApiGroup     device.ApiGroup
	ServerApiGroup     server.ApiGroup
	FileApiGroup       file.ApiGroup
	TaskApiGroup       task.ApiGroup
	SlbApiGroup        slb.ApiGroup
	ConfigApiGroup     config.ApiGroup
	CmdbApiGroup       cmdb.ApiGroup
	MonitApiGroup      monit.ApiGroup
	FileHandleApiGroup fileHandle.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
