package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"mft-fs/safe-remotefs/safe-remotefscomms"
	"mft-fs/safe-remotefs/server"
	"net"
)

func main() {

	rootDirectory := flag.String("rootDirectory", "", "Root Directory")
	flag.Parse()

	lis, e := net.Listen("tcp", ":8000")
	if e != nil {
		log.Fatalf("Error: %v\n", e)
	}
	defer lis.Close()

	s := grpc.NewServer()
	safe_remotefscomms.RegisterRemoteFSCommsServer(s, server.NewServerHandler(*rootDirectory))
	s.Serve(lis)
}
