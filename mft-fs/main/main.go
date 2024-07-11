package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"mft-fs/basicfs"

	"github.com/jacobsa/fuse"
	"github.com/jacobsa/fuse/fuseutil"
)

func main() {

	// reading commandline args
	mountDirectory := flag.String("mountDirectory", "", "mount directory")
	rootDirectory := flag.String("rootDirectory", "", "root directory")
	flag.Parse()

	fmt.Println(fmt.Sprintf("mountDirectory: %s, rootDirectory: %s", *mountDirectory, *rootDirectory))
	// create an appropriate file system
	// printer("started")
	fs, _ := basicfs.NewBasicFS(*rootDirectory)
	server := fuseutil.NewFileSystemServer(&fs)

	// mount the filesystem
	cfg := &fuse.MountConfig{
		ReadOnly:    false,
		DebugLogger: log.New(os.Stderr, "fuse: ", 0),
		ErrorLogger: log.New(os.Stderr, "fuse: ", 0),
	}
	mfs, err := fuse.Mount(*mountDirectory, server, cfg)
	if err != nil {
		log.Fatalf("Mount: %v", err)
	}

	// wait for it to be unmounted
	if err = mfs.Join(context.Background()); err != nil {
		log.Fatalf("Join: %v", err)
	}
}
