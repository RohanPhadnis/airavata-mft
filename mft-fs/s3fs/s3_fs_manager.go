package main

// package s3fs

import (
	"encoding/json"
	"fmt"
	"github.com/jacobsa/fuse"
	"github.com/jacobsa/fuse/fuseops"
	"mft-fs/abstractfs"
	"net/http"
	"os"
	"time"
)

func main() {
	_, e := NewFSManager("./cache", "http://localhost:8000")
	if e != nil {
		fmt.Println(e)
	}
}

type FSManager struct {
	abstractfs.FSManager
	inodeInfo      map[fuseops.InodeID]*abstractfs.FileInfo
	handleToInode  map[fuseops.HandleID]fuseops.InodeID
	cacheDirectory string
	updateTime     time.Time
	serverURL      string
	currentInode   fuseops.InodeID
	maxInode       fuseops.InodeID
}

func NewFSManager(cacheDirectory string, serverURL string) (*FSManager, error) {
	output := &FSManager{
		inodeInfo:      make(map[fuseops.InodeID]*abstractfs.FileInfo),
		handleToInode:  make(map[fuseops.HandleID]fuseops.InodeID),
		cacheDirectory: cacheDirectory,
		serverURL:      serverURL,
		updateTime:     time.Unix(0, 0),
	}
	e := output.update()
	if e != nil {
		return nil, e
	}
	output.allocateInodes()
	return output, nil
}

const addDeviceURI = "/add_device"
const getDirectoryInfoURI = "/get_directory_info"
const allocateInodesURI = "/allocate_inodes"
const addInodeURI = "/add_inode"
const removeInodeURI = "/remove_inode"
const requestReadURI = "/request_read"
const ackReadURI = "/ack_read"
const requestWriteURI = "/request_write"
const ackWriteURI = "/ack_write"

type getResponseStruct struct {
	Name             string
	Path             string
	Inode            uint64
	Children         []uint64
	ChildrenIndexMap map[string]uint64
	Parent           uint64
	Nlink            uint32
	Size             uint64
	Mode             int
	Atime            uint64
	Mtime            uint64
	Ctime            uint64
	Crtime           uint64
	Uid              uint32
	Gid              uint32
	DirentType       uint8
	WriteTime        uint64
}

func (manager *FSManager) update() error {
	fmt.Println(manager.serverURL + getDirectoryInfoURI)
	response, e := http.Get(manager.serverURL + getDirectoryInfoURI)
	defer response.Body.Close()
	if e != nil {
		return e
	}
	var buff []byte = make([]byte, 12000)
	bytesRead, _ := response.Body.Read(buff)
	fmt.Println(string(buff[:bytesRead]))
	var data []getResponseStruct
	e = json.Unmarshal(buff[:bytesRead], &data)
	if e != nil {
		return e
	}
	manager.updateTime = time.Now()
	return nil
}

func (manager *FSManager) uploadFile()     {}
func (manager *FSManager) downloadFile()   {}
func (manager *FSManager) allocateInodes() {}
func (manager *FSManager) requestRead()    {}
func (manager *FSManager) ackRead()        {}
func (manager *FSManager) requestWrite()   {}
func (manager *FSManager) ackWrite()       {}

func (manager *FSManager) GetLength() uint64 {
	return uint64(len(manager.inodeInfo))
}

func (manager *FSManager) GetSize() uint64 {
	return uint64(4096)
}

func (manager *FSManager) GetInfo(inode fuseops.InodeID) (*abstractfs.FileInfo, error) {
	e := manager.update()
	if e != nil {
		return nil, e
	}
	info, ok := manager.inodeInfo[inode]
	if !ok {
		return nil, fuse.ENOENT
	}
	return info, nil
}

func (manager *FSManager) RmDir(inode fuseops.InodeID) error {
	return nil
}

func (manager *FSManager) CloseFile(inode fuseops.InodeID, file *os.File, write bool) {
	defer file.Close()
	if write {
		manager.uploadFile()
	}
}
