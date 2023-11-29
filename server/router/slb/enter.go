package slb

type RouterGroup struct {
	SlbClusterRouter
	SlbCertRouter
	SlbAccesslistRouter
	SlbUpstreamRouter
	SlbDomainRouter
}
