package server

import (
	"context"
	"github.com/jacobsa/fuse/fuseops"
	"mft-fs/osfsmanager"
	"mft-fs/safe-remotefs/safe-remotefscomms"
	"os"
	"time"
)

type Server struct {
	safe_remotefscomms.UnimplementedRemoteFSCommsServer
	manager *osfsmanager.OSFSManager
}

func NewServerHandler(root string) *Server {
	output := &Server{
		manager: osfsmanager.NewOSFSManager(root),
	}
	return output
}

func (server *Server) GetSizeComm(ctx context.Context, in *safe_remotefscomms.Empty) (*safe_remotefscomms.UintMsg, error) {
	size, e := server.manager.GetSize()
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.UintMsg{Data: size}, nil
}

func (server *Server) GetLengthComm(ctx context.Context, in *safe_remotefscomms.Empty) (*safe_remotefscomms.UintMsg, error) {
	length, e := server.manager.GetLength()
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.UintMsg{Data: length}, nil
}

func (server *Server) GetInfoComm(ctx context.Context, in *safe_remotefscomms.UintMsg) (*safe_remotefscomms.FileInfoMsg, error) {
	info, e := server.manager.GetInfo(fuseops.InodeID(in.Data))
	if e != nil {
		return nil, e
	}
	output := &safe_remotefscomms.FileInfoMsg{
		ChildrenIndexMap: make(map[string]uint32),
	}
	safe_remotefscomms.ConvertToComm(info, output)
	return output, nil
}

func (server *Server) SetInfoComm(ctx context.Context, in *safe_remotefscomms.SetInfoParamsMsg) (*safe_remotefscomms.Empty, error) {
	var uidptr *uint32 = nil
	var gidptr *uint32 = nil
	var sizeptr *uint64 = nil
	var modeptr *os.FileMode = nil
	var atimeptr *time.Time = nil
	var mtimeptr *time.Time = nil
	if in.Uid != -1 {
		*uidptr = uint32(in.Uid)
	}
	if in.Gid != -1 {
		*gidptr = uint32(in.Gid)
	}
	if in.Mode != -1 {
		*modeptr = os.FileMode(in.Mode)
	}
	if in.Size != -1 {
		*sizeptr = uint64(in.Size)
	}
	if !in.Atime.AsTime().Equal(time.Time{}) {
		*atimeptr = in.Atime.AsTime()
	}
	if !in.Mtime.AsTime().Equal(time.Time{}) {
		*mtimeptr = in.Mtime.AsTime()
	}
	e := server.manager.SetInfo(fuseops.InodeID(in.Inode), uidptr, gidptr, sizeptr, modeptr, atimeptr, mtimeptr)
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.Empty{}, nil
}

func (server *Server) DeleteComm(ctx context.Context, in *safe_remotefscomms.UintMsg) (*safe_remotefscomms.Empty, error) {
	e := server.manager.Delete(fuseops.InodeID(in.Data))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.Empty{}, nil
}

func (server *Server) MkDirComm(ctx context.Context, in *safe_remotefscomms.MkInodeMsg) (*safe_remotefscomms.UintMsg, error) {
	inode, e := server.manager.MkDir(fuseops.InodeID(in.Parent), in.Name, os.FileMode(in.Mode))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.UintMsg{Data: uint64(inode)}, nil
}

func (server *Server) GenerateHandleComm(ctx context.Context, in *safe_remotefscomms.UintMsg) (*safe_remotefscomms.UintMsg, error) {
	handle, e := server.manager.GenerateHandle(fuseops.InodeID(in.Data))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.UintMsg{Data: uint64(handle)}, nil
}

func (server *Server) CreateFileComm(ctx context.Context, in *safe_remotefscomms.MkInodeMsg) (*safe_remotefscomms.UintMsg, error) {
	inode, e := server.manager.CreateFile(fuseops.InodeID(in.Parent), in.Name, os.FileMode(in.Mode))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.UintMsg{Data: uint64(inode)}, nil
}

func (server *Server) RmDirComm(ctx context.Context, in *safe_remotefscomms.UintMsg) (*safe_remotefscomms.Empty, error) {
	e := server.manager.RmDir(fuseops.InodeID(in.Data))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.Empty{}, nil
}

func (server *Server) DeleteHandleComm(ctx context.Context, in *safe_remotefscomms.UintMsg) (*safe_remotefscomms.Empty, error) {
	e := server.manager.DeleteHandle(fuseops.HandleID(in.Data))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.Empty{}, nil
}

func (server *Server) SyncFileComm(ctx context.Context, in *safe_remotefscomms.UintMsg) (*safe_remotefscomms.Empty, error) {
	e := server.manager.SyncFile(fuseops.InodeID(in.Data))
	if e != nil {
		return nil, e
	}
	return &safe_remotefscomms.Empty{}, nil
}

func (server *Server) WriteAtComm(ctx context.Context, in *safe_remotefscomms.ContentMsg) (*safe_remotefscomms.UintMsg, error) {
	n, _ := server.manager.WriteAt(fuseops.InodeID(in.Inode), in.Data, in.Off)
	return &safe_remotefscomms.UintMsg{Data: uint64(n)}, nil
}

func (server *Server) ReadAtComm(ctx context.Context, in *safe_remotefscomms.ContentMsg) (*safe_remotefscomms.UintMsg, error) {
	n, _ := server.manager.ReadAt(fuseops.InodeID(in.Inode), in.Data, in.Off)
	return &safe_remotefscomms.UintMsg{Data: uint64(n)}, nil
}