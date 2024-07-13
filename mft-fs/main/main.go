package main

import (
	"context"
	"github.com/jacobsa/fuse"
	"github.com/jacobsa/fuse/fuseutil"
	"log"
	"mft-fs/abstractfs"
	"mft-fs/osfsmanager"
	"os"
)

func main() {

	manager := osfsmanager.NewOSFSManager("./test")
	fs, _ := abstractfs.NewAbstractFS(manager)
	server := fuseutil.NewFileSystemServer(&fs)

	// mount the filesystem
	cfg := &fuse.MountConfig{
		ReadOnly:    false,
		DebugLogger: log.New(os.Stderr, "fuse: ", 0),
		ErrorLogger: log.New(os.Stderr, "fuse: ", 0),
	}
	mfs, err := fuse.Mount("./mount", server, cfg)
	if err != nil {
		log.Fatalf("Mount: %v", err)
	}

	// wait for it to be unmounted
	if err = mfs.Join(context.Background()); err != nil {
		log.Fatalf("Join: %v", err)
	}
}
