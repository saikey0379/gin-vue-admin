package slb

type ServiceGroup struct {
	SlbClusterService
	SlbCertService
	SlbAccesslistService
	SlbUpstreamService
	SlbDomainService
}
