package osInstall

type ServiceGroup struct {
	OsInstallSystemService
	OsInstallConfigPxeService
	OsInstallConfigKickstartService
	OsInstallQueueService
	OsInstallLogService
}
