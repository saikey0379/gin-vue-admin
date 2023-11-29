package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/idc"
	"github.com/flipped-aurora/gin-vue-admin/server/model/osInstall"

	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/server"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/slb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmdb"
	"github.com/flipped-aurora/gin-vue-admin/server/model/monit"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, idc.IdcInfo{}, idc.IdcCabinet{}, idc.IdcRoom{}, device.DeviceBareMetal{}, osInstall.OsInstallConfigPxe{}, osInstall.OsInstallConfigKickstart{}, idc.IdcIpSegment{}, idc.IdcIpSubnet{}, idc.IdcIpPreempt{}, osInstall.OsInstallQueue{}, osInstall.OsInstallLog{}, server.ServerDiscovery{}, server.ServerInfo{}, file.FileConfig{}, file.FileBinary{}, file.FileCommon{}, file.FileCommand{}, file.FileScript{}, file.FileCommadBlacklist{}, task.TaskParallel{}, task.TaskTimed{}, task.TaskCrontab{}, task.TaskRealtime{}, slb.SlbCluster{}, slb.SlbCert{}, slb.SlbAccesslist{}, slb.SlbUpstream{}, slb.SlbDomain{}, cmdb.CmdbIpSegment{}, cmdb.CmdbIpSubnet{}, cmdb.CmdbIpPreempt{}, cmdb.CmdbRegion{}, monit.RuleLabel{}, monit.RuleAnnotation{}, monit.RuleGroup{}, monit.PrometheusCluster{}, monit.RuleRecord{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
