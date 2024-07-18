package osfsmanager

import (
	"crypto/rand"
	"fmt"
	"github.com/jacobsa/fuse"
	"github.com/jacobsa/fuse/fuseops"
	"github.com/jacobsa/fuse/fuseutil"
	"golang.org/x/sys/unix"
	"mft-fs/abstractfs"
	"mft-fs/datastructures"
	"os"
	"syscall"
	"time"
)

type OSFSManager struct {
	abstractfs.FSManager
	root          string
	inodeInfo     map[fuseops.InodeID]*abstractfs.FileInfo
	handleToInode map[fuseops.HandleID]fuseops.InodeID
	inodeCounter  fuseops.InodeID
}

type parentChildPair struct {
	parent fuseops.InodeID
	child  string
}

func (manager *OSFSManager) bfs(root string) {
	q := datastructures.NewQueue()
	q.Enqueue(&(parentChildPair{
		parent: 0,
		child:  root,
	}))
	for !q.IsEmpty() {
		current := q.Dequeue().(*parentChildPair)
		var direntType fuseutil.DirentType
		stat, e := os.Stat(current.child)
		if e != nil {
			fmt.Println("error:")
			fmt.Println(e.Error())
			continue
		}
		if stat.IsDir() {
			dir, e := os.ReadDir(current.child)
			if e != nil {
				continue
			}
			for _, file := range dir {
				q.Enqueue(&(parentChildPair{
					parent: manager.inodeCounter,
					child:  fmt.Sprintf("%s/%s", current.child, file.Name()),
				}))
			}
			direntType = fuseutil.DT_Directory
		} else {
			direntType = fuseutil.DT_File
		}
		info := abstractfs.NewFileInfo(stat.Name(), current.child, manager.inodeCounter, current.parent, direntType)
		manager.inodeInfo[manager.inodeCounter] = &info
		parentInfo, ok := manager.inodeInfo[current.parent]
		if ok {
			parentInfo.ChildrenIndexMap[info.Name] = len(parentInfo.Children)
			parentInfo.Children = append(parentInfo.Children, manager.inodeCounter)
		}
		manager.updateInfo(manager.inodeInfo[manager.inodeCounter])
		manager.inodeCounter++
	}
}

func NewOSFSManager(root string) *OSFSManager {
	manager := &OSFSManager{
		root:          root,
		inodeInfo:     make(map[fuseops.InodeID]*abstractfs.FileInfo),
		handleToInode: make(map[fuseops.HandleID]fuseops.InodeID),
		inodeCounter:  1,
	}
	manager.bfs(root)
	return manager
}

func (manager *OSFSManager) GenerateHandle(inode fuseops.InodeID) (fuseops.HandleID, error) {
	var buff [8]byte
	_, e := rand.Read(buff[:])
	if e != nil {
		return 0, e
	}
	var output uint64 = 0
	for i := 0; i < len(buff); i++ {
		output = output | uint64(buff[i]<<(8*i))
	}
	handle := fuseops.HandleID(output)
	_, ok := manager.handleToInode[handle]
	if ok {
		return manager.GenerateHandle(inode)
	}
	info, ok := manager.inodeInfo[inode]
	if !ok {
		return 0, fuse.ENOENT
	}
	info.Handle = handle
	manager.handleToInode[handle] = inode
	return handle, nil
}

/*

todo
	update
	fillAttributes
*/

func (manager *OSFSManager) updateInfo(info *abstractfs.FileInfo) error {

	fileInfo, e := os.Stat(info.Path)
	if e != nil {
		return e
	}
	stats := fileInfo.Sys().(*syscall.Stat_t)

	info.Size = uint64(fileInfo.Size())

	info.Nlink = 1

	info.Mode = fileInfo.Mode()

	/*info.Atime = time.Unix(stats.Atimespec.Sec, stats.Atimespec.Nsec)
	info.Mtime = time.Unix(stats.Mtimespec.Sec, stats.Mtimespec.Nsec)
	info.Ctime = time.Unix(stats.Ctimespec.Sec, stats.Ctimespec.Nsec)
	info.Crtime = time.Unix(stats.Birthtimespec.Sec, stats.Birthtimespec.Nsec)*/
	var timeStats unix.Stat_t
	e = unix.Stat(info.Path, &timeStats)
	if e != nil {
		return e
	}
	info.Atime = time.Unix(timeStats.Atim.Sec, timeStats.Atim.Nsec)
	info.Mtime = time.Unix(timeStats.Mtim.Sec, timeStats.Mtim.Nsec)
	info.Ctime = time.Unix(timeStats.Ctim.Sec, timeStats.Ctim.Nsec)
	info.Crtime = time.Unix(timeStats.Ctim.Sec, timeStats.Ctim.Nsec)

	info.Uid = stats.Uid
	info.Gid = stats.Gid

	return nil
}

func (manager *OSFSManager) GetSize() uint64 {
	stat, _ := os.Stat(manager.root)
	return uint64(stat.Size())
}

func (manager *OSFSManager) GetLength() uint64 {
	return uint64(len(manager.inodeInfo))
}

func (manager *OSFSManager) GetInfo(inode fuseops.InodeID) (*abstractfs.FileInfo, error) {
	info, ok := manager.inodeInfo[inode]
	if !ok {
		return nil, fuse.ENOENT
	}
	return info, nil
}

func (manager *OSFSManager) SetInfo(inode fuseops.InodeID, uidptr *uint32, gidptr *uint32, sizeptr *uint64, modeptr *os.FileMode, atimeptr *time.Time, mtimeptr *time.Time) error {
	info, e := manager.GetInfo(inode)
	if e != nil {
		return e
	}

	if modeptr != nil {
		e = os.Chmod(info.Path, *modeptr)
		if e != nil {
			return e
		}
	}

	uid := -1
	gid := -1
	if uidptr != nil {
		uid = int(*uidptr)
	}
	if gidptr != nil {
		gid = int(*gidptr)
	}
	e = os.Chown(info.Path, uid, gid)
	if e != nil {
		return e
	}

	atime := time.Time{}
	mtime := time.Time{}
	if atimeptr != nil {
		atime = *atimeptr
	}
	if mtimeptr != nil {
		mtime = *mtimeptr
	}
	e = os.Chtimes(info.Path, atime, mtime)
	if e != nil {
		return e
	}

	return manager.updateInfo(info)
}

func (manager *OSFSManager) MkDir(parent fuseops.InodeID, name string, mode os.FileMode) (fuseops.InodeID, error) {
	parentInfo, e := manager.GetInfo(parent)
	if e != nil {
		return 0, e
	}

	path := parentInfo.Path + "/" + name
	e = os.Mkdir(path, mode)
	if e != nil {
		return 0, e
	}

	inode := manager.inodeCounter //fuseops.InodeID(osInfo.Sys().(*syscall.Stat_t).Ino)
	manager.inodeCounter++
	infoObj := abstractfs.NewFileInfo(name, path, inode, parent, fuseutil.DT_Directory)
	manager.inodeInfo[inode] = &infoObj
	e = manager.updateInfo(&infoObj)
	if e != nil {
		return 0, e
	}

	parentInfo.Children = append(parentInfo.Children, inode)
	parentInfo.ChildrenIndexMap[name] = len(parentInfo.Children) - 1

	return inode, nil
}

func (manager *OSFSManager) CreateFile(parent fuseops.InodeID, name string, mode os.FileMode) (fuseops.InodeID, error) {
	parentInfo, e := manager.GetInfo(parent)
	if e != nil {
		return 0, e
	}

	path := parentInfo.Path + "/" + name
	file, e := os.Create(path)
	defer file.Close()
	if e != nil {
		return 0, e
	}

	e = file.Chmod(mode)
	if e != nil {
		return 0, e
	}

	inode := manager.inodeCounter //fuseops.InodeID(osInfo.Sys().(*syscall.Stat_t).Ino)
	manager.inodeCounter++
	infoObj := abstractfs.NewFileInfo(name, path, inode, parent, fuseutil.DT_File)
	manager.inodeInfo[inode] = &infoObj
	e = manager.updateInfo(&infoObj)
	if e != nil {
		return 0, e
	}

	parentInfo.Children = append(parentInfo.Children, inode)
	parentInfo.ChildrenIndexMap[name] = len(parentInfo.Children) - 1

	return inode, nil
}

func (manager *OSFSManager) RmDir(inode fuseops.InodeID) error {
	return nil
}

func (manager *OSFSManager) Delete(inode fuseops.InodeID) error {

	// first delete any reference in parents
	// then delete self

	info, e := manager.GetInfo(inode)
	if e != nil {
		return e
	}

	parentInfo, e := manager.GetInfo(info.Parent)
	if e != nil {
		return e
	}

	parentInfo.Children = datastructures.Remove(parentInfo.Children, parentInfo.ChildrenIndexMap[info.Name])
	target := parentInfo.ChildrenIndexMap[info.Name]
	for name, index := range parentInfo.ChildrenIndexMap {
		if index > target {
			parentInfo.ChildrenIndexMap[name] = index - 1
		}
	}
	delete(parentInfo.ChildrenIndexMap, info.Name)

	delete(manager.inodeInfo, info.Inode)

	return nil
}

func (manager *OSFSManager) DeleteHandle(handle fuseops.HandleID) {
	delete(manager.handleToInode, handle)
}

func (manager *OSFSManager) SyncFile(inode fuseops.InodeID) error {
	info, e := manager.GetInfo(inode)
	if e != nil {
		return e
	}
	file, e := os.OpenFile(info.Path, os.O_RDWR, info.Mode)
	if e != nil {
		return e
	}
	defer file.Close()
	e = file.Sync()
	if e != nil {
		return e
	}
	return nil
}

func (manager *OSFSManager) ReadAt(inode fuseops.InodeID, data []byte, off int64) (int, error) {
	info, e := manager.GetInfo(inode)
	if e != nil {
		return 0, e
	}

	file, e := os.OpenFile(info.Path, os.O_RDWR, info.Mode)
	if e != nil {
		return 0, e
	}
	defer file.Close()

	n, e := file.ReadAt(data, off)
	return n, e
}

func (manager *OSFSManager) WriteAt(inode fuseops.InodeID, data []byte, off int64) (int, error) {
	info, e := manager.GetInfo(inode)
	if e != nil {
		return 0, e
	}

	file, e := os.OpenFile(info.Path, os.O_RDWR, info.Mode)
	if e != nil {
		return 0, e
	}
	defer file.Close()

	n, e := file.WriteAt(data, off)
	return n, e
}
