package file

type RouterGroup struct {
	FileConfigRouter
	FileBinaryRouter
	FileCommonRouter
	FileCommandRouter
	FileScriptRouter
	FileCommadBlacklistRouter
}
