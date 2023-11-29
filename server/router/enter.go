package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/router/config"
	"github.com/flipped-aurora/gin-vue-admin/server/router/device"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/file"
	"github.com/flipped-aurora/gin-vue-admin/server/router/fileHandle"
	"github.com/flipped-aurora/gin-vue-admin/server/router/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/router/monit"
	"github.com/flipped-aurora/gin-vue-admin/server/router/osInstall"
	"github.com/flipped-aurora/gin-vue-admin/server/router/server"
	"github.com/flipped-aurora/gin-vue-admin/server/router/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/task"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	OsInstall  osInstall.RouterGroup
	Idc        idc.RouterGroup
	Device     device.RouterGroup
	Server     server.RouterGroup
	File       file.RouterGroup
	FileHandle fileHandle.RouterGroup
	Task       task.RouterGroup
	Slb        slb.RouterGroup
	Config     config.RouterGroup
	Cmdb       cmdb.RouterGroup
	Monit      monit.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
