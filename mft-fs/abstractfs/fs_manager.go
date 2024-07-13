package abstractfs

import (
	"github.com/jacobsa/fuse/fuseops"
	"github.com/jacobsa/fuse/fuseutil"
	"os"
	"time"
)

type FSManager interface {
	GetSize() uint64
	GetLength() uint64
	GetInfo(id fuseops.InodeID) (*FileInfo, error)
	SetInfo(id fuseops.InodeID, uidptr *uint32, gidptr *uint32, sizeptr *uint64, modeptr *os.FileMode, atimeptr *time.Time, mtimeptr *time.Time) error
	Delete(inode fuseops.InodeID) error
	MkDir(parent fuseops.InodeID, name string, mode os.FileMode) (fuseops.InodeID, error)
	GenerateHandle(id fuseops.InodeID) (fuseops.HandleID, error)
	CreateFile(parent fuseops.InodeID, name string, mode os.FileMode) (fuseops.InodeID, error)
	RmDir(inode fuseops.InodeID) error
	DeleteHandle(handle fuseops.HandleID)
	GetFile(inode fuseops.InodeID) (*os.File, error)
	CloseFile(inode fuseops.InodeID, file *os.File)
	SyncFile(inode fuseops.InodeID, file *os.File) error
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
	}
}
