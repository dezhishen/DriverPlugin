package model

import "github.com/OpenListTeam/OpenList/v5/layers/file"

// DriverPlugin 插件接口
// Wasm插件需要实现该接口
type DriverPlugin interface {
	// Name 插件名称
	Name() string
	// CopyFile 复制文件
	CopyFile(sources []string, targets []string) ([]*file.BackFileAction, error)
	// MoveFile 移动文件
	MoveFile(sources []string, targets []string) ([]*file.BackFileAction, error)
	// ListFile 列举文件
	ListFile(path []string, opt *file.ListFileOption) ([]*file.HostFileObject, error)
	// FindFile 搜索文件
	FindFile(path []string, opt *file.FindFileOption) ([]*file.HostFileObject, error)
	// Download 获取文件
	Download(path []string, opt *file.DownloadOption) ([]*file.LinkFileObject, error)
	// Uploader 上传文件
	Uploader(path []string, opt *file.UploaderOption) ([]*file.BackFileAction, error)
	// KillFile 删除文件
	RemoveFile(path []string, opt *file.KillFileOption) ([]*file.BackFileAction, error)
	// MakeFile 创建文件
	MakeFile(path []string, opt *file.MakeFileOption) ([]*file.BackFileAction, error)
	// MakePath 创建目录
	MakePath(path []string, opt *file.MakeFileOption) ([]*file.BackFileAction, error)
}
