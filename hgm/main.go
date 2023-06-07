package main

import (
	"github.com/Bitspark/go-bitnode/api/wsApi"
	"github.com/Bitspark/go-bitnode/bitnode"
	"github.com/Bitspark/go-bitnode/library"
	"log"
	"os"
)

func main() {
	localAddress := os.Getenv("BITNODE_LOCAL_ADDRESS")
	remoteNodeAddress := os.Getenv("BITNODE_REMOTE_ADDRESS")

	node := bitnode.NewNode()
	dom := bitnode.NewDomain()
	node.AddMiddlewares(library.GetMiddlewares())

	// Prepare node connections.
	nodeConns := wsApi.NewNodeConns(node, remoteNodeAddress)

	// Prepare node.
	localDom, err := dom.AddDomain("local")
	if err != nil {
		log.Fatal(err)
	}
	if err := localDom.LoadFromFile("./local.yml"); err != nil {
		log.Fatal(err)
	}
	if err := localDom.Compile(); err != nil {
		log.Fatal(err)
	}

	// Create BlankSparkable system.
	spb, err := dom.GetSparkable("local.BlankSparkable")
	if err != nil {
		log.Fatal(err)
	}
	sys, err := node.PrepareSystem(bitnode.Credentials{}, *spb)
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance for the sparkable.
	s := &BlankSparkable{
		System: sys,
	}

	// Add the custom BlankSparkable implementation.
	if err := s.init(); err != nil {
		log.Fatal(err)
	}

	// Make BlankSparkable instance the root hologram.
	node.SetSystem(sys.Native())

	// Create server.
	server := wsApi.NewServer(nodeConns, localAddress)

	// Start server.
	log.Fatal(server.Listen())
}
