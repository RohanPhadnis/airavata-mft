package abstractfs

import (
	"github.com/jacobsa/fuse/fuseops"
	"github.com/jacobsa/fuse/fuseutil"
	"os"
	"time"
)

type FSManager interface {
	GetSize() (uint64, error)
	GetLength() (uint64, error)
	GetInfo(id fuseops.InodeID) (*FileInfo, error)
	SetInfo(id fuseops.InodeID, uidptr *uint32, gidptr *uint32, sizeptr *uint64, modeptr *os.FileMode, atimeptr *time.Time, mtimeptr *time.Time) error
	Delete(inode fuseops.InodeID) error
	MkDir(parent fuseops.InodeID, name string, mode os.FileMode) (fuseops.InodeID, error)
	GenerateHandle(id fuseops.InodeID) (fuseops.HandleID, error)
	CreateFile(parent fuseops.InodeID, name string, mode os.FileMode) (fuseops.InodeID, error)
	RmDir(inode fuseops.InodeID) error
	DeleteHandle(handle fuseops.HandleID) error
	SyncFile(inode fuseops.InodeID) error
	WriteAt(inode fuseops.InodeID, data []byte, off int64) (n int, err error)
	ReadAt(inode fuseops.InodeID, data []byte, off int64) (n int, err error)
	Destroy() error
}

type FileInfo struct {
	Name             string
	Path             string
	Inode            fuseops.InodeID
	Children         []fuseops.InodeID
	ChildrenIndexMap map[string]int
	Parent           fuseops.InodeID
	Nlink            uint32
	Size             uint64
	Mode             os.FileMode
	Atime            time.Time
	Mtime            time.Time
	Ctime            time.Time
	Crtime           time.Time
	Uid              uint32
	Gid              uint32
	DirentType       fuseutil.DirentType
	Handle           fuseops.HandleID
	Cache            bool
	CacheTime        time.Time
	WriteTime        time.Time
}

func NewFileInfo(name string, path string, inode fuseops.InodeID, parent fuseops.InodeID, direntType fuseutil.DirentType) FileInfo {
	return FileInfo{
		Name:             name,
		Children:         make([]fuseops.InodeID, 0),
		ChildrenIndexMap: make(map[string]int),
		Parent:           parent,
		Inode:            inode,
		Path:             path,
		DirentType:       direntType,
		Cache:            false,
		CacheTime:        time.Unix(0, 0),
		WriteTime:        time.Now(),
	}
}
