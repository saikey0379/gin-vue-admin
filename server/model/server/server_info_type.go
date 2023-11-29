package server

type NicInfo struct {
	Name string
	Mac  string
	Ip   string
}
type NicDevice struct {
	Id    string
	Model string
}
type CpuInfo struct {
	Id    string
	Model string
	Core  string
}
type DiskInfo struct {
	Name string
	Size string
}
type MemoryInfo struct {
	Name string
	Size string
	Type string
}
type GpuInfo struct {
	Id     string
	Model  string
	Memory string
}
type MotherboardInfo struct {
	Name  string
	Model string
}
